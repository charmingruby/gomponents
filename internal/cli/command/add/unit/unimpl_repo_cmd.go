package unit

import (
	"github.com/charmingruby/bob/internal/cli/input"
	"github.com/charmingruby/bob/internal/component/atom"
	"github.com/charmingruby/bob/internal/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunUnimplRepo(m filesystem.Manager) *cobra.Command {
	var (
		module       string
		modelName    string
		databaseName string
	)

	cmd := &cobra.Command{
		Use:   "unimpl-repo",
		Short: "Generates a new unimplemented repository",
		Run: func(cmd *cobra.Command, args []string) {
			if err := parseUnimplRepoInput(module, modelName, databaseName); err != nil {
				panic(err)
			}

			repository := atom.MakeUnimplementedRepository(m, module, modelName, databaseName)

			if err := m.GenerateFile(repository); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module")
	cmd.Flags().StringVarP(&modelName, "model", "n", "", "model to be managed by the repository")
	cmd.Flags().StringVarP(&databaseName, "database", "d", "", "database that will implement the repository")

	return cmd
}

func parseUnimplRepoInput(module, modelName, databaseName string) error {
	args := []input.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "model name",
			Value:      modelName,
			IsRequired: true,
		},
		{
			FieldName:  "database",
			Value:      databaseName,
			IsRequired: true,
		},
	}

	return input.Validate(args)
}
