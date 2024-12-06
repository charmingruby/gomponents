package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/constant"
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeModelComponent(m filesystem.Manager, module, name string) filesystem.File {
	prepareDirectoriesForModel(m, module)

	return component.New(component.ComponentInput{
		DestinationDirectory: scaffold.CorePath(m.ModuleDirectory(module), []string{scaffold.MODEL_PACKAGE}),
		Package:              module,
		Name:                 name,
		HasTest:              true,
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.MODEL_TEMPLATE,
		TemplateData: structure.NewDefaultData(name),
		FileName:     name,
	})
}

func prepareDirectoriesForModel(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE, scaffold.MODEL_PACKAGE},
	)
}
