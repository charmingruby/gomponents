package core

import (
	"github.com/charmingruby/bob/internal/command/gen/atom"
	"github.com/charmingruby/bob/internal/command/gen/molecule/service"
	"github.com/charmingruby/bob/internal/command/shared/filesystem"
	"github.com/charmingruby/bob/internal/command/shared/opt"
	"github.com/charmingruby/bob/internal/command/shared/scaffold"
)

func MakeCore(m filesystem.Manager, module, database string) {
	prepareDirectoriesForCore(m, module)

	sampleActor := module

	service.MakeService(m, sampleActor, module)

	repository := atom.MakeRepositoryComponent(m, module, sampleActor, opt.POSTGRES_DATABASE_OPTION)
	if err := m.GenerateFile(repository); err != nil {
		panic(err)
	}

	persistenceRepository := atom.MakePersistenceRepositoryComponent(m, module, sampleActor, database)
	if err := m.GenerateFile(persistenceRepository); err != nil {
		panic(err)
	}

	model := atom.MakeModelComponent(m, module, sampleActor)
	if err := m.GenerateFile(model); err != nil {
		panic(err)
	}
}

func prepareDirectoriesForCore(m filesystem.Manager, module string) {
	m.GenerateNestedDirectories(
		m.ModuleDirectory(module),
		[]string{scaffold.CORE_PACKAGE},
	)
}
