package networks

import (
	"github.com/alexandrpaliy/gcorelabscloud-go/gcore/network/v1/extensions"

	"github.com/alexandrpaliy/gcorelabscloud-go/client/flags"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/networks/v1/client"
	"github.com/alexandrpaliy/gcorelabscloud-go/client/utils"

	"github.com/urfave/cli/v2"
)

var aliasText = "alias is mandatory argument"

var extensionListCommand = cli.Command{
	Name:     "list",
	Usage:    "List extensions",
	Category: "extension",
	Action: func(c *cli.Context) error {
		client, err := client.NewNetworkClientV1(c)
		if err != nil {
			_ = cli.ShowAppHelp(c)
			return cli.NewExitError(err, 1)
		}
		results, err := extensions.ListAll(client)
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		utils.ShowResults(results, c.String("format"))
		return nil
	},
}

var extensionGetCommand = cli.Command{
	Name:      "show",
	Usage:     "Get extension information",
	ArgsUsage: "<alias>",
	Category:  "extension",
	Action: func(c *cli.Context) error {
		extensionID, err := flags.GetFirstStringArg(c, aliasText)
		if err != nil {
			_ = cli.ShowCommandHelp(c, "show")
			return err
		}
		client, err := client.NewNetworkClientV1(c)
		if err != nil {
			_ = cli.ShowAppHelp(c)
			return cli.NewExitError(err, 1)
		}
		ext, err := extensions.Get(client, extensionID).Extract()
		if err != nil {
			return cli.NewExitError(err, 1)
		}
		utils.ShowResults(ext, c.String("format"))
		return nil
	},
}

var extensionCommands = cli.Command{
	Name:  "extension",
	Usage: "GCloud neutron extensions API",
	Subcommands: []*cli.Command{
		&extensionListCommand,
		&extensionGetCommand,
	},
}
