package db

import (
	"errors"
	"fmt"
	"github.com/ice-cream-heaven/log"
	"gorm.io/gorm/logger"
	"time"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	mysqlC "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Client struct {
	db *gorm.DB
}

func New(c *Config, tables ...interface{}) (*Client, error) {
	p := &Client{}

	c.apply()

	if c.Logger == nil {
		c.Logger = logger.Discard
	}

	var d gorm.Dialector
	switch c.Type {
	case "sqlite":
		d = newSqlite(c)

	case "mysql":
		d = mysql.New(mysql.Config{
			DSN: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.Username, c.Password, c.Address, c.Port, c.Name),
			DSNConfig: &mysqlC.Config{
				Timeout:                 time.Second * 5,
				ReadTimeout:             time.Second * 30,
				WriteTimeout:            time.Second * 30,
				AllowAllFiles:           true,
				AllowCleartextPasswords: true,
				AllowNativePasswords:    true,
				AllowOldPasswords:       true,
				CheckConnLiveness:       true,
				ClientFoundRows:         true,
				ColumnsWithAlias:        true,
				InterpolateParams:       true,
				MultiStatements:         true,
				ParseTime:               true,
			},
			Conn:                          nil,
			DefaultStringSize:             500,
			DontSupportNullAsDefaultValue: true,
		})

	case "postgres":
		d = postgres.New(postgres.Config{
			DSN:                  fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Address, c.Port, c.Username, c.Password, c.Name),
			PreferSimpleProtocol: true,
			WithoutReturning:     false,
			Conn:                 nil,
		})

	case "sqlserver":
		d = sqlserver.Open(fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s", c.Username, c.Password, c.Address, c.Port, c.Name))

	default:
		return nil, errors.New("unknown database")
	}

	var err error
	p.db, err = gorm.Open(d, &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: &schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: true,
		},
		FullSaveAssociations:                     false,
		Logger:                                   c.Logger,
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		AllowGlobalUpdate:                        true,
		CreateBatchSize:                          100,
	})
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	if c.Debug {
		p.db = p.db.Debug()
	}

	conn, err := p.db.DB()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	err = conn.Ping()
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	switch c.Type {
	case "sqlite":
		// 自动减少存储文件大小
		err = p.db.Session(&gorm.Session{
			NewDB: true,
		}).Exec("PRAGMA auto_vacuum = 1").Error
		if err != nil {
			log.Errorf("err:%v", err)
			return nil, err
		}
	}

	err = p.AutoMigrate(tables...)
	if err != nil {
		log.Errorf("err:%v", err)
		return nil, err
	}

	return p, nil
}

func (p *Client) AutoMigrate(dst ...interface{}) error {
	return p.db.AutoMigrate(dst...)
}

func (p *Client) NewScoop() *Scoop {
	return NewScoop(p.db)
}
