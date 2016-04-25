package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/harborapp/harbor-api/config"
	"github.com/harborapp/harbor-api/router"
	"github.com/harborapp/harbor-api/server"
)

// Server provides the sub-command to start the API server.
func Server() cli.Command {
	return cli.Command{
		Name:  "server",
		Usage: "Start the Harbor server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:        "db-driver",
				Value:       "mysql",
				Usage:       "Database driver selection",
				EnvVar:      "HARBOR_DB_DRIVER",
				Destination: &config.Database.Driver,
			},
			cli.StringFlag{
				Name:        "db-name",
				Value:       "harbor",
				Usage:       "Name for database connection",
				EnvVar:      "HARBOR_DB_NAME",
				Destination: &config.Database.Name,
			},
			cli.StringFlag{
				Name:        "db-username",
				Value:       "root",
				Usage:       "Username for database connection",
				EnvVar:      "HARBOR_DB_USERNAME",
				Destination: &config.Database.Username,
			},
			cli.StringFlag{
				Name:        "db-password",
				Value:       "root",
				Usage:       "Password for database connection",
				EnvVar:      "HARBOR_DB_PASSWORD",
				Destination: &config.Database.Password,
			},
			cli.StringFlag{
				Name:        "db-host",
				Value:       "localhost:3306",
				Usage:       "Host for database connection",
				EnvVar:      "HARBOR_DB_HOST",
				Destination: &config.Database.Host,
			},
			cli.StringFlag{
				Name:        "addr",
				Value:       ":8080",
				Usage:       "Address to bind the server",
				EnvVar:      "HARBOR_SERVER_ADDR",
				Destination: &config.Server.Addr,
			},
			cli.StringFlag{
				Name:        "cert",
				Value:       "",
				Usage:       "Path to SSL cert",
				EnvVar:      "HARBOR_SERVER_CERT",
				Destination: &config.Server.Cert,
			},
			cli.StringFlag{
				Name:        "key",
				Value:       "",
				Usage:       "Path to SSL key",
				EnvVar:      "HARBOR_SERVER_KEY",
				Destination: &config.Server.Key,
			},
			cli.StringFlag{
				Name:        "root",
				Value:       "/",
				Usage:       "Root folder of the app",
				EnvVar:      "HARBOR_SERVER_ROOT",
				Destination: &config.Server.Root,
			},
			cli.StringFlag{
				Name:        "storage",
				Value:       "storage/",
				Usage:       "Folder for storing uploads",
				EnvVar:      "SOLDER_SERVER_STORAGE",
				Destination: &config.Server.Storage,
			},
		},
		Action: func(c *cli.Context) {
			srv := server.Load()
			srv.Run(router.Load())
		},
	}
}