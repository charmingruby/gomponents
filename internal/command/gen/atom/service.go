package atom

import (
	"github.com/charmingruby/bob/internal/command/gen/atom/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func ServicePath() string {
	return "core/service"
}

func MakeServiceComponent(sourceDirectory, module, name string) filesystem.File {
	return component.New(component.ComponentInput{
		Module:               module,
		Name:                 name,
		Suffix:               "service",
		DestinationDirectory: filesystem.ModulePath(sourceDirectory, module, ServicePath()),
	}).Componetize(component.ComponetizeInput{
		TemplateName: constant.SERVICE_TEMPLATE,
		TemplateData: structure.NewDefaultData(name),
		FileName:     name,
		FileSuffix:   "service",
	})
}
