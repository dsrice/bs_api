package ug

import (
	"app/infra/genarator"
	"github.com/dave/jennifer/jen"
	"path"
)

func CreateUsecase(cg *genarator.CreateGenerator) error {
	cg.BasePath = "/go/src/app/usecases/"

	err := createCi(cg)
	if err != nil {
		return err
	}

	err = createImp(cg)
	if err != nil {
		return err
	}

	return nil
}

func createCi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ui")

	f.Type().Id(cg.In).Interface()

	f.Save(path.Join(cg.BasePath, "ui", cg.In+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("usecases")

	f.ImportName("app/usecases/ui", "ui")
	f.ImportName("app/repositoriesri", "ri")

	f.Type().Id(cg.Fn).Struct()

	f.Func().Id("New"+cg.In).Params(
		jen.Id("repo").Qual("app/repositories/ri", "InRepository"),
	).Qual("app/usecases/ui", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn).Values()),
	)

	f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}