package exec

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	backend "github.com/quantize-io/tfctl/cmd/backend"
	util "github.com/quantize-io/tfctl/cmd/util"
)

func RunDestroy(cmd *cobra.Command, args []string) {

	// Load the config file
	config, err := util.LoadConfigFromFile(viper.GetString("config-path"), viper.GetString("env-name"))
	if err != nil {
		logrus.Fatal(err)
	}

	// Configure the backend
	logrus.Infof("processing destroy for env: %v", config.Metadata.Name)
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

	// Execute the destroy
	params := make([]string, 0)
	params = append(params, "destroy")
	params = append(params, "-auto-approve")

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
