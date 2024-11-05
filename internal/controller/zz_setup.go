// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	cluster "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/cluster"
	node "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/node"
	autolaunchgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/autolaunchgroup"
	eipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/eipassociate"
	instance "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/instance"
	interfaceattach "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/interfaceattach"
	servergroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/servergroup"
	volumeattach "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/volumeattach"
	globaleip "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/globaleip"
	globaleipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/globaleipassociate"
	globalinternetbandwidth "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/globalinternetbandwidth"
	vpcbandwidth "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/vpcbandwidth"
	vpcbandwidthassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/vpcbandwidthassociate"
	vpceip "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/vpceip"
	vpceipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/vpceipassociate"
	vpceipv3associate "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/vpceipv3associate"
	vpcinternetgateway "github.com/huaweicloud/provider-huaweicloud/internal/controller/eip/vpcinternetgateway"
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
	addressgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/addressgroup"
	flowlog "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/flowlog"
	networkacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/networkacl"
	networkinterface "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/networkinterface"
	peeringconnection "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/peeringconnection"
	peeringconnectionaccepter "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/peeringconnectionaccepter"
	route "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/route"
	routetable "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/routetable"
	secgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/secgroup"
	secgrouprule "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/secgrouprule"
	subnet "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/subnet"
	subnetworkinterface "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/subnetworkinterface"
	trafficmirrorfilter "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/trafficmirrorfilter"
	trafficmirrorfilterrule "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/trafficmirrorfilterrule"
	trafficmirrorsession "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/trafficmirrorsession"
	vip "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/vip"
	vipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/vipassociate"
	vpc "github.com/huaweicloud/provider-huaweicloud/internal/controller/vpc/vpc"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		cluster.Setup,
		node.Setup,
		autolaunchgroup.Setup,
		eipassociate.Setup,
		instance.Setup,
		interfaceattach.Setup,
		servergroup.Setup,
		volumeattach.Setup,
		globaleip.Setup,
		globaleipassociate.Setup,
		globalinternetbandwidth.Setup,
		vpcbandwidth.Setup,
		vpcbandwidthassociate.Setup,
		vpceip.Setup,
		vpceipassociate.Setup,
		vpceipv3associate.Setup,
		vpcinternetgateway.Setup,
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
		addressgroup.Setup,
		flowlog.Setup,
		networkacl.Setup,
		networkinterface.Setup,
		peeringconnection.Setup,
		peeringconnectionaccepter.Setup,
		route.Setup,
		routetable.Setup,
		secgroup.Setup,
		secgrouprule.Setup,
		subnet.Setup,
		subnetworkinterface.Setup,
		trafficmirrorfilter.Setup,
		trafficmirrorfilterrule.Setup,
		trafficmirrorsession.Setup,
		vip.Setup,
		vipassociate.Setup,
		vpc.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
