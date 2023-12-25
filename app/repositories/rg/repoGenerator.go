package rg

import (
	"app/infra/genarator"
	"fmt"
	"github.com/dave/jennifer/jen"
)

func CreateRepository(cg *genarator.CreateGenerator) error {
	cg.BasePath = "/go/src/app/repositories/"

	err := createRi(cg)
	if err != nil {
		return err
	}

	err = createImp(cg)
	if err != nil {
		return err
	}

	return nil
}

func createRi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ri")

	f.Type().Id(cg.In).Interface()

	fmt.Printf("%#v", f)

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("repositories")

	f.ImportName("app/infra/database/connection", "connection")
	f.ImportName("app/repositories/ri", "ri")

	f.Type().Id(cg.Fn + "Imp").Struct(
		jen.Id("con").Op("*").Qual("app/infra/database/connection", "Connection"),
	)

	f.Func().Id("New"+cg.In).Params(
		jen.Id("con").Op("*").Qual("app/infra/database/connection", "Connection"),
	).Qual("app/repositories/ri", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn + "Imp").Values(
			jen.Dict{jen.Id("con"): jen.Id("con")}),
		),
	)

	fmt.Printf("%#v", f)

	//f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}