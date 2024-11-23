package single

import (
	"fmt"

	"github.com/charmingruby/bob/config"
	"github.com/charmingruby/bob/internal/command/shared/constant"
	"github.com/charmingruby/bob/internal/command/shared/fs"
	"github.com/charmingruby/bob/internal/command/shared/validator"
	"github.com/spf13/cobra"
)

const (
	SERVICE_TEMPLATE = "service"

	DEFAULT_SERVICE_PKG = "service"
)

func RunService(cfg config.Configuration) *cobra.Command {
	var (
		module string
		name   string
		pkg    string
	)

	cmd := &cobra.Command{
		Use:   "service",
		Short: "Generates a new service",
		Run: func(cmd *cobra.Command, args []string) {
			arguments, err := validateServiceArgs(module, name, pkg)
			if err != nil {
				panic(err)
			}

			component := makeServiceComponent(
				cfg.BaseConfiguration.RootDir,
				cfg.BaseConfiguration.SourceDir,
				arguments[0].Value,
				arguments[1].Value,
				arguments[2].Value,
			)

			file := fs.File{
				CommandType:          constant.GENERATE_COMMAND,
				TemplateName:         SERVICE_TEMPLATE,
				TemplateData:         component.Data,
				FileName:             component.Name,
				FileSuffix:           component.Package.Name,
				ResourceName:         component.Name,
				ResourceSuffix:       component.Suffix,
				DestinationDirectory: component.Directory,
				HasTest:              component.HasTest,
			}

			if err := fs.GenerateFile(file); err != nil {
				panic(err)
			}
		},
	}

	cmd.Flags().StringVarP(&module, "module", "m", "", "module name")
	cmd.Flags().StringVarP(&name, "name", "n", "", "service name")
	cmd.Flags().StringVarP(&pkg, "pkg", "p", DEFAULT_SERVICE_PKG, "service package")

	return cmd
}

func makeServiceComponent(rootDir, srcDir, module, name, pkg string) Single {
	component := New(SingleInput{
		Module:      module,
		Name:        name,
		PackageName: pkg,
		Suffix:      pkg,
		HasTest:     false,
	}, WithDefaultTemplate())

	// source_dir/module/core/service/name_service.go
	directory := fmt.Sprintf("%s/%s/%s/core/%s",
		rootDir,
		srcDir,
		module,
		component.Package.Name,
	)

	component.Directory = directory

	return *component
}

func validateServiceArgs(
	module string,
	name string,
	pkg string,
) ([]*validator.Arg, error) {
	args := []*validator.Arg{
		{
			FieldName:  "module",
			Value:      module,
			IsRequired: true,
		},
		{
			FieldName:  "name",
			Value:      name,
			IsRequired: true,
		},
		{
			FieldName: "pkg",
			Value:     pkg,
		},
	}

	if err := validator.ValidateArgsList(args); err != nil {
		return nil, err
	}

	return args, nil
}
