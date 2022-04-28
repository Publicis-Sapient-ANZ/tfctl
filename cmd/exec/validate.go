package exec

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	util "github.com/quantize-io/tfctl/cmd/util"
)

func RunValidate(cmd *cobra.Command, args []string) {

	// Load the config file
	config, err := util.LoadConfigFromFile(viper.GetString("config-path"), viper.GetString("env-name"))
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
