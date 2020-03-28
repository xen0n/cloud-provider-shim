package shim

import (
	"context"

	"github.com/sirupsen/logrus"
	v1 "k8s.io/api/core/v1"
	cloudprovider "k8s.io/cloud-provider"
)

type ShimLoadBalancer struct {
	logger logrus.Logger
}

var _ cloudprovider.LoadBalancer = (*ShimLoadBalancer)(nil)

// TODO: Break this up into different interfaces (LB, etc) when we have more than one type of service
// GetLoadBalancer returns whether the specified load balancer exists, and
// if so, what its status is.
// Implementations must treat the *v1.Service parameter as read-only and not modify it.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (b *ShimLoadBalancer) GetLoadBalancer(ctx context.Context, clusterName string, service *v1.Service) (status *v1.LoadBalancerStatus, exists bool, err error) {
	b.logger.WithFields(logrus.Fields{
		"clusterName": clusterName,
		"service":     service,
	}).Debug("GetLoadBalancer called")

	panic("not implemented") // TODO: Implement
}

// GetLoadBalancerName returns the name of the load balancer. Implementations must treat the
// *v1.Service parameter as read-only and not modify it.
func (b *ShimLoadBalancer) GetLoadBalancerName(ctx context.Context, clusterName string, service *v1.Service) string {
	b.logger.WithFields(logrus.Fields{
		"clusterName": clusterName,
		"service":     service,
	}).Debug("GetLoadBalancerName called")

	panic("not implemented") // TODO: Implement
}

// EnsureLoadBalancer creates a new load balancer 'name', or updates the existing one. Returns the status of the balancer
// Implementations must treat the *v1.Service and *v1.Node
// parameters as read-only and not modify them.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (b *ShimLoadBalancer) EnsureLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) (*v1.LoadBalancerStatus, error) {
	b.logger.WithFields(logrus.Fields{
		"clusterName": clusterName,
		"service":     service,
		"nodes":       nodes,
	}).Debug("EnsureLoadBalancer called")

	panic("not implemented") // TODO: Implement
}

// UpdateLoadBalancer updates hosts under the specified load balancer.
// Implementations must treat the *v1.Service and *v1.Node
// parameters as read-only and not modify them.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (b *ShimLoadBalancer) UpdateLoadBalancer(ctx context.Context, clusterName string, service *v1.Service, nodes []*v1.Node) error {
	b.logger.WithFields(logrus.Fields{
		"clusterName": clusterName,
		"service":     service,
		"nodes":       nodes,
	}).Debug("UpdateLoadBalancer called")

	panic("not implemented") // TODO: Implement
}

// EnsureLoadBalancerDeleted deletes the specified load balancer if it
// exists, returning nil if the load balancer specified either didn't exist or
// was successfully deleted.
// This construction is useful because many cloud providers' load balancers
// have multiple underlying components, meaning a Get could say that the LB
// doesn't exist even if some part of it is still laying around.
// Implementations must treat the *v1.Service parameter as read-only and not modify it.
// Parameter 'clusterName' is the name of the cluster as presented to kube-controller-manager
func (b *ShimLoadBalancer) EnsureLoadBalancerDeleted(ctx context.Context, clusterName string, service *v1.Service) error {
	b.logger.WithFields(logrus.Fields{
		"clusterName": clusterName,
		"service":     service,
	}).Debug("EnsureLoadBalancerDeleted called")

	panic("not implemented") // TODO: Implement
}
