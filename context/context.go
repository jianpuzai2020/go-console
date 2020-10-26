package context

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"go.uber.org/zap"
	"main/config"
	_ "github.com/go-sql-driver/mysql"
)

var err error

type Context struct {
	Conf   			 *config.Config
	DB     			 *xorm.Engine
	DBWrite 		 *xorm.Engine
	DB2     		 *xorm.Engine
	DB2Write 		 *xorm.Engine
	Logger *zap.Logger
}

type ConfigDB struct {
	URL     string
	ShowSQL bool
}

func New() *Context {
	r := new(Context)
	r.initConf()

	r.DB = newDB(r, ConfigDB{URL: r.Conf.DbUrl,ShowSQL:r.Conf.DbShow})
	r.DBWrite = newDB(r, ConfigDB{URL:r.Conf.DbWriteUrl,ShowSQL:r.Conf.DbWriteShow})
	r.DB2 = newDB2(r, ConfigDB{URL: r.Conf.Db2Url, ShowSQL:r.Conf.Db2Show})
	r.DB2Write = newDB2(r, ConfigDB{URL: r.Conf.Db2WriteUrl, ShowSQL:r.Conf.Db2WriteShow})

	return r
}

func (c *Context) initConf() {
	cfg := new(config.Config)
	cfg.LoadServer()
	cfg.LoadDb()
	cfg.LoadDb2()
	cfg.LoadLog()
	c.Conf = cfg
}

func newDB(c *Context, d ConfigDB) *xorm.Engine {
	db, err := xorm.NewEngine("mysql", d.URL)
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	db.SetMaxOpenConns(c.Conf.DbMaxOpenConns)	//最大打开连接数
	db.SetMaxIdleConns(c.Conf.DbMaxIdleConns)	//连接池的空闲数
	db.SetConnMaxLifetime(c.Conf.DbConnMaxLifetime)

	if err = db.Ping(); err != nil {
		c.Logger.Fatal(err.Error())
	}

	db.ShowSQL(d.ShowSQL)
	return db
}

func newDB2(c *Context,d ConfigDB) *xorm.Engine  {
	db, err := xorm.NewEngine("mysql", d.URL)
	if err != nil {
		c.Logger.Fatal(err.Error())
	}

	db.SetMaxOpenConns(c.Conf.Db2MaxOpenConns) //最大打开连接数
	db.SetMaxIdleConns(c.Conf.Db2MaxIdleConns) //连接池的空闲数
	db.SetConnMaxLifetime(c.Conf.Db2ConnMaxLifetime)
	if err = db.Ping(); err != nil {
		c.Logger.Fatal(err.Error())
	}

	db.ShowSQL(d.ShowSQL)
	return db
}

func (c *Context) Recover()  {
	if err := recover(); err != nil {
		fmt.Printf("%s\n", err)
		c.Logger.Error(fmt.Sprintf("%s", err))
	}
}

func (c *Context) Close() {
	c.DB.Close()
	c.DBWrite.Close()
}
