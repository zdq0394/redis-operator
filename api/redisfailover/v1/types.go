package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisFailover represents a Redis failover
type RedisFailover struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisFailoverSpec   `json:"spec"`
	Status            RedisFailoverStatus `json:"status"`
}

// RedisFailoverSpec represents a Redis failover spec
type RedisFailoverSpec struct {
	Redis    RedisSettings    `json:"redis,omitempty"`
	Sentinel SentinelSettings `json:"sentinel,omitempty"`
}

// RedisSettings defines the specification of the redis cluster
type RedisSettings struct {
	Image             string                      `json:"image,omitempty"`
	Replicas          int32                       `json:"replicas,omitempty"`
	Resources         corev1.ResourceRequirements `json:"resources,omitempty"`
	CustomConfig      []string                    `json:"customConfig,omitempty"`
	Command           []string                    `json:"command,omitempty"`
	ShutdownConfigMap string                      `json:"shutdownConfigMap,omitempty"`
	Storage           RedisStorage                `json:"storage,omitempty"`
	Exporter          RedisExporter               `json:"exporter,omitempty"`
	Perceptron        RedisPerceptron             `json:"perceptron,omitempty"`
	Affinity          *corev1.Affinity            `json:"affinity,omitempty"`
	SecurityContext   *corev1.PodSecurityContext  `json:"securityContext,omitempty"`
	Tolerations       []corev1.Toleration         `json:"tolerations,omitempty"`
	HostPort          int32                       `json:"hostport,omitempty"`
	Password          string                      `json:"password,omitempty"`
}

// SentinelSettings defines the specification of the sentinel cluster
type SentinelSettings struct {
	Image           string                      `json:"image,omitempty"`
	Replicas        int32                       `json:"replicas,omitempty"`
	Resources       corev1.ResourceRequirements `json:"resources,omitempty"`
	CustomConfig    []string                    `json:"customConfig,omitempty"`
	Command         []string                    `json:"command,omitempty"`
	Affinity        *corev1.Affinity            `json:"affinity,omitempty"`
	SecurityContext *corev1.PodSecurityContext  `json:"securityContext,omitempty"`
	Tolerations     []corev1.Toleration         `json:"tolerations,omitempty"`
	HostPort        int32                       `json:"hostport,omitempty"`
}

// RedisExporter defines the specification for the redis exporter
type RedisExporter struct {
	Enabled  bool   `json:"enabled,omitempty"`
	Image    string `json:"image,omitempty"`
	HostPort int32  `json:"hostport,omitempty"`
}

// RedisPerceptron defines the specification for the redis perceptron
type RedisPerceptron struct {
	Enabled      bool   `json:"enabled,omitempty"`
	Image        string `json:"image,omitempty"`
	ProxyURL     string `json:"proxyURL,omitempty"`
	RegisterPort int32  `json:"registerPort,omitempty"`
	TTL          string `json:"ttl,omitempty"`
	MAXConn      string `json:"maxConn,omitempty"`
}

// RedisStorage defines the structure used to store the Redis Data
type RedisStorage struct {
	KeepAfterDeletion     bool                          `json:"keepAfterDeletion,omitempty"`
	EmptyDir              *corev1.EmptyDirVolumeSource  `json:"emptyDir,omitempty"`
	PersistentVolumeClaim *corev1.PersistentVolumeClaim `json:"persistentVolumeClaim,omitempty"`
}

// RedisNode defines the structure used to store the Redis Node Info
type RedisNode struct {
	PodIP    string `json:"podIP,omitempty"`
	HostIP   string `json:"hostIP,omitempty"`
	IsMaster bool   `json:"isMaster,omitempty"`
}

// SentinelNode defines the structure used to store the Sentinel Node Info
type SentinelNode struct {
	PodIP  string `json:"podIP,omitempty"`
	HostIP string `json:"hostIP,omitempty"`
}

// RedisFailoverStatus defines the structure used to store the RedisFailoverStatus Info
type RedisFailoverStatus struct {
	RedisNodes    []RedisNode    `json:"redisNodes,omitempty"`
	SentinelNodes []SentinelNode `json:"sentinelNodes,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RedisFailoverList represents a Redis failover list
type RedisFailoverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []RedisFailover `json:"items"`
}
