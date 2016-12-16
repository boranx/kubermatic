package aws

import (
	"bytes"
	"errors"
	"fmt"
	"text/template"
	"time"

	"github.com/golang/glog"
	ktemplate "github.com/kubermatic/api/template"
	"golang.org/x/net/context"

	"encoding/base64"

	sdk "github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/kubermatic/api"
	"github.com/kubermatic/api/provider"
)

const (
	// VPCCidrBlock is the default CIDR block in the VpcSubnet
	VPCCidrBlock = "10.10.0.0/16"
	// SubnetCidrBlock is the default CIDR for the VPC
	SubnetCidrBlock = "10.10.10.0/24"
)

const (
	accessKeyIDAnnotationKey     = "acccess-key-id"
	secretAccessKeyAnnotationKey = "secret-access-key"
	sshKeyNameKey                = "ssh-key-fingerprint"
	subnetIDKey                  = "subnet-id"
	vpcIDKey                     = "vpc-id"
	internetGatewayIDKey         = "internet-gateway-id"
	routeTableIDKey              = "route-table-id"
)

const (
	awsFilterName    = "Name"
	awsFilterState   = "instance-state-name"
	awsFilterRunning = "running"
	awsFilterPending = "pending"
)

const (
	tplPath = "template/coreos/aws-cloud-config-node.yaml"
)

var (
	defaultKubermaticClusterNameTagKey = "kubermatic-cluster-name"
	defaultKubermaticClusterIDTagKey   = "kubermatic-cluster-id"
)

type aws struct {
	datacenters map[string]provider.DatacenterMeta
}

// NewCloudProvider returns a new aws provider.
func NewCloudProvider(datacenters map[string]provider.DatacenterMeta) provider.CloudProvider {
	return &aws{
		datacenters: datacenters,
	}
}

func createVpc(svc *ec2.EC2) (*ec2.Vpc, error) {
	vReq := &ec2.CreateVpcInput{
		CidrBlock:       sdk.String(VPCCidrBlock),
		InstanceTenancy: sdk.String(ec2.TenancyDefault),
	}
	vpcOut, err := svc.CreateVpc(vReq)
	if err != nil {
		return nil, err
	}

	return vpcOut.Vpc, nil
}

func createSubnet(svc *ec2.EC2, vpc *ec2.Vpc) (*ec2.Subnet, error) {
	sOut, err := svc.CreateSubnet(&ec2.CreateSubnetInput{
		CidrBlock: sdk.String(SubnetCidrBlock),
		VpcId:     vpc.VpcId,
	})
	if err != nil {
		return nil, err
	}

	return sOut.Subnet, nil
}

func createInternetGateway(svc *ec2.EC2, vpc *ec2.Vpc) (*ec2.InternetGateway, error) {
	igOut, err := svc.CreateInternetGateway(&ec2.CreateInternetGatewayInput{})
	if err != nil {
		return nil, err
	}

	_, err = svc.AttachInternetGateway(&ec2.AttachInternetGatewayInput{
		InternetGatewayId: igOut.InternetGateway.InternetGatewayId,
		VpcId:             vpc.VpcId,
	})
	if err != nil {
		return nil, err
	}

	return igOut.InternetGateway, nil
}

func addRoute(svc *ec2.EC2, vpc *ec2.Vpc, gateway *ec2.InternetGateway) (*ec2.RouteTable, error) {
	rtOut, err := svc.DescribeRouteTables(&ec2.DescribeRouteTablesInput{
		Filters: []*ec2.Filter{
			{Name: sdk.String("vpc-id"), Values: []*string{vpc.VpcId}},
		},
	})
	if err != nil {
		return nil, err
	}

	if len(rtOut.RouteTables) != 1 {
		return nil, errors.New("Could not find main RouteTable")
	}

	_, err = svc.CreateRoute(&ec2.CreateRouteInput{
		GatewayId:            gateway.InternetGatewayId,
		DestinationCidrBlock: sdk.String("0.0.0.0/0"),
		RouteTableId:         rtOut.RouteTables[0].RouteTableId,
	})
	if err != nil {
		return nil, err
	}

	return rtOut.RouteTables[0], nil
}

func createTags(svc *ec2.EC2, vpc *ec2.Vpc, gateway *ec2.InternetGateway, subnet *ec2.Subnet, routeTable *ec2.RouteTable, cluster *api.Cluster) error {
	_, err := svc.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{vpc.VpcId, subnet.SubnetId, gateway.InternetGatewayId, routeTable.RouteTableId},
		Tags: []*ec2.Tag{
			{
				Key:   sdk.String(defaultKubermaticClusterIDTagKey),
				Value: sdk.String(cluster.Metadata.UID),
			},
			{
				Key:   sdk.String(defaultKubermaticClusterNameTagKey),
				Value: sdk.String(cluster.Metadata.Name),
			},
			{
				Key:   sdk.String(awsFilterName),
				Value: sdk.String(fmt.Sprintf("kubermatic-%s", cluster.Metadata.Name)),
			},
		},
	})

	return err
}

func (a *aws) InitializeCloudSpec(cluster *api.Cluster) error {
	if cluster.Spec.Cloud.AWS.VPCId != "" {
		return nil
	}

	svc, err := a.getSession(cluster)
	if err != nil {
		return err
	}

	vpc, err := createVpc(svc)
	if err != nil {
		return err
	}
	cluster.Spec.Cloud.AWS.VPCId = *vpc.VpcId

	subnet, err := createSubnet(svc, vpc)
	if err != nil {
		return err
	}
	cluster.Spec.Cloud.AWS.SubnetID = *subnet.SubnetId

	gateway, err := createInternetGateway(svc, vpc)
	if err != nil {
		return err
	}
	cluster.Spec.Cloud.AWS.InternetGatewayID = *gateway.InternetGatewayId

	routeTable, err := addRoute(svc, vpc, gateway)
	if err != nil {
		return err
	}
	cluster.Spec.Cloud.AWS.RouteTableID = *routeTable.RouteTableId

	err = createTags(svc, vpc, gateway, subnet, routeTable, cluster)
	if err != nil {
		return err
	}

	return nil
}

func (*aws) MarshalCloudSpec(cs *api.CloudSpec) (map[string]string, error) {
	return map[string]string{
		accessKeyIDAnnotationKey:     cs.AWS.AccessKeyID,
		secretAccessKeyAnnotationKey: cs.AWS.SecretAccessKey,
		sshKeyNameKey:                cs.AWS.SSHKeyName,
		subnetIDKey:                  cs.AWS.SubnetID,
		vpcIDKey:                     cs.AWS.VPCId,
		internetGatewayIDKey:         cs.AWS.InternetGatewayID,
		routeTableIDKey:              cs.AWS.RouteTableID,
	}, nil
}

func (*aws) UnmarshalCloudSpec(annotations map[string]string) (*api.CloudSpec, error) {
	spec := &api.CloudSpec{
		AWS: &api.AWSCloudSpec{},
	}
	var ok bool
	if spec.AWS.AccessKeyID, ok = annotations[accessKeyIDAnnotationKey]; !ok {
		return nil, errors.New("no access key ID found")
	}

	if spec.AWS.SecretAccessKey, ok = annotations[secretAccessKeyAnnotationKey]; !ok {
		return nil, errors.New("no secret key found")
	}

	if spec.AWS.SubnetID, ok = annotations[subnetIDKey]; !ok {
		return nil, errors.New("no subnet ID found")
	}

	if spec.AWS.VPCId, ok = annotations[vpcIDKey]; !ok {
		return nil, errors.New("no vpc ID found")
	}

	if spec.AWS.InternetGatewayID, ok = annotations[internetGatewayIDKey]; !ok {
		return nil, errors.New("no internet gateway ID found")
	}

	if spec.AWS.RouteTableID, ok = annotations[routeTableIDKey]; !ok {
		return nil, errors.New("no route table ID found")
	}

	if spec.AWS.SSHKeyName, ok = annotations[sshKeyNameKey]; !ok {
		return nil, errors.New("no route table ID found")
	}

	return spec, nil
}

func (a *aws) userData(
	buf *bytes.Buffer,
	instanceName string,
	node *api.NodeSpec,
	clusterState *api.Cluster,
	dc provider.DatacenterMeta,
	key *api.KeyCert,
) error {
	data := ktemplate.Data{
		DC:                node.DatacenterName,
		ClusterName:       clusterState.Metadata.Name,
		SSHAuthorizedKeys: []string{},
		EtcdURL:           clusterState.Address.EtcdURL,
		APIServerURL:      clusterState.Address.URL,
		Region:            dc.Spec.AWS.Region,
		Name:              instanceName,
		ClientKey:         key.Key.Base64(),
		ClientCert:        key.Cert.Base64(),
		RootCACert:        clusterState.Status.RootCA.Cert.Base64(),
		ApiserverPubSSH:   clusterState.Status.ApiserverSSH,
		ApiserverToken:    clusterState.Address.Token,
		FlannelCIDR:       clusterState.Spec.Cloud.Network.Flannel.CIDR,
	}

	tpl, err := template.
		New("aws-cloud-config-node.yaml").
		Funcs(ktemplate.FuncMap).
		ParseFiles(tplPath)
	if err != nil {
		return err
	}
	return tpl.Execute(buf, data)
}

func (a *aws) CreateNodes(ctx context.Context, cluster *api.Cluster, node *api.NodeSpec, num int) ([]*api.Node, error) {
	dc, ok := a.datacenters[node.DatacenterName]
	if !ok || dc.Spec.AWS == nil {
		return nil, fmt.Errorf("invalid datacenter %q", node.DatacenterName)
	}
	if node.AWS.Type == "" {
		return nil, errors.New("no AWS node type specified")
	}
	svc, err := a.getSession(cluster)
	if err != nil {
		return nil, err
	}
	var createdNodes []*api.Node
	var buf bytes.Buffer
	for i := 0; i < num; i++ {
		buf.Reset()
		id := provider.ShortUID(5)
		instanceName := fmt.Sprintf("kubermatic-%s-%s", cluster.Metadata.Name, id)

		clientKC, err := cluster.CreateKeyCert(instanceName, []string{})
		if err != nil {
			return createdNodes, err
		}

		if err = a.userData(&buf, instanceName, node, cluster, dc, clientKC); err != nil {
			return createdNodes, err
		}
		netSpec := []*ec2.InstanceNetworkInterfaceSpecification{
			{
				DeviceIndex:              sdk.Int64(0), // eth0
				AssociatePublicIpAddress: sdk.Bool(true),
				DeleteOnTermination:      sdk.Bool(true),
				SubnetId:                 sdk.String(cluster.Spec.Cloud.AWS.SubnetID),
			},
		}

		instanceRequest := &ec2.RunInstancesInput{
			ImageId:           sdk.String(dc.Spec.AWS.AMI),
			MaxCount:          sdk.Int64(1),
			MinCount:          sdk.Int64(1),
			InstanceType:      sdk.String(node.AWS.Type),
			UserData:          sdk.String(base64.StdEncoding.EncodeToString(buf.Bytes())),
			KeyName:           sdk.String(cluster.Spec.Cloud.AWS.SSHKeyName),
			NetworkInterfaces: netSpec,
		}

		newNode, err := launch(svc, instanceName, instanceRequest, cluster)

		if err != nil {
			return createdNodes, err
		}
		createdNodes = append(createdNodes, newNode)
	}
	return createdNodes, nil
}

func (a *aws) Nodes(ctx context.Context, cluster *api.Cluster) ([]*api.Node, error) {
	svc, err := a.getSession(cluster)
	if err != nil {
		return nil, err
	}

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{{
			// TODO: Direct Tag filtering
			Name: sdk.String(awsFilterState),
			Values: []*string{
				sdk.String(awsFilterRunning),
				sdk.String(awsFilterPending),
			},
		}},
	}

	resp, err := svc.DescribeInstances(params)
	if err != nil {
		return nil, err
	}

	nodes := make([]*api.Node, 0, len(resp.Reservations))
	for _, n := range resp.Reservations {
		for _, instance := range n.Instances {
			var isOwner bool
			var name string
			for _, tag := range instance.Tags {
				if *tag.Key == defaultKubermaticClusterIDTagKey && *tag.Value == cluster.Metadata.UID {
					isOwner = true
				}
				if *tag.Key == awsFilterName {
					name = *tag.Value
				}
			}
			if isOwner {
				nodes = append(nodes, createNode(name, instance))
			}
		}
	}
	return nodes, nil
}

func (a *aws) DeleteNodes(ctx context.Context, cluster *api.Cluster, UIDs []string) error {
	svc, err := a.getSession(cluster)
	if err != nil {
		return err
	}

	awsInstanceIds := make([]*string, len(UIDs))
	for i := 0; i < len(UIDs); i++ {
		awsInstanceIds[i] = sdk.String(UIDs[i])
	}

	terminateRequest := &ec2.TerminateInstancesInput{
		InstanceIds: awsInstanceIds,
	}

	_, err = svc.TerminateInstances(terminateRequest)
	return err
}

func (a *aws) getSession(cluster *api.Cluster) (*ec2.EC2, error) {
	awsSpec := cluster.Spec.Cloud.GetAWS()
	config := sdk.NewConfig()
	dc, found := a.datacenters[cluster.Spec.Cloud.DatacenterName]
	if !found || dc.Spec.AWS == nil {
		return nil, fmt.Errorf("can't find datacenter %s", cluster.Spec.Cloud.DatacenterName)
	}
	config = config.WithRegion(dc.Spec.AWS.Region)
	config = config.WithCredentials(credentials.NewStaticCredentials(awsSpec.AccessKeyID, awsSpec.SecretAccessKey, ""))
	// TODO: specify retrycount
	config = config.WithMaxRetries(3)
	return ec2.New(session.New(config)), nil
}

func createNode(name string, instance *ec2.Instance) *api.Node {

	privateIP := ""
	publicIP := ""
	if instance.PrivateIpAddress != nil {
		privateIP = *instance.PrivateIpAddress
	}

	if instance.PublicIpAddress != nil {
		publicIP = *instance.PublicIpAddress
	}

	return &api.Node{
		Metadata: api.Metadata{
			UID:  *instance.InstanceId,
			Name: name,
		},
		Status: api.NodeStatus{
			Addresses: map[string]string{
				"public":  publicIP,
				"private": privateIP,
			},
		},
		Spec: api.NodeSpec{
			DatacenterName: *instance.Placement.AvailabilityZone,
			AWS: &api.AWSNodeSpec{
				Type: *instance.InstanceType,
			},
		},
	}
}

func launch(client *ec2.EC2, name string, instance *ec2.RunInstancesInput, cluster *api.Cluster) (*api.Node, error) {
	serverReq, err := client.RunInstances(instance)
	if err != nil {
		return nil, err
	}

	_, err = client.CreateTags(&ec2.CreateTagsInput{
		Resources: []*string{serverReq.Instances[0].InstanceId},
		Tags: []*ec2.Tag{
			{
				Key:   sdk.String(defaultKubermaticClusterIDTagKey),
				Value: sdk.String(cluster.Metadata.UID),
			},
			{
				Key:   sdk.String(defaultKubermaticClusterNameTagKey),
				Value: sdk.String(cluster.Metadata.Name),
			},
			{
				Key:   sdk.String(awsFilterName),
				Value: sdk.String(name),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// Allow unchecked source/destination addresses for flannel
	_, err = client.ModifyInstanceAttribute(&ec2.ModifyInstanceAttributeInput{
		SourceDestCheck: &ec2.AttributeBooleanValue{
			Value: sdk.Bool(true),
		},
		InstanceId: serverReq.Instances[0].InstanceId,
	})
	if err != nil {
		return nil, err
	}

	return createNode(name, serverReq.Instances[0]), nil
}

func (a *aws) doCleanUpAWS(c *api.Cluster) error {
	svc, err := a.getSession(c)
	if err != nil {
		return err
	}

	// alive tests for living instaces
	alive := func() (bool, error) {
		ress, erre := svc.DescribeInstances(&ec2.DescribeInstancesInput{
			Filters: []*ec2.Filter{{
				Name:   sdk.String("tag-value"),
				Values: []*string{sdk.String(c.Metadata.UID)},
			}},
		})
		if erre != nil {
			return true, err
		}

		// Look for living instaces
		for _, res := range ress.Reservations {
			if len(res.Instances) > 0 {
				return true, nil
			}
		}
		return false, nil
	}

	// Wait for nodes to terminatle
	for {
		instancesAlive, erre := alive()
		if erre != nil {
			return erre
		}
		if !instancesAlive {
			break
		}
		time.Sleep(time.Second * 45)
	}

	if c.Spec.Cloud.AWS.RouteTableID != "" {
		_, err = svc.DeleteRouteTable(&ec2.DeleteRouteTableInput{
			RouteTableId: sdk.String(c.Spec.Cloud.AWS.RouteTableID),
		})
		if err != nil {
			glog.V(2).Infof("Failed to delete RouteTable %s during aws-cleanup for cluster %s : %v", c.Spec.Cloud.AWS.RouteTableID, c.Metadata.Name, err)
		}
	}

	if c.Spec.Cloud.AWS.InternetGatewayID != "" && c.Spec.Cloud.AWS.VPCId != "" {
		_, err = svc.DetachInternetGateway(&ec2.DetachInternetGatewayInput{
			InternetGatewayId: sdk.String(c.Spec.Cloud.AWS.InternetGatewayID),
			VpcId:             sdk.String(c.Spec.Cloud.AWS.VPCId),
		})
		if err != nil {
			glog.V(2).Infof("Failed to detach InternetGateway %s from VPC %s during aws-cleanup for cluster %s : %v", c.Spec.Cloud.AWS.InternetGatewayID, c.Spec.Cloud.AWS.VPCId, c.Metadata.Name, err)
		}
	}

	if c.Spec.Cloud.AWS.SubnetID != "" {
		_, err = svc.DeleteSubnet(&ec2.DeleteSubnetInput{
			SubnetId: sdk.String(c.Spec.Cloud.AWS.SubnetID),
		})
		if err != nil {
			glog.V(2).Infof("Failed to delete Subnet %s during aws-cleanup for cluster %s : %v", c.Spec.Cloud.AWS.SubnetID, c.Metadata.Name, err)
		}
	}

	if c.Spec.Cloud.AWS.InternetGatewayID != "" {
		_, err = svc.DeleteInternetGateway(&ec2.DeleteInternetGatewayInput{
			InternetGatewayId: sdk.String(c.Spec.Cloud.AWS.InternetGatewayID),
		})
		if err != nil {
			glog.V(2).Infof("Failed to delete InternetGateway %s during aws-cleanup for cluster %s : %v", c.Spec.Cloud.AWS.InternetGatewayID, c.Metadata.Name, err)
		}
	}

	if c.Spec.Cloud.AWS.VPCId != "" {
		_, err = svc.DeleteVpc(&ec2.DeleteVpcInput{
			VpcId: sdk.String(c.Spec.Cloud.AWS.VPCId),
		})
		if err != nil {
			glog.V(2).Infof("Failed to delete VPC %s during aws-cleanup for cluster %s : %v", c.Spec.Cloud.AWS.VPCId, c.Metadata.Name, err)
		}
	}
	return nil
}

func (a *aws) CleanUp(c *api.Cluster) error {
	go a.doCleanUpAWS(c)
	return nil
}
