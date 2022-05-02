package exec

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	config "github.com/quantize-io/tfctl/cmd/config"
)

func RunValidate(cmd *cobra.Command, args []string) {

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

	// Build the config set
	buildPath, err := BuildConfigSet(*config)
	if err != nil {
		logrus.Fatal(err)

	}

	// Execute the validation
	params := make([]string, 0)
	params = append(params, "validate")

	commandConfig := CommandConfig{
		Command:          "terraform",
		Paramaters:       params,
		WorkingDirecotry: buildPath,
	}
	_, err2 := ExecuteCommand(commandConfig)
	if err2 != nil {
		logrus.Fatal(err)
	}

}
