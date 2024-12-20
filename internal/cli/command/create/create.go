package create

import (
	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/shared/filesystem"

	"github.com/spf13/cobra"
)

type Command struct {
	cmd *cobra.Command
	fs  filesystem.Manager
}

func New(cmd *cobra.Command, config config.Configuration) *Command {
	return &Command{
		cmd: cmd,
		fs:  filesystem.New(config),
	}
}

func (c *Command) Setup() {
	var goVersion string

	cmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a new project",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseCreateInput(goVersion); err != nil {
				panic(err)
			}

			organism.PerformSetup(c.fs, goVersion)
		},
	}

	cmd.Flags().StringVarP(&goVersion, "golang version", "v", "1.23.3", "golang version for setup, by default, it will be 1.23.3")

	c.cmd.AddCommand(cmd)
}

func parseCreateInput(goVersion string) error {
	args := []input.Arg{
		{
			FieldName: "go version",
			Value:     goVersion,
		},
	}

	return input.Validate(args)
}
