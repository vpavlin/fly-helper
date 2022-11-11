package secrets

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"encoding/base64"

	"github.com/superfly/flyctl/client"
	"github.com/vpavlin/fly-helper/internal/multierror"
)

func (s Secrets) PrepSetSecrets() (map[string]string, error) {
	result := make(map[string]string)
	for _, e := range s.Input {
		content, err := ioutil.ReadFile(e.Path)
		if err != nil {
			return result, err
		}
		result[ToSecretName(PREFIX, e.Name)] = base64.StdEncoding.EncodeToString([]byte(content))
	}

	return result, nil
}

func (s Secrets) Push(fly *client.Client, appName string) error {
	prep, err := s.PrepSetSecrets()
	if err != nil {
		return err
	}

	if len(prep) == 0 {
		return nil
	}

	_, err = fly.API().SetSecrets(context.Background(), appName, prep)
	if err != nil {
		return err
	}

	return nil
}

func (s Secrets) WriteSecrets() error {
	var errors multierror.Multierror
	for _, e := range s.Output {
		name := ToSecretName(PREFIX, e.Name)

		content, ok := os.LookupEnv(name)
		if !ok {
			errors = append(errors, fmt.Errorf("Failed to get env var %s", name))
		}

		data, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			errors = append(errors, err)
		}
		err = ioutil.WriteFile(e.Path, data, os.FileMode(0600))
		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors.ToError()
}
