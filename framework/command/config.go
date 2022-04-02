package command

import (
	"errors"
	"fmt"

	"webman/framework/cobra"
	"webman/framework/contract"

	"github.com/kr/pretty"
)

func initConfigCmd() *cobra.Command {
	configCmd.AddCommand(configGetCmd)
	return configCmd
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

var configGetCmd = &cobra.Command{
	Use:     "get",
	Short:   "获取配置命令",
	Example: "webman config get \"path\"",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		if len(args) != 1 {
			return errors.New("参数错误")
		}
		path := args[0]
		conf := configService.Get(path)
		if conf == nil {
			return errors.New("配置路径 " + path + " 不存在")
		}
		fmt.Printf("%# v \n", pretty.Formatter(conf))
		return nil
	},
}
