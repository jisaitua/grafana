package main

import (
	"os"
	"runtime"
	"strconv"

	"github.com/grafana/grafana/pkg/cmd"
	"github.com/grafana/grafana/pkg/log"
	"github.com/grafana/grafana/pkg/setting"

	"github.com/codegangsta/cli"
)

var version = "master"
var commit = "NA"
var buildstamp string

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	buildstampInt64, _ := strconv.ParseInt(buildstamp, 10, 64)

	setting.BuildVersion = version
	setting.BuildCommit = commit
	setting.BuildStamp = buildstampInt64

	app := cli.NewApp()
	app.Name = "Grafana Backend"
	app.Usage = "grafana web"
	app.Version = version
	app.Commands = []cli.Command{cmd.Web, cmd.ImportJson,
		cmd.ListAccounts, cmd.CreateAccount, cmd.DeleteAccount,
		cmd.ListDataSources, cmd.CreateDataSource, cmd.DescribeDataSource, cmd.DeleteDataSource}
	app.Flags = append(app.Flags, []cli.Flag{
		cli.StringFlag{
			Name:  "config",
			Value: "grafana.ini",
			Usage: "path to config file",
		},
	}...)
	app.Run(os.Args)

	log.Close()
}
