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

// CategrafglobalSpec defines the desired state of Categrafglobal
type CategrafglobalSpec struct {
	Global     Global     `json:"global"`
	Log        Log        `json:"log"`
	WriterOpt  WriterOpt  `json:"writer_opt"`
	Writers    []Writers  `json:"writers"`
	HTTP       HTTP       `json:"http"`
	Ibex       Ibex       `json:"ibex"`
	Heartbeat  Heartbeat  `json:"heartbeat"`
	Prometheus Prometheus `json:"prometheus"`
}
type Global struct {
	PrintConfigs bool     `json:"print_configs"`
	Hostname     string   `json:"hostname"`
	OmitHostname bool     `json:"omit_hostname"`
	Interval     int      `json:"interval"`
	Providers    []string `json:"providers"`
	Concurrency  int      `json:"concurrency"`
}
type Log struct {
	FileName   string `json:"file_name"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	LocalTime  bool   `json:"local_time"`
	Compress   bool   `json:"compress"`
}
type WriterOpt struct {
	Batch    int `json:"batch"`
	ChanSize int `json:"chan_size"`
}
type Writers struct {
	URL                 string   `json:"url"`
	TLSMinVersion       string   `json:"tls_min_version,omitempty"`
	TLSCa               string   `json:"tls_ca,omitempty"`
	TLSCert             string   `json:"tls_cert,omitempty"`
	TLSKey              string   `json:"tls_key,omitempty"`
	InsecureSkipVerify  bool     `json:"insecure_skip_verify,omitempty"`
	BasicAuthUser       string   `json:"basic_auth_user,omitempty"`
	BasicAuthPass       string   `json:"basic_auth_pass,omitempty"`
	Headers             []string `json:"headers,omitempty"`
	Timeout             int      `json:"timeout"`
	DialTimeout         int      `json:"dial_timeout"`
	MaxIdleConnsPerHost int      `json:"max_idle_conns_per_host"`
}
type HTTP struct {
	Enable             bool   `json:"enable"`
	Address            string `json:"address"`
	PrintAccess        bool   `json:"print_access"`
	RunMode            string `json:"run_mode"`
	IgnoreHostname     bool   `json:"ignore_hostname"`
	IgnoreGlobalLabels bool   `json:"ignore_global_labels"`
}
type Ibex struct {
	Enable   bool     `json:"enable"`
	Interval string   `json:"interval"`
	Servers  []string `json:"servers"`
	MetaDir  string   `json:"meta_dir"`
}
type Heartbeat struct {
	Enable              bool     `json:"enable"`
	URL                 string   `json:"url"`
	Interval            int      `json:"interval"`
	BasicAuthUser       string   `json:"basic_auth_user"`
	BasicAuthPass       string   `json:"basic_auth_pass"`
	Headers             []string `json:"headers,omitempty"`
	Timeout             int      `json:"timeout"`
	DialTimeout         int      `json:"dial_timeout"`
	MaxIdleConnsPerHost int      `json:"max_idle_conns_per_host"`
}
type Prometheus struct {
	Enable           bool   `json:"enable"`
	ScrapeConfigFile string `json:"scrape_config_file"`
	LogLevel         string `json:"log_level"`
	WalStoragePath   string `json:"wal_storage_path,omitempty"`
	WalMinDuration   int    `json:"wal_min_duration,omitempty"`
}

// CategrafglobalStatus defines the observed state of Categrafglobal
type CategrafglobalStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// Categrafglobal is the Schema for the categrafglobals API
type Categrafglobal struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CategrafglobalSpec   `json:"spec,omitempty"`
	Status CategrafglobalStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// CategrafglobalList contains a list of Categrafglobal
type CategrafglobalList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Categrafglobal `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Categrafglobal{}, &CategrafglobalList{})
}
