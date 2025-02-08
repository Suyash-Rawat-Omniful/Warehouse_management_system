package appinit

import (
	"context"
	"service2/redis"
	"time"

	"github.com/omniful/go_commons/config"
	"github.com/omniful/go_commons/db/sql/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Initialize(ctx context.Context) {
	initializeDB(ctx)
	redis.Start()
}

func initializeDB(ctx context.Context) {
	debugMode := config.GetBool(ctx, "postgresql.debugMode")
	connIdleTimeout := 10 * time.Minute
	slavesConfig := make([]postgres.DBConfig, 0)
	masterConfig := postgres.DBConfig{
		Host:               "localhost",
		Port:               "5432",
		Username:           "sample_user",
		Password:           "root",
		Dbname:             "warehouse_management_system",
		MaxOpenConnections: 10,
		MaxIdleConnections: 2,
		ConnMaxLifetime:    connIdleTimeout,
		DebugMode:          debugMode,
	}

	DBConfig := postgres.InitializeDBInstance(masterConfig, &slavesConfig)
	DB = DBConfig.GetMasterDB(ctx)

}
