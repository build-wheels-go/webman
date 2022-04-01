package config

import (
	"bytes"
	"fmt"
	"log"

	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"webman/framework"
	"webman/framework/contract"

	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"gopkg.in/yaml.v3"
)

var _ contract.Config = (*WmConfig)(nil)

type WmConfig struct {
	// 服务容器
	c framework.Container
	// 配置文件夹
	folder string
	// 路径分隔符，默认.
	separator string
	// 配置文件读写锁
	lock sync.RWMutex
	// 所有环境变量
	envs map[string]string
	// 配置文件结构
	configs map[string]interface{}
	// 配置文件的原始信息
	raws map[string][]byte
}

func NewWmConfig(params ...interface{}) (interface{}, error) {
	container := params[0].(framework.Container)
	configFolder := params[1].(string)
	envs := params[2].(map[string]string)
	// 判断文件是否存在
	if _, err := os.Stat(configFolder); os.IsNotExist(err) {
		return nil, errors.New("folder " + configFolder + " not exist: " + err.Error())
	}

	wmConfig := &WmConfig{
		c:         container,
		folder:    configFolder,
		separator: ".",
		envs:      envs,
		configs:   map[string]interface{}{},
		raws:      map[string][]byte{},
		lock:      sync.RWMutex{},
	}

	files, err := ioutil.ReadDir(configFolder)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	for _, file := range files {
		fileName := file.Name()
		if err := wmConfig.loadConfigFile(configFolder, fileName); err != nil {
			log.Println(err)
			continue
		}
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}
	err = watcher.Add(configFolder)
	if err != nil {
		return nil, err
	}

	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()

		select {
		case ev := <-watcher.Events:
			path, _ := filepath.Abs(ev.Name)
			index := strings.Index(path, string(os.PathSeparator))
			folder := path[:index]
			fileName := path[index:]
			if ev.Op&fsnotify.Create == fsnotify.Create {
				log.Println("创建文件：", ev.Name)
				_ = wmConfig.loadConfigFile(folder, fileName)
			}
			if ev.Op&fsnotify.Write == fsnotify.Write {
				log.Println("写入文件：", ev.Name)
				_ = wmConfig.loadConfigFile(folder, fileName)
			}
			if ev.Op&fsnotify.Remove == fsnotify.Remove {
				log.Println("移除文件：", ev.Name)
				_ = wmConfig.removeConfigFile(fileName)
			}
		case err := <-watcher.Errors:
			log.Println("error:", err)
			return
		}
	}()

	return wmConfig, nil
}

func (conf *WmConfig) loadConfigFile(folder string, file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()

	s := strings.Split(file, ".")
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		fileName := s[0]
		// 读取文件内容
		bf, err := ioutil.ReadFile(filepath.Join(folder, file))
		if err != nil {
			return err
		}
		// 环境变量替换
		bf = replace(bf, conf.envs)
		// 解析yaml文件
		c := make(map[string]interface{})
		if err := yaml.Unmarshal(bf, &c); err != nil {
			return err
		}
		conf.configs[fileName] = c
		conf.raws[fileName] = bf

		if fileName == "app" && conf.c.IsBind(contract.AppKey) {
			if p, ok := c["path"]; ok {
				appService := conf.c.MustMake(contract.AppKey).(contract.App)
				appService.LoadAppConfig(cast.ToStringMapString(p))
			}
		}
	}
	return nil
}

// 删除配置文件
func (conf *WmConfig) removeConfigFile(file string) error {
	conf.lock.Lock()
	defer conf.lock.Unlock()

	s := strings.Split(file, ".")
	if len(s) == 2 && (s[1] == "yaml" || s[1] == "yml") {
		name := s[0]
		delete(conf.configs, name)
		delete(conf.raws, name)
	}
	return nil
}

// 占位替换
func replace(content []byte, envs map[string]string) []byte {
	if envs == nil {
		return content
	}
	mapping := func(name []byte, envs map[string]string) []byte {
		args := bytes.SplitN(bytes.TrimSpace(name), []byte{':'}, 2)
		if v, ok := envs[string(args[0])]; ok {
			return []byte(v)
		} else if len(args) > 1 { // default value
			return args[1]
		}
		return nil
	}

	r := regexp.MustCompile(`\${(.*?)}`)
	re := r.FindAllSubmatch(content, -1)
	for _, i := range re {
		if len(i) == 2 {
			content = bytes.ReplaceAll(content, i[0], mapping(i[1], envs))
		}
	}
	return content
}

func searchMap(source map[string]interface{}, path []string) interface{} {
	if len(path) == 0 {
		return source
	}
	// 判断是否有下一个路径
	next, ok := source[path[0]]
	if ok {
		if len(path) == 1 {
			return next
		}
		switch next.(type) {
		case map[interface{}]interface{}:
			return searchMap(cast.ToStringMap(next), path[1:])
		case map[string]interface{}:
			return searchMap(next.(map[string]interface{}), path[1:])
		default:
			return nil
		}
	}
	return nil
}

// 获取配置项
func (conf *WmConfig) find(key string) interface{} {
	conf.lock.RLock()
	defer conf.lock.RUnlock()

	return searchMap(conf.configs, strings.Split(key, conf.separator))
}

func (conf *WmConfig) IsExist(key string) bool {
	return conf.find(key) != nil
}

func (conf *WmConfig) Get(key string) interface{} {
	return conf.find(key)
}

func (conf *WmConfig) GetBool(key string) bool {
	return cast.ToBool(conf.find(key))
}

func (conf *WmConfig) GetString(key string) string {
	return cast.ToString(conf.find(key))
}

func (conf *WmConfig) GetInt(key string) int {
	return cast.ToInt(conf.find(key))
}
