package config

import (
	"fmt"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
	"sigs.k8s.io/yaml"
)

func ValidateConfig(configFilePath string) error {

	buf, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		logrus.Warnf("unable to read configuration yaml file: %v", configFilePath)
		return fmt.Errorf("unable to read configuration yaml file: %v", configFilePath)
	}
	y, err := yaml.YAMLToJSON(buf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return err
	}

	schemaLoader := gojsonschema.NewStringLoader(schemav1beta1)
	documentLoader := gojsonschema.NewStringLoader(string(y))

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		logrus.Fatal(err.Error())
	}

	if !result.Valid() {
		logrus.Printf("The configuration is not valid. see errors below:\n")
		for _, desc := range result.Errors() {
			logrus.Printf("- %s\n", desc)
		}
		return fmt.Errorf("config file does not validate")
	}

	return nil

}
