// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/cluster"
	node "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/node"
	bandwidth "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/bandwidth"
	bandwidthassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/bandwidthassociate"
	eip "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/eip"
	eipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/eipassociate"
	eipv3associate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/eipv3associate"
	internetbandwidth "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/internetbandwidth"
	internetgateway "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/internetgateway"
	accesskey "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/accesskey"
	acl "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/acl"
	agency "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/agency"
	group "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/group"
	groupmembership "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/groupmembership"
	grouproleassignment "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/grouproleassignment"
	loginpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/loginpolicy"
	passwordpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/passwordpolicy"
	project "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/project"
	protectionpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/protectionpolicy"
	provider "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/provider"
	providerconversion "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/providerconversion"
	role "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/role"
	roleassignment "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/roleassignment"
	user "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/user"
	userroleassignment "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/userroleassignment"
	usertoken "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/usertoken"
	virtualmfadevice "github.com/huaweicloud/provider-huaweicloud/internal/controller/iam/virtualmfadevice"
	bucket "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucket"
	bucketacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketacl"
	bucketobject "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobject"
	bucketobjectacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobjectacl"
	providerconfig "github.com/huaweicloud/provider-huaweicloud/internal/controller/providerconfig"
	secgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/secgroup"
	secgrouprule "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/secgrouprule"
	subnet "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/subnet"
	vpc "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		node.Setup,
		bandwidth.Setup,
		bandwidthassociate.Setup,
		eip.Setup,
		eip.Setup,
		eipassociate.Setup,
		eipassociate.Setup,
		eipv3associate.Setup,
		internetbandwidth.Setup,
		internetgateway.Setup,
		accesskey.Setup,
		acl.Setup,
		agency.Setup,
		group.Setup,
		groupmembership.Setup,
		grouproleassignment.Setup,
		loginpolicy.Setup,
		passwordpolicy.Setup,
		project.Setup,
		protectionpolicy.Setup,
		provider.Setup,
		providerconversion.Setup,
		role.Setup,
		roleassignment.Setup,
		user.Setup,
		userroleassignment.Setup,
		usertoken.Setup,
		virtualmfadevice.Setup,
		bucket.Setup,
		bucketacl.Setup,
		bucketobject.Setup,
		bucketobjectacl.Setup,
		providerconfig.Setup,
		secgroup.Setup,
		secgrouprule.Setup,
		subnet.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
