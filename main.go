package main

import (
	"context"
	appinit "service2/init"
	"service2/router"
	"strconv"
	"time"

	"github.com/omniful/go_commons/db/sql/migration"
	"github.com/omniful/go_commons/http"
	"github.com/omniful/go_commons/log"
)

const (
	modeMigration  = "migration"
	upMigration    = "up"
	downMigration  = "down"
	forceMigration = "force"
)

func main() {
	// err := config.Init(time.Second * 10)
	// if err != nil {
	// 	log.Panicf("Error while initialising config, err: %v", err)
	// 	panic(err)
	// }

	// ctx, err := config.TODOContext()
	// if err != nil {
	// 	log.Panicf("Error while getting context from config, err: %v", err)
	// 	panic(err)
	// }
	ctx := context.TODO()
	runMigration(ctx, "up", "0")
	appinit.Initialize(ctx)

	// runHttpServer(ctx)
	server := http.InitializeServer(":8081", 10*time.Second, 10*time.Second, 70*time.Second)
	log.Infof("Starting server on port 8081")

	err := router.Initialize(ctx, server)
	if err != nil {
		log.Panicf("Error while initialising router, err: %v", err)
		panic(err)
	}

	err = server.StartServer("Tenant-service")
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

}

func runMigration(ctx context.Context, migrationType string, number string) {
	database := "warehouse_management_system"
	mysqlWriteHost := "localhost"
	mysqlWritePort := "5432"
	mysqlWritePassword := "root"
	mysqlWriterUsername := "sample_user"

	m, err := migration.InitializeMigrate("file://deployment/migration", "postgres://"+mysqlWriteHost+":"+mysqlWritePort+"/"+database+"?user="+mysqlWriterUsername+"&password="+mysqlWritePassword+"&sslmode=disable")
	if err != nil {
		panic(err)
	}

	switch migrationType {
	case upMigration:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	case downMigration:
		err = m.Down()
		if err != nil {
			panic(err)
		}
		break
	case forceMigration:
		version, parseErr := strconv.Atoi(number)
		if parseErr != nil {
			panic(parseErr)
		}

		err = m.ForceVersion(version)
		if err != nil {
			return
		}
		break
	default:
		err = m.Up()
		if err != nil {
			panic(err)
		}
		break
	}
}
