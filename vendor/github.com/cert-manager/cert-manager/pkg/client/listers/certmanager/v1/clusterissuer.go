/*
Copyright The cert-manager Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/cert-manager/cert-manager/pkg/apis/certmanager/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterIssuerLister helps list ClusterIssuers.
// All objects returned here must be treated as read-only.
type ClusterIssuerLister interface {
	// List lists all ClusterIssuers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterIssuer, err error)
	// Get retrieves the ClusterIssuer from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterIssuer, error)
	ClusterIssuerListerExpansion
}

// clusterIssuerLister implements the ClusterIssuerLister interface.
type clusterIssuerLister struct {
	listers.ResourceIndexer[*v1.ClusterIssuer]
}

// NewClusterIssuerLister returns a new ClusterIssuerLister.
func NewClusterIssuerLister(indexer cache.Indexer) ClusterIssuerLister {
	return &clusterIssuerLister{listers.New[*v1.ClusterIssuer](indexer, v1.Resource("clusterissuer"))}
}