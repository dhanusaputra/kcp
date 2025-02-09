//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by kcp code-generator. DO NOT EDIT.

package v1alpha1

import (
	"context"

	kcpclient "github.com/kcp-dev/apimachinery/v2/pkg/client"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/watch"

	apiresourcev1alpha1 "github.com/kcp-dev/kcp/sdk/apis/apiresource/v1alpha1"
	apiresourcev1alpha1client "github.com/kcp-dev/kcp/sdk/client/clientset/versioned/typed/apiresource/v1alpha1"
)

// NegotiatedAPIResourcesClusterGetter has a method to return a NegotiatedAPIResourceClusterInterface.
// A group's cluster client should implement this interface.
type NegotiatedAPIResourcesClusterGetter interface {
	NegotiatedAPIResources() NegotiatedAPIResourceClusterInterface
}

// NegotiatedAPIResourceClusterInterface can operate on NegotiatedAPIResources across all clusters,
// or scope down to one cluster and return a apiresourcev1alpha1client.NegotiatedAPIResourceInterface.
type NegotiatedAPIResourceClusterInterface interface {
	Cluster(logicalcluster.Path) apiresourcev1alpha1client.NegotiatedAPIResourceInterface
	List(ctx context.Context, opts metav1.ListOptions) (*apiresourcev1alpha1.NegotiatedAPIResourceList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
}

type negotiatedAPIResourcesClusterInterface struct {
	clientCache kcpclient.Cache[*apiresourcev1alpha1client.ApiresourceV1alpha1Client]
}

// Cluster scopes the client down to a particular cluster.
func (c *negotiatedAPIResourcesClusterInterface) Cluster(clusterPath logicalcluster.Path) apiresourcev1alpha1client.NegotiatedAPIResourceInterface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return c.clientCache.ClusterOrDie(clusterPath).NegotiatedAPIResources()
}

// List returns the entire collection of all NegotiatedAPIResources across all clusters.
func (c *negotiatedAPIResourcesClusterInterface) List(ctx context.Context, opts metav1.ListOptions) (*apiresourcev1alpha1.NegotiatedAPIResourceList, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).NegotiatedAPIResources().List(ctx, opts)
}

// Watch begins to watch all NegotiatedAPIResources across all clusters.
func (c *negotiatedAPIResourcesClusterInterface) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.clientCache.ClusterOrDie(logicalcluster.Wildcard).NegotiatedAPIResources().Watch(ctx, opts)
}
