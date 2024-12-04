package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func ModelPath() string {
	return "core/model"
}

func MakeModelComponent(m filesystem.Manager, module, name string) filesystem.File {
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
