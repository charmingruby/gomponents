package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/component/input"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/spf13/cobra"
)

func ModelPath() string {
	return "core/model"
}

func RunModel(m component.Manager) *cobra.Command {
	var (
		module string
		name   string
	)

	cmd := &cobra.Command{
		Use:   "model",
		Short: "Generates a new model",
		Run: func(cmd *cobra.Command, args []string) {
			if err := input.ValidateDefaultCommandInput(module, name); err != nil {
				panic(err)
			}

			if err := filesystem.GenerateFile(MakeModelComponent(
				m,
				module,
				name,
			)); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "model name")

	return cmd
}

func MakeModelComponent(m component.Manager, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		DestinationDirectory: m.AppendToModuleDirectory(module, ModelPath()),
		Module:               module,
		Name:                 name,
		HasTest:              true,
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.MODEL_TEMPLATE,
		TemplateData: structure.NewDefaultData(name),
		FileName:     name,
	})
}
