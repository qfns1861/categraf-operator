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
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CategrafglobalSpec defines the desired state of Categrafglobal
type CategrafglobalSpec struct {
	// from config.json
	Global     Global           `json:"global" toml:"global"`
	WriterOpt  WriterOpt        `json:"writer_opt,omitempty" toml:"writer_opt"`
	Writers    []WriterOption   `json:"writers,omitempty" toml:"writers"`
	HTTP       *HTTP            `json:"http,omitempty" toml:"http"`
	Interval   Duration         `json:"interval,omitempty" toml:"interval"`
	Prometheus *Prometheus      `json:"prometheus,omitempty" toml:"prometheus"`
	Ibex       *IbexConfig      `json:"ibex,omitempty" toml:"ibex"`
	Heartbeat  *HeartbeatConfig `json:"heartbeat,omitempty" toml:"heartbeat"`
	Log        Log              `json:"log,omitempty" toml:"log"`

	HTTPProviderConfig *HTTPProviderConfig `json:"http_provider"`
}

type Global struct {
	PrintConfigs bool              `json:"print_configs" toml:"print_configs"`
	Hostname     string            `json:"hostname" toml:"hostname"`
	OmitHostname bool              `json:"omit_hostname" toml:"omit_hostname"`
	Labels       map[string]string `json:"labels" toml:"labels"`
	Precision    string            `json:"precision" toml:"precision"`
	Providers    []string          `json:"providers" toml:"providers"`
	Concurrency  int               `json:"concurrency" toml:"concurrency"`
}

type Log struct {
	FileName   string `json:"file_name" toml:"file_name"`
	MaxSize    int    `json:"max_size" toml:"max_size"`
	MaxAge     int    `json:"max_age" toml:"max_age"`
	MaxBackups int    `json:"max_backups" toml:"max_backups"`
	LocalTime  bool   `json:"local_time" toml:"local_time"`
	Compress   bool   `json:"compress" toml:"compress"`
}

type WriterOpt struct {
	Batch    int `json:"batch" toml:"batch"`
	ChanSize int `json:"chan_size" toml:"chan_size"`
}

type WriterOption struct {
	Url                 string   `json:"url" toml:"url"`
	BasicAuthUser       string   `json:"basic_auth_user" toml:"basic_auth_user"`
	BasicAuthPass       string   `json:"basic_auth_pass" toml:"basic_auth_pass"`
	Headers             []string `json:"headers" toml:"headers"`
	Timeout             int64    `json:"timeout" toml:"timeout"`
	DialTimeout         int64    `json:"dial_timeout" toml:"dial_timeout"`
	MaxIdleConnsPerHost int      `json:"max_idle_conns_per_host" toml:"max_idle_conns_per_host"`
	ClientConfig    `json:"tls,omitempty" toml:"tls"`
}

type HTTP struct {
	Enable             bool   `json:"enable" toml:"enable"`
	Address            string `json:"address" toml:"address"`
	PrintAccess        bool   `json:"print_access" toml:"print_access"`
	RunMode            string `json:"run_mode" toml:"run_mode"`
	IgnoreHostname     bool   `json:"ignore_hostname" toml:"ignore_hostname"`
	IgnoreGlobalLabels bool   `json:"ignore_global_labels" toml:"ignore_global_labels"`
	CertFile           string `json:"cert_file" toml:"cert_file"`
	KeyFile            string `json:"key_file" toml:"key_file"`
	ReadTimeout        int    `json:"read_timeout" toml:"read_timeout"`
	WriteTimeout       int    `json:"write_timeout" toml:"write_timeout"`
	IdleTimeout        int    `json:"idle_timeout" toml:"idle_timeout"`
}

type IbexConfig struct {
	Interval Duration `json:"interval,omitempty" toml:"interval"`
	Enable   bool     `json:"enable" toml:"enable"`
	MetaDir  string   `json:"meta_dir" toml:"meta_dir"`
	Servers  []string `json:"servers" toml:"servers"`
}

type HTTPProviderConfig struct {
	Interval       int64    `json:"interval,omitempty" toml:"interval"`
	RemoteUrl      string   `json:"remote_url" toml:"remote_url"`
	Headers        []string `json:"headers" toml:"headers"`
	AuthUsername   string   `json:"basic_auth_user" toml:"basic_auth_user"`
	AuthPassword   string   `json:"basic_auth_pass" toml:"basic_auth_pass"`
	Timeout        int      `json:"timeout" toml:"timeout"`
	ReloadInterval int      `json:"reload_interval" toml:"reload_interval"`
}

type HeartbeatConfig struct {
	Enable              bool     `json:"enable" toml:"enable"`
	Url                 string   `json:"url" toml:"url"`
	BasicAuthUser       string   `json:"basic_auth_user" toml:"basic_auth_user"`
	BasicAuthPass       string   `json:"basic_auth_pass" toml:"basic_auth_pass"`
	Headers             []string `json:"headers" toml:"headers"`
	Timeout             int64    `json:"timeout" toml:"timeout"`
	DialTimeout         int64    `json:"dial_timeout" toml:"dial_timeout"`
	MaxIdleConnsPerHost int      `json:"max_idle_conns_per_host" toml:"max_idle_conns_per_host"`

	HTTPProxy        `json:"http_proxy" toml:"http_proxy"`
	ClientConfig `json:"tls,omitempty" toml:"tls"`
}

type HTTPProxy struct {
	ProxyURL string `json:"proxy_url" toml:"proxy_url"`
}

type Duration time.Duration

type Prometheus struct {
	Enable            bool     `json:"enable" toml:"enable"`
	LogLevel          string   `json:"log_level" toml:"loglever"`
	ScrapeConfigFile  string   `json:"scrape_config_file" toml:"scrapeConfigFile"`
	WebAddress        string   `json:"web_address" toml:"webAddress"`
	StoragePath       string   `json:"wal_storage_path" toml:"storagePath"`
	RetentionSize     string   `json:"retention_size" toml:"retentionSize"`
	MinBlockDuration  Duration `json:"min_block_duration,omitempty" toml:"min_block_duration,omitempty"`
	MaxBlockDuration  Duration `json:"max_block_duration,omitempty" toml:"max_block_duration,omitempty"`
	RetentionDuration Duration `json:"retention_time_duration,omitempty" toml:"retention_time_duration,omitempty"`
}

type ClientConfig struct {
	UseTLS             bool   `json:"use_tls,omitempty" toml:"use_tls"`
	TLSCA              string `json:"tls_ca,omitempty" toml:"tls_ca"`
	TLSCert            string `json:"tls_cert,omitempty" toml:"tls_cert"`
	TLSKey             string `json:"tls_key,omitempty" toml:"tls_key"`
	TLSKeyPwd          string `json:"tls_key_pwd,omitempty" toml:"tls_key_pwd"`
	InsecureSkipVerify bool   `json:"insecure_skip_verify,omitempty" toml:"insecure_skip_verify"`
	ServerName         string `json:"tls_server_name,omitempty" toml:"tls_server_name"`
	TLSMinVersion      string `json:"tls_min_version,omitempty" toml:"tls_min_version"`
	TLSMaxVersion      string `json:"tls_max_version,omitempty" toml:"tls_max_version"`
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
