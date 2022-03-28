package command

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
	"webman/framework/cobra"
	"webman/framework/contract"
)

func initAppCmd() *cobra.Command {
	appCmd.AddCommand(appStartCmd)
	return appCmd
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "应用控制命令",
	Long:  "应用控制命令，包含应用启动、关闭、重启、查询等功能",
	RunE: func(cmd *cobra.Command, args []string) error {
		_ = cmd.Help()
		return nil
	},
}

var appStartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动web服务",
	RunE: func(cmd *cobra.Command, args []string) error {
		// 从command中获取服务容器
		container := cmd.GetContainer()
		// 从服务容器中获取kernel服务实例
		kernelService := container.MustMake(contract.KernelKey).(contract.Kernel)
		// 从服务实例中获取http引擎
		core := kernelService.HttpEngine()
		// 创建http服务
		server := http.Server{
			Handler: core,
			Addr:    ":8888",
		}

		go func() {
			_ = server.ListenAndServe()
		}()

		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
		<-quit

		timeoutCtx, clean := context.WithTimeout(context.Background(), 5*time.Second)
		defer clean()

		if err := server.Shutdown(timeoutCtx); err != nil {
			log.Fatal("Server Shutdown:", err)
		}
		return nil
	},
}
