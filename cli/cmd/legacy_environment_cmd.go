package cmd

import (
	"github.com/spf13/cobra"
)

var environmentApplyFile string

var environmentCmd = &cobra.Command{
	GroupID:    cmdGroupConfig.ID,
	Use:        "environment",
	Short:      "Manage your qualitytrace environments",
	Long:       "Manage your qualitytrace environments",
	Deprecated: "Please use `qualitytrace (apply|delete|export|get|list) environment` commands instead.",
	PreRun:     setupCommand(),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PostRun: teardownCommand,
}

var environmentApplyCmd = &cobra.Command{
	Use:        "apply",
	Short:      "Create or update an environment to Qualitytrace",
	Long:       "Create or update an environment to Qualitytrace",
	Deprecated: "Please use `qualitytrace apply environment --file [path]` command instead.",
	PreRun:     setupCommand(),
	Run: func(cmd *cobra.Command, args []string) {
		// call new apply command
		applyParams.DefinitionFile = dataStoreApplyFile
		applyCmd.Run(applyCmd, []string{"environment"})
	},
	PostRun: teardownCommand,
}

func init() {
	rootCmd.AddCommand(environmentCmd)

	environmentApplyCmd.PersistentFlags().StringVarP(&environmentApplyFile, "file", "f", "", "file containing the environment configuration")
	environmentCmd.AddCommand(environmentApplyCmd)
}
