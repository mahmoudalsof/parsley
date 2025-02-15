package main

import (
	"errors"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Import struct {
	name string
	path string
}

type Type struct {
	arr bool
	ptr bool
	pkg string
	typ string
}

type Struct struct {
	name   string
	fields []Field
}

type Field struct {
	name string
	tag  string
	typ  Type
}

type Define struct {
	name string
	typ  Type
}

func newImport(at *ast.ImportSpec) Import {
	n, p := "", at.Path.Value
	if at.Name != nil {
		n = at.Name.Name
	} else {
		toks := strings.Split(strings.Trim(at.Path.Value, "\""), "/")
		n = toks[len(toks)-1]
	}
	return Import{
		name: n,
		path: p,
	}
}

func getType(e ast.Expr) (t Type, err error) {
	if arr, ok := e.(*ast.ArrayType); ok {
		t.arr = true
		e = arr.Elt
	} else if ptr, ok := e.(*ast.StarExpr); ok {
		t.ptr = true
		e = ptr.X
	}

	switch st := e.(type) {
	case *ast.Ident:
		t.typ = st.Name
	case *ast.SelectorExpr:
		t.typ = st.Sel.Name
		if pkg, ok := st.X.(*ast.Ident); ok {
			t.pkg = pkg.Name
		} else {
			err = errors.New("unsuported type")
		}
	default:
		err = errors.New("unsupported type")
	}
	return
}

func newField(f *ast.Field) (Field, error) {
	if len(f.Names) == 0 || f.Names[0].Name == "" {
		return Field{}, errors.New("field has no name")
	}

	ts, err := getType(f.Type)
	if err != nil {
		return Field{}, err
	}

	tag := ""
	if f.Tag != nil {
		tag = f.Tag.Value
	}
	return Field{
		name: f.Names[0].Name,
		tag:  tag,
		typ:  ts,
	}, nil
}

func newStruct(name string, public bool, st *ast.StructType) Struct {
	flds := []Field{}
	for _, f := range st.Fields.List {
		if fl, err := newField(f); err == nil {
			if ast.IsExported(fl.name) || public {
				flds = append(flds, fl)
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return Struct{
		name:   name,
		fields: flds,
	}
}

type Parser struct {
	PkgDir    string
	PkgName   string
	Structs   []Struct
	Defines   []Define
	Imports   []Import
	AllTypes  bool
	AllPublic bool
}

type visitor struct {
	*Parser

	skip bool
	json bool
	publ bool
}

const (
	optionTag = "parsley:"
	jsonTag   = "json"
	skipTag   = "skip"
	publicTag = "public"
)

func (v *visitor) handleComment(comments *ast.CommentGroup) {
	if comments == nil {
		return
	}

	for _, c := range comments.List {
		comment := c.Text
		if len(comment) < 3 {
			return
		}

		switch comment[1] {
		case '/':
			comment = comment[2:]
		case '*':
			comment = comment[2 : len(comment)-2]
		}

		for _, comment := range strings.Split(comment, "\n") {
			comment = strings.TrimSpace(comment)
			if strings.HasPrefix(comment, optionTag) {
				opts := strings.TrimPrefix(comment, optionTag)
				toks := strings.Split(opts, ",")
				for _, e := range toks {
					switch e {
					case jsonTag:
						v.json = true
					case skipTag:
						v.skip = true
					case publicTag:
						v.publ = true
					}
				}
			}
		}
	}
}

func (v *visitor) Visit(n ast.Node) (w ast.Visitor) {
	switch n := n.(type) {
	case *ast.Package:
		return v

	case *ast.ImportSpec:
		imp := newImport(n)
		for _, e := range v.Imports {
			if e.name == imp.name {
				if e.path != imp.path {
					panic("conflicting import names")
				} else {
					return v
				}
			}
		}
		v.Imports = append(v.Imports, newImport(n))

	case *ast.GenDecl:
		return v

	case *ast.File:
		v.PkgName = n.Name.String()
		return v

	case *ast.CommentGroup:
		v.handleComment(n)
		return v

	case *ast.TypeSpec:
		if !v.skip && (v.json || v.AllTypes) {
			name := n.Name.String()
			switch t := n.Type.(type) {
			case *ast.StructType:
				st := newStruct(name, v.publ || v.AllPublic, t)
				v.Structs = append(v.Structs, st)
			case *ast.Ident, *ast.SelectorExpr, *ast.ArrayType:
				if ts, err := getType(n.Type); err == nil {
					v.Defines = append(v.Defines, Define{
						name: name,
						typ:  ts,
					})
				}
			}
		}

		v.skip, v.json, v.publ = false, false, false
	}

	return nil
}

func (p *Parser) Parse(fname string, isDir bool) (err error) {
	info, err := os.Stat(fname)
	if err != nil {
		log.Fatal(err)
	}
	if info.IsDir() {
		p.PkgDir = fname
	} else {
		p.PkgDir = filepath.Dir(fname)
	}

	fset := token.NewFileSet()
	if isDir {
		packages, err := parser.ParseDir(
			fset,
			fname,
			excludeTestFiles,
			parser.ParseComments,
		)
		if err != nil {
			return err
		}
		for _, pckg := range packages {
			ast.Walk(&visitor{Parser: p}, pckg)
		}
	} else {
		f, err := parser.ParseFile(
			fset,
			fname,
			nil,
			parser.ParseComments,
		)
		if err != nil {
			return err
		}
		ast.Walk(&visitor{Parser: p}, f)
	}

	// Sort structs for consistent code generation
	sort.Slice(p.Structs, func(i, j int) bool {
		return p.Structs[i].name < p.Structs[j].name
	})
	sort.Slice(p.Defines, func(i, j int) bool {
		return p.Defines[i].name < p.Defines[j].name
	})
	sort.Slice(p.Imports, func(i, j int) bool {
		return p.Imports[i].name < p.Imports[j].name
	})

	return nil
}

func excludeTestFiles(fi os.FileInfo) bool {
	return !strings.HasSuffix(fi.Name(), "_test.go")
}
