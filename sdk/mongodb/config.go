package mongodb

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Config ...
type Config struct {
	Username    string        `json:"username" mapstructure:"username"`
	Password    string        `json:"password" mapstructure:"password"`
	Database    string        `json:"database" mapstructure:"database"`
	Clusters    string        `json:"clusters" mapstructure:"clusters"`
	Timeout     int32         `json:"timeout" mapstructure:"timeout"`
	Options     *MongoOptions `json:"options" mapstructure:"options"`
	AutoMigrate bool          `json:"auto_migrate" mapstructure:"auto_migrate"`
}

// MongoOptions ...
type MongoOptions struct {
	ReplicaSet string `json:"replica_set" mapstructure:"replica_set"`
	SSL        bool   `json:"ssl" mapstructure:"ssl"`
	AuthSource string
}

// GenConnectString ...
func (m *Config) GenConnectString() string {
	if m.Timeout == 0 {
		m.Timeout = 10 // default
	}
	if strings.HasPrefix(m.Clusters, "mongodb://") {
		return m.Clusters
	}

	clusters := strings.Split(m.Clusters, ",")
	nClusterArr := []string{}
	for _, c := range clusters {
		nClusterArr = append(nClusterArr, strings.TrimSpace(c))
	}
	clustersStr := strings.Join(nClusterArr, ",")
	connectStr := fmt.Sprintf("mongodb://%s:%s@%s/%s", m.Username, m.Password, clustersStr, m.Database)
	if len(m.Username) == 0 {
		connectStr = fmt.Sprintf("mongodb://%s/%s", clustersStr, m.Database)
	}

	if m.Options == nil {
		return connectStr
	}
	ops := url.Values{}
	if len(m.Options.ReplicaSet) > 0 {
		ops.Set("replicaSet", m.Options.ReplicaSet)
	}
	if m.Options.SSL {
		ops.Set("ssl", strconv.FormatBool(true))
	}
	if m.Options.AuthSource != "" {
		ops.Set("authSource", m.Options.AuthSource)
	}
	if len(ops.Encode()) > 0 {
		connectStr = fmt.Sprintf("%s?%s", connectStr, ops.Encode())
	}
	return connectStr
}
