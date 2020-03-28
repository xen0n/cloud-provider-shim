package foo

import (
	cloudprovider "k8s.io/cloud-provider"
)

type foo struct{}

var _ cloudprovider.Interface = foo{}

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping or run custom controllers specific to the cloud provider.
// Any tasks started here should be cleaned up when the stop channel closes.
func (p foo) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	// TODO: currently no-op
}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (p foo) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	return nil, false
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (p foo) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (p foo) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (p foo) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (p foo) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (p foo) ProviderName() string {
	// TODO: return something that can be configured
	return "shim"
}

// HasClusterID returns true if a ClusterID is required and set
func (p foo) HasClusterID() bool {
	return true
}
