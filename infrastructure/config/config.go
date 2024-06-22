package config

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"strings"
)

var Cfg Config

func init() {
	// 定义命令行参数
	envFlag := flag.String("c", "dev", "environment (dev or prod)")
	flag.Parse()

	// 根据环境变量选择配置文件路径
	var confPath string
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting currentDir path:", err)
		return
	}
	rootPath := getProjectRootPath(currentDir)
	confDirPath := filepath.Join(rootPath, "config")

	fmt.Println("Current working directory:", confDirPath)
	switch *envFlag {
	case "dev":
		confPath = filepath.Join(confDirPath, "dev-config.toml")
	case "prod":
		confPath = filepath.Join(confDirPath, "prod-config.toml")
	default:
		confPath = filepath.Join(confDirPath, "dev-config.toml")
	}

	// 加载配置文件
	if _, err := toml.DecodeFile(confPath, &Cfg); err != nil {
		log.Fatalf("InitConfig err: %v", err)
		return
	}
	log.Infof("Loaded config from: %s", confPath)
}

func getProjectRootPath(path string) string {
	// 找到 "readLater-backend" 字符串的结束位置
	index := strings.Index(path, "readLater-backend")
	if index == -1 {
		// 如果没有找到，直接返回原始路径
		return path
	}

	// 从头部开始截取到 "readLater-backend" 之后的部分
	endIndex := index + len("readLater-backend")
	return path[:endIndex]
}

type Config struct {
	HTTPServerConf HttpServerConfig `toml:"http_server"`
	MysqlConf      MySqlConfig      `toml:"mysql_db"`
	RedisConf      RedisConfig      `toml:"redis_db"`
}

type HttpServerConfig struct {
	Enable   bool `toml:"enable"`
	GPort    int  `toml:"gport"`
	MPort    int  `toml:"mport"`
	WTimeout int  `toml:"wTimeout"`
	RTimeout int  `toml:"wTimeout"`
}

type RedisConfig struct {
	ReadAddr     string `toml:"read_addr"`
	WriteAddr    string `toml:"write_addr"`
	DB           int    `toml:"db"`
	Passwd       string `toml:"passwd"`
	MaxIdel      int    `toml:"max_idel"`
	MaxActive    int    `toml:"max_active"`
	IdelTimeout  int    `toml:"idle_timeout"`
	ConnTimeout  int    `toml:"conn_timeout"`
	WriteTimeout int    `toml:"write_timeout"`
	ReadTimeout  int    `toml:"read_timeout"`
}

type DyConfig struct {
	ClientKey    string `toml:"client_key"`
	ClientSecret string `toml:"client_secret"`
	GrantType    string `toml:"grant_type"`
}

type MQConfConfig struct {
	Topics        []string `toml:"topics"`
	Brokers       []string `toml:"brokers"`
	ConsumerGroup string   `toml:"group"`
}

type MySqlConfig struct {
	DBName      string `toml:"db_name"`
	Host        string `toml:"host"`
	Port        int64  `toml:"port"`
	User        string `toml:"user"`
	Passwd      string `toml:"passwd"`
	TablePrefix string `toml:"table_prefix"`
}
