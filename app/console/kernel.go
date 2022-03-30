package console

import (
	"time"
	"webman/app/console/command/demo"
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
func addAppCommands(rootCmd *cobra.Command) {
	// rootCmd.AddCronCommand("* * * * * *", demo.FooCmd)
	rootCmd.AddDistributedCronCommand("test", "*/5 * * * * *", demo.FooCmd, 2*time.Second)
}
