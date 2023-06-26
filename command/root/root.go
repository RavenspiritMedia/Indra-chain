package root

import (
	"fmt"
	"os"

	"github.com/RavenspiritMedia/Indra-chain/command/backup"
	"github.com/RavenspiritMedia/Indra-chain/command/genesis"
	"github.com/RavenspiritMedia/Indra-chain/command/helper"
	"github.com/RavenspiritMedia/Indra-chain/command/ibft"
	"github.com/RavenspiritMedia/Indra-chain/command/license"
	"github.com/RavenspiritMedia/Indra-chain/command/monitor"
	"github.com/RavenspiritMedia/Indra-chain/command/peers"
	"github.com/RavenspiritMedia/Indra-chain/command/secrets"
	"github.com/RavenspiritMedia/Indra-chain/command/server"
	"github.com/RavenspiritMedia/Indra-chain/command/status"
	"github.com/RavenspiritMedia/Indra-chain/command/txpool"
	"github.com/RavenspiritMedia/Indra-chain/command/version"
	"github.com/RavenspiritMedia/Indra-chain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Polygon Edge is a framework for building Ethereum-compatible Blockchain networks",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
