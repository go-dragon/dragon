package repository

import (
	"dragon/core/dragon/conf"
	"dragon/core/dragon/dlogger"
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql" //导入mysql驱动
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"regexp"
	"time"
)

var (
	GormDB *gorm.DB //master db
)

// Logger sql logger
type Logger struct {
	logger.Writer
}

func (l Logger) Printf(s string, i ...interface{}) {
	s = fmt.Sprintf(s, i...)
	// 日志打印
	res, _ := regexp.MatchString("(Error)|(SLOW SQL)", s)

	// if sql error
	if res {
		dlogger.SqlError(s)
	} else {
		dlogger.SqlInfo(s)
	}
}

//init db
func InitDB() {
	var err error
	var dsnMaster string
	var logHandler logger.Interface
	if conf.Env == "dev" {
		logHandler = logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			Colorful:      true,
			LogLevel:      logger.Info,
		})
	} else {
		// other env write log
		logHandler = logger.New(Logger{}, logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		})
	}

	//mysql master
	dsnMaster = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=%s&loc=Local", //loc set the timezone
		viper.GetString("database.mysql.master.user"), viper.GetString("database.mysql.master.password"), viper.GetString("database.mysql.master.host"), viper.GetString("database.mysql.master.port"), viper.GetString("database.mysql.master.database"), viper.GetString("database.mysql.master.charset"), viper.GetString("database.mysql.master.timeout"))

	//gorm realizes mysql reconnect
	GormDB, err = gorm.Open(mysql.New(mysql.Config{
		DriverName:                "mysql",
		DSN:                       dsnMaster,
		Conn:                      nil,
		SkipInitializeWithVersion: false,
		DefaultStringSize:         0,
		DisableDatetimePrecision:  false,
		DontSupportRenameIndex:    false,
		DontSupportRenameColumn:   false,
	}), &gorm.Config{
		SkipDefaultTransaction:                   true,
		NamingStrategy:                           nil,
		Logger:                                   logHandler,
		NowFunc:                                  nil,
		DryRun:                                   false,
		PrepareStmt:                              true,
		DisableAutomaticPing:                     false,
		DisableForeignKeyConstraintWhenMigrating: false,
		AllowGlobalUpdate:                        false,
		ClauseBuilders:                           nil,
		ConnPool:                                 nil,
		Dialector:                                nil,
		Plugins:                                  nil,
	})
	if err != nil {
		log.Fatalln(err)
	}
	sqlDb, err := GormDB.DB()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("mysql maxIdle conns:", viper.GetInt("database.mysql.master.maxidle"))
	log.Println("mysql maxOpenConn conns:", viper.GetInt("database.mysql.master.maxconn"))
	sqlDb.SetMaxIdleConns(viper.GetInt("database.mysql.master.maxidle"))
	sqlDb.SetMaxOpenConns(viper.GetInt("database.mysql.master.maxconn"))

	sqlDb.SetConnMaxIdleTime(10 * 60 * time.Second) // 连接最大空闲时间为10分钟，超过则释放重新生成
	sqlDb.SetConnMaxLifetime(30 * 60 * time.Second) // 每过30分钟，重新建立数据库连接
}
