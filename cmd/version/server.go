package version

import (
	"fmt"
	"github.com/spf13/cobra"
	"rm/common"
)

var (
	configYml string
	port      string
	mode      string
	StartCmd  = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "rmserver version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Printf("module: %s\n", common.Module)
	fmt.Printf("desc: %s\n", common.Desc)
	fmt.Printf("version: %s\n", common.Version)
	fmt.Printf("build: %s\n", common.Build)
	return nil
}
