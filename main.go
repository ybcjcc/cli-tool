package main

import (
	"cli-tools/api"
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	app := &cli.App{
		Name: "cli tool using golang",
		Commands: []*cli.Command{
			{
				Name:    "timestamp to datetime",
				Aliases: []string{"t2d"},
				Usage:   "unix timestamp to datetime",
				Action: func(c *cli.Context) error {
					in := c.Args().Get(0)
					ts, _ := strconv.ParseInt(in, 10, 64)
					t := time.Unix(ts, 0)
					fmt.Println(t.Format("2006-01-02 15:04:05"))
					return nil
				},
			},
			{
				Name:    "datetime to timestamp",
				Aliases: []string{"d2t"},
				Usage:   "formatted datetime to timestamp",
				Action: func(c *cli.Context) error {
					in := c.Args().Get(0)
					loc, _ := time.LoadLocation("Local")
					var err error
					var t time.Time
					if len(in) != 10 {
						t, err = time.ParseInLocation("2006-01-02 15:04:05", in, loc)
					} else {
						t, err = time.ParseInLocation("2006-01-02", in, loc)
					}
					if err == nil {
						fmt.Println(t.Unix())
					}
					return nil
				},
			},
			{
				Name:    "stock hang qing",
				Aliases: []string{"hq"},
				Usage:   "stock hang qing",
				Action: func(c *cli.Context) error {
					in := c.Args().Get(0)
					return api.Stock(in)
				},
			},
			{
				Name:    "youdao translate",
				Aliases: []string{"yd"},
				Usage:   "youdao translate tool",
				Action: func(c *cli.Context) error {
					in := c.Args().Get(0)
					return api.Youdao(in)
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
