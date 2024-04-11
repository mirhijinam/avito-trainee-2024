package main

import (
	"context"
	"log"

	_ "github.com/lib/pq"
	"github.com/mirhijinam/avito-trainee-2024/internal/config"
	"github.com/mirhijinam/avito-trainee-2024/internal/pkg/db"
)

func main() {
	dbCfg, err := config.GetDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	_ = db.MustOpenDB(ctx, dbCfg)

	srvCfg := config.GetServerConfig()

	log.Println("Server has been successfully started on the port " + srvCfg.HTTPPort)

	for {
	}

}
