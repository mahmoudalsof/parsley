package reader

import "time"

func (r *Reader) timeSeq(idx int) (res []time.Time, err error) {
	var bs []byte
	if bs, err = r.GetByteArray(); err == nil {
		if r.Next() {
			res, err = r.timeSeq(idx + 1)
		} else {
			res = make([]time.Time, idx+1)
		}
		if err == nil {
			res[idx], err = time.Parse(time.RFC3339, string(bs))
		}
	}
	return
}

func (r *Reader) GetTimes() (res []time.Time, err error) {
	if err = r.OpenArray(); err == nil {
		if res, err = r.timeSeq(0); err == nil {
			err = r.CloseArray()
		}
	}
	return
}

func (r *Reader) GetTime() (time.Time, error) {
	rpos := r.pos
	if bs, err := r.GetByteArray(); err != nil {
		return time.Time{}, err
	} else if len(bs) == 0 {
		return time.Time{}, NewUnknownTimeFormatError(string(bs), rpos)
	} else if bs[0] >= '0' && bs[0] <= '9' {
		if tm, err := time.Parse(time.RFC3339Nano, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC822, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC822Z, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.Kitchen, string(bs)); err == nil {
			return tm, nil
		} else {
			return tm, NewUnknownTimeFormatError(string(bs), rpos)
		}
	} else {
		if tm, err := time.Parse(time.ANSIC, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.UnixDate, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RubyDate, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC850, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC1123, string(bs)); err == nil {
			return tm, nil
		} else if tm, err := time.Parse(time.RFC1123Z, string(bs)); err == nil {
			return tm, nil
		} else {
			return tm, NewUnknownTimeFormatError(string(bs), rpos)
		}
	}
}

func (r *Reader) GetTimePtr() (res *time.Time, err error) {
	if v, err := r.GetTime(); err == nil {
		res = &v
	}
	return
}
