package exec

import (
	"fmt"

	"github.com/quantize-io/tfctl/cmd/model"
	"github.com/sirupsen/logrus"
)

func ExecuteInit(workingDir string, config model.Config) (result CommandResult, err error) {

	params := make([]string, 0)
	params = append(params, "init")

	initCommandConfig := CommandConfig{
		Command:          "terraform",
		Paramaters:       params,
		WorkingDirecotry: workingDir,
	}

	logrus.Info(fmt.Printf("Starting terraform init for environment: %s", workingDir))
	_, error := ExecuteCommand(initCommandConfig)
	if error != nil {
		logrus.Error(error)
		return result, error
	}

	return result, nil

}
