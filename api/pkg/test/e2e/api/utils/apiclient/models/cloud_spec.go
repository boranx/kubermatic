// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudSpec CloudSpec mutually stores access data to a cloud provider.
// swagger:model CloudSpec
type CloudSpec struct {

	// DatacenterName where the users 'cloud' lives in.
	DatacenterName string `json:"dc,omitempty"`

	// alibaba
	Alibaba *AlibabaCloudSpec `json:"alibaba,omitempty"`

	// aws
	Aws *AWSCloudSpec `json:"aws,omitempty"`

	// azure
	Azure *AzureCloudSpec `json:"azure,omitempty"`

	// bringyourown
	Bringyourown BringYourOwnCloudSpec `json:"bringyourown,omitempty"`

	// digitalocean
	Digitalocean *DigitaloceanCloudSpec `json:"digitalocean,omitempty"`

	// fake
	Fake *FakeCloudSpec `json:"fake,omitempty"`

	// gcp
	Gcp *GCPCloudSpec `json:"gcp,omitempty"`

	// hetzner
	Hetzner *HetznerCloudSpec `json:"hetzner,omitempty"`

	// kubevirt
	Kubevirt *KubevirtCloudSpec `json:"kubevirt,omitempty"`

	// openstack
	Openstack *OpenstackCloudSpec `json:"openstack,omitempty"`

	// packet
	Packet *PacketCloudSpec `json:"packet,omitempty"`

	// vsphere
	Vsphere *VSphereCloudSpec `json:"vsphere,omitempty"`
}

// Validate validates this cloud spec
func (m *CloudSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAlibaba(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDigitalocean(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFake(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGcp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHetzner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubevirt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOpenstack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePacket(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVsphere(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CloudSpec) validateAlibaba(formats strfmt.Registry) error {

	if swag.IsZero(m.Alibaba) { // not required
		return nil
	}

	if m.Alibaba != nil {
		if err := m.Alibaba.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alibaba")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateDigitalocean(formats strfmt.Registry) error {

	if swag.IsZero(m.Digitalocean) { // not required
		return nil
	}

	if m.Digitalocean != nil {
		if err := m.Digitalocean.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("digitalocean")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateFake(formats strfmt.Registry) error {

	if swag.IsZero(m.Fake) { // not required
		return nil
	}

	if m.Fake != nil {
		if err := m.Fake.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fake")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateGcp(formats strfmt.Registry) error {

	if swag.IsZero(m.Gcp) { // not required
		return nil
	}

	if m.Gcp != nil {
		if err := m.Gcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("gcp")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateHetzner(formats strfmt.Registry) error {

	if swag.IsZero(m.Hetzner) { // not required
		return nil
	}

	if m.Hetzner != nil {
		if err := m.Hetzner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("hetzner")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateKubevirt(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubevirt) { // not required
		return nil
	}

	if m.Kubevirt != nil {
		if err := m.Kubevirt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("kubevirt")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateOpenstack(formats strfmt.Registry) error {

	if swag.IsZero(m.Openstack) { // not required
		return nil
	}

	if m.Openstack != nil {
		if err := m.Openstack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("openstack")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validatePacket(formats strfmt.Registry) error {

	if swag.IsZero(m.Packet) { // not required
		return nil
	}

	if m.Packet != nil {
		if err := m.Packet.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("packet")
			}
			return err
		}
	}

	return nil
}

func (m *CloudSpec) validateVsphere(formats strfmt.Registry) error {

	if swag.IsZero(m.Vsphere) { // not required
		return nil
	}

	if m.Vsphere != nil {
		if err := m.Vsphere.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("vsphere")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CloudSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudSpec) UnmarshalBinary(b []byte) error {
	var res CloudSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
