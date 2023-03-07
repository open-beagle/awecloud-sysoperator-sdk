/*
 * @Description:
 * @Autor: gaoshiyao
 * @Date: 2021-08-23 12:07:53
 * @LastEditors: gaoshiyao
 * @LastEditTime: 2021-09-13 14:35:34
 */
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BgClusterSpec defines the desired state of BgCluster
// +k8s:openapi-gen=true
type BgClusterSpec struct {
	Alias        string `json:"alias,omitempty" protobuf:"bytes,1,opt,name=alias"`
	APIServer    string `json:"apiserver" protobuf:"bytes,2,opt,name=apiserver"`
	MetricServer string `json:"metricserver,omitempty" protobuf:"bytes,3,opt,name=metricserver"`
	SAToken      string `json:"satoken" protobuf:"bytes,4,opt,name=satoken"`
	// gaoshiyao
	Status       string `json:"status,omitempty" protobuf:"bytes,5,opt,name=status"`           // 是否在线 up | down
	ClusterType  string `json:"clusterType,omitempty" protobuf:"bytes,6,opt,name=clusterType"` // 集群类型 core | member
	City         string `json:"city,omitempty" protobuf:"bytes,7,opt,name=city"`
	ComputerName string `json:"computerName,omitempty" protobuf:"bytes,8,opt,name=computerName"`
	Node         string `json:"node,omitempty" protobuf:"bytes,9,opt,name=node"`
	Cpu          string `json:"cpu,omitempty" protobuf:"bytes,10,opt,name=cpu"`
	Memory       string `json:"memory,omitempty" protobuf:"bytes,11,opt,name=memory"`
	Fs           string `json:"fs,omitempty" protobuf:"bytes,12,opt,name=fs"`
	Des          string `json:"des,omitempty" protobuf:"bytes,13,opt,name=des"`
	Secret       string `json:"secret,omitempty" protobuf:"bytes,1,opt,name=secret"`
}

// BgClusterStatus defines the observed state of BgCluster
// +k8s:openapi-gen=true
type BgClusterStatus struct {
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

// BgCluster is the Schema for the bgclusters API
// +kubebuilder:subresource:status
// +kubebuilder:resource:path=bgclusters,scope=Cluster
type BgCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	Spec   BgClusterSpec   `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
	Status BgClusterStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// BgClusterList contains a list of BgCluster
type BgClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items           []BgCluster `json:"items" protobuf:"bytes,2,rep,name=items"`
}

func init() {
	SchemeBuilder.Register(&BgCluster{}, &BgClusterList{})
}
