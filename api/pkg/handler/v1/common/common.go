package common

import (
	"github.com/prometheus/client_golang/prometheus"

	"github.com/kubermatic/kubermatic/api/pkg/crd/kubermatic/v1"
	"github.com/kubermatic/kubermatic/api/pkg/presets"
	"github.com/kubermatic/kubermatic/api/pkg/provider"
	"github.com/kubermatic/kubermatic/api/pkg/version"
)

// OIDCConfiguration is a struct that holds
// OIDC provider configuration data, read from command line arguments
type OIDCConfiguration struct {
	// URL holds OIDC Issuer URL address
	URL string
	// ClientID holds OIDC ClientID
	ClientID string
	// ClientSecret holds OIDC ClientSecret
	ClientSecret string
	// CookieHashKey is required, used to authenticate the cookie value using HMAC
	// It is recommended to use a key with 32 or 64 bytes.
	CookieHashKey string
	// CookieSecureMode if true then cookie received only with HTTPS otherwise with HTTP.
	CookieSecureMode bool
	// OfflineAccessAsScope if true then "offline_access" scope will be used
	// otherwise 'access_type=offline" query param will be passed
	OfflineAccessAsScope bool
}

// UpdateManager specifies a set of methods to handle cluster versions & updates
type UpdateManager interface {
	GetVersion(from, clusterType string) (*version.MasterVersion, error)
	GetMasterVersions(string) ([]*version.MasterVersion, error)
	GetDefault() (*version.MasterVersion, error)
	AutomaticUpdate(from, clusterType string) (*version.MasterVersion, error)
	GetPossibleUpdates(from, clusterType string) ([]*version.MasterVersion, error)
}

// PresetsManager specifies a set of methods to handle presets for specific provider
type PresetsManager interface {
	GetPresets() *presets.Presets
	SetCloudCredentials(credentialName string, cloud v1.CloudSpec, dc provider.DatacenterMeta) (*v1.CloudSpec, error)
}

// ServerMetrics defines metrics used by the API.
type ServerMetrics struct {
	HTTPRequestsTotal          *prometheus.CounterVec
	HTTPRequestsDuration       *prometheus.HistogramVec
	InitNodeDeploymentFailures *prometheus.CounterVec
}

// IsBringYourOwnProvider determines whether the spec holds BringYourOwn provider
func IsBringYourOwnProvider(spec v1.CloudSpec) (bool, error) {
	providerName, err := provider.ClusterCloudProviderName(spec)
	if err != nil {
		return false, err
	}
	return providerName == provider.BringYourOwnCloudProvider, nil
}
