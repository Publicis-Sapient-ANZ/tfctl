package exec

import (
	"fmt"
	"os"
	"path/filepath"

	cp "github.com/otiai10/copy"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/quantize-io/tfctl/cmd/model"
)

func BuildConfigSet(config model.Config) (string, error) {

	var workingDir = viper.GetString("config-path")
	var buildDir = fmt.Sprintf("%v/build/%v", workingDir, config.Metadata.Environment)

	logrus.Infof("building config set at: %v", buildDir)

	// Create the buld dir if not exists
	err := os.MkdirAll(buildDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	// Clean out the build dir from previous runs (only tf related files)
	d, err := os.Open(buildDir)
	if err != nil {
		return "", err
	}
	defer d.Close()

	files, err := d.Readdir(-1)
	if err != nil {
		return "", err
	}

	for _, file := range files {
		if file.Mode().IsRegular() {
			switch filepath.Ext(file.Name()) {
			case ".tf":
				os.Remove(fmt.Sprintf("%v/%v", buildDir, file.Name()))
			case ".tfvars":
				os.Remove(fmt.Sprintf("%v/%v", buildDir, file.Name()))
			case ".json":
				os.Remove(fmt.Sprintf("%v/%v", buildDir, file.Name()))
			case ".yaml":
				os.Remove(fmt.Sprintf("%v/%v", buildDir, file.Name()))
			}
		}
	}

	// Copy in the new consolidated config set
	err = cp.Copy(fmt.Sprintf("%v/terraform", workingDir), buildDir)
	if err != nil {
		return "", err
	}

	err = cp.Copy(fmt.Sprintf("%v/environments/%v", workingDir, config.Metadata.Environment), buildDir)
	if err != nil {
		return "", err
	}

	return buildDir, nil

}
