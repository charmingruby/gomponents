package rest_component

import (
	restConstant "github.com/charmingruby/bob/internal/command/gen/molecule/rest/constant"
	"github.com/charmingruby/bob/internal/command/gen/molecule/rest/structure"
	"github.com/charmingruby/bob/internal/command/shared/component"
	"github.com/charmingruby/bob/internal/command/shared/component/constant"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
)

func MakeRequestHelperComponent(m filesystem.Manager) filesystem.File {
	pkg := "rest"

	if err := m.GenerateNestedDirectories(
		m.ModuleDirectory(constant.SHARED_MODULE),
		[]string{"transport", pkg},
	); err != nil {
		panic(err)
	}

	return component.New(component.ComponentInput{
		Module: pkg,
		DestinationDirectory: m.AppendToModuleDirectory(
			constant.SHARED_MODULE,
			"transport/rest",
		),
	}).Componetize(component.ComponetizeInput{
		TemplateName: restConstant.REST_REQUEST_HELPER_TEMPLATE,
		TemplateData: structure.NewRequestHelperData(m.DependencyPath()),
		FileName:     "request",
	})
}
