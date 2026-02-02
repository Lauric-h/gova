package cmd

import (
	"fmt"
	"gova/internal/domain"

	"github.com/spf13/cobra"
)

var monthCmd = &cobra.Command{
	Use:   "month",
	Short: "Visualize monthly stats",
	Long:  "Visualize monthly stats",
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

		shouldGetLast, err := cmd.Flags().GetBool("last")
		if err != nil {
			return err
		}

		period := domain.CreateMonth(shouldGetLast)
		activitiesSummary, err := appCtx.StatService.ListActivities(period)
		if err != nil {
			return err
		}

		fmt.Println("Du", period.StartDay.Format("02/01/2006"), "au", period.EndDay.Format("02/01/2006"))
		if len(activitiesSummary) == 0 {
			fmt.Println("Pas d'activités sur cette période pour le moment")
			return nil
		}

		for _, activity := range activitiesSummary {
			fmt.Printf("Activité %s (%d): %.1f km, %.1fh, %dm de dénivelé positif\n",
				activity.SportType.String(),
				activity.Count,
				activity.GetDistanceInKm(),
				activity.GetDurationInHours(),
				activity.TotalAscent,
			)
		}

		return nil
	},
}

func init() {
	monthCmd.Flags().BoolP("last", "l", false, "Get last month's stats")
	rootCmd.AddCommand(monthCmd)
}
