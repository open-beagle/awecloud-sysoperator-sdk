/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// BgUser ， 用户对象，仅解决登录的问题
// 说明： 与用户BgUser有关系的对象 ， BgRoleUser ， BgGroupUser
//        移除BgUser时，要移除与之相关的绑定，角色与Group

// BgUserSpec defines the desired state of BgUser
// +k8s:openapi-gen=true
type BgUserSpec struct {
	Alias    string         `json:"alias"`
	Phone    string         `json:"phone,omitempty"`
	Email    string         `json:"email"`
	Avatar   string         `json:"avatar,omitempty"`
	Disable  bool           `json:"disable,omitempty"`
	Password string         `json:"password,omitempty"`
	Default  *BgUserDefault `json:"default,omitempty"`
	UserType string         `json:"userType"` // 用户类型 sa:超管 tenant:租户  user:普通用户
}

// BgUserStatus defines the observed state of BgUser
// +k8s:openapi-gen=true
type BgUserStatus struct {
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty"`

	// Token ， 会Watch ServiceAccount.Secrets，随之变化
	// +optional
	Token string `json:"token,omitempty"`

	// 分组
	// +optional
	Groups map[string]BgUserGroup `json:"groups,omitempty"`

	// 角色
	// +optional
	Roles []string `json:"roles,omitempty"`

	// 用户设置密码以后，密码Password将会从Spec从删除加密保存在Status之中
	// +optional
	Password string `json:"password,omitempty"`

	// 在用户进入垃圾回收时，设置Disable=true
	// +optional
	Disable bool `json:"disable,omitempty"`

	// 用户最后一次登录时间
	// +optional
	LastLogin *metav1.Time `json:"lastLogin,omitempty"`
}

// BgUserDefault defines the observed state of BgUser
// +k8s:openapi-gen=true
type BgUserDefault struct {
	Group   string `json:"group,omitempty"`
	Cluster string `json:"cluster,omitempty"`
}

// BgUserGroup defines the observed state of BgUser
// +k8s:openapi-gen=true
type BgUserGroup struct {
	Name       string   `json:"name"`
	Alias      string   `json:"alias,omitempty"`
	Des        string   `json:"des,omitempty"`
	Role       string   `json:"role"`
	Default    string   `json:"namespace,omitempty"`
	Namespaces []string `json:"namespaces,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgUser is the Schema for the bgusers API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgusers,scope=Cluster
type BgUser struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BgUserSpec   `json:"spec,omitempty"`
	Status BgUserStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgUserList contains a list of BgUser
type BgUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BgUser `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BgUser{}, &BgUserList{})
}
