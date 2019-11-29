package persist

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	"github.com/spf13/viper"
)

func NewDB() (*gorm.DB, error) {
	drvr := viper.GetString("connstr.drvr")
	dial := viper.GetString("connstr.dial")
	host := viper.GetString("connstr.host")
	port := viper.GetString("connstr.port")
	user := viper.GetString("connstr.user")
	pass := viper.GetString("connstr.pass")
	dbse := viper.GetString("connstr.dbse")

	connstr := fmt.Sprintf("%s://%s:%s@%s:%s?database=%s", dial, user, pass, host, port, dbse)
	db, err := gorm.Open(drvr, connstr)

	if err != nil {
		return nil, err
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)

	return db, nil
}
