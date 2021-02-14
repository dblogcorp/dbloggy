package main

import "github.com/dblogcorp/dbloggy/cmd/dblog-sso/app"

func main() {
	cmd := app.NewCommand()
	if err := cmd.Execute(); err != nil {
		panic(err)
	}
}
