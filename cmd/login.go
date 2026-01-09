package cmd

import (
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Strava",
	Long:  "Login to Strava",
	RunE: func(cmd *cobra.Command, args []string) error {
		if appCtx == nil || appCtx.AuthService == nil {
			return fmt.Errorf("strava Auth Service is not initialized")
		}

		fmt.Println("Login called")
		fmt.Println("Checking for config file...")

		loginUrl := stravaClient.BuildAuthURL()

		fmt.Println("Open this URL in your browser to login to Strava", loginUrl)
		err := exec.Command("open", loginUrl).Start()
		if err != nil {
			fmt.Println("Could not open browser", err.Error())
		}

		// -> -> User does not click on accept -> abort
		// -> -> User clicks on accept
		// -> -> -> Get code from browser
		code := "todo"
		err = appCtx.AuthService.GetTokenFromCode(code)
		if err != nil {
			return fmt.Errorf("could not get token from code: %s", err.Error())
		}

		return nil

		// -> -> -> -> close browser tab
		// SUCCESS
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
