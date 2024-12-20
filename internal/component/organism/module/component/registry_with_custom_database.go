package component

import (
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/component/organism/module/constant"
	"github.com/charmingruby/bob/internal/component/organism/module/data"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeRegistryWithCustomDatabase(m filesystem.Manager, module, repositoryModel, database string) filesystem.File {
	prepareDirectoriesForRegistryWithCustomDatabase(m, module)

	path := m.SourceDirectory + "/" + module

	return base.New(base.ComponentInput{
		Package:              module,
		DestinationDirectory: path,
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.MODULE_WITH_CUSTOM_DATABASE_TEMPLATE,
			TemplateData: data.NewModuleWithDatabaseData(
				m.DependencyPath(),
				module,
				database,
				repositoryModel,
			),
			FileName: module,
		})
}

func prepareDirectoriesForRegistryWithCustomDatabase(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.SourceDirectory,
		[]string{module},
	)
}
