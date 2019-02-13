// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AuthenticationLister helps list Authentications.
type AuthenticationLister interface {
	// List lists all Authentications in the indexer.
	List(selector labels.Selector) (ret []*v1.Authentication, err error)
	// Get retrieves the Authentication from the index for a given name.
	Get(name string) (*v1.Authentication, error)
	AuthenticationListerExpansion
}

// authenticationLister implements the AuthenticationLister interface.
type authenticationLister struct {
	indexer cache.Indexer
}

// NewAuthenticationLister returns a new AuthenticationLister.
func NewAuthenticationLister(indexer cache.Indexer) AuthenticationLister {
	return &authenticationLister{indexer: indexer}
}

// List lists all Authentications in the indexer.
func (s *authenticationLister) List(selector labels.Selector) (ret []*v1.Authentication, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Authentication))
	})
	return ret, err
}

// Get retrieves the Authentication from the index for a given name.
func (s *authenticationLister) Get(name string) (*v1.Authentication, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("authentication"), name)
	}
	return obj.(*v1.Authentication), nil
}
