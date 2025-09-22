package main

import (
	"github.com/giancarlosisasi/gemini-cli-clone/cmd"
	"github.com/rs/zerolog/log"
)

func main() {
	if err := cmd.Execute(); err != nil {
		log.Debug().Err(err).Msg("error on main function.")
		log.Fatal().Msg("error to init the gemini-cli clone.")
	}
}
