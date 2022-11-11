package config

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"github.com/vpavlin/fly-helper/internal/secrets"
)

type Config struct {
	AppName string
	Secrets secrets.Secrets
}

func NewConfigFromCommand(cmd *cobra.Command) (*Config, error) {
	configEnv, err := cmd.Flags().GetString("config-env")
	if err != nil {
		return nil, err
	}

	if len(configEnv) > 0 {
		return NewConfigFromEnv(configEnv)
	}

	configFile, err := cmd.Flags().GetString("config")
	if err != nil {
		return nil, err
	}

	return NewConfigFromFile(configFile)

}

func NewConfigFromEnv(envvar string) (*Config, error) {
	content, ok := os.LookupEnv(envvar)
	if !ok {
		return nil, fmt.Errorf("Failed to load the env var %s", envvar)
	}

	data, err := base64.StdEncoding.DecodeString(content)
	if err != nil {
		return nil, err
	}

	return NewConfig(data)
}

func NewConfigFromFile(path string) (*Config, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	return NewConfig(data)
}

func NewConfig(data []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(data, &config)

	return &config, err
}
