package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/tijanadmi/food_order/api"
	db "github.com/tijanadmi/food_order/db/sqlc"
	"github.com/tijanadmi/food_order/util"
)

var interruptSignals = []os.Signal{
	os.Interrupt,
	syscall.SIGTERM,
	syscall.SIGINT,
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	ctx, stop := signal.NotifyContext(context.Background(), interruptSignals...)
	defer stop()

	connPool, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to db")
	}

	//runDBMigration(config.MigrationURL, config.DBSource)

	store := db.NewStore(connPool)

	runGinServer(config, store)
}

// func runDBMigration(migrationURL string, dbSource string) {
// 	migration, err := migrate.New(migrationURL, dbSource)
// 	if err != nil {
// 		log.Fatal().Err(err).Msg("cannot create new migrate instance")
// 	}

// 	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
// 		log.Fatal().Err(err).Msg("failed to run migrate up")
// 	}

// 	log.Info().Msg("db migrated successfully")
// }

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create server")
	}

	// Configure CORS middleware
	// corsHandler := cors.New(cors.Options{
	// 	AllowedOrigins:   []string{"http://localhost:5173"}, // Ovde dozvoli specifične ili sve origne
	// 	AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
	// 	AllowedHeaders:   []string{"Origin", "Content-Type", "Authorization", "Cookie"},
	// 	ExposedHeaders:   []string{"Content-Length", "Set-Cookie"},
	// 	AllowCredentials: true,
	// }).Handler(mux)

	err = server.Start(config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start server")
	}
}
