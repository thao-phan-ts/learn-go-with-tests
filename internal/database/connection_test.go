package database

import (
	"log"
	"testing"
)

func TestConnection(t *testing.T) {
	var DBConfig Config
	DBConfig.Username = "sdet_automation"
	DBConfig.Password = "bEmX3SM8wJ4NR7Wv2AZc"
	DBConfig.Host = "dop-rs7-release-mysql.ckzhnzkexzci.ap-southeast-1.rds.amazonaws.com"
	DBConfig.Port = 3306
	DBConfig.Database = "stg_digital_journey"

	db, err := ConnectMySQL(&DBConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

}
