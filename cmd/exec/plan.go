package exec

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	backend "github.com/quantize-io/tfctl/cmd/backend"
	util "github.com/quantize-io/tfctl/cmd/util"
)

func RunPlan(cmd *cobra.Command, args []string) {

	// Load the config file
	config, err := util.LoadConfigFromFile(viper.GetString("config-path"), viper.GetString("env-name"))
	if err != nil {
		logrus.Fatal(err)
	}

	// Configure the backend
	logrus.Infof("processing plan for env: %v", config.Metadata.Name)
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

	// Execute the plan
	params := make([]string, 0)
	params = append(params, "plan")

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
