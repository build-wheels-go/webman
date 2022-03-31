package env

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"
	"webman/framework/contract"
)

var _ contract.Env = (*WmEnvService)(nil)

type WmEnvService struct {
	// .env文件所在目录
	folder string
	// 保存所有的环境
	envs map[string]string
}

func NewWmEnvService(params ...interface{}) (interface{}, error) {
	if len(params) != 1 {
		return nil, errors.New("NewWmEnvService param error")
	}

	envFolder := params[0].(string)
	wmEnv := &WmEnvService{
		folder: envFolder,
		envs:   map[string]string{"APP_ENV": contract.EnvDevelopment},
	}

	file := filepath.Join(envFolder, ".env")

	fi, err := os.Open(file)
	if err == nil {
		defer fi.Close()

		br := bufio.NewReader(fi)
		for {
			line, _, err := br.ReadLine()
			if err == io.EOF {
				break
			}
			s := bytes.SplitN(line, []byte{'='}, 2)
			if len(s) < 2 {
				continue
			}

			key := string(s[0])
			val := string(s[1])
			wmEnv.envs[key] = val
		}
	}

	//获取当前的环境变量
	for _, e := range os.Environ() {
		segments := strings.SplitN(e, "=", 2)
		if len(segments) < 2 {
			continue
		}
		wmEnv.envs[segments[0]] = segments[1]
	}
	return wmEnv, nil
}

func (env *WmEnvService) AppEnv() string {
	appEnv := env.Get("APP_ENV")
	if appEnv == "" {
		appEnv = contract.EnvDevelopment
	}

	return appEnv
}

func (env *WmEnvService) IsExist(key string) bool {
	_, ok := env.envs[key]
	return ok
}

func (env *WmEnvService) Get(key string) string {
	if val, ok := env.envs[key]; ok {
		return val
	}
	return ""
}

func (env *WmEnvService) All() map[string]string {
	return env.envs
}
