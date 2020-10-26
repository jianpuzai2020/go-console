package config

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	Cfg *ini.File
)

type Config struct {
	TcpPort           string
	DbTcp             string
	LogTcp            string
	LogFlag           bool
	DbUrl             string
	DbShow            bool
	DbMaxOpenConns    int
	DbMaxIdleConns    int
	DbConnMaxLifetime time.Duration
	DbWriteUrl        string
	DbWriteShow       bool
	Db2Name 		  string
	Db2Url 			  string
	Db2Show 		  bool
	Db2WriteUrl 	  string
	Db2WriteShow 	  bool
	Db2MaxOpenConns   int
	Db2MaxIdleConns   int
	Db2ConnMaxLifetime time.Duration
	ProcessWorkerSize  int
	ProcessConnectSize int
	LogLevel 		   string
	LogFilename 	   string
	LogMaxSize 		   int
	LogMaxBackups      int
	LogMaxAge 		   int
}

func init() {

	var err error
	Cfg, err = ini.Load("env.ini")
	if err != nil {
		log.Fatal("Fail to parse 'env.ini': %v", err)
	}
}

func (c *Config) LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': #{err}")
	}

	c.TcpPort = sec.Key("TCP_PORT").MustString("")
	c.LogTcp = sec.Key("LOG_TCP").MustString("")
	c.LogFlag = sec.Key("LOG_FLAG").MustBool(false)
	c.DbTcp = sec.Key("DB_TCP").MustString("")
	c.ProcessWorkerSize = sec.Key("PROCESS_WORKER_SIZE").MustInt(1)
	c.ProcessConnectSize = sec.Key("PROCESS_CONNECT_SIZE").MustInt(100)
}

func (c *Config) LoadDb() {
	db, err := Cfg.GetSection("db")
	if err != nil {
		log.Fatalf("Fail to get section 'db': #{err}")
	}

	c.DbUrl = db.Key("URL").MustString("")
	c.DbShow = db.Key("SHOW_SQL").MustBool(false)
	c.DbMaxOpenConns = db.Key("MAX_OPEN_CONNS").MustInt(100)
	c.DbMaxIdleConns = db.Key("MAX_IDLE_CONNS").MustInt(20)
	c.DbConnMaxLifetime = time.Duration(db.Key("CONN_MAX_LIFETIME").MustInt(30)) * time.Second

	dbw, err := Cfg.GetSection("db_write")
	if err != nil {
		log.Fatalf("Fail to get section 'db_write': %v", err)
	}
	c.DbWriteUrl = dbw.Key("URL").MustString("")
	c.DbWriteShow = dbw.Key("SHOW_SQL").MustBool(false)
}

func  (c *Config) LoadDb2() {
	db, err := Cfg.GetSection("db2")
	if err != nil {
		log.Fatalf("Fail to get section 'db': %v", err)
	}

	c.Db2Name = db.Key("DB_NAME").MustString("")
	c.Db2Url = db.Key("URL").MustString("")
	c.Db2Show = db.Key("SHOW_SQL").MustBool(false)
	c.Db2MaxOpenConns = db.Key("MAX_OPEN_CONNS").MustInt(100)
	c.Db2MaxIdleConns = db.Key("MAX_IDLE_CONNS").MustInt(20)
	c.Db2ConnMaxLifetime = time.Duration(db.Key("CONN_MAX_LIFETIME").MustInt(30)) * time.Second

	dbw, err := Cfg.GetSection("db2_write")
	if err != nil {
		log.Fatalf("Fail to get section 'db_write': %v", err)
	}

	c.Db2WriteUrl = dbw.Key("URL").MustString("")
	c.Db2WriteShow = dbw.Key("SHOW_SQL").MustBool(false)
}

func  (c *Config) LoadLog() {
	sec, err := Cfg.GetSection("log")
	if err != nil {
		log.Fatalf("Fail to get section 'log': %v", err)
	}

	c.LogFilename = sec.Key("FILE_NAME").MustString("/tmp/go.log")
	c.LogLevel = sec.Key("LOG_LEVEL").MustString("info")
	c.LogMaxSize = sec.Key("MAX_SIZE").MustInt(1024)
	c.LogMaxBackups = sec.Key("MAX_BACKUPS").MustInt(5)
	c.LogMaxAge = sec.Key("MAX_AGE").MustInt(168)
}