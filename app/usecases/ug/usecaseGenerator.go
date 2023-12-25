package ug

import (
	"app/infra/genarator"
	"fmt"
	"github.com/dave/jennifer/jen"
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

	fmt.Printf("%#v", f)
	//	f.Save(path.Join(cg.BasePath, "ui", cg.Fn+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("usecases")

	f.ImportName("app/usecases/ui", "ui")

	f.Type().Id(cg.Fn + "Imp").Struct()

	f.Func().Id("New"+cg.In).Params().Qual("app/usecases/ci", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn + "Imp").Values()),
	)

	fmt.Printf("%#v", f)

	//f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}