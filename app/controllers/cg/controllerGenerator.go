package cg

import (
	"app/infra/genarator"
	"fmt"
	"github.com/dave/jennifer/jen"
)

func CreateController(cg *genarator.CreateGenerator) error {
	cg.BasePath = "/go/src/app/controllers/"

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
	f := jen.NewFile("ci")

	f.Type().Id(cg.In).Interface()

	fmt.Printf("%#v", f)
	//	f.Save(path.Join(cg.BasePath, "ci", cg.Fn+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("controllers")

	f.ImportName("app/controllers/ci", "ci")

	f.Type().Id(cg.Fn + "Imp").Struct()

	f.Func().Id("New"+cg.In).Params().
		Qual("app/controllers/ci", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn + "Imp").Values(
			jen.Dict{jen.Id("con"): jen.Id("con")}),
		),
	)

	fmt.Printf("%#v", f)

	//f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}