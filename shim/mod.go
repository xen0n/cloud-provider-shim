package shim

import (
	cloudprovider "k8s.io/cloud-provider"
)

type CloudProvider struct {
	Name string

	LB *ShimLoadBalancer
}

var _ cloudprovider.Interface = (*CloudProvider)(nil)

// Initialize provides the cloud with a kubernetes client builder and may spawn goroutines
// to perform housekeeping or run custom controllers specific to the cloud provider.
// Any tasks started here should be cleaned up when the stop channel closes.
func (p *CloudProvider) Initialize(clientBuilder cloudprovider.ControllerClientBuilder, stop <-chan struct{}) {
	// TODO: currently no-op
}

// LoadBalancer returns a balancer interface. Also returns true if the interface is supported, false otherwise.
func (p *CloudProvider) LoadBalancer() (cloudprovider.LoadBalancer, bool) {
	if p.LB == nil {
		return nil, false
	}
	return p.LB, true
}

// Instances returns an instances interface. Also returns true if the interface is supported, false otherwise.
func (p *CloudProvider) Instances() (cloudprovider.Instances, bool) {
	return nil, false
}

// Zones returns a zones interface. Also returns true if the interface is supported, false otherwise.
func (p *CloudProvider) Zones() (cloudprovider.Zones, bool) {
	return nil, false
}

// Clusters returns a clusters interface.  Also returns true if the interface is supported, false otherwise.
func (p *CloudProvider) Clusters() (cloudprovider.Clusters, bool) {
	return nil, false
}

// Routes returns a routes interface along with whether the interface is supported.
func (p *CloudProvider) Routes() (cloudprovider.Routes, bool) {
	return nil, false
}

// ProviderName returns the cloud provider ID.
func (p *CloudProvider) ProviderName() string {
	return p.Name
}

// HasClusterID returns true if a ClusterID is required and set
func (p *CloudProvider) HasClusterID() bool {
	return true
}
