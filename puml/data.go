package puml

import (
	"github.com/worldiety/dddl/parser"
	"github.com/worldiety/dddl/plantuml"
)

func Data(doc *parser.Doc, data *parser.Data) *plantuml.Diagram {
	diag := plantuml.NewDiagram()

	if len(data.ChoiceTypes()) > 0 {
		diag.Add(plantuml.NewInterface(data.Name.Name).NoteRight(plantuml.NewNote(Data2Str(data))))
		for _, choice := range data.ChoiceTypes() {
			choiceName := TypeDeclToStr(choice)
			choiceData := doc.DataByName(choiceName)
			if choiceData == nil {
				diag.Add(plantuml.NewClass(choiceName).Extends(data.Name.Name))
			} else {
				diag.Add(Class(doc, choiceData).Extends(data.Name.Name))
			}
		}

	} else {
		diag.Add(Class(doc, data))
	}

	return diag
}

func Class(doc *parser.Doc, data *parser.Data) *plantuml.Class {
	c := plantuml.NewClass(data.Name.Name)
	for _, declaration := range data.FieldTypes() {
		c.AddAttrs(plantuml.Attr{
			Visibility: plantuml.Public,
			Name:       typeDeclToLinkStr(declaration),
		})
	}

	return c
}

func TypeDeclToStr(decl *parser.TypeDeclaration) string {
	tmp := decl.Name.Name
	if len(decl.Params) > 0 {
		tmp += "<"
		for i, param := range decl.Params {
			tmp += TypeDeclToStr(param)
			if i < len(decl.Params)-1 {
				tmp += ", "
			}
		}
		tmp += ">"
	}

	return tmp
}

func Data2Str(data *parser.Data) string {
	tmp := data.Name.Name + " = \n"
	for i, declaration := range data.ChoiceTypes() {
		tmp += typeDeclToLinkStr(declaration)
		if i < len(data.ChoiceTypes())-1 {
			tmp += "\noder "
		}
	}

	for i, declaration := range data.FieldTypes() {
		tmp += typeDeclToLinkStr(declaration)
		if i < len(data.FieldTypes())-1 {
			tmp += "\nund "
		}
	}

	return tmp
}