package config

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"

	model "github.com/quantize-io/tfctl/cmd/model"
)

func LoadConfigFromFile(filePath string) (*model.Config, error) {

	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		logrus.Warnf("unable to read configuration yaml file: %v", filePath)
		return nil, fmt.Errorf("unable to read configuration yaml file: %v", filePath)
	}

	config := &model.Config{}

	err1 := yaml.Unmarshal(buf, config)
	if err1 != nil {
		logrus.Warnf("unable to parse yaml file: %v", filePath)
		return nil, fmt.Errorf("unable to parse yaml file: %v", filePath)
	}

	logrus.Infof("loaded config file for: %v", config.Metadata.Name)

	return config, nil

}
