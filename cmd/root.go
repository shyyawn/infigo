package cmd

import (
	"fmt"
	c "github.com/logrusorgru/aurora"
	log "github.com/shyyawn/go-to/logging"
	"github.com/shyyawn/infigo/cmd/seoCheck"
	"github.com/spf13/cobra"
	//config "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "infigo",
	Short: "infigo is a cli to fix random stuff, just a cli for me to keep repeated tasks",
	Long:  "infigo is a cli to fix random stuff, just a cli for me to keep repeated tasks",
	Run:   infigo,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(seoCheck.Cmd)
}

func infigo(cmd *cobra.Command, args []string) {
	fmt.Print(c.Green("Welcome to Infigo"))
}
