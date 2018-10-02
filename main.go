package main

import (
	"os"

	"github.com/gwenker/matrix-rollup-dice/rollupdice"
	"github.com/sirupsen/logrus"
	cli "gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()

	logrus.SetLevel(logrus.DebugLevel)

	app.Commands = []cli.Command{{
		Name:  "start",
		Usage: "start the matrix rollup dice daemon",
		Flags: []cli.Flag{
			cli.StringFlag{Name: "core-uri", Value: "localhost:8080",
				EnvVar: "MATRIX_CORE_URI"},
		},
		Action: func(c *cli.Context) error { return rollupdice.Start(c.String("core-uri")) },
	}}

	if err := app.Run(os.Args); err != nil {
		logrus.Errorf("%+v", err)
	}
}
