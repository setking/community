package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configCmd 是 `myapp config` 父命令
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "管理配置文件",
	Long:  `读取、写入、列出配置文件中的键值对。`,
}

// config set key=value [key=value ...]
var configSetCmd = &cobra.Command{
	Use:   "set KEY=VALUE [KEY=VALUE ...]",
	Short: "写入一个或多个配置项到 yaml 文件",
	Example: `  myapp config set host=0.0.0.0
  myapp config set port=9090 debug=true log=warn
  myapp config set -c ./prod.yaml host=10.0.0.1 port=443`,
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cfgFile, _ := cmd.Flags().GetString("config")
		if cfgFile == "" {
			cfgFile = "config.yaml"
		}

		// 解析所有 key=value 参数
		for _, arg := range args {
			parts := strings.SplitN(arg, "=", 2)
			if len(parts) != 2 {
				return fmt.Errorf("格式错误: %q，应为 KEY=VALUE", arg)
			}
			key, raw := parts[0], parts[1]

			// 自动推断类型：bool → int → float → string
			viper.Set(key, parseValue(raw))
			fmt.Printf("  set  %s = %s\n", key, raw)
		}

		// 写回 yaml 文件
		viper.SetConfigFile(cfgFile)

		// 先尝试读取已有内容（合并，不覆盖其他 key）
		_ = viper.ReadInConfig()

		if err := viper.WriteConfigAs(cfgFile); err != nil {
			return fmt.Errorf("写入配置文件失败: %w", err)
		}
		fmt.Printf("\n已写入 → %s\n", cfgFile)
		return nil
	},
}

// config get key [key ...]
var configGetCmd = &cobra.Command{
	Use:     "get KEY [KEY ...]",
	Short:   "读取一个或多个配置项的值",
	Example: `  myapp config get host port debug`,
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		for _, key := range args {
			if !viper.IsSet(key) {
				fmt.Printf("  %-12s (未设置)\n", key)
			} else {
				fmt.Printf("  %-12s = %v\n", key, viper.Get(key))
			}
		}
		return nil
	},
}

// config list
var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出配置文件中所有键值",
	RunE: func(cmd *cobra.Command, args []string) error {
		all := viper.AllSettings()
		if len(all) == 0 {
			fmt.Println("（配置为空）")
			return nil
		}
		fmt.Println("=== 当前全部配置 ===")
		printMap("", all)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configListCmd)
	rootCmd.AddCommand(configCmd)
}

// ── 工具函数 ──────────────────────────────────────────────────────────────

// parseValue 按 bool → int → float64 → string 顺序推断类型
func parseValue(s string) interface{} {
	if b, err := strconv.ParseBool(s); err == nil {
		return b
	}
	if i, err := strconv.ParseInt(s, 10, 64); err == nil {
		return i
	}
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		return f
	}
	return s
}

// printMap 递归打印嵌套 map（支持 viper 中的嵌套配置）
func printMap(prefix string, m map[string]interface{}) {
	for k, v := range m {
		fullKey := k
		if prefix != "" {
			fullKey = prefix + "." + k
		}
		if nested, ok := v.(map[string]interface{}); ok {
			printMap(fullKey, nested)
		} else {
			fmt.Printf("  %-20s = %v\n", fullKey, v)
		}
	}
}
