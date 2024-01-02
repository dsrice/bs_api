package cg

import (
	"app/infra/genarator"
	"bufio"
	"fmt"
	"github.com/dave/jennifer/jen"
	"os"
	"path"
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

	err = createImpTest(cg)
	if err != nil {
		return err
	}

	err = addInterface(cg)
	if err != nil {
		return err
	}

	return nil
}

func createCi(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("ci")

	f.Type().Id(cg.In).Interface()

	f.Save(path.Join(cg.BasePath, "ci", cg.In+".go"))

	return nil
}

func createImp(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("controllers")

	f.ImportName("app/controllers/ci", "ci")
	f.ImportName("app/usecases/ui", "ui")

	f.Type().Id(cg.Fn).Struct()

	f.Func().Id("New"+cg.In).Params(
		jen.Id("uc").Qual("app/usecases/ui", "InUsecase"),
	).
		Qual("app/controllers/ci", cg.In).Block(
		jen.Return(jen.Op("&").Id(cg.Fn).Values()),
	)

	f.Save(path.Join(cg.BasePath, cg.Fn+".go"))

	return nil
}

func createImpTest(cg *genarator.CreateGenerator) error {
	f := jen.NewFile("controllers_test")

	f.Save(path.Join(cg.BasePath, cg.Fn+"_test.go"))

	return nil
}

func addInterface(cg *genarator.CreateGenerator) error {
	path := path.Join(cg.BasePath, "ci", "inController.go")

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0775)

	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}

	for scanner.Scan() {
		// ここで一行ずつ処理
		t := scanner.Text()

		if t == "}" {
			add := fmt.Sprintf("	%s %s", cg.In, cg.In)
			lines = append(lines, add)
		}

		lines = append(lines, t)
	}

	f, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0775)

	if err != nil {
		return err
	}
	defer f.Close()

	for _, line := range lines {
		f.WriteString(line + "\n")
	}

	return nil
}