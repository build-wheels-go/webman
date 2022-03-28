package demo

import (
	"log"
	"webman/framework/cobra"
)

var FooCmd = &cobra.Command{
	Use:   "foo",
	Short: "示例command",
	RunE: func(cmd *cobra.Command, args []string) error {
		log.Println("execute foo command")
		return nil
	},
}
