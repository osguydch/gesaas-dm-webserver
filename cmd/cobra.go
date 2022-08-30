package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"rm/cmd/api"
	"rm/cmd/version"
)

var rootCmd = &cobra.Command{
	Use:          "rmserver",
	Short:        "RmServer",
	SilenceUsage: true,
	Long:         `WebUI server`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	rootCmd.AddCommand(api.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

//Execute : apply commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
