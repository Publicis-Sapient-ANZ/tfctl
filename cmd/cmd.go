package cmd

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	exec "github.com/quantize-io/tfctl/cmd/exec"
)

var v string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "tfctl",
	Short: "cftl terraform command line",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	RootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		if err := setUpLogs(os.Stdout, v); err != nil {
			return err
		}
		return nil
	}

	RootCmd.PersistentFlags().StringVarP(&v, "verbosity", "v", logrus.InfoLevel.String(), "Log level (debug, info, warn, error, fatal, panic")

	RootCmd.PersistentFlags().String("env-name", "", "The environment for which you wish to perform the terraform operation on")
	RootCmd.MarkPersistentFlagRequired("env-name") //nolint

	viper.BindPFlag("env-name", RootCmd.PersistentFlags().Lookup("env-name")) //nolint

	RootCmd.PersistentFlags().String("config-path", "", "Pointer to the root of the configuration config set (if you are not running the tfctl from the working directory)")
	viper.BindPFlag("config-path", RootCmd.PersistentFlags().Lookup("config-path")) //nolint

	RootCmd.AddCommand(PlanCmd)
	RootCmd.AddCommand(ValidateCmd)
	RootCmd.AddCommand(ApplyCmd)
	RootCmd.AddCommand(DestroyCmd)

}

var ValidateCmd = &cobra.Command{
	Use:   "validate",
	Short: "perform a terraform validate",
	Run: func(cmd *cobra.Command, args []string) {
		exec.RunValidate(cmd, args)
	},
}

var PlanCmd = &cobra.Command{
	Use:   "plan",
	Short: "perform a terraform plan",
	Run: func(cmd *cobra.Command, args []string) {
		exec.RunPlan(cmd, args)
	},
}

var ApplyCmd = &cobra.Command{
	Use:   "apply",
	Short: "perform a terraform apply",
	Run: func(cmd *cobra.Command, args []string) {
		exec.RunApply(cmd, args)
	},
}

var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "perform a terraform destroy",
	Run: func(cmd *cobra.Command, args []string) {
		exec.RunDestroy(cmd, args)
	},
}

func setUpLogs(out io.Writer, level string) error {

	customFormatter := new(logrus.TextFormatter)
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	customFormatter.FullTimestamp = true
	logrus.SetFormatter(customFormatter)
	logrus.SetOutput(out)
	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		return err
	}
	logrus.SetLevel(lvl)
	return nil
}
