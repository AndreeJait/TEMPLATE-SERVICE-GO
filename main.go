package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AndreeJait/GO-ANDREE-UTILITIES/logs"
	"github.com/AndreeJait/GO-ANDREE-UTILITIES/util/echokit"
	mm "github.com/AndreeJait/GO-ANDREE-UTILITIES/util/migration"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/config"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/di"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/infrastructure"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/interfaces/controller"
	"github.com/AndreeJait/TEMPLATE-SERVICE-GO/shared"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/pkg/errors"
	swagger "github.com/swaggo/echo-swagger"
	"github.com/urfave/cli"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	container, err := di.Container()

	if err != nil {
		panic(err)
	}

	err = container.Invoke(func(conf *config.EnvConfiguration, controllerHolder controller.Holder, closer *shared.Closer, migration mm.Tool, infrastructureHolder infrastructure.Holder, sharedHolder shared.Holder) error {

		app := cli.NewApp()

		// - migration command
		app.Commands = append(app.Commands, NewMigrationCommand(sharedHolder.Logger, closer, migration)...)
		app.Commands = append(app.Commands, NewRunCommand(conf, controllerHolder, infrastructureHolder, closer, sharedHolder)...)
		if err := app.Run(os.Args); err != nil {
			sharedHolder.Logger.Errorf("%s", err)
			return err
		}
		return nil
	})

	_ = container.Invoke(func(logger logs.Logger) error {
		if err != nil {
			logger.Error(err)
		}
		return nil
	})

}

func NewMigrationCommand(logger logs.Logger, closer *shared.Closer, migration mm.Tool) []cli.Command {
	return []cli.Command{
		{
			Name:  "migrate",
			Usage: "migrate sql and nosql database",
			After: func(c *cli.Context) error {
				logger.Infof("closing resource: %+v", c)
				closer.Close()
				return nil
			},
			Subcommands: []cli.Command{
				{
					Name:   "check",
					Usage:  "check nosql migration",
					Action: migrationCheck(logger, migration),
				},
				{
					Name:   "up",
					Usage:  "update migration to latest version",
					Action: migrationUp(logger, migration),
				},
				{
					Name:   "down",
					Usage:  "downgrade migration to previous version",
					Action: migrationDown(logger, migration),
				},
				{
					Name:   "init",
					Usage:  "initialize migration table for no sql",
					Action: migrationInit(logger, migration),
				},
			},
		},
	}
}

func migrationCheck(logger logs.Logger, migration mm.Tool) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		logContext(logger, c)

		if err := migration.Check(); err != nil {
			return errors.Wrap(err, "mongodb is not migrated")
		}

		logger.Info("database already up to date!")
		return nil
	}
}

func migrationUp(logger logs.Logger, migration mm.Tool) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		logContext(logger, c)

		if err := migration.Up(); err != nil {
			return errors.Wrap(err, "failed to up migrate mongodb database")
		}

		return nil
	}
}

func migrationDown(logger logs.Logger, migration mm.Tool) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		logContext(logger, c)

		if err := migration.Down(); err != nil {
			return errors.Wrap(err, "failed to down migrate mongodb database")
		}

		return nil
	}
}

func migrationInit(logger logs.Logger, migration mm.Tool) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		logContext(logger, c)

		_ = migration.Initialize()
		return nil
	}
}

func NewRunCommand(conf *config.EnvConfiguration, controllerHolder controller.Holder, infrastructureHolder infrastructure.Holder, closer *shared.Closer, sharedHolder shared.Holder) []cli.Command {
	return []cli.Command{
		{
			Name:  "run",
			Usage: "run as http service",
			Action: func(c *cli.Context) error {
				ec := echo.New()

				ec.Use(middleware.CORSWithConfig(middleware.CORSConfig{
					Skipper:      middleware.DefaultSkipper,
					AllowOrigins: []string{"*"},
					AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
				}))
				ec.Use(middleware.Recover())
				ec.Validator = sharedHolder.Validator

				wrapper, _ := echokit.NewLoggerWrapper(sharedHolder.Logger)
				ec.Logger = wrapper
				ec.Logger.SetLevel(log.INFO)

				ec.Static("/assets", "assets")

				ec.GET("/swagger/*", swagger.WrapHandler)
				ec.GET("/debug/pprof/*", echo.WrapHandler(http.DefaultServeMux))

				// Auth

				sig := make(chan os.Signal)
				signal.Notify(sig, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

				go func() {
					if err := ec.Start(fmt.Sprintf(":%d", conf.ServerPort)); err != nil {
						sharedHolder.Logger.Errorf("failed to start echo http server %s", err)
						sig <- syscall.SIGINT
					}
				}()

				<-sig

				_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
				defer cancel()

				sharedHolder.Logger.Info("closing resource: %+v", c)
				closer.Close()

				return nil
			},
		},
	}
}

func logContext(logger logs.Logger, c *cli.Context) {
	logger.Info("cli context : %+v", c)
}
