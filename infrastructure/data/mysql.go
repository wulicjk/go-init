package data

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"readLater-backend/infrastructure/config"
)

const (
	MYSQL_CONNECT_TIMEOUT = 5000 //ms
	MYSQL_READ_TIMEOUT    = 5000 //ms
	MYSQL_WRITE_TIMEOUT   = 5000 //ms
)

var MysqlDB *gorm.DB

func init() {
	cfg := config.Cfg.MysqlConf
	mysqlRes := NewMysqlClient(cfg.Host, cfg.Port, cfg.User, cfg.Passwd, cfg.DBName, cfg.TablePrefix)
	db, err := mysqlRes.Open()
	if err != nil {
		return
	}
	MysqlDB = db
	return
}

type MysqlClient struct {
	Host         string
	Port         int64
	UserName     string
	Passwd       string
	DBName       string
	ConnTimeout  int
	ReadTimeout  int
	WriteTimeout int
	DSNName      string
	TablePrefix  string
	db           *gorm.DB
}

func NewMysqlClient(host string, port int64, userName string, passwd string, dbName string, tablePrefix string) *MysqlClient {
	mc := &MysqlClient{
		Host:         host,
		Port:         port,
		UserName:     userName,
		Passwd:       passwd,
		DBName:       dbName,
		TablePrefix:  tablePrefix,
		ConnTimeout:  MYSQL_CONNECT_TIMEOUT,
		ReadTimeout:  MYSQL_READ_TIMEOUT,
		WriteTimeout: MYSQL_WRITE_TIMEOUT,
	}
	return mc
}

func (mc *MysqlClient) Open() (db *gorm.DB, err error) {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mc.UserName, mc.Passwd, "tcp", mc.Host, mc.Port, mc.DBName)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   mc.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}
	return db, nil
}
