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
	v1 "github.com/cert-manager/cert-manager/pkg/apis/acme/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// OrderLister helps list Orders.
// All objects returned here must be treated as read-only.
type OrderLister interface {
	// List lists all Orders in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Order, err error)
	// Orders returns an object that can list and get Orders.
	Orders(namespace string) OrderNamespaceLister
	OrderListerExpansion
}

// orderLister implements the OrderLister interface.
type orderLister struct {
	listers.ResourceIndexer[*v1.Order]
}

// NewOrderLister returns a new OrderLister.
func NewOrderLister(indexer cache.Indexer) OrderLister {
	return &orderLister{listers.New[*v1.Order](indexer, v1.Resource("order"))}
}

// Orders returns an object that can list and get Orders.
func (s *orderLister) Orders(namespace string) OrderNamespaceLister {
	return orderNamespaceLister{listers.NewNamespaced[*v1.Order](s.ResourceIndexer, namespace)}
}

// OrderNamespaceLister helps list and get Orders.
// All objects returned here must be treated as read-only.
type OrderNamespaceLister interface {
	// List lists all Orders in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Order, err error)
	// Get retrieves the Order from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Order, error)
	OrderNamespaceListerExpansion
}

// orderNamespaceLister implements the OrderNamespaceLister
// interface.
type orderNamespaceLister struct {
	listers.ResourceIndexer[*v1.Order]
}