package secrets

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"encoding/base64"

	"github.com/sirupsen/logrus"
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

	logrus.Infof("Pushing %d secrets for %s", len(prep), appName)

	_, err = fly.API().SetSecrets(context.Background(), appName, prep)
	if err != nil {
		return err
	}

	return nil
}

func (s Secrets) WriteSecrets() error {
	var multierrors multierror.Multierror
	for _, e := range s.Output {
		name := ToSecretName(PREFIX, e.Name)

		content, ok := os.LookupEnv(name)
		if !ok {
			multierrors = append(multierrors, fmt.Errorf("Failed to get env var %s", name))
		}

		data, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			multierrors = append(multierrors, err)
		}

		dir := filepath.Dir(e.Path)
		_, err = os.Stat(dir)
		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				err = os.MkdirAll(dir, 0700)
				if err != nil {
					multierrors = append(multierrors, err)
				}
			} else {
				multierrors = append(multierrors, err)
			}
		}
		err = ioutil.WriteFile(e.Path, data, os.FileMode(0600))
		if err != nil {
			multierrors = append(multierrors, err)
		}
	}

	return multierrors.ToError()
}
