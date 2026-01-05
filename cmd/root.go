package cmd

import (
	"fmt"
	"gova/internal/client"
	"gova/internal/config"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var shouldGetLast bool
var stravaClient *client.Client

var rootCmd = &cobra.Command{
	Use:   "gova",
	Short: "gova is a CLI tool to visualize your Strava stats",
	Long:  "gova is a CLI tool to visualize your weekly and monthly Strava stats for running and trail running",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	stravaClient = client.NewClient(cfg.BaseURL, cfg.StravaToken)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
