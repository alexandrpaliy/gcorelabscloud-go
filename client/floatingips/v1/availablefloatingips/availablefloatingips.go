package availablefloatingips

import (
	"github.com/alexandrpaliy/gcorelabscloud-go/client/floatingips/v1/client"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/utils"
	"github.com/alexandrpaliy/gcorelabscloud-go/gcore/floatingip/v1/floatingips"
	"github.com/urfave/cli/v2"
)

var availableFloatingIPListSubCommand = cli.Command{
	Name:     "list",
	Usage:    "Available floating ips list",
	Category: "availablefloatingip",
	Action: func(c *cli.Context) error {
		client, err := client.NewAvailableFloatingIPClientV1(c)
		if err != nil {
			_ = cli.ShowAppHelp(c)
			return cli.NewExitError(err, 1)
		}
		results, err := floatingips.ListAll(client, nil)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		utils.ShowResults(results, c.String("format"))
		return nil
	},
}

var AvailableFloatingIPCommands = cli.Command{
	Name:  "available",
	Usage: "GCloud available floating ips API",
	Subcommands: []*cli.Command{
		&availableFloatingIPListSubCommand,
	},
}
