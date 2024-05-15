package cmd

import (
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	GroupID: cmdGroupConfig.ID,
	Use:     "server",
	Short:   "Manage your qualitytrace server",
	Long:    "Manage your qualitytrace server",
	PreRun:  setupCommand(SkipVersionMismatchCheck()),
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
	PostRun: teardownCommand,
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
