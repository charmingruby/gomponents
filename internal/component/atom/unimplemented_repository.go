package atom

import (
	"github.com/charmingruby/bob/internal/component/atom/constant"
	"github.com/charmingruby/bob/internal/component/atom/data"
	"github.com/charmingruby/bob/internal/component/base"
	"github.com/charmingruby/bob/internal/shared/definition"
	"github.com/charmingruby/bob/internal/shared/filesystem"
)

func MakeUnimplementedRepository(m filesystem.Manager, module, name, database string) filesystem.File {
	prepareDirectoriesForUnimplementedRepository(m, module, database)

	return base.New(base.ComponentInput{
		DestinationDirectory: definition.PersistencePath(m.ModuleDirectory(module), []string{database}),
		Package:              module,
		Name:                 name,
		Suffix:               "repository",
	}).Componetize(
		definition.ADD_COMMAND,
		base.ComponetizeInput{
			TemplateName: constant.REPOSITORY_UNIMPLEMENTED_TEMPLATE,
			TemplateData: data.NewUnimplementedRepositoryData(m.DependencyPath(), module, name, database),
			FileName:     name,
			FileSuffix:   "repository",
		})
}

func prepareDirectoriesForUnimplementedRepository(m filesystem.Manager, module, database string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{definition.PERSISTENCE_PACKAGE, database},
	)
}
