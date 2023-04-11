package main

// Writes the header for the source file
func (g *Generator) WriteHeader() {
	g.buf.WriteString("// Code generated by parsley for scanning JSON strings. DO NOT EDIT.\n")
	g.buf.WriteString("package " + g.pkgName + "\n\n")
}

// Writes imports for the source file
func (g *Generator) WriteImports(pkgs map[string]string) {
	g.buf.WriteString("import (\n")
	for name, path := range pkgs {
		g.buf.WriteString(name + path + "\n")
	}
	g.buf.WriteString(")\n\n")
	g.buf.WriteString("var _ *reader.Reader\n")
	g.buf.WriteString("var _ *writer.Writer\n\n")
}

// Implements functions for a define and writes it to the buffer
func (g *Generator) WriteDefine(df Define) {
	name := df.name
	di := NewDefineInfo(df)

	code := "func (o *" + name + ")DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {\n" +
		createDecodeDefineBody(di) +
		"    return\n" +
		"}\n\n" +
		"func (o *" + name + ")sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []" + name + ", err error) {\n" +
		"    var e " + name + "\n" +
		"    if err = e.DecodeObjectPJSON(r, filter); err == nil {\n" +
		"        if !r.Next() {\n" +
		"            res = make([]" + name + ", idx+1)\n" +
		"            res[idx] = e\n" +
		"            return\n" +
		"        } else if res, err = o.sequencePJSON(r, filter, idx + 1); err == nil {\n" +
		"            res[idx] = e\n" +
		"        }\n" +
		"    }\n" +
		"    return\n" +
		"}\n\n" +
		"func (o *" + name + ") DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []" + name + ", err error) {\n" +
		"    if err = r.OpenArray(); err == nil {\n" +
		"        if res, err = o.sequencePJSON(r, filter, 0); err == nil {\n" +
		"            err = r.CloseArray()\n" +
		"        }\n" +
		"    }\n" +
		"    return\n" +
		"}\n\n" +
		"func (o *" + name + ") EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {\n" +
		"    if o == nil {\n" +
		"        w.Raw(\"null\")\n" +
		"    }\n" +
		"    " + createEncodeDefineBody(di) + "\n" +
		"}\n\n" +
		"func (o *" + name + ") EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []" + name + ") {\n" +
		"    if slc == nil {\n" +
		"        w.Raw(\"null\")\n" +
		"    }\n" +
		"    w.Byte('[')\n" +
		"    if len(slc) > 0 {\n" +
		"        slc[0].EncodeObjectPJSON(w, filter)\n" +
		"        for i:=1; i<len(slc); i++ {\n" +
		"            w.Byte(',')\n" +
		"            slc[i].EncodeObjectPJSON(w, filter)\n" +
		"        }\n" +
		"    }\n" +
		"    w.Byte(']')\n" +
		"}\n\n" +
		"func (o *" + name + ") ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {\n" +
		"    if o == nil {\n" +
		"        return 4, 0\n" +
		"    }\n" +
		"    return " + createDefineLengthBody(di) + "\n" +
		"}\n\n" +
		"func (o *" + name + ") SliceLengthPJSON(filter []parse.Filter, slc []" + name + ") (bytes int, volatile int) {\n" +
		"    for _, obj := range slc {\n" +
		"        b, v := obj.ObjectLengthPJSON(filter)\n" +
		"        bytes += b+1\n" +
		"        volatile += v\n" +
		"    }\n" +
		"    if bytes == 0 {\n" +
		"        return 2, 0\n" +
		"    } else {\n" +
		"        return bytes+1, volatile\n" +
		"    }\n" +
		"}\n\n"

	g.buf.WriteString(code)
}

func (g *Generator) WriteStruct(st Struct) {
	name := st.name
	fis := []FieldInfo{}
	for _, f := range st.fields {
		fis = append(fis, NewFieldInfo(f, g.defaultCase))
	}

	code := "func (o *" + name + ")DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {\n" +
		createFilterHeader(fis) +
		"    var key []byte\n" +
		"    _ = key\n" +
		"    err = r.OpenObject()\n" +
		"    if r.GetType() != reader.TerminatorToken {\n" +
		"        for err == nil {\n" +
		"            if key, err = r.GetKey(); err == nil {\n" +
		"                if r.IsNull() {\n" +
		"                    r.SkipNull()\n" +
		"                } else {\n" +
		createDecodeObjectBody(fis) +
		"                }\n" +
		"                if err == nil && !r.Next() {\n" +
		"                    break\n" +
		"                }\n" +
		"            }\n" +
		"        }\n" +
		"    }\n" +
		"    if err == nil {\n" +
		"        err = r.CloseObject()\n" +
		"    }\n" +
		"    return\n" +
		"}\n\n" +
		"func (o *" + name + ")sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []" + name + ", err error) {\n" +
		"    var e " + name + "\n" +
		"    if err = e.DecodeObjectPJSON(r, filter); err == nil {\n" +
		"        if !r.Next() {\n" +
		"            res = make([]" + name + ", idx+1)\n" +
		"            res[idx] = e\n" +
		"            return\n" +
		"        } else if res, err = o.sequencePJSON(r, filter, idx + 1); err == nil {\n" +
		"            res[idx] = e\n" +
		"        }\n" +
		"    }\n" +
		"    return\n" +
		"}\n\n" +
		"func (o *" + name + ") DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []" + name + ", err error) {\n" +
		"    if err = r.OpenArray(); err == nil {\n" +
		"        if res, err = o.sequencePJSON(r, filter, 0); err == nil {\n" +
		"            err = r.CloseArray()\n" +
		"        }\n" +
		"    }\n" +
		"    return\n" +
		"}\n\n" +
		"func (o *" + name + ") EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {\n" +
		"    if o == nil {\n" +
		"        w.Raw(\"null\")\n" +
		"    } else {\n" +
		createFilterHeader(fis) +
		"        w.Byte('{')\n" +
		createEncodeObjectBody(fis) +
		"        w.Byte('}')\n" +
		"    }\n" +
		"}\n\n" +
		"func (o *" + name + ") EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []" + name + ") {\n" +
		"    if slc == nil {\n" +
		"        w.Raw(\"null\")\n" +
		"    } else if len(slc) == 0 {\n" +
		"        w.Raw(\"[]\")\n" +
		"    } else {\n" +
		"        w.Byte('[')\n" +
		"        slc[0].EncodeObjectPJSON(w, filter)\n" +
		"        for i:=1; i<len(slc); i++  {\n" +
		"            w.Byte(',')\n" +
		"            slc[i].EncodeObjectPJSON(w, filter)\n" +
		"        }\n" +
		"        w.Byte(']')\n" +
		"    }\n" +
		"}\n\n" +
		"func (o *" + name + ") ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {\n" +
		"    if o == nil {\n" +
		"        return 4, 0\n" +
		"    } else {\n" +
		createFilterHeader(fis) +
		createObjectLengthBody(fis) +
		"        if bytes == 0 {\n" +
		"            return 2, 0\n" +
		"        } else {\n" +
		"            return bytes+1, volatile\n" +
		"        }\n" +
		"    }\n" +
		"}\n\n" +
		"func (o *" + name + ") SliceLengthPJSON(filter []parse.Filter, slc []" + name + ") (bytes int, volatile int) {\n" +
		"    for _, obj := range slc {\n" +
		"        b, v := obj.ObjectLengthPJSON(filter)\n" +
		"        bytes, volatile = bytes+b+1, volatile+v\n" +
		"    }\n" +
		"    if bytes == 0 {\n" +
		"        return 2, 0\n" +
		"    } else {\n" +
		"        return bytes+1, volatile\n" +
		"    }\n" +
		"}\n\n"

	g.buf.WriteString(code)
}
