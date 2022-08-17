/*
 * @Description:
 * @Autor: gaoshiyao
 * @Date: 2021-08-23 12:07:53
 * @LastEditors: gaoshiyao
 * @LastEditTime: 2021-09-02 19:46:12
 */
// NOTE: Boilerplate only.  Ignore this file.

// Package v1 contains API Schema definitions for the bcc v1 API group
// +k8s:deepcopy-gen=package,register
// +groupName=bcc.bd-apaas.com
package v1

import (
	"github.com/open-beagle/awecloud-sysoperator-sdk/pkg/traefik/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion   = schema.GroupVersion{Group: "bcc.bd-apaas.com", Version: "v1"}
	SchemeGatewayVersion = schema.GroupVersion{Group: "bcc.bd-apaas.com", Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	SchemeGatewayBuilder = runtime.NewSchemeBuilder(addGateWayTypes)
)

// Adds the list of known types to the given scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&BgUser{},
		&BgUserList{},
		&BgToken{},
		&BgTokenList{},
		&BgStorage{},
		&BgStorageList{},
		&BgSetting{},
		&BgSettingList{},
		&BgRole{},
		&BgRoleUser{},
		&BgRoleList{},
		&BgRoleUserList{},
		&BgNodeAdd{},
		&BgNodeAddList{},
		&BgNamespace{},
		&BgNamespaceList{},
		&BgMenu{},
		&BgMenuList{},
		&BgLoginLog{},
		&BgLoginLogList{},
		&BgGroupUser{},
		&BgGroupUserList{},
		&BgGroupNamespace{},
		&BgGroupNamespaceList{},
		&BgGroup{},
		&BgGroupList{},
		&BgCluster{},
		&BgClusterList{},
		&BgCloudShell{},
		&BgCloudShellList{},
		&BgAppStore{},
		&BgAppStoreList{},
		&BgAppDeploy{},
		&BgAppDeployList{},
		&BgIngressHost{},
		&BgIngressHostList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}

func addGateWayTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGatewayVersion,
		&v1alpha1.IngressHost{},
		&v1alpha1.IngressRoute{},
		&v1alpha1.IngressRouteList{},
		&v1alpha1.IngressRouteTCP{},
		&v1alpha1.IngressRouteTCPList{},
		&v1alpha1.IngressRouteUDP{},
		&v1alpha1.IngressRouteUDPList{},
		&v1alpha1.Middleware{},
		&v1alpha1.MiddlewareList{},
		&v1alpha1.TLSOption{},
		&v1alpha1.TLSOptionList{},
		&v1alpha1.TLSStore{},
		&v1alpha1.TLSStoreList{},
		&v1alpha1.BeagleService{},
		&v1alpha1.BeagleServiceList{},
	)
	metav1.AddToGroupVersion(scheme, SchemeGatewayVersion)
	return nil
}
