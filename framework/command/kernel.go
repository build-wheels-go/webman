package command

import (
	"webman/framework"
	"webman/framework/cobra"
)

func InitRootCommand(container framework.Container) *cobra.Command {
	// 为根命令设置服务容器
	rootCmd.SetContainer(container)

	addKernelCommands(rootCmd)
	return rootCmd
}

var rootCmd = &cobra.Command{
	// 定义根命令的关键字
	Use: "webman",
	// 简短介绍
	Short: "webman 命令",
	// 详细介绍
	Long: "webman 框架的命令行工具，使用这个命令行工具能方便的执行框架自带命令，也能方便编写业务命令",
	// 根命令执行函数
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.InitDefaultHelpFlag()
		return cmd.Help()
	},
	CompletionOptions: cobra.CompletionOptions{DisableDefaultCmd: true},
}

// 添加框架核心命令
func addKernelCommands(root *cobra.Command) {
	root.AddCommand(initAppCommand())
}
