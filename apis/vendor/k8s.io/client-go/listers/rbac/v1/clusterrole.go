/*
Copyright The Kubernetes Authors.

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
	v1 "k8s.io/api/rbac/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/listers"
	"k8s.io/client-go/tools/cache"
)

// ClusterRoleLister helps list ClusterRoles.
// All objects returned here must be treated as read-only.
type ClusterRoleLister interface {
	// List lists all ClusterRoles in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterRole, err error)
	// Get retrieves the ClusterRole from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterRole, error)
	ClusterRoleListerExpansion
}

// clusterRoleLister implements the ClusterRoleLister interface.
type clusterRoleLister struct {
	listers.ResourceIndexer[*v1.ClusterRole]
}

// NewClusterRoleLister returns a new ClusterRoleLister.
func NewClusterRoleLister(indexer cache.Indexer) ClusterRoleLister {
	return &clusterRoleLister{listers.New[*v1.ClusterRole](indexer, v1.Resource("clusterrole"))}
}
