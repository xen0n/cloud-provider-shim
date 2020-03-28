package cloudprovidershim

//go:generate protoc -I proto -I third_party --gogo_out=plugins=grpc:. ./proto/v1alpha1/lb.proto
