package cobra

import (
	"log"
	"time"
	"webman/framework/contract"

	"github.com/robfig/cron/v3"
)

func (c *Command) AddDistributedCronCommand(serviceName string, spec string, cmd *Command, holdTime time.Duration) {
	root := c.Root()

	if root.Cron == nil {
		root.Cron = cron.New(cron.WithParser(cron.NewParser(cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow | cron.Descriptor)))
		root.CronSpecs = []CronSpec{}
	}

	root.CronSpecs = append(root.CronSpecs, CronSpec{
		Type:        "distributed-cron",
		Cmd:         cmd,
		Spec:        spec,
		ServiceName: serviceName,
	})

	appService := root.GetContainer().MustMake(contract.AppKey).(contract.App)
	distributedService := root.GetContainer().MustMake(contract.DistributedKey).(contract.Distributed)
	appID := appService.AppID()

	var cronCmd Command

	cronCmd = *cmd
	ctx := root.Context()
	cronCmd.args = []string{}
	cronCmd.SetParentNull()

	_, _ = root.Cron.AddFunc(spec, func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
			}
		}()
		// 选举appID
		selectAppID, err := distributedService.Select(serviceName, appID, holdTime)
		if err != nil {
			return
		}
		if selectAppID != appID {
			return
		}
		if err := cronCmd.ExecuteContext(ctx); err != nil {
			log.Println(err)
		}
	})
}
