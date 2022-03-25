package console

import (
	"webman/framework"
	"webman/framework/cobra"
	"webman/framework/command"
)

func RunCommand(container framework.Container) error {
	// root command
	rootCmd := command.InitRootCommand(container)

	addAppCommands(rootCmd)
	// 执行命令
	return rootCmd.Execute()
}

// 添加业务命令
func addAppCommands(root *cobra.Command) {

}
