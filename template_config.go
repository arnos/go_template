package main

import (
	"github.com/BurntSushi/toml"
	"github.com/rs/zerolog/log"
	"os"
)

func templateConfiguration() {
	configurationFile := os.Getenv("ConfigurationFile")

	if "" != configurationFile {
		if _, tomlErr := toml.DecodeFile(configurationFile, &Config); tomlErr != nil {
			log.Printf("Error reading the configuration file.  Error is: %s", tomlErr.Error())
		}
  } else {
     log.Error().Msgf("failed to import configuration file")
  }
}
