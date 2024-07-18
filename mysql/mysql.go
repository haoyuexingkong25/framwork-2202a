package mysql

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlClient(handler func(db *gorm.DB) error) error {

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		//user, pass, hort, port, dbname,
		viper.GetString("mysql.user"),
		viper.GetString("mysql.pass"),
		viper.GetString("mysql.hort"),
		viper.GetString("mysql.port"),
		viper.GetString("mysql.dbname"),
	)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	cli.Begin()
	db, err := cli.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	return handler(cli)

}

func OpensBegin(handle func(db *gorm.DB) error) error {
	return MysqlClient(func(db *gorm.DB) error {
		tx := db.Begin()
		var err error
		defer func() {
			if err == nil {
				tx.Commit()
			} else {
				tx.Rollback()
			}
		}()
		return handle(tx)
	})
}
