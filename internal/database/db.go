package database

import (
	"context"
	"fmt"
	"github.com/jinglanghe/go-start/internal/config"
	"github.com/jinglanghe/go-start/internal/dao"
	"github.com/jinglanghe/go-start/utils/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var (
	db *gorm.DB
)

func GetDb(ctx context.Context) *gorm.DB {
	if ctx == nil {
		ctx = context.Background()
	}
	return db.WithContext(ctx)
}

func InitDb(config *config.AppConfig) (*gorm.DB, func()) {
	db = NewGormDB(config)

	clearFunc := func() {
		if db != nil {
			sqlDB, err := db.DB()
			if err != nil {
				log.Error(err).Msg("close db connection failed")
				return
			}
			sqlDB.Close()
			log.Info().Msg("closed db connection")
		}
	}

	if err := dao.AutoMigrate(db); err != nil {
		log.Fatal().Err(err).Send()
		return nil, clearFunc
	}

	go PingDb(db, config)

	log.Info().Msgf("Successfully connected to db, type: %v", config.Database.ServerType)
	return db, clearFunc

}

func NewGormDB(c *config.AppConfig) *gorm.DB {
	var dsn string
	var dialector gorm.Dialector
	switch c.Database.ServerType {

	case "postgres":
		dsn = c.Database.DSN()
		dialector = postgres.Open(dsn)

	default:
		panic(fmt.Errorf("unkown database type: %s", c.Database.ServerType))
	}

	dbConfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "tb_",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	}

	CreateDatabase(c)

	gdb, err := gorm.Open(dialector, dbConfig)
	if err != nil {
		log.Fatal().Err(err).Send()
		return nil
	}

	if c.Database.Debug {
		gdb = gdb.Debug()
	}

	sqlDB, err := gdb.DB()
	if err != nil {
		log.Fatal().Err(err).Send()
		return nil
	}

	sqlDB.SetMaxIdleConns(c.Database.MaxIdleConn)
	sqlDB.SetMaxOpenConns(c.Database.MaxOpenConn)
	sqlDB.SetConnMaxLifetime(time.Duration(c.Database.ConnMaxLifetime) * time.Second)
	return gdb
}

func CreateDatabase(c *config.AppConfig) {
	dbConf := c.Database
	// create database if not exists
	preDsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=postgres sslmode=%s",
		dbConf.Host, dbConf.Port, dbConf.Username, dbConf.Password, dbConf.SslMode)

	switch c.Database.ServerType {

	case "postgres":
		gdb, err := gorm.Open(postgres.Open(preDsn), &gorm.Config{})
		if err != nil {
			log.Fatal().Err(err).Send()
		}

		exit := 0
		res1 := gdb.Table("pg_database").Select("count(1)").
			Where("datname = ?", dbConf.DbName).Scan(&exit)
		if res1.Error != nil {
			log.Fatal().Err(err).Send()
		}

		if exit == 0 {
			log.Info().Msgf("trying to create database: %s", dbConf.DbName)
			res2 := gdb.Exec(fmt.Sprintf("CREATE DATABASE %s", dbConf.DbName))
			if res2.Error != nil {
				log.Fatal().Err(err).Send()
			}
		}
	}

}

func PingDb(db *gorm.DB, c *config.AppConfig) {
	var count = 0
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal().Err(err).Send()
	}

	for {
		time.Sleep(time.Duration(c.Database.DbPing) * time.Second)
		if _, err = sqlDb.Exec("create table if not exists t1(c1 int)"); err != nil {
			log.Error(err).Msg("create test table error")
			sqlDb.Close()
			if db, err = gorm.Open(postgres.Open(c.Database.DSN()), &gorm.Config{}); err != nil {
				log.Error(err).Msg("reconnect database error")
			}
		}
		sqlDb, _ = db.DB()
		err = sqlDb.Ping()
		if err != nil {
			log.Error(err).Msg("db ping error")
		}
		count++
		if count%20 == 0 {
			// print heart beat in each 20 cycles
			log.Error(err).Int("count", count).Msg("db ping!")
		}
		if count > 10000 {
			count = 0
		}
	}
}
