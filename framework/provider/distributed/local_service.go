package distributed

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"time"
	"webman/framework"
	"webman/framework/contract"
)

var _ contract.Distributed = (*LocalDistributedService)(nil)

type LocalDistributedService struct {
	container framework.Container
}

func NewLocalDistributedService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("param error")
	}

	container := params[0].(framework.Container)
	return &LocalDistributedService{container: container}, nil
}

func (s *LocalDistributedService) Select(serviceName, appID string, holdTime time.Duration) (selectAppID string, err error) {
	appService := s.container.MustMake(contract.AppKey).(contract.App)
	lockFile := filepath.Join(appService.RuntimeFolder(), "distribute_"+serviceName)
	lock, err := os.OpenFile(lockFile, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return "", err
	}
	//获取文件锁,如果已经被占用则返回当前的appId
	if err = syscall.Flock(int(lock.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		selectAppIDByt, err := ioutil.ReadAll(lock)
		if err != nil {
			return "", err
		}
		return string(selectAppIDByt), nil
	}

	go func() {
		defer func() {
			// 释放文件锁
			syscall.Flock(int(lock.Fd()), syscall.LOCK_UN)
			// 释放文件
			lock.Close()
			// 删除文件
			os.Remove(lockFile)
		}()
		// 创建有效时间计时器
		timer := time.NewTimer(holdTime)
		// 阻塞等待计时器结束
		<-timer.C
	}()

	if _, err = lock.WriteString(appID); err != nil {
		return "", err
	}

	return appID, nil
}
