package metasearch

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"storj.io/storj/satellite/metainfo"
)

const (
	defaulConfigPath        = "/root/.local/share/storj/metasearch/config.yaml"
	defaultDatabase         = "cockroach://root@cockroachdb-public:26257/master?sslmode=disable"
	defaultMetainfoDatabase = "cockroach://root@cockroachdb-public:26257/master?sslmode=disable"
	defaultEndpoint         = ":9998"
)

type Config struct {
	Database string
	Metainfo metainfo.Config
	Log      zap.Config
	Endpoint string
}

func (c *Config) Read() {
	c.setDefault()
	c.setConfig()
	c.setEnv()
}

func (c *Config) setDefault() {
	c.Log = zap.NewDevelopmentConfig()
	c.Database = defaultDatabase
	c.Metainfo.DatabaseURL = defaultMetainfoDatabase
	c.Endpoint = defaultEndpoint
}

func (c *Config) setConfig() {
	c.Log = zap.NewDevelopmentConfig()
	yamlData, err := os.ReadFile(defaulConfigPath)
	if err != nil {
		fmt.Println(err)
	} else {
		err = yaml.Unmarshal(yamlData, c)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (c *Config) setEnv() {
	envDatabase := os.Getenv("STORJ_DATABASE")
	if envDatabase != "" {
		c.Database = envDatabase
	}
	envMetainfoDatabaseURL := os.Getenv("STORJ_METAINFO_DATABASE_URL")
	if envMetainfoDatabaseURL != "" {
		c.Metainfo.DatabaseURL = envMetainfoDatabaseURL
	}
	envEndpoint := os.Getenv("STORJ_METASEARCH_ENDPOINT")
	if envEndpoint != "" {
		c.Endpoint = envEndpoint
	}
}
