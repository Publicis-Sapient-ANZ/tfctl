package exec

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	backend "github.com/quantize-io/tfctl/cmd/backend"
	config "github.com/quantize-io/tfctl/cmd/config"
)

func RunApply(cmd *cobra.Command, args []string) {

	var filePath string = fmt.Sprintf("%v/environments/%v/%v.tfctl.yaml", viper.GetString("config-path"), viper.GetString("env-name"), viper.GetString("env-name"))

	// Validate the config schema
	err := config.ValidateConfig(filePath)
	if err != nil {
		logrus.Fatal("config file does not validate")
	}

	// Load the config file
	config, err := config.LoadConfigFromFile(filePath)
	if err != nil {
		logrus.Fatal(err)
	}

	// Configure the backend
	logrus.Infof("processing apply for env: %v", config.Metadata.Name)
	err = backend.InitBackend(config)
	if err != nil {
		logrus.Fatal(err)
	}

	// Build the config set
	buildPath, err := BuildConfigSet(*config)
	if err != nil {
		logrus.Fatal(err)

	}

	// Initialise
	_, err1 := ExecuteInit(buildPath, *config)
	if err1 != nil {
		logrus.Fatal(err)
	}

	// Execute the apply
	params := make([]string, 0)
	params = append(params, "apply")
	params = append(params, "-input=false")
	params = append(params, "-auto-approve")

	commandConfig := CommandConfig{
		Command:          "terraform",
		Paramaters:       params,
		WorkingDirecotry: buildPath,
	}
	commandResult, err2 := ExecuteCommand(commandConfig)
	if err2 != nil {
		logrus.Fatal(err)
	}

	// Exit with the code captured by the command
	os.Exit(commandResult.exitCode)

}
