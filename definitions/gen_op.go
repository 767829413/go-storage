package definitions

import (
	"fmt"
	"github.com/Xuanwo/gg"
	"github.com/Xuanwo/templateutils"
	log "github.com/sirupsen/logrus"
)

type genOperation struct {
	g *gg.Generator
}

func GenerateOperation(path string) {
	gop := &genOperation{
		g: gg.New(),
	}

	gop.generateHeader()

	for _, ns := range []Namespace{Service{}, Storage{}} {
		gop.generateOperation(ns)
		gop.generateStub(ns)
		gop.generateDefaultPairs(ns)
		gop.generateFeatures(ns)
	}

	err := gop.g.WriteFile(path)
	if err != nil {
		log.Fatalf("generate to %s: %v", path, err)
	}
}

func (gop *genOperation) generateHeader() {
	f := gop.g.NewGroup()

	f.AddLineComment("Code generated by go generate cmd/definitions; DO NOT EDIT.")
	f.AddPackage("types")
	f.NewImport().
		AddPath("context").
		AddPath("io").
		AddPath("net/http").
		AddPath("time")

}

func (gop *genOperation) generateOperation(ns Namespace) {
	f := gop.g.NewGroup()

	nsName := ns.Name()
	nsDisplayName := templateutils.ToPascal(nsName) + "r"
	ops := ns.Operations()

	inter := f.NewInterface(nsDisplayName)
	// Generate String()
	inter.NewFunction("String").AddResult("", "string")
	// Generate Features()
	inter.NewFunction("Features").AddResult("", templateutils.ToPascal(nsName)+"Features")

	for _, op := range ops {
		pname := templateutils.ToPascal(op.Name)

		inter.AddLineComment("%s %s", pname, op.Description)
		gop := inter.NewFunction(pname)

		for _, p := range op.Params {
			gop.AddParameter(p.Name, p.Type.FullName("types"))
		}
		for _, r := range op.Results {
			gop.AddResult(r.Name, r.Type.FullName("types"))
		}

		// We need to generate XxxWithContext functions if not local.
		if !op.Local {
			inter.AddLineComment("%sWithContext %s", pname, op.Description)
			gop := inter.NewFunction(pname + "WithContext")

			// Insert context param.
			gop.AddParameter("ctx", "context.Context")
			for _, p := range op.Params {
				gop.AddParameter(p.Name, p.Type.FullName("types"))
			}
			for _, r := range op.Results {
				gop.AddResult(r.Name, r.Type.FullName("types"))
			}
		}
		// Insert an empty for different functions.
		inter.AddLine()
	}
	inter.NewFunction("mustEmbedUnimplemented" + nsDisplayName)

}
func (gop *genOperation) generateStub(ns Namespace) {
	f := gop.g.NewGroup()
	nsName := ns.Name()

	nsDisplayName := templateutils.ToPascal(nsName) + "r"
	ops := ns.Operations()

	stubName := "Unimplemented" + nsDisplayName
	f.AddLineComment("%s must be embedded to have forward compatible implementations.", stubName)
	f.NewStruct(stubName)

	f.NewFunction("mustEmbedUnimplemented"+nsDisplayName).
		WithReceiver("s", stubName).
		AddBody(gg.Line())
	f.NewFunction("String").
		WithReceiver("s", stubName).
		AddResult("", "string").
		AddBody(
			gg.Return(gg.Lit(stubName)))
	f.NewFunction("Features").
		WithReceiver("s", stubName).
		AddResult("fe", templateutils.ToPascal(nsName)+"Features").
		AddBody(gg.Return())

	for _, op := range ops {
		pname := templateutils.ToPascal(op.Name)

		gop := f.NewFunction(pname).
			WithReceiver("s", stubName)

		for _, p := range op.Params {
			gop.AddParameter(p.Name, p.Type.FullName("types"))
		}
		for _, r := range op.Results {
			gop.AddResult(r.Name, r.Type.FullName("types"))
		}
		// If not local, we need to add and set error
		if !op.Local {
			gop.AddBody(gg.S(`err = NewOperationNotImplementedError("%s")`, op.Name))
		}
		gop.AddBody(gg.Return())

		// We need to generate XxxWithContext functions if not local.
		if !op.Local {
			gop := f.NewFunction(pname+"WithContext").
				WithReceiver("s", stubName)

			// Insert context param.
			gop.AddParameter("ctx", "context.Context")
			for _, p := range op.Params {
				gop.AddParameter(p.Name, p.Type.FullName("types"))
			}
			for _, r := range op.Results {
				gop.AddResult(r.Name, r.Type.FullName("types"))
			}
			gop.AddBody(
				gg.S(`err = NewOperationNotImplementedError("%s")`, op.Name),
				gg.Return(),
			)
		}
	}
}
func (gop *genOperation) generateDefaultPairs(ns Namespace) {
	f := gop.g.NewGroup()

	nsName := ns.Name()
	nsNameP := templateutils.ToPascal(nsName)

	structName := fmt.Sprintf("Default%sPairs", nsNameP)
	f.AddLineComment("%s is the default pairs for %s.", structName, nsNameP)
	sf := f.NewStruct(structName)
	sf.AddLine()
	for _, op := range ns.Operations() {
		sf.AddField(templateutils.ToPascal(op.Name), "[]Pair")
	}
}

func (gop *genOperation) generateFeatures(ns Namespace) {
	f := gop.g.NewGroup()

	nsName := ns.Name()

	structName := templateutils.ToPascal(nsName) + "Features"
	f.AddLineComment("%s indicates features supported by servicer.", structName)
	sf := f.NewStruct(structName)
	sf.AddLine()
	sf.AddLineComment("operation features")
	for _, op := range ns.Operations() {
		sf.AddField(templateutils.ToPascal(op.Name), "bool")
	}
	sf.AddLineComment("operation-related features")
	for _, fe := range SortFeatures(FeaturesArray) {
		if fe.HasNamespace(nsName) {
			sf.AddField(templateutils.ToPascal(fe.Name), "bool")
		}
	}

	f.NewFunction("Has").
		WithReceiver("s", structName).
		AddParameter("name", "string").
		AddResult("", "bool").
		AddBody(gg.Embed(func() gg.Node {
			g := gg.NewGroup()

			s := g.NewSwitch("name")
			for _, op := range ns.Operations() {
				s.NewCase(gg.Lit(op.Name)).
					AddBody(gg.Return(
						gg.S("s.%s", templateutils.ToPascal(op.Name))))
			}
			for _, fe := range SortFeatures(FeaturesArray) {
				if fe.HasNamespace(nsName) {
					s.NewCase(gg.Lit(fe.Name)).
						AddBody(gg.Return(
							gg.S("s.%s", templateutils.ToPascal(fe.Name))))
				}
			}
			s.NewDefault().AddBody("return false")

			return g
		}))
}