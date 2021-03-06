package fake

import (
	v1 "github.com/openshift/client-go/authorization/clientset/versioned/typed/authorization/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeAuthorizationV1 struct {
	*testing.Fake
}

func (c *FakeAuthorizationV1) ClusterPolicies() v1.ClusterPolicyInterface {
	return &FakeClusterPolicies{c}
}

func (c *FakeAuthorizationV1) ClusterPolicyBindings() v1.ClusterPolicyBindingInterface {
	return &FakeClusterPolicyBindings{c}
}

func (c *FakeAuthorizationV1) ClusterRoles() v1.ClusterRoleInterface {
	return &FakeClusterRoles{c}
}

func (c *FakeAuthorizationV1) ClusterRoleBindings() v1.ClusterRoleBindingInterface {
	return &FakeClusterRoleBindings{c}
}

func (c *FakeAuthorizationV1) LocalResourceAccessReviews(namespace string) v1.LocalResourceAccessReviewInterface {
	return &FakeLocalResourceAccessReviews{c, namespace}
}

func (c *FakeAuthorizationV1) LocalSubjectAccessReviews(namespace string) v1.LocalSubjectAccessReviewInterface {
	return &FakeLocalSubjectAccessReviews{c, namespace}
}

func (c *FakeAuthorizationV1) Policies(namespace string) v1.PolicyInterface {
	return &FakePolicies{c, namespace}
}

func (c *FakeAuthorizationV1) PolicyBindings(namespace string) v1.PolicyBindingInterface {
	return &FakePolicyBindings{c, namespace}
}

func (c *FakeAuthorizationV1) ResourceAccessReviews() v1.ResourceAccessReviewInterface {
	return &FakeResourceAccessReviews{c}
}

func (c *FakeAuthorizationV1) Roles(namespace string) v1.RoleInterface {
	return &FakeRoles{c, namespace}
}

func (c *FakeAuthorizationV1) RoleBindings(namespace string) v1.RoleBindingInterface {
	return &FakeRoleBindings{c, namespace}
}

func (c *FakeAuthorizationV1) RoleBindingRestrictions(namespace string) v1.RoleBindingRestrictionInterface {
	return &FakeRoleBindingRestrictions{c, namespace}
}

func (c *FakeAuthorizationV1) SelfSubjectRulesReviews(namespace string) v1.SelfSubjectRulesReviewInterface {
	return &FakeSelfSubjectRulesReviews{c, namespace}
}

func (c *FakeAuthorizationV1) SubjectAccessReviews() v1.SubjectAccessReviewInterface {
	return &FakeSubjectAccessReviews{c}
}

func (c *FakeAuthorizationV1) SubjectRulesReviews(namespace string) v1.SubjectRulesReviewInterface {
	return &FakeSubjectRulesReviews{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeAuthorizationV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
