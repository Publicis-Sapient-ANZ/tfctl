package util

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v3"

	model "github.com/quantize-io/tfctl/cmd/model"
)

func LoadConfigFromFile(configPath string, envName string) (*model.Config, error) {

	var filepath string = fmt.Sprintf("%v/environments/%v/%v.tfctl.yaml", configPath, envName, envName)

	buf, err := ioutil.ReadFile(filepath)
	if err != nil {
		logrus.Warnf("unable to read configuration yaml file: %v", filepath)
		return nil, fmt.Errorf("unable to read configuration yaml file: %v", filepath)
	}

	config := &model.Config{}

	err1 := yaml.Unmarshal(buf, config)
	if err1 != nil {
		logrus.Warnf("unable to parse yaml file: %v", filepath)
		return nil, fmt.Errorf("unable to parse yaml file: %v", filepath)
	}

	logrus.Infof("loaded config file for: %v", config.Metadata.Name)

	return config, nil

}
