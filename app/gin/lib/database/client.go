package lib/database

import (
	"github.com/kantaroso/kanta-workspace/config"
)

type DatabaseClient interface {
	query()
}

type DatabaseClient struct {
	DatabaseClient
	config config.Database
}

func GetDatabaseClient(){

}
