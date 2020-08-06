// Copyright 2020 arugal, zhangwei24@apache.org
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/armon/go-socks5"
	"github.com/urfave/cli"
)

var (
	cmdStart = cli.Command{
		Name:  "start",
		Usage: "start socks5 server",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:     "addr, a",
				Usage:    "socks5 listens for addresses",
				Required: false,
				EnvVar:   "SOCKS5_SERVER_ADDR",
				Value:    ":55055",
			},
			cli.StringFlag{
				Name:     "username",
				Usage:    "socks5 server username",
				Required: false,
				Value:    "admin",
			},
			cli.StringFlag{
				Name:     "password, pwd",
				Usage:    "socks5 server password",
				Required: false,
				Value:    "admin",
			},
		},
		Action: func(ctx *cli.Context) error {
			addr := ctx.String("addr")
			username := ctx.String("username")
			password := ctx.String("password")

			creds := socks5.StaticCredentials{
				username: password,
			}

			cator := socks5.UserPassAuthenticator{Credentials: creds}

			conf := &socks5.Config{
				AuthMethods: []socks5.Authenticator{cator},
			}
			server, err := socks5.New(conf)
			if err != nil {
				return err
			}

			if err := server.ListenAndServe("tcp", addr); err != nil {
				return err
			}
			return nil
		},
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "socks5-server"
	app.Usage = "https://github.com/arugal/socks5-server"
	app.Compiled = time.Now()
	app.Copyright = "(c) " + strconv.Itoa(time.Now().Year()) + " arugal"
	app.Description = "Socks5 Server"

	app.Commands = []cli.Command{
		cmdStart,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
