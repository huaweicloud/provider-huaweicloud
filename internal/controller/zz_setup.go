// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	authorization "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/authorization"
	bandwidthpackage "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/bandwidthpackage"
	centralnetwork "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/centralnetwork"
	centralnetworkattachment "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/centralnetworkattachment"
	centralnetworkconnectionbandwidthassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/centralnetworkconnectionbandwidthassociate"
	centralnetworkpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/centralnetworkpolicy"
	centralnetworkpolicyapply "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/centralnetworkpolicyapply"
	connection "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/connection"
	globalconnectionbandwidth "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/globalconnectionbandwidth"
	globalconnectionbandwidthassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/globalconnectionbandwidthassociate"
	interregionbandwidth "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/interregionbandwidth"
	networkinstance "github.com/huaweicloud/provider-huaweicloud/internal/controller/cc/networkinstance"
	addon "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/addon"
	chart "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/chart"
	cluster "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/cluster"
	clusterlogconfig "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/clusterlogconfig"
	clusterupgrade "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/clusterupgrade"
	namespace "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/namespace"
	node "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/node"
	nodeattach "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/nodeattach"
	nodepool "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/nodepool"
	nodepoolnodesadd "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/nodepoolnodesadd"
	pvc "github.com/huaweicloud/provider-huaweicloud/internal/controller/cce/pvc"
	alarmrule "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/alarmrule"
	alarmtemplate "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/alarmtemplate"
	dashboard "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/dashboard"
	dashboardwidget "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/dashboardwidget"
	eventreport "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/eventreport"
	oneclickalarm "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/oneclickalarm"
	resourcegroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/ces/resourcegroup"
	account "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/account"
	backup "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/backup"
	bigkeyanalysis "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/bigkeyanalysis"
	customtemplate "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/customtemplate"
	diagnosistask "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/diagnosistask"
	hotkeyanalysis "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/hotkeyanalysis"
	instance "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/instance"
	instancerestore "github.com/huaweicloud/provider-huaweicloud/internal/controller/dcs/instancerestore"
	rabbitmqexchange "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqexchange"
	rabbitmqexchangeassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqexchangeassociate"
	rabbitmqinstance "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqinstance"
	rabbitmqplugin "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqplugin"
	rabbitmqqueue "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqqueue"
	rabbitmqqueuemessageclear "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqqueuemessageclear"
	rabbitmquser "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmquser"
	rabbitmqvhost "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rabbitmqvhost"
	rocketmqconsumergroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rocketmqconsumergroup"
	rocketmqinstance "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rocketmqinstance"
	rocketmqmigrationtask "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rocketmqmigrationtask"
	rocketmqtopic "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rocketmqtopic"
	rocketmquser "github.com/huaweicloud/provider-huaweicloud/internal/controller/dms/rocketmquser"
	autolaunchgroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/autolaunchgroup"
	eipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/eipassociate"
	instanceecs "github.com/huaweicloud/provider-huaweicloud/internal/controller/ecs/instance"
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
	snapshot "github.com/huaweicloud/provider-huaweicloud/internal/controller/evs/snapshot"
	snapshotrollback "github.com/huaweicloud/provider-huaweicloud/internal/controller/evs/snapshotrollback"
	volume "github.com/huaweicloud/provider-huaweicloud/internal/controller/evs/volume"
	volumetransfer "github.com/huaweicloud/provider-huaweicloud/internal/controller/evs/volumetransfer"
	volumetransferaccepter "github.com/huaweicloud/provider-huaweicloud/internal/controller/evs/volumetransferaccepter"
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
	authorizationmodelarts "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/authorization"
	dataset "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/dataset"
	datasetversion "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/datasetversion"
	model "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/model"
	network "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/network"
	notebook "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/notebook"
	notebookmountstorage "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/notebookmountstorage"
	resourcepool "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/resourcepool"
	service "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/service"
	workspace "github.com/huaweicloud/provider-huaweicloud/internal/controller/modelarts/workspace"
	dnatrule "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/dnatrule"
	gateway "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/gateway"
	privatednatrule "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/privatednatrule"
	privategateway "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/privategateway"
	privatesnatrule "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/privatesnatrule"
	privatetransitip "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/privatetransitip"
	snatrule "github.com/huaweicloud/provider-huaweicloud/internal/controller/nat/snatrule"
	bucket "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucket"
	bucketacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketacl"
	bucketobject "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobject"
	bucketobjectacl "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketobjectacl"
	bucketpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketpolicy"
	bucketreplication "github.com/huaweicloud/provider-huaweicloud/internal/controller/obs/bucketreplication"
	providerconfig "github.com/huaweicloud/provider-huaweicloud/internal/controller/providerconfig"
	backuprds "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/backup"
	crossregionbackupstrategy "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/crossregionbackupstrategy"
	databaselogsshrinking "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/databaselogsshrinking"
	extendloglink "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/extendloglink"
	instancerds "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/instance"
	instanceeipassociate "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/instanceeipassociate"
	ltslog "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/ltslog"
	mysqlaccount "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/mysqlaccount"
	mysqlbinlog "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/mysqlbinlog"
	mysqldatabase "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/mysqldatabase"
	mysqldatabaseprivilege "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/mysqldatabaseprivilege"
	mysqldatabasetablerestore "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/mysqldatabasetablerestore"
	parametergroup "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/parametergroup"
	pgaccount "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgaccount"
	pgaccountprivileges "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgaccountprivileges"
	pgaccountroles "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgaccountroles"
	pgdatabase "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgdatabase"
	pgdatabaseprivilege "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgdatabaseprivilege"
	pghba "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pghba"
	pgplugin "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgplugin"
	pgpluginparameter "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgpluginparameter"
	pgpluginupdate "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgpluginupdate"
	pgsqllimit "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/pgsqllimit"
	primarystandbyswitch "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/primarystandbyswitch"
	readreplicainstance "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/readreplicainstance"
	restore "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/restore"
	sqlaudit "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/sqlaudit"
	sqlserveraccount "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/sqlserveraccount"
	sqlserverdatabase "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/sqlserverdatabase"
	sqlserverdatabaseprivilege "github.com/huaweicloud/provider-huaweicloud/internal/controller/rds/sqlserverdatabaseprivilege"
	logtank "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/logtank"
	messagedetection "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/messagedetection"
	messagepublish "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/messagepublish"
	messagetemplate "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/messagetemplate"
	subscription "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/subscription"
	subscriptionfilterpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/subscriptionfilterpolicy"
	topic "github.com/huaweicloud/provider-huaweicloud/internal/controller/smn/topic"
	imageautosync "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/imageautosync"
	imagepermissions "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/imagepermissions"
	imageretentionpolicy "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/imageretentionpolicy"
	imagetrigger "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/imagetrigger"
	organization "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/organization"
	organizationpermissions "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/organizationpermissions"
	repository "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/repository"
	repositorysharing "github.com/huaweicloud/provider-huaweicloud/internal/controller/swr/repositorysharing"
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
		authorization.Setup,
		bandwidthpackage.Setup,
		centralnetwork.Setup,
		centralnetworkattachment.Setup,
		centralnetworkconnectionbandwidthassociate.Setup,
		centralnetworkpolicy.Setup,
		centralnetworkpolicyapply.Setup,
		connection.Setup,
		globalconnectionbandwidth.Setup,
		globalconnectionbandwidthassociate.Setup,
		interregionbandwidth.Setup,
		networkinstance.Setup,
		addon.Setup,
		chart.Setup,
		cluster.Setup,
		clusterlogconfig.Setup,
		clusterupgrade.Setup,
		namespace.Setup,
		node.Setup,
		nodeattach.Setup,
		nodepool.Setup,
		nodepoolnodesadd.Setup,
		pvc.Setup,
		alarmrule.Setup,
		alarmtemplate.Setup,
		dashboard.Setup,
		dashboardwidget.Setup,
		eventreport.Setup,
		oneclickalarm.Setup,
		resourcegroup.Setup,
		account.Setup,
		backup.Setup,
		bigkeyanalysis.Setup,
		customtemplate.Setup,
		diagnosistask.Setup,
		hotkeyanalysis.Setup,
		instance.Setup,
		instancerestore.Setup,
		rabbitmqexchange.Setup,
		rabbitmqexchangeassociate.Setup,
		rabbitmqinstance.Setup,
		rabbitmqplugin.Setup,
		rabbitmqqueue.Setup,
		rabbitmqqueuemessageclear.Setup,
		rabbitmquser.Setup,
		rabbitmqvhost.Setup,
		rocketmqconsumergroup.Setup,
		rocketmqinstance.Setup,
		rocketmqmigrationtask.Setup,
		rocketmqtopic.Setup,
		rocketmquser.Setup,
		autolaunchgroup.Setup,
		eipassociate.Setup,
		instanceecs.Setup,
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
		snapshot.Setup,
		snapshotrollback.Setup,
		volume.Setup,
		volumetransfer.Setup,
		volumetransferaccepter.Setup,
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
		authorizationmodelarts.Setup,
		dataset.Setup,
		datasetversion.Setup,
		model.Setup,
		network.Setup,
		notebook.Setup,
		notebookmountstorage.Setup,
		resourcepool.Setup,
		service.Setup,
		workspace.Setup,
		dnatrule.Setup,
		gateway.Setup,
		privatednatrule.Setup,
		privategateway.Setup,
		privatesnatrule.Setup,
		privatetransitip.Setup,
		snatrule.Setup,
		bucket.Setup,
		bucketacl.Setup,
		bucketobject.Setup,
		bucketobjectacl.Setup,
		bucketpolicy.Setup,
		bucketreplication.Setup,
		providerconfig.Setup,
		backuprds.Setup,
		crossregionbackupstrategy.Setup,
		databaselogsshrinking.Setup,
		extendloglink.Setup,
		instancerds.Setup,
		instanceeipassociate.Setup,
		ltslog.Setup,
		mysqlaccount.Setup,
		mysqlbinlog.Setup,
		mysqldatabase.Setup,
		mysqldatabaseprivilege.Setup,
		mysqldatabasetablerestore.Setup,
		parametergroup.Setup,
		pgaccount.Setup,
		pgaccountprivileges.Setup,
		pgaccountroles.Setup,
		pgdatabase.Setup,
		pgdatabaseprivilege.Setup,
		pghba.Setup,
		pgplugin.Setup,
		pgpluginparameter.Setup,
		pgpluginupdate.Setup,
		pgsqllimit.Setup,
		primarystandbyswitch.Setup,
		readreplicainstance.Setup,
		restore.Setup,
		sqlaudit.Setup,
		sqlserveraccount.Setup,
		sqlserverdatabase.Setup,
		sqlserverdatabaseprivilege.Setup,
		logtank.Setup,
		messagedetection.Setup,
		messagepublish.Setup,
		messagetemplate.Setup,
		subscription.Setup,
		subscriptionfilterpolicy.Setup,
		topic.Setup,
		imageautosync.Setup,
		imagepermissions.Setup,
		imageretentionpolicy.Setup,
		imagetrigger.Setup,
		organization.Setup,
		organizationpermissions.Setup,
		repository.Setup,
		repositorysharing.Setup,
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
