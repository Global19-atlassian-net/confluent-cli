package cluster

import (
	"github.com/spf13/cobra"

	pcmd "github.com/confluentinc/cli/internal/pkg/cmd"
)

type command struct {
	*pcmd.StateFlagCommand
	prerunner  pcmd.PreRunner
	metaClient Metadata
}

// New returns the Cobra command for `cluster`.
func New(prerunner pcmd.PreRunner, metaClient Metadata) *cobra.Command {
	cmd := &command{
		StateFlagCommand: pcmd.NewAnonymousStateFlagCommand(&cobra.Command{
			Use:   "cluster",
			Short: "Retrieve metadata about Confluent Platform clusters.",
		}, prerunner, SubcommandFlags),
		prerunner:  prerunner,
		metaClient: metaClient,
	}
	cmd.init()
	return cmd.Command
}

func (c *command) init() {
	c.AddCommand(NewDescribeCommand(c.prerunner, c.metaClient))
	c.AddCommand(NewListCommand(c.prerunner))
	c.AddCommand(NewRegisterCommand(c.prerunner))
	c.AddCommand(NewUnregisterCommand(c.prerunner))
}
