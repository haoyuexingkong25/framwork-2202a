package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MysqlClient(hort, port, user, pass, dbname string, handler func(db *gorm.DB) error) error {

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, hort, port, dbname,
	)
	cli, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db, err := cli.DB()
	if err != nil {
		return err
	}
	defer db.Close()

	return handler(cli)

}
