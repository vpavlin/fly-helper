package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/superfly/flyctl/client"
	"github.com/vpavlin/fly-helper/internal/fly"
	"github.com/vpavlin/fly-helper/internal/secrets"
)

const DEFAULT_CONFIG_ENV_NAME = "FLY_HELPER_CONFG_ENV"

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

	if len(config.AppName) == 0 {
		flytoml, err := fly.NewFlyToml("./fly.toml")
		if err != nil {
			return nil, err
		}

		config.AppName = flytoml.App
		logrus.Infof("Using App name '%s' from fly.toml", config.AppName)
	}

	return &config, err
}

func (c Config) ToEnv() (string, string, error) {
	data, err := json.Marshal(c)
	if err != nil {
		return "", "", err
	}

	content := base64.StdEncoding.EncodeToString(data)

	return DEFAULT_CONFIG_ENV_NAME, content, nil
}

func (c Config) Push(fly *client.Client) error {
	name, content, err := c.ToEnv()
	if err != nil {
		return err
	}

	secretMap := map[string]string{
		name: content,
	}

	logrus.Infof("Pushing flyhelper config for app %s into %s", c.AppName, name)

	_, err = fly.API().SetSecrets(context.Background(), c.AppName, secretMap)
	return err
}
