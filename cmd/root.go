package cmd

import (
	"context"
	"fmt"
	"gova/internal/config"
	"gova/internal/service"
	"gova/internal/strava"
	"os"

	"github.com/spf13/cobra"
)

type AppContext struct {
	Config       *config.Config
	StravaClient *strava.Client
	StatService  *service.StatService
	AuthService  *service.AuthService
}

const appCtxKey = "appCtx"

var rootCmd = &cobra.Command{
	Use:   "gova",
	Short: "gova is a CLI tool to visualize your Strava stats",
	Long:  "gova is a CLI tool to visualize your weekly and monthly Strava stats for running and trail running",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load()
		if err != nil {
			return err
		}

		oAuthClient := strava.NewOauthClient(cfg)
		authService := service.NewAuthService(oAuthClient)

		stravaClient := strava.NewClient(cfg, authService)
		appCtx := &AppContext{
			Config:       cfg,
			AuthService:  authService,
			StravaClient: stravaClient,
			StatService:  service.NewStatService(stravaClient),
		}

		cmd.SetContext(context.WithValue(cmd.Context(), appCtxKey, appCtx))

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
