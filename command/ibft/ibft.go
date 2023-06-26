package ibft

import (
	"github.com/RavenspiritMedia/Indra-chain/command/helper"
	"github.com/RavenspiritMedia/Indra-chain/command/ibft/candidates"
	"github.com/RavenspiritMedia/Indra-chain/command/ibft/propose"
	"github.com/RavenspiritMedia/Indra-chain/command/ibft/quorum"
	"github.com/RavenspiritMedia/Indra-chain/command/ibft/snapshot"
	"github.com/RavenspiritMedia/Indra-chain/command/ibft/status"
	_switch "github.com/RavenspiritMedia/Indra-chain/command/ibft/switch"
	"github.com/spf13/cobra"
)

func GetCommand() *cobra.Command {
	ibftCmd := &cobra.Command{
		Use:   "ibft",
		Short: "Top level IBFT command for interacting with the IBFT consensus. Only accepts subcommands.",
	}

	helper.RegisterGRPCAddressFlag(ibftCmd)

	registerSubcommands(ibftCmd)

	return ibftCmd
}

func registerSubcommands(baseCmd *cobra.Command) {
	baseCmd.AddCommand(
		// ibft status
		status.GetCommand(),
		// ibft snapshot
		snapshot.GetCommand(),
		// ibft propose
		propose.GetCommand(),
		// ibft candidates
		candidates.GetCommand(),
		// ibft switch
		_switch.GetCommand(),
		// ibft quorum
		quorum.GetCommand(),
	)
}
