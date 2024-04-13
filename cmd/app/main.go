package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/mirhijinam/avito-trainee-2024/internal/api"
	"github.com/mirhijinam/avito-trainee-2024/internal/config"
	"github.com/mirhijinam/avito-trainee-2024/internal/pkg/db"
	"github.com/mirhijinam/avito-trainee-2024/internal/repository"
	"github.com/mirhijinam/avito-trainee-2024/internal/service"
)

func main() {
	dbCfg, err := config.GetDBConfig()
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	conn := db.MustOpenDB(ctx, dbCfg)

	srvCfg := config.GetServerConfig()

	br := repository.New(conn)
	lruSz, err := strconv.Atoi(os.Getenv("LRU_CACHE_SIZE"))
	fmt.Println("debug! lrucache size:", lruSz)
	bs := service.New(br, lruSz)
	h := api.New(bs)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			fmt.Fprintln(w, "Hi there!")
		} else {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("GET /user_banner", h.GetBanner())
	mux.HandleFunc("GET /banner", h.GetBannerList())
	mux.HandleFunc("POST /banner", h.CreateBanner())
	mux.HandleFunc("PATCH /banner/{id}", h.UpdateBanner())
	mux.HandleFunc("DELETE /banner/{id}", h.DeleteBanner())

	srv := http.Server{
		Addr:    srvCfg.HTTPPort,
		Handler: mux,
	}
	log.Println("Server is starting on port " + srvCfg.HTTPPort)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
