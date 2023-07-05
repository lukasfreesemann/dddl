package parser

import "golang.org/x/exp/slices"

type Context struct {
	node
	KeywordContext *KeywordContext `@@`
	Name           *Ident          `@@`
	ToDo           *ToDo           `( "{" @@?`
	Definition     *Definition     `@@? `
	Elements       []*TypeDecl     `@@* "}" )?`
}

func ContextOf(root Node) *Context {
	for root != nil {
		if wf, ok := root.(*Context); ok {
			return wf
		}
		root = root.Parent()
	}

	return nil
}

func (n *Context) DeclarationsByName(name string) []Node {
	var res []Node
	for _, element := range n.Elements {
		if element.DataType != nil && element.DataType.Name.Value == name {
			res = append(res, element.DataType)
		}

		if element.Workflow != nil && element.Workflow.Name.Value == name {
			res = append(res, element.Workflow)
		}
	}

	return res
}

func (n *Context) DeclaredName() *Ident {
	return n.Name
}

func (n *Context) DataTypes() []*Data {
	var res []*Data
	for _, element := range n.Elements {
		if element.DataType != nil {
			res = append(res, element.DataType)
		}
	}

	slices.SortFunc(res, func(a, b *Data) bool {
		return a.Name.Value < b.Name.Value
	})

	return res
}

func (n *Context) Workflows() []*Workflow {
	var res []*Workflow
	for _, element := range n.Elements {
		if element.Workflow != nil {
			res = append(res, element.Workflow)
		}
	}

	slices.SortFunc(res, func(a, b *Workflow) bool {
		return a.Name.Value < b.Name.Value
	})

	return res
}

func (n *Context) Children() []Node {
	var res []Node
	res = append(res, n.KeywordContext, n.Name)
	if n.ToDo != nil {
		res = append(res, n.ToDo)
	}

	for _, element := range n.Elements {
		res = append(res, element)
	}

	if n.Definition != nil {
		res = append(res, n.Definition)
	}

	return res
}