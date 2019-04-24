package conf

import (
	"fmt"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"os"
)

var ConfigTOML iris.Configuration

var Config GlobalConf

type GlobalConf struct {
	Port int32
	Ip   string
	Addr string
	Env  string //运行环境
}

func init() {
	var env string
	switch os.Getenv("GO_RUN_ENV") {
	case "prod":
		env = "prod"
	case "stage":
		env = "stage"
	default:
		env = "dev"
	}

	ConfigTOML = iris.TOML("./conf/iris.tml")
	other := ConfigTOML.Other
	Config = GlobalConf{}
	for k, v := range other {
		if k == "Port" {
			Config.Port = v.(int32)
			continue
		}
		if k == "Ip" {
			Config.Ip = v.(string)
			continue
		}
	}
	Config.Addr = fmt.Sprintf("%s:%d", Config.Ip, Config.Port)
	Config.Env = env
	golog.Infof("项目配置初始化成功:=(%v)", Config)
}
