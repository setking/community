package cmd

import (
	"fmt"
	"runtime"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "输出版本信息",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("myapp  version : 1.0.0\n")
		fmt.Printf("go     version : %s\n", runtime.Version())
		fmt.Printf("os/arch        : %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}
