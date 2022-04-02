package command

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"syscall"
	"time"
	"webman/framework/cobra"
	"webman/framework/contract"
	"webman/framework/util"

	"github.com/sevlyar/go-daemon"
)

var cronDaemon = false

func initCronCmd() *cobra.Command {
	cronStartCmd.Flags().BoolVarP(&cronDaemon, "daemon", "d", false, "start serve daemon")
	cronCmd.AddCommand(cronListCmd)
	cronCmd.AddCommand(cronStartCmd)
	cronCmd.AddCommand(cronStopCmd)
	cronCmd.AddCommand(cronRestartCmd)
	cronCmd.AddCommand(cronStateCmd)
	return cronCmd
}

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "定时任务相关命令",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}
		return nil
	},
}

var cronListCmd = &cobra.Command{
	Use:   "list",
	Short: "列出所有的定时任务",
	RunE: func(cmd *cobra.Command, args []string) error {
		cronSpecs := cmd.Root().CronSpecs
		ps := [][]string{}
		for _, spec := range cronSpecs {
			line := []string{spec.Type, spec.Spec, spec.Cmd.Use, spec.Cmd.Short, spec.ServiceName}
			ps = append(ps, line)
		}
		util.PrettyPrint(ps)
		return nil
	},
}

var cronStartCmd = &cobra.Command{
	Use:   "start",
	Short: "启动cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		// 获取容器中的app服务
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")
		serverLogFile := filepath.Join(appService.LogFolder(), "cron.log")
		currentFolder := appService.BaseFolder()
		// daemon 模式
		if cronDaemon {
			ctx := &daemon.Context{
				PidFileName: serverPidFile,
				PidFilePerm: 0664,
				LogFileName: serverLogFile,
				LogFilePerm: 0640,
				// 设置工作路径
				WorkDir: currentFolder,
				Umask:   027,
				// 子进程参数
				Args: []string{"", "cron", "start", "--daemon=true"},
			}
			// 启动子进程，d不为空表示父进程，为空表示子进程
			d, err := ctx.Reborn()
			if err != nil {
				return err
			}
			if d != nil {
				//父进程打印信息，不做任何操作
				fmt.Println("cron serve started,pid:", d.Pid)
				fmt.Println("log file:", serverLogFile)
				return nil
			}

			defer func(ctx *daemon.Context) {
				_ = ctx.Release()
			}(ctx)

			fmt.Println("daemon started")
			cmd.Root().Cron.Run()
			return nil
		}

		// not daemon mode
		fmt.Println("start cron job")
		pidContent := strconv.Itoa(os.Getpid())
		fmt.Println("[PID]", pidContent)
		if err := ioutil.WriteFile(serverPidFile, []byte(pidContent), 0664); err != nil {
			return err
		}
		cmd.Root().Cron.Run()
		return nil
	},
}

var cronStopCmd = &cobra.Command{
	Use:   "stop",
	Short: "停止cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		pidContent, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}
		if len(pidContent) > 0 {
			pid, err := strconv.Atoi(string(pidContent))
			if err != nil {
				return err
			}
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				return err
			}
			if err := ioutil.WriteFile(serverPidFile, []byte{}, 0664); err != nil {
				return err
			}
			fmt.Println("stop pid:", pid)
		}
		return nil
	},
}

var cronRestartCmd = &cobra.Command{
	Use:   "restart",
	Short: "重启cron常驻进程",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		pidContent, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}
		if len(pidContent) > 0 {
			pid, err := strconv.Atoi(string(pidContent))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
					return err
				}
				for i := 0; i < 10; i++ {
					if !util.CheckProcessExist(pid) {
						break
					}
					time.Sleep(10 * time.Second)
				}
				fmt.Println("kill process:", pidContent)
			}
		}
		cronDaemon = true
		return cronStartCmd.RunE(cmd, args)
	},
}

var cronStateCmd = &cobra.Command{
	Use:   "state",
	Short: "cron常驻进程状态",
	RunE: func(cmd *cobra.Command, args []string) error {
		container := cmd.GetContainer()
		appService := container.MustMake(contract.AppKey).(contract.App)

		serverPidFile := filepath.Join(appService.RuntimeFolder(), "cron.pid")

		pidContent, err := ioutil.ReadFile(serverPidFile)
		if err != nil {
			return err
		}
		if len(pidContent) > 0 {
			pid, err := strconv.Atoi(string(pidContent))
			if err != nil {
				return err
			}
			if util.CheckProcessExist(pid) {
				fmt.Println("cron server started,pid:", pidContent)
				return nil
			}
		}
		fmt.Println("no cron server start")
		return nil
	},
}
