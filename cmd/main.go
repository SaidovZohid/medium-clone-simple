package main

import (
	"fmt"
	"log"

	"github.com/SaidovZohid/medium-clone-simple/config"
	"github.com/SaidovZohid/medium-clone-simple/server"
	"github.com/SaidovZohid/medium-clone-simple/storage"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load(".")

	psqlUrl := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Database,
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatalf("Failed to connect to postgres: %v", err)
	}

	log.Println("Postgres Connection successfully done!")

	strg := storage.NewStorage(psqlConn)

	router := server.NewServer(&server.Options{
		Strg: strg,
	})

	if err = router.Run(cfg.Port); err != nil {
		log.Fatalf("Failed to run to server: %v", err)
	}
	fmt.Println("Hello World!")
}
