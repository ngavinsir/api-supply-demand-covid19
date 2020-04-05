package cmd

import (
	"github.com/ngavinsir/api-supply-demand-covid19/database"
	"github.com/ngavinsir/api-supply-demand-covid19/handler"
	"github.com/spf13/cobra"
)

var cmdAdmin = &cobra.Command{
	Use:  "admin [email password]",
	Args: cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		db, err := database.InitDB()
		if err != nil {
			panic(err)
		}
		defer db.Close()

		api := handler.NewAPI(db)
		api.Cmd(args)
	},
}
