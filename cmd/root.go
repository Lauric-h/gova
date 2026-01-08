package cmd

import (
	"fmt"
	"gova/internal/config"
	"gova/internal/service"
	"gova/internal/strava"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var shouldGetLast bool
var statService *service.StatService
var cfg *config.Config

var rootCmd = &cobra.Command{
	Use:   "gova",
	Short: "gova is a CLI tool to visualize your Strava stats",
	Long:  "gova is a CLI tool to visualize your weekly and monthly Strava stats for running and trail running",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	var err error
	cfg, err = config.Load()
	if err != nil {
		log.Fatal(err)
	}

	stravaClient := strava.NewClient(cfg.StravaToken, cfg.ClientId)
	statService = service.NewStatService(stravaClient)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
