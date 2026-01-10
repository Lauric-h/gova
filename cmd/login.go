package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Strava",
	Long:  "Login to Strava",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context().Value(appCtxKey)
		if ctx == nil {
			return fmt.Errorf("ctx is nil")
		}

		appCtx, ok := ctx.(*AppContext)
		if !ok {
			return fmt.Errorf("invalid appContext")
		}

		if appCtx.AuthService == nil {
			return fmt.Errorf("strava Auth Service is not initialized")
		}

		fmt.Println("Starting authentication with Strava...")
		fmt.Println("Checking for existing credentials...")

		// Start HTTP Server
		loginUrl := appCtx.AuthService.BuildLoginUrl()
		fmt.Println("\nOpen this URL in your browser to authenticate:")
		fmt.Println(loginUrl)
		oAuthResult, err := appCtx.AuthService.StartOAuthFlow()
		if err != nil {
			return err
		}

		err = appCtx.AuthService.GetTokenFromCode(oAuthResult.Code)
		if err != nil {
			return fmt.Errorf("could not get token from code: %s", err.Error())
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
