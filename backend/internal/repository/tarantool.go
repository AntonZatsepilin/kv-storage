package repository

import "github.com/sirupsen/logrus"

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
	SSLmode  string
}

func NewTarantoolDB(cfg Config) () {
	logrus.Info("Connecting to the database")
	db, err := 

	logrus.Info("Checking the connection to the database")
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}