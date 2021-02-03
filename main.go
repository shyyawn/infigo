package main

import (
	_ "github.com/logrusorgru/aurora"
	config "github.com/shyyawn/go-to/config"
	_ "github.com/shyyawn/go-to/logging"
	log "github.com/shyyawn/go-to/logging"
	"github.com/shyyawn/infigo/cmd"
	_ "github.com/spf13/cobra"
	"path/filepath"
)

func main() {
	log.Info("Starting Infigo")

	appPath, err := filepath.Abs("./")
	if err != nil {
		log.Fatal("Some issue getting the directory path")
	}

	config.Init(appPath)
	cmd.Execute()
}
