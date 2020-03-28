package foo

import (
	cloudprovider "k8s.io/cloud-provider"
)

type foo struct{}

var _ cloudprovider.Interface = foo{}
