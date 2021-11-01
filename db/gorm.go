package db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	MainDb         *gorm.DB
	DatabaseConfig DbConfig
)

type DbConfig struct {
	Dbtype   string
	Host     string
	Port     int
	Name     string
	Username string
	Password string
}

func InitDatabaseConfig(cfg *viper.Viper) {
	DatabaseConfig.Port = cfg.GetInt("port")
	DatabaseConfig.Dbtype = cfg.GetString("dbtype")
	DatabaseConfig.Host = cfg.GetString("host")
	DatabaseConfig.Name = cfg.GetString("name")
	DatabaseConfig.Username = cfg.GetString("username")
	DatabaseConfig.Password = cfg.GetString("password")
}

func InitDBConnect() *gorm.DB {
	if DatabaseConfig.Dbtype != "mysql" {
		log.Println("db type unknow")
	}
	var err error
	conn := GetConnect()
	var con *gorm.DB
	if DatabaseConfig.Dbtype == "mysql" {
		con, err = gorm.Open(mysql.Open(conn), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			NamingStrategy:                           schema.NamingStrategy{SingularTable: true},
		})
	} else {
		panic("db type unknow")
	}
	if err != nil {
		log.Fatalf("%s connect error %v", DatabaseConfig.Dbtype, err)
	} else {
		log.Printf("%s connect success!", DatabaseConfig.Dbtype)
	}
	if con.Error != nil {
		log.Fatalf("database error %v", con.Error)
	}
	// con.LogMode(true)
	return con
}

func GetConnect() string {
	return fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=1000ms",
		DatabaseConfig.Username,
		DatabaseConfig.Password,
		DatabaseConfig.Host,
		DatabaseConfig.Port,
		DatabaseConfig.Name,
	)
}

func RecycleDatabaseConfig() {
	DatabaseConfig = DbConfig{}
}

func CloseDB() {
	if dbSQL, ok := MainDb.DB(); ok != nil {
		defer dbSQL.Close()
	}
}

type GormDB struct {
	*gorm.DB
	gdbDone bool
}

func BeginTx() *GormDB {
	txn := MainDb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	return &GormDB{txn, false}
}

func (c *GormDB) DbCommit() {
	if c.gdbDone {
		return
	}
	tx := c.Commit()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}

func (c *GormDB) DbRollback() {
	if c.gdbDone {
		return
	}
	tx := c.Rollback()
	c.gdbDone = true
	if err := tx.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
}
