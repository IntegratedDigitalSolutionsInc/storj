package metasearch

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"storj.io/storj/satellite/metainfo"
)

const (
	defaulConfigPath = "/root/.local/share/storj/metasearch/config.yaml"
	defaultEndpoint  = ":9998"
)

type Config struct {
	Database string
	Metainfo metainfo.Config
	Log      zap.Config
	Endpoint string
}

func (c *Config) Read() {
	c.Log = zap.NewDevelopmentConfig()
	yamlData, err := os.ReadFile(defaulConfigPath)
	if err != nil {
		fmt.Println(err)
	}
	c.Endpoint = defaultEndpoint

	err = yaml.Unmarshal(yamlData, c)
	if err != nil {
		fmt.Println(err)
	}

	c.Database = os.Getenv("STORJ_DATABASE")
	c.Metainfo.DatabaseURL = os.Getenv("STORJ_METAINFO_DATABASE_URL")
	c.Endpoint = os.Getenv("STORJ_METASEARCH_ENDPOINT")
}
