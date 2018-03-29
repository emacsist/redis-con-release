package main

import (
	"os"
	"github.com/Sirupsen/logrus"
	"io/ioutil"
	"path"
	"encoding/json"
	"./config"
	"time"
	"github.com/gomodule/redigo/redis"
	"strconv"
	"strings"
	"bufio"
)

func main() {
	var appConfig config.AppConfig
	readJSON("app.json", &appConfig)
	for {
		for _, e := range appConfig.Redis {
			go releaseIdelConnection(e, appConfig.Idle)
		}
		// 一天处理一次就好.
		time.Sleep(time.Hour * 24)
	}
	forever := make(chan bool)
	<-forever
}

func releaseIdelConnection(e config.RedisEle, maxIdle int) {
	con, err := redis.Dial("tcp", e.Host+":"+strconv.Itoa(e.Port))
	if err != nil {
		logrus.Errorf("connect error %v:%v", e.Host, e.Port)
		return
	}
	defer con.Close()
	if len(e.Passwd) > 0 {
		if _, err := con.Do("auth", e.Passwd); err != nil {
			return
		}
	}
	reply, err := redis.String(con.Do("client", "list"))
	if err != nil {
		logrus.Warnf("get client list error %v", err.Error())
		return
	}
	scanner := bufio.NewScanner(strings.NewReader(reply))
	for scanner.Scan() {
		// 每一行的数据类似如下
		// id=2 addr=127.0.0.1:52432 fd=7 name= age=1210 idle=1207 flags=N db=0 sub=0 psub=0 multi=-1 qbuf=0 qbuf-free=0 obl=0 oll=0 omem=0 events=r cmd=info
		client := scanner.Text()

		// 以下两个单位为 秒
		idle := 0
		addr := ""

		fields := strings.Fields(client)

		for _, f := range fields {
			kv := strings.Split(f, "=")
			k := kv[0]
			v := kv[1]

			k = strings.ToLower(k)

			if k == "addr" {
				addr = v
			} else if k == "idle" {
				idle, _ = strconv.Atoi(v)
			}
		}

		// 大于指定的空闲时间, 则 kill 掉
		if idle > maxIdle {
			reply, _ := redis.String(con.Do("client", "kill", addr))
			logrus.Warnf("kill %v, result %v", addr, reply)
		}
	}
}

// readJSON: 从 filePath 里读取数据, 并转换为 jsonObject 对象
func readJSON(filePath string, jsonObject interface{}) {
	appPath, e := os.Executable()
	if e != nil {
		logrus.Errorf("File error: %v\n", e)
		os.Exit(1)
	}

	file, e := ioutil.ReadFile(path.Dir(appPath) + "/" + filePath)
	if e != nil {
		logrus.Errorf("File error: %v\n", e)
		os.Exit(1)
	}
	e = json.Unmarshal(file, jsonObject)
	if e != nil {
		logrus.Errorf("invalid json data error: %v\n", e)
		os.Exit(1)
	}
}
