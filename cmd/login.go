package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to Strava",
	Long:  "Login to Strava",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("login called")
		fmt.Println("checking for config file...")
		// Get client_id from Config file
		// If no info -> abort
		// TODO MOVE TO CFG LOAD
		clientId := os.Getenv("STRAVA_CLIENT_ID")
		if clientId == "" {
			log.Fatal("clientId is required")
		}

		// -> open browser to URL with scope
		// TODO STORE URL SOMEWHERE ELSE
		url := fmt.Sprintf("http://www.strava.com/oauth/authorize?client_id=%s&response_type=code&redirect_uri=http://localhost/exchange_token&approval_prompt=force&scope=read,activity:read_all",
			clientId,
		)
		err := exec.Command("open", url).Start()
		if err != nil {
			fmt.Println("Could not open browser", err.Error())
			fmt.Println("Open this URL in your browser", url)
		}

		// -> -> User does not click on accept -> abort
		// -> -> User clicks on accept
		// -> -> -> Get code from browser
		// -> -> -> POST to exchange code for token
		// -> -> -> No right scope -> abort
		// -> -> -> Store token + refresh token + exp date in config file
		// SUCCESS
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
}
