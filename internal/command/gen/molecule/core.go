package molecule

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/spf13/cobra"
)

func RunCore(m component.Manager) *cobra.Command {
	var (
		module string
	)

	cmd := &cobra.Command{
		Use:   "core",
		Short: "Generates a new core molecule",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateOnlyModuleCommandInput(module); err != nil {
				panic(err)
			}

			MakeCore(m, module)
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")

	return cmd
}

func MakeCore(m component.Manager, module string) {
	if err := filesystem.GenerateDirectory(m.ModuleDirectory(module), "core"); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateMultipleDirectories(
		component.ModulePath(m.SourceDirectory, module, CorePath()),
		[]string{"service", "model", "repository"},
	); err != nil {
		panic(err)
	}

	sampleActor := module

	service.MakeService(m, sampleActor, module)

	if err := filesystem.GenerateFile(atom.MakeRepositoryComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}

	if err := filesystem.GenerateFile(atom.MakeModelComponent(m, module, sampleActor)); err != nil {
		panic(err)
	}
}

func CorePath() string {
	return "core"
}