package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var meCmd = &cobra.Command{
	Use:   "me",
	Short: "Visualize profile info",
	Long:  "Visualize profile info",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context().Value(appCtxKey)
		if ctx == nil {
			return fmt.Errorf("ctx is nil")
		}

		appCtx, ok := ctx.(*AppContext)
		if !ok {
			return fmt.Errorf("invalid appContext")
		}

		if appCtx.StatService == nil {
			return fmt.Errorf("application not initialized, run 'gova login' first")
		}

		athleteSummary, err := appCtx.StatService.GetAthleteSummary()
		if err != nil {
			return fmt.Errorf("failed to fetch athlete summary: %w", err)
		}

		fmt.Println(athleteSummary)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(meCmd)
}
