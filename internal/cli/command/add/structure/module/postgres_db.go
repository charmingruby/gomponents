package module

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/organism"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunPostgresDB(m filesystem.Manager) *cobra.Command {
	var (
		module    string
		modelName string
		tableName string
	)

	cmd := &cobra.Command{
		Use:   "postgres-db",
		Short: "Generates a module with postgres database",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parsePostgresDBInput(module, modelName, tableName); err != nil {
				panic(err)
			}

			organism.PerformModuleWithPostgresDatabase(m, module, modelName, tableName)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "modelName", "n", "", "base model name")
	cmd.Flags().StringVarP(&tableName, "tableName", "t", "", "base table name")

	return cmd
}

func parsePostgresDBInput(module, model, tableName string) error {
	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "model name",
			Value:      model,
			IsRequired: true,
		},
		{
			FieldName:  "table name",
			Value:      tableName,
			IsRequired: true,
		},
	}

	return input.Validate(args)
}
