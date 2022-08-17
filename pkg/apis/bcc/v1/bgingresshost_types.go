/*
 * @Description:
 * @Autor: gaoshiyao
 * @Date: 2021-09-02 10:29:02
 * @LastEditors: gaoshiyao
 * @LastEditTime: 2021-09-02 11:13:49
 */
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgIngressHostSpec defines the desired state of BgIngressHost
type BgIngressHostSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	Group     string `json:"group,omitempty"`
	Namespace string `json:"namespace,omitempty"`
	Lb        string `json:"lb,omitempty"`
	Host      string `json:"host,omitempty"`
	Approval  string `json:"approval,omitempty"`  // 审批状态
	ApplyUser string `json:"applyUser,omitempty"` // 申请人
}

// BgIngressHostStatus defines the observed state of BgIngressHost
type BgIngressHostStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
	// 创建人
	// +optional
	Creator string `json:"creator,omitempty" protobuf:"bytes,1,opt,name=creator"`

	// 创建时间
	// +optional
	Created metav1.Time `json:"created,omitempty" protobuf:"bytes,2,opt,name=created"`

	// 运行状态， Running , Completed
	// +optional
	Phase BgPhase `json:"phase,omitempty" protobuf:"bytes,3,opt,name=phase,casttype=BgPhase"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgIngressHost is the Schema for the bgingresshosts API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgingresshosts,scope=Cluster
type BgIngressHost struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BgIngressHostSpec   `json:"spec,omitempty"`
	Status BgIngressHostStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgIngressHostList contains a list of BgIngressHost
type BgIngressHostList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BgIngressHost `json:"items"`
}
