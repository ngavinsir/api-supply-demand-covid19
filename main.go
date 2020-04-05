package main

import (
	"github.com/ngavinsir/api-supply-demand-covid19/cmd"
)

//go:generate sqlboiler --wipe psql

func main() {
	cmd.Execute()
}
