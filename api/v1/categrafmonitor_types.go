/*
Copyright 2024.

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

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CategrafMonitorSpec defines the desired state of CategrafMonitor
type CategrafMonitorSpec struct {
	Node         string         `json:"node"`
	OtherFile    []OtherFile    `json:"otherFile,omitempty"`
	Certificate  []Certificate  `json:"certificate,omitempty"`
	MonitorLite  []MonitorLite  `json:"monitorLite,omitempty"`
	MonitorSuper []MonitorSuper `json:"monitorSuper,omitempty"`
}

type OtherFile struct {
	Name string `json:"name"`
	Data string `json:"data"`
}

type Certificate struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
type MonitorLite struct {
	Name string `json:"name"`
	Data string `json:"data"`
}
type MonitorSuper struct {
	Name      string     `json:"name"`
	Interval  int        `json:"interval" toml:"interval"`
	Mappings  string     `json:"mappings,omitempty" toml:"mappings"`
	Instances []Instance `json:"instances" toml:"instances"`
}
type Instance struct {
	Data string `json:"data" toml:"data"`
}

// CategrafMonitorStatus defines the observed state of CategrafMonitor
type CategrafMonitorStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// CategrafMonitor is the Schema for the categrafmonitors API
type CategrafMonitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CategrafMonitorSpec   `json:"spec,omitempty"`
	Status CategrafMonitorStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CategrafMonitorList contains a list of CategrafMonitor
type CategrafMonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CategrafMonitor `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CategrafMonitor{}, &CategrafMonitorList{})
}
