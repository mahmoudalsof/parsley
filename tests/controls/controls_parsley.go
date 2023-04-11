// Code generated by parsley for scanning JSON strings. DO NOT EDIT.
package controls

import (
	parse "github.com/Soreing/parsley"
	reader "github.com/Soreing/parsley/reader"
	writer "github.com/Soreing/parsley/writer"
)

var _ *reader.Reader
var _ *writer.Writer

func (o *EmptyObject) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	var key []byte
	_ = key
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					err = r.Skip()
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *EmptyObject) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []EmptyObject, err error) {
	var e EmptyObject
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]EmptyObject, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *EmptyObject) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []EmptyObject, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *EmptyObject) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	} else {
		w.Byte('{')
		w.Byte('}')
	}
}

func (o *EmptyObject) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []EmptyObject) {
	if slc == nil {
		w.Raw("null")
	} else if len(slc) == 0 {
		w.Raw("[]")
	} else {
		w.Byte('[')
		slc[0].EncodeObjectPJSON(w, filter)
		for i := 1; i < len(slc); i++ {
			w.Byte(',')
			slc[i].EncodeObjectPJSON(w, filter)
		}
		w.Byte(']')
	}
}

func (o *EmptyObject) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	} else {
		if bytes == 0 {
			return 2, 0
		} else {
			return bytes + 1, volatile
		}
	}
}

func (o *EmptyObject) SliceLengthPJSON(filter []parse.Filter, slc []EmptyObject) (bytes int, volatile int) {
	for _, obj := range slc {
		b, v := obj.ObjectLengthPJSON(filter)
		bytes, volatile = bytes+b+1, volatile+v
	}
	if bytes == 0 {
		return 2, 0
	} else {
		return bytes + 1, volatile
	}
}

func (o *EscapedField) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	c := [1]bool{}
	if filter == nil {
		for i := range c {
			c[i] = true
		}
	} else {
		for i := range filter {
			k := filter[i].Field
			if k == "soɯə \"value\"" {
				c[0] = true
			}
		}
	}
	var key []byte
	_ = key
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					if string(key) == "soɯə \"value\"" && c[0] {
						o.Value, err = r.GetString()
					} else {
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *EscapedField) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []EscapedField, err error) {
	var e EscapedField
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]EscapedField, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *EscapedField) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []EscapedField, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *EscapedField) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	} else {
		c := [1]bool{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "soɯə \"value\"" {
					c[0] = true
				}
			}
		}
		w.Byte('{')
		off := 1
		if c[0] {
			w.Raw(",\"soɯə \\\"value\\\"\":"[off:])
			w.String(o.Value)
			off = 0
		}
		w.Byte('}')
	}
}

func (o *EscapedField) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []EscapedField) {
	if slc == nil {
		w.Raw("null")
	} else if len(slc) == 0 {
		w.Raw("[]")
	} else {
		w.Byte('[')
		slc[0].EncodeObjectPJSON(w, filter)
		for i := 1; i < len(slc); i++ {
			w.Byte(',')
			slc[i].EncodeObjectPJSON(w, filter)
		}
		w.Byte(']')
	}
}

func (o *EscapedField) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	} else {
		c := [1]bool{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "soɯə \"value\"" {
					c[0] = true
				}
			}
		}
		if c[0] {
			b, v := writer.StringLen(o.Value)
			bytes, volatile = bytes+b+20, volatile+v
		}
		if bytes == 0 {
			return 2, 0
		} else {
			return bytes + 1, volatile
		}
	}
}

func (o *EscapedField) SliceLengthPJSON(filter []parse.Filter, slc []EscapedField) (bytes int, volatile int) {
	for _, obj := range slc {
		b, v := obj.ObjectLengthPJSON(filter)
		bytes, volatile = bytes+b+1, volatile+v
	}
	if bytes == 0 {
		return 2, 0
	} else {
		return bytes + 1, volatile
	}
}

func (o *PrivateField) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	var key []byte
	_ = key
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					err = r.Skip()
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *PrivateField) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []PrivateField, err error) {
	var e PrivateField
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]PrivateField, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *PrivateField) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []PrivateField, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *PrivateField) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	} else {
		w.Byte('{')
		w.Byte('}')
	}
}

func (o *PrivateField) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []PrivateField) {
	if slc == nil {
		w.Raw("null")
	} else if len(slc) == 0 {
		w.Raw("[]")
	} else {
		w.Byte('[')
		slc[0].EncodeObjectPJSON(w, filter)
		for i := 1; i < len(slc); i++ {
			w.Byte(',')
			slc[i].EncodeObjectPJSON(w, filter)
		}
		w.Byte(']')
	}
}

func (o *PrivateField) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	} else {
		if bytes == 0 {
			return 2, 0
		} else {
			return bytes + 1, volatile
		}
	}
}

func (o *PrivateField) SliceLengthPJSON(filter []parse.Filter, slc []PrivateField) (bytes int, volatile int) {
	for _, obj := range slc {
		b, v := obj.ObjectLengthPJSON(filter)
		bytes, volatile = bytes+b+1, volatile+v
	}
	if bytes == 0 {
		return 2, 0
	} else {
		return bytes + 1, volatile
	}
}

func (o *PublicField) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	c := [1]bool{}
	if filter == nil {
		for i := range c {
			c[i] = true
		}
	} else {
		for i := range filter {
			k := filter[i].Field
			if k == "field" {
				c[0] = true
			}
		}
	}
	var key []byte
	_ = key
	err = r.OpenObject()
	if r.GetType() != reader.TerminatorToken {
		for err == nil {
			if key, err = r.GetKey(); err == nil {
				if r.IsNull() {
					r.SkipNull()
				} else {
					if string(key) == "field" && c[0] {
						o.field, err = r.GetString()
					} else {
						err = r.Skip()
					}
				}
				if err == nil && !r.Next() {
					break
				}
			}
		}
	}
	if err == nil {
		err = r.CloseObject()
	}
	return
}

func (o *PublicField) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []PublicField, err error) {
	var e PublicField
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]PublicField, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *PublicField) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []PublicField, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *PublicField) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	} else {
		c := [1]bool{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "field" {
					c[0] = true
				}
			}
		}
		w.Byte('{')
		off := 1
		if c[0] {
			w.Raw(",\"field\":"[off:])
			w.String(o.field)
			off = 0
		}
		w.Byte('}')
	}
}

func (o *PublicField) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []PublicField) {
	if slc == nil {
		w.Raw("null")
	} else if len(slc) == 0 {
		w.Raw("[]")
	} else {
		w.Byte('[')
		slc[0].EncodeObjectPJSON(w, filter)
		for i := 1; i < len(slc); i++ {
			w.Byte(',')
			slc[i].EncodeObjectPJSON(w, filter)
		}
		w.Byte(']')
	}
}

func (o *PublicField) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	} else {
		c := [1]bool{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "field" {
					c[0] = true
				}
			}
		}
		if c[0] {
			b, v := writer.StringLen(o.field)
			bytes, volatile = bytes+b+9, volatile+v
		}
		if bytes == 0 {
			return 2, 0
		} else {
			return bytes + 1, volatile
		}
	}
}

func (o *PublicField) SliceLengthPJSON(filter []parse.Filter, slc []PublicField) (bytes int, volatile int) {
	for _, obj := range slc {
		b, v := obj.ObjectLengthPJSON(filter)
		bytes, volatile = bytes+b+1, volatile+v
	}
	if bytes == 0 {
		return 2, 0
	} else {
		return bytes + 1, volatile
	}
}
