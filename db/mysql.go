package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	Log "github.com/sirupsen/logrus"
	"go-chain/config"
)

func Init() *dbr.Session {

	session := getSession()
	return session

}

func getSession() *dbr.Session {

	db, err := dbr.Open("mysql",
		config.USER+":"+config.PASSWORD+"@tcp("+config.HOST+":"+config.PORT+")/"+config.DB,
		nil)

	if err != nil {
		Log.Error(err)
	} else {
		session := db.NewSession(nil)
		return session
	}
	return nil
}
