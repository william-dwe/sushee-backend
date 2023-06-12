package main

import (
	"sushee-backend/config"
	"sushee-backend/db"
	"sushee-backend/server"

	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Sushee API Backend Started")
	log.Info().Msg("CONFIG: " + config.Config.ENVConfig.Mode)

	dbErr := db.Connect()
	if dbErr != nil {
		log.Fatal().Msg("error connecting to DB")
	}

	server.Init()

	log.Info().Msg("Sushee API Backend is ready!")
}
