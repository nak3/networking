/*
Copyright 2020 The Knative Authors

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

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1alpha1 "sigs.k8s.io/service-apis/apis/v1alpha1"
)

// TLSRouteLister helps list TLSRoutes.
type TLSRouteLister interface {
	// List lists all TLSRoutes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.TLSRoute, err error)
	// TLSRoutes returns an object that can list and get TLSRoutes.
	TLSRoutes(namespace string) TLSRouteNamespaceLister
	TLSRouteListerExpansion
}

// tLSRouteLister implements the TLSRouteLister interface.
type tLSRouteLister struct {
	indexer cache.Indexer
}

// NewTLSRouteLister returns a new TLSRouteLister.
func NewTLSRouteLister(indexer cache.Indexer) TLSRouteLister {
	return &tLSRouteLister{indexer: indexer}
}

// List lists all TLSRoutes in the indexer.
func (s *tLSRouteLister) List(selector labels.Selector) (ret []*v1alpha1.TLSRoute, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TLSRoute))
	})
	return ret, err
}

// TLSRoutes returns an object that can list and get TLSRoutes.
func (s *tLSRouteLister) TLSRoutes(namespace string) TLSRouteNamespaceLister {
	return tLSRouteNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TLSRouteNamespaceLister helps list and get TLSRoutes.
type TLSRouteNamespaceLister interface {
	// List lists all TLSRoutes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.TLSRoute, err error)
	// Get retrieves the TLSRoute from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.TLSRoute, error)
	TLSRouteNamespaceListerExpansion
}

// tLSRouteNamespaceLister implements the TLSRouteNamespaceLister
// interface.
type tLSRouteNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all TLSRoutes in the indexer for a given namespace.
func (s tLSRouteNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.TLSRoute, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.TLSRoute))
	})
	return ret, err
}

// Get retrieves the TLSRoute from the indexer for a given namespace and name.
func (s tLSRouteNamespaceLister) Get(name string) (*v1alpha1.TLSRoute, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("tlsroute"), name)
	}
	return obj.(*v1alpha1.TLSRoute), nil
}
