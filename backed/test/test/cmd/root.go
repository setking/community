package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "myapp",
	Short: "myapp 是一个示例 CLI 工具",
	Long: `myapp 是一个使用 Cobra + pflag + Viper 构建的示例 CLI 工具。

配置优先级（从高到低）：
  1. 命令行 Flag
  2. 环境变量（前缀 MYAPP_）
  3. 配置文件（默认 ./config.yaml）
  4. 默认值`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return bindFlags(cmd)
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		// -v / --version
		if viper.GetBool("version") {
			fmt.Println("myapp version 1.0.0")
			return nil
		}

		// 打印当前生效的配置
		fmt.Println("=== 当前生效配置 ===")
		fmt.Printf("  config  : %s\n", viper.GetString("config"))
		fmt.Printf("  host    : %s\n", viper.GetString("host"))
		fmt.Printf("  port    : %d\n", viper.GetInt("port"))
		fmt.Printf("  debug   : %v\n", viper.GetBool("debug"))
		fmt.Printf("  log     : %s\n", viper.GetString("log"))
		return nil
	},
}

// Execute 是程序入口，由 main.go 调用
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	// ── Persistent flags（全局，所有子命令都继承） ──────────────────────
	pf := rootCmd.PersistentFlags()

	pf.StringP("config", "c", "config.yaml", "配置文件路径")
	pf.BoolP("debug", "d", false, "开启 debug 模式")
	pf.StringP("log", "l", "info", "日志级别 (debug|info|warn|error)")

	// ── Local flags（仅 root 命令） ────────────────────────────────────
	lf := rootCmd.Flags()

	lf.BoolP("version", "v", false, "输出版本信息")
	lf.StringP("host", "H", "127.0.0.1", "服务监听地址")
	lf.IntP("port", "p", 8080, "服务监听端口")

	// ── 注册子命令 ─────────────────────────────────────────────────────
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(versionCmd)
}

// initConfig 在命令执行前由 cobra.OnInitialize 调用，负责加载配置文件
func initConfig() {
	// 从 flag 读取 -c 指定的配置文件（此时 viper 还没绑定 flag，直接读 pflag）
	cfgFile, _ := rootCmd.PersistentFlags().GetString("config")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
	}

	// 环境变量：MYAPP_HOST、MYAPP_PORT、MYAPP_DEBUG …
	viper.SetEnvPrefix("MYAPP")
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Printf("使用配置文件: %s\n\n", viper.ConfigFileUsed())
	}
}

// bindFlags 将当前命令的所有 pflag 绑定到 viper（只绑定用户显式传入的 flag）
// 放在 PersistentPreRunE 而不是 init()，是为了让子命令也能正确绑定自己的 flag
func bindFlags(cmd *cobra.Command) error {
	cmd.Flags().VisitAll(func(f *pflag.Flag) {
		// 将 flag 名中的 "-" 统一转为 "." 作为 viper key（可选）
		viperKey := strings.ReplaceAll(f.Name, "-", ".")

		// 绑定：让 viper 知道这个 flag 的存在
		_ = viper.BindPFlag(viperKey, f)

		// 如果用户没有在命令行显式传入该 flag，但 viper 已有值（来自配置文件/环境变量），
		// 则用 viper 的值覆盖 pflag 默认值，使 flag.Value 与 viper 保持一致
		if !f.Changed && viper.IsSet(viperKey) {
			val := viper.Get(viperKey)
			_ = cmd.Flags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	// 同样处理 PersistentFlags
	cmd.InheritedFlags().VisitAll(func(f *pflag.Flag) {
		viperKey := strings.ReplaceAll(f.Name, "-", ".")
		_ = viper.BindPFlag(viperKey, f)
		if !f.Changed && viper.IsSet(viperKey) {
			val := viper.Get(viperKey)
			_ = cmd.InheritedFlags().Set(f.Name, fmt.Sprintf("%v", val))
		}
	})

	return nil
}
