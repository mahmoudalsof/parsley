// Code generated by parsley for scanning JSON strings. DO NOT EDIT.
package models

import (
	parse "github.com/Soreing/parsley"
	reader "github.com/Soreing/parsley/reader"
	externals "github.com/Soreing/parsley/tests/externals"
	writer "github.com/Soreing/parsley/writer"
)

var _ *reader.Reader
var _ *writer.Writer

func (o *Employee) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	c := [7]bool{}
	f := [7][]parse.Filter{}
	if filter == nil {
		for i := range c {
			c[i] = true
		}
	} else {
		for i := range filter {
			k := filter[i].Field
			if k == "id" {
				c[0] = true
			} else if k == "person" {
				c[1], f[1] = true, filter[i].Filter
			} else if k == "devices" {
				c[2], f[2] = true, filter[i].Filter
			} else if k == "isActive" {
				c[3] = true
			} else if k == "rating" {
				c[4] = true
			} else if k == "lineManager" {
				c[5], f[5] = true, filter[i].Filter
			} else if k == "tags" {
				c[6] = true
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
					if string(key) == "id" && c[0] {
						o.Id, err = r.GetString()
					} else if string(key) == "person" && c[1] {
						err = o.Person.DecodeObjectPJSON(r, f[1])
					} else if string(key) == "devices" && c[2] {
						o.Devices, err = (*externals.Device)(nil).DecodeSlicePJSON(r, f[2])
					} else if string(key) == "isActive" && c[3] {
						o.IsActive, err = r.GetBool()
					} else if string(key) == "rating" && c[4] {
						o.Rating, err = r.GetFloat64()
					} else if string(key) == "lineManager" && c[5] {
						o.LineManager = &Employee{}
						err = o.LineManager.DecodeObjectPJSON(r, f[5])
					} else if string(key) == "tags" && c[6] {
						o.Tags, err = r.GetStrings()
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

func (o *Employee) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []Employee, err error) {
	var e Employee
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]Employee, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *Employee) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []Employee, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *Employee) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	} else {
		c := [7]bool{}
		f := [7][]parse.Filter{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "id" {
					c[0] = true
				} else if k == "person" {
					c[1], f[1] = true, filter[i].Filter
				} else if k == "devices" {
					c[2], f[2] = true, filter[i].Filter
				} else if k == "isActive" {
					c[3] = true
				} else if k == "rating" {
					c[4] = true
				} else if k == "lineManager" {
					c[5], f[5] = true, filter[i].Filter
				} else if k == "tags" {
					c[6] = true
				}
			}
		}
		w.Byte('{')
		off := 1
		if c[0] {
			w.Raw(",\"id\":"[off:])
			w.String(o.Id)
			off = 0
		}
		if c[1] {
			w.Raw(",\"person\":"[off:])
			o.Person.EncodeObjectPJSON(w, f[1])
			off = 0
		}
		if c[2] {
			w.Raw(",\"devices\":"[off:])
			(*externals.Device)(nil).EncodeSlicePJSON(w, f[2], o.Devices)
			off = 0
		}
		if c[3] {
			w.Raw(",\"isActive\":"[off:])
			w.Bool(o.IsActive)
			off = 0
		}
		if c[4] {
			w.Raw(",\"rating\":"[off:])
			w.Float64(o.Rating)
			off = 0
		}
		if c[5] {
			w.Raw(",\"lineManager\":"[off:])
			o.LineManager.EncodeObjectPJSON(w, f[5])
			off = 0
		}
		if c[6] {
			w.Raw(",\"tags\":"[off:])
			w.Strings(o.Tags)
			off = 0
		}
		w.Byte('}')
	}
}

func (o *Employee) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []Employee) {
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

func (o *Employee) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	} else {
		c := [7]bool{}
		f := [7][]parse.Filter{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "id" {
					c[0] = true
				} else if k == "person" {
					c[1], f[1] = true, filter[i].Filter
				} else if k == "devices" {
					c[2], f[2] = true, filter[i].Filter
				} else if k == "isActive" {
					c[3] = true
				} else if k == "rating" {
					c[4] = true
				} else if k == "lineManager" {
					c[5], f[5] = true, filter[i].Filter
				} else if k == "tags" {
					c[6] = true
				}
			}
		}
		if c[0] {
			b, v := writer.StringLen(o.Id)
			bytes, volatile = bytes+b+6, volatile+v
		}
		if c[1] {
			b, v := o.Person.ObjectLengthPJSON(f[1])
			bytes, volatile = bytes+b+10, volatile+v
		}
		if c[2] {
			b, v := (*externals.Device)(nil).SliceLengthPJSON(f[2], o.Devices)
			bytes, volatile = bytes+b+11, volatile+v
		}
		if c[3] {
			bytes += writer.BoolLen(o.IsActive) + 12
		}
		if c[4] {
			bytes += writer.Float64Len(o.Rating) + 10
		}
		if c[5] {
			b, v := o.LineManager.ObjectLengthPJSON(f[5])
			bytes, volatile = bytes+b+15, volatile+v
		}
		if c[6] {
			b, v := writer.StringsLen(o.Tags)
			bytes, volatile = bytes+b+8, volatile+v
		}
		if bytes == 0 {
			return 2, 0
		} else {
			return bytes + 1, volatile
		}
	}
}

func (o *Employee) SliceLengthPJSON(filter []parse.Filter, slc []Employee) (bytes int, volatile int) {
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

func (o *Person) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	c := [3]bool{}
	if filter == nil {
		for i := range c {
			c[i] = true
		}
	} else {
		for i := range filter {
			k := filter[i].Field
			if k == "fname" {
				c[0] = true
			} else if k == "lname" {
				c[1] = true
			} else if k == "dob" {
				c[2] = true
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
					if string(key) == "fname" && c[0] {
						o.Fname, err = r.GetString()
					} else if string(key) == "lname" && c[1] {
						o.Lname, err = r.GetString()
					} else if string(key) == "dob" && c[2] {
						o.DOB, err = r.GetTime()
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

func (o *Person) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []Person, err error) {
	var e Person
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]Person, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *Person) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []Person, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *Person) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	} else {
		c := [3]bool{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "fname" {
					c[0] = true
				} else if k == "lname" {
					c[1] = true
				} else if k == "dob" {
					c[2] = true
				}
			}
		}
		w.Byte('{')
		off := 1
		if c[0] {
			w.Raw(",\"fname\":"[off:])
			w.String(o.Fname)
			off = 0
		}
		if c[1] {
			w.Raw(",\"lname\":"[off:])
			w.String(o.Lname)
			off = 0
		}
		if c[2] {
			w.Raw(",\"dob\":"[off:])
			w.Time(o.DOB)
			off = 0
		}
		w.Byte('}')
	}
}

func (o *Person) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []Person) {
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

func (o *Person) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	} else {
		c := [3]bool{}
		if filter == nil {
			for i := range c {
				c[i] = true
			}
		} else {
			for i := range filter {
				k := filter[i].Field
				if k == "fname" {
					c[0] = true
				} else if k == "lname" {
					c[1] = true
				} else if k == "dob" {
					c[2] = true
				}
			}
		}
		if c[0] {
			b, v := writer.StringLen(o.Fname)
			bytes, volatile = bytes+b+9, volatile+v
		}
		if c[1] {
			b, v := writer.StringLen(o.Lname)
			bytes, volatile = bytes+b+9, volatile+v
		}
		if c[2] {
			bytes += writer.TimeLen(o.DOB) + 7
		}
		if bytes == 0 {
			return 2, 0
		} else {
			return bytes + 1, volatile
		}
	}
}

func (o *Person) SliceLengthPJSON(filter []parse.Filter, slc []Person) (bytes int, volatile int) {
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

func (o *EmployeeList) DecodeObjectPJSON(r *reader.Reader, filter []parse.Filter) (err error) {
	*o, err = (*Employee)(nil).DecodeSlicePJSON(r, filter)
	return
}

func (o *EmployeeList) sequencePJSON(r *reader.Reader, filter []parse.Filter, idx int) (res []EmployeeList, err error) {
	var e EmployeeList
	if err = e.DecodeObjectPJSON(r, filter); err == nil {
		if !r.Next() {
			res = make([]EmployeeList, idx+1)
			res[idx] = e
			return
		} else if res, err = o.sequencePJSON(r, filter, idx+1); err == nil {
			res[idx] = e
		}
	}
	return
}

func (o *EmployeeList) DecodeSlicePJSON(r *reader.Reader, filter []parse.Filter) (res []EmployeeList, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = o.sequencePJSON(r, filter, 0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (o *EmployeeList) EncodeObjectPJSON(w *writer.Writer, filter []parse.Filter) {
	if o == nil {
		w.Raw("null")
	}
	(*Employee)(nil).EncodeSlicePJSON(w, filter, *o)

}

func (o *EmployeeList) EncodeSlicePJSON(w *writer.Writer, filter []parse.Filter, slc []EmployeeList) {
	if slc == nil {
		w.Raw("null")
	}
	w.Byte('[')
	if len(slc) > 0 {
		slc[0].EncodeObjectPJSON(w, filter)
		for i := 1; i < len(slc); i++ {
			w.Byte(',')
			slc[i].EncodeObjectPJSON(w, filter)
		}
	}
	w.Byte(']')
}

func (o *EmployeeList) ObjectLengthPJSON(filter []parse.Filter) (bytes int, volatile int) {
	if o == nil {
		return 4, 0
	}
	return (*Employee)(nil).SliceLengthPJSON(filter, *o)
}

func (o *EmployeeList) SliceLengthPJSON(filter []parse.Filter, slc []EmployeeList) (bytes int, volatile int) {
	for _, obj := range slc {
		b, v := obj.ObjectLengthPJSON(filter)
		bytes += b + 1
		volatile += v
	}
	if bytes == 0 {
		return 2, 0
	} else {
		return bytes + 1, volatile
	}
}
