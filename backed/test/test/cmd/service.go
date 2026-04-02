package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "启动 HTTP 服务",
	Long:  `启动一个 HTTP 服务，地址和端口由 --host / --port 或配置文件决定。`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		// 子命令需要显式调用 bindFlags（cobra 不会自动向上冒泡 PersistentPreRunE）
		return bindFlags(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		host := viper.GetString("host")
		port := viper.GetInt("port")
		debug := viper.GetBool("debug")
		log := viper.GetString("log")

		fmt.Printf("启动服务 → http://%s:%d\n", host, port)
		fmt.Printf("  debug : %v\n", debug)
		fmt.Printf("  log   : %s\n", log)
		return nil
	},
}

func init() {
	// serve 子命令独有的 flag（不影响其他子命令）
	serveCmd.Flags().IntP("workers", "w", 4, "工作线程数")
	_ = viper.BindPFlag("workers", serveCmd.Flags().Lookup("workers"))
}
