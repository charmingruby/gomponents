package unit

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunService(m filesystem.Manager) *cobra.Command {
	var (
		module      string
		serviceName string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseServiceInput(module, serviceName); err != nil {
				panic(err)
			}

			service := atom.MakeService(m, module, serviceName)

			if err := m.GenerateFile(service); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&serviceName, "name", "n", "", "service name")

	return cmd
}

func parseServiceInput(module, serviceName string) error {
	inputs := []input.Arg{
		{
			FieldName:  "module",
			IsRequired: true,
			Value:      module,
			Type:       input.StringType,
		},
		{
			FieldName:  "service name",
			IsRequired: true,
			Value:      serviceName,
			Type:       input.StringType,
		},
	}

	return input.Validate(inputs)
}
