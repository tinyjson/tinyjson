package example_generates

import (
	"bytes"
	lexer "github.com/tinyjson/lexer"
	externallib "github.com/tinyjson/tinyjson/example_generates/external_lib"
	"strconv"
)

func tinyjsonMarshalC39f6f78a15d523b(w *bytes.Buffer, this *externallib.ExternalClass) {
	w.WriteString("{")
	w.WriteString("\"Key\":")
	w.WriteString(strconv.Quote(string(this.Key)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC39f6f78a15d523b(lex *lexer.Lexer, this *externallib.ExternalClass) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key":
				v107, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key = (string)(v107)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func tinyjsonMarshalC56ec3f2525632186(w *bytes.Buffer, this *AnonymousAnonymousStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString(strconv.Quote(string(this.AnonymousStruct.SimpleStruct.Key1)))
	w.WriteString(",")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.AnonymousStruct.SimpleStruct.Key2)))
	w.WriteString(",")
	w.WriteString("\"Key3\":")
	w.WriteString(strconv.Quote(string(this.AnonymousStruct.Key3)))
	w.WriteString(",")
	w.WriteString("\"Key4\":")
	w.WriteString(strconv.Quote(string(this.Key4)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC56ec3f2525632186(lex *lexer.Lexer, this *AnonymousAnonymousStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				v111, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.AnonymousStruct.SimpleStruct.Key1 = (string)(v111)

			case "Key2":
				v113, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.AnonymousStruct.SimpleStruct.Key2 = (string)(v113)

			case "Key3":
				v115, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.AnonymousStruct.Key3 = (string)(v115)

			case "Key4":
				v117, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key4 = (string)(v117)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *AnonymousAnonymousStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC56ec3f2525632186(w, this)
	return w.Bytes(), nil
}

func (this *AnonymousAnonymousStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC56ec3f2525632186(lex, this)
}

func tinyjsonMarshalC57e9d1860d1d68d8(w *bytes.Buffer, this *AnonymousStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString(strconv.Quote(string(this.SimpleStruct.Key1)))
	w.WriteString(",")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.SimpleStruct.Key2)))
	w.WriteString(",")
	w.WriteString("\"Key3\":")
	w.WriteString(strconv.Quote(string(this.Key3)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC57e9d1860d1d68d8(lex *lexer.Lexer, this *AnonymousStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				v59, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.SimpleStruct.Key1 = (string)(v59)

			case "Key2":
				v61, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.SimpleStruct.Key2 = (string)(v61)

			case "Key3":
				v63, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key3 = (string)(v63)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *AnonymousStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC57e9d1860d1d68d8(w, this)
	return w.Bytes(), nil
}

func (this *AnonymousStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC57e9d1860d1d68d8(lex, this)
}

func tinyjsonMarshalC68255aaf95e94627(w *bytes.Buffer, this *DoubleArray) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	w.WriteString("[")
	for v64, v65 := range *this {
		if v64 > 0 {
			w.WriteString(",")
		}
		if v65 == nil {
			w.WriteString("null")
			break
		}
		w.WriteString("[")
		for v66, v67 := range v65 {
			if v66 > 0 {
				w.WriteString(",")
			}
			w.WriteString(strconv.Quote(string(v67)))
		}
		w.WriteString("]")
	}
	w.WriteString("]")
}

func tinyjsonUnmarshalC68255aaf95e94627(lex *lexer.Lexer, this *DoubleArray) error {
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ArrayIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	} else {
		*this = make([][]string, 0)
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		for {
			if lex.Controls[0] == lexer.ArrayOut {
				lex.Controls = lex.Controls[1:]

				break
			}
			var v69 []string
			if lex.Controls[0] == lexer.Nil {
				v69 = nil
				lex.Controls = lex.Controls[1:]
			} else if lex.Controls[0] != lexer.ArrayIn {
				lex.SkipValue()
				return lexer.ErrorUnexpectedType
			} else {
				v69 = make([]string, 0)
				lex.Controls = lex.Controls[1:]
				lex.Actions = lex.Actions[4:]
				for {
					if lex.Controls[0] == lexer.ArrayOut {
						lex.Controls = lex.Controls[1:]

						break
					}
					v71, err := lex.ReadString()
					if err != nil {
						return err
					}

					v69 = append(v69, string(v71))
				}
			}

			*this = append(*this, v69)
		}
	}

	return nil
}

func (this *DoubleArray) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC68255aaf95e94627(w, this)
	return w.Bytes(), nil
}

func (this *DoubleArray) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC68255aaf95e94627(lex, this)
}

func tinyjsonMarshalC1408d2ac22c4d294(w *bytes.Buffer, this *IntPtr) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	v125 := *this
	w.WriteString(strconv.FormatInt(int64(*v125), 10))
}

func tinyjsonUnmarshalC1408d2ac22c4d294(lex *lexer.Lexer, this *IntPtr) error {
	v127, err := lex.ReadInt()
	if err == lexer.ErrorNilValue {
		err = nil
		*this = nil
		return nil
	}
	v128 := int(v127)
	v129 := &v128
	*this = IntPtr(v129)
	if err != nil {
		return err
	}
	return nil
}

func tinyjsonMarshalCc697f48392907a0(w *bytes.Buffer, this *IntPtrPtr) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	v1 := *this
	if *v1 == nil {
		w.WriteString("null")
		return
	}
	v2 := *v1
	w.WriteString(strconv.FormatInt(int64(*v2), 10))
}

func tinyjsonUnmarshalCc697f48392907a0(lex *lexer.Lexer, this *IntPtrPtr) error {
	v4, err := lex.ReadInt()
	if err == lexer.ErrorNilValue {
		err = nil
		*this = nil
		return nil
	}
	v5 := int(v4)
	v6 := &v5
	v7 := &v6
	*this = IntPtrPtr(v7)
	if err != nil {
		return err
	}
	return nil
}

func tinyjsonMarshalC1b6cffa2ba517936(w *bytes.Buffer, this *MapFloatString) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	var v47 bool
	w.WriteString("{")
	for v45, v46 := range *this {
		if v47 {
			w.WriteString(",")
		} else {
			v47 = true
		}
		w.WriteString("\"" + strconv.FormatFloat(float64(v45), 'g', -1, 64) + "\"")
		w.WriteString(":")
		w.WriteString(strconv.Quote(string(v46)))
	}
	w.WriteString("}")
}

func tinyjsonUnmarshalC1b6cffa2ba517936(lex *lexer.Lexer, this *MapFloatString) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[float64]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v49, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			v50, _ := strconv.ParseFloat(v49, 64)
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v52, err := lex.ReadString()
			if err != nil {
				return err
			}

			(*this)[float64(v50)] = string(v52)
		}
	}
	return nil
}

func (this *MapFloatString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1b6cffa2ba517936(w, this)
	return w.Bytes(), nil
}

func (this *MapFloatString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1b6cffa2ba517936(lex, this)
}

func tinyjsonMarshalC866cb397916001e(w *bytes.Buffer, this *MapIntString) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	var v93 bool
	w.WriteString("{")
	for v91, v92 := range *this {
		if v93 {
			w.WriteString(",")
		} else {
			v93 = true
		}
		w.WriteString("\"" + strconv.FormatInt(int64(v91), 10) + "\"")
		w.WriteString(":")
		w.WriteString(strconv.Quote(string(v92)))
	}
	w.WriteString("}")
}

func tinyjsonUnmarshalC866cb397916001e(lex *lexer.Lexer, this *MapIntString) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[int]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v95, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			v96, _ := strconv.ParseInt(v95, 10, 64)
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v98, err := lex.ReadString()
			if err != nil {
				return err
			}

			(*this)[int(v96)] = string(v98)
		}
	}
	return nil
}

func (this *MapIntString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC866cb397916001e(w, this)
	return w.Bytes(), nil
}

func (this *MapIntString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC866cb397916001e(lex, this)
}

func tinyjsonMarshalC490bd268b68e6a3f(w *bytes.Buffer, this *MapMap) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	var v74 bool
	w.WriteString("{")
	for v72, v73 := range *this {
		if v74 {
			w.WriteString(",")
		} else {
			v74 = true
		}
		w.WriteString(strconv.Quote(string(v72)))
		w.WriteString(":")
		if v73 == nil {
			w.WriteString("null")
			break
		}
		var v77 bool
		w.WriteString("{")
		for v75, v76 := range v73 {
			if v77 {
				w.WriteString(",")
			} else {
				v77 = true
			}
			w.WriteString(strconv.Quote(string(v75)))
			w.WriteString(":")
			w.WriteString(strconv.Quote(string(v76)))
		}
		w.WriteString("}")
	}
	w.WriteString("}")
}

func tinyjsonUnmarshalC490bd268b68e6a3f(lex *lexer.Lexer, this *MapMap) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[string]map[string]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v79, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v80 := map[string]string{}
			data := lex.Data()
			if lex.Controls[0] == lexer.Nil {
				v80 = nil
				lex.Controls = lex.Controls[1:]
			} else if lex.Controls[0] != lexer.ObjectIn {
				lex.SkipValue()
			} else {
				lex.Controls = lex.Controls[1:]
				lex.Actions = lex.Actions[4:]
				v80 = make(map[string]string, 0)
				for {
					if lex.Controls[0] == lexer.ObjectOut {
						lex.Controls = lex.Controls[1:]
						break
					}
					v81, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
					lex.Controls = lex.Controls[1:]
					lex.Actions = lex.Actions[2:]
					v83, err := lex.ReadString()
					if err != nil {
						return err
					}

					v80[v81] = string(v83)
				}
			}
			(*this)[v79] = v80
		}
	}
	return nil
}

func (this *MapMap) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC490bd268b68e6a3f(w, this)
	return w.Bytes(), nil
}

func (this *MapMap) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC490bd268b68e6a3f(lex, this)
}

func tinyjsonMarshalC592e17f7b068d9db(w *bytes.Buffer, this *MapStringString) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	var v27 bool
	w.WriteString("{")
	for v25, v26 := range *this {
		if v27 {
			w.WriteString(",")
		} else {
			v27 = true
		}
		w.WriteString(strconv.Quote(string(v25)))
		w.WriteString(":")
		w.WriteString(strconv.Quote(string(v26)))
	}
	w.WriteString("}")
}

func tinyjsonUnmarshalC592e17f7b068d9db(lex *lexer.Lexer, this *MapStringString) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
	} else {
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		*this = make(map[string]string, 0)
		for {
			if lex.Controls[0] == lexer.ObjectOut {
				lex.Controls = lex.Controls[1:]
				break
			}
			v29, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			v31, err := lex.ReadString()
			if err != nil {
				return err
			}

			(*this)[v29] = string(v31)
		}
	}
	return nil
}

func (this *MapStringString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC592e17f7b068d9db(w, this)
	return w.Bytes(), nil
}

func (this *MapStringString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC592e17f7b068d9db(lex, this)
}

func tinyjsonMarshalC365a858149c6e2d1(w *bytes.Buffer, this *MyArray) {
	if *this == nil {
		w.WriteString("null")
		return
	}
	w.WriteString("[")
	for v86, v87 := range *this {
		if v86 > 0 {
			w.WriteString(",")
		}
		w.WriteString(strconv.Quote(string(v87)))
	}
	w.WriteString("]")
}

func tinyjsonUnmarshalC365a858149c6e2d1(lex *lexer.Lexer, this *MyArray) error {
	if lex.Controls[0] == lexer.Nil {
		*this = nil
		lex.Controls = lex.Controls[1:]
	} else if lex.Controls[0] != lexer.ArrayIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	} else {
		*this = make([]string, 0)
		lex.Controls = lex.Controls[1:]
		lex.Actions = lex.Actions[4:]
		for {
			if lex.Controls[0] == lexer.ArrayOut {
				lex.Controls = lex.Controls[1:]

				break
			}
			v90, err := lex.ReadString()
			if err != nil {
				return err
			}

			*this = append(*this, string(v90))
		}
	}

	return nil
}

func (this *MyArray) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC365a858149c6e2d1(w, this)
	return w.Bytes(), nil
}

func (this *MyArray) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC365a858149c6e2d1(lex, this)
}

func tinyjsonMarshalC380704bb7b4d7c03(w *bytes.Buffer, this *MyExternal) {
	tinyjsonMarshalC39f6f78a15d523b(w, (*externallib.ExternalClass)(this))
}

func tinyjsonUnmarshalC380704bb7b4d7c03(lex *lexer.Lexer, this *MyExternal) error {
	tinyjsonUnmarshalC39f6f78a15d523b(lex, (*externallib.ExternalClass)(this))
	return nil
}

func (this *MyExternal) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC380704bb7b4d7c03(w, this)
	return w.Bytes(), nil
}

func (this *MyExternal) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC380704bb7b4d7c03(lex, this)
}

func tinyjsonMarshalC268447a4189deb99(w *bytes.Buffer, this *MyExternalAlias) {
	tinyjsonMarshalC39f6f78a15d523b(w, (*externallib.ExternalClass)(this))
}

func tinyjsonUnmarshalC268447a4189deb99(lex *lexer.Lexer, this *MyExternalAlias) error {
	tinyjsonUnmarshalC39f6f78a15d523b(lex, (*externallib.ExternalClass)(this))
	return nil
}

func (this *MyExternalAlias) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC268447a4189deb99(w, this)
	return w.Bytes(), nil
}

func (this *MyExternalAlias) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC268447a4189deb99(lex, this)
}

func tinyjsonMarshalC2e3108dabb158644(w *bytes.Buffer, this *MyFloat) {
	w.WriteString(strconv.FormatFloat(float64(*this), 'g', -1, 64))
}

func tinyjsonUnmarshalC2e3108dabb158644(lex *lexer.Lexer, this *MyFloat) error {
	v54, err := lex.ReadFloat()
	*this = MyFloat(v54)
	if err != nil {
		return err
	}
	return nil
}

func (this *MyFloat) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC2e3108dabb158644(w, this)
	return w.Bytes(), nil
}

func (this *MyFloat) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC2e3108dabb158644(lex, this)
}

func tinyjsonMarshalC1a714cf86b83d0e2(w *bytes.Buffer, this *MyInt) {
	w.WriteString(strconv.FormatInt(int64(*this), 10))
}

func tinyjsonUnmarshalC1a714cf86b83d0e2(lex *lexer.Lexer, this *MyInt) error {
	v56, err := lex.ReadInt()
	*this = MyInt(v56)
	if err != nil {
		return err
	}
	return nil
}

func (this *MyInt) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1a714cf86b83d0e2(w, this)
	return w.Bytes(), nil
}

func (this *MyInt) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1a714cf86b83d0e2(lex, this)
}

func tinyjsonMarshalC55104dc76695721d(w *bytes.Buffer, this *MyString) {
	w.WriteString(strconv.Quote(string(*this)))
}

func tinyjsonUnmarshalC55104dc76695721d(lex *lexer.Lexer, this *MyString) error {
	v85, err := lex.ReadString()
	if err != nil {
		return err
	}
	*this = (MyString)(v85)

	return nil
}

func (this *MyString) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC55104dc76695721d(w, this)
	return w.Bytes(), nil
}

func (this *MyString) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC55104dc76695721d(lex, this)
}

func tinyjsonMarshalC25845c95d4491d1b(w *bytes.Buffer, this *MyStringAlias) {
	tinyjsonMarshalC55104dc76695721d(w, (*MyString)(this))
}

func tinyjsonUnmarshalC25845c95d4491d1b(lex *lexer.Lexer, this *MyStringAlias) error {
	tinyjsonUnmarshalC55104dc76695721d(lex, (*MyString)(this))
	return nil
}

func (this *MyStringAlias) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC25845c95d4491d1b(w, this)
	return w.Bytes(), nil
}

func (this *MyStringAlias) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC25845c95d4491d1b(lex, this)
}

func tinyjsonMarshalC6ec34c367674cb74(w *bytes.Buffer, this *MyUInt16) {
	w.WriteString(strconv.FormatUint(uint64(*this), 10))
}

func tinyjsonUnmarshalC6ec34c367674cb74(lex *lexer.Lexer, this *MyUInt16) error {
	v44, err := lex.ReadInt()
	*this = MyUInt16(v44)
	if err != nil {
		return err
	}
	return nil
}

func (this *MyUInt16) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC6ec34c367674cb74(w, this)
	return w.Bytes(), nil
}

func (this *MyUInt16) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC6ec34c367674cb74(lex, this)
}

func tinyjsonMarshalC3c04951aa42655d9(w *bytes.Buffer, this *SimpleStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString(strconv.Quote(string(this.Key1)))
	w.WriteString(",")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.Key2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC3c04951aa42655d9(lex *lexer.Lexer, this *SimpleStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				v37, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key1 = (string)(v37)

			case "Key2":
				v39, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key2 = (string)(v39)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *SimpleStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC3c04951aa42655d9(w, this)
	return w.Bytes(), nil
}

func (this *SimpleStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC3c04951aa42655d9(lex, this)
}

func tinyjsonMarshalC2584c47f2cdf5b8a(w *bytes.Buffer, this *StructA1) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.A)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC2584c47f2cdf5b8a(lex *lexer.Lexer, this *StructA1) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v34, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A = (string)(v34)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA1) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC2584c47f2cdf5b8a(w, this)
	return w.Bytes(), nil
}

func (this *StructA1) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC2584c47f2cdf5b8a(lex, this)
}

func tinyjsonMarshalC30b95ff183c471d4(w *bytes.Buffer, this *StructA2) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.A)))
	w.WriteString(",")
	w.WriteString("\"c\":")
	w.WriteString(strconv.Quote(string(this.StructC.C)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC30b95ff183c471d4(lex *lexer.Lexer, this *StructA2) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v133, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A = (string)(v133)

			case "c":
				v135, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructC.C = (string)(v135)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA2) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC30b95ff183c471d4(w, this)
	return w.Bytes(), nil
}

func (this *StructA2) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC30b95ff183c471d4(lex, this)
}

func tinyjsonMarshalC1bf98be2a9d78d73(w *bytes.Buffer, this *StructA3) {
	w.WriteString("{")
	w.WriteString("\"a3\":")
	w.WriteString(strconv.Quote(string(this.A3)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC1bf98be2a9d78d73(lex *lexer.Lexer, this *StructA3) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a3":
				v10, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A3 = (string)(v10)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA3) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1bf98be2a9d78d73(w, this)
	return w.Bytes(), nil
}

func (this *StructA3) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1bf98be2a9d78d73(lex, this)
}

func tinyjsonMarshalC2606cd2b57d29245(w *bytes.Buffer, this *StructA4) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.A)))
	w.WriteString(",")
	w.WriteString("\"d\":")
	w.WriteString(strconv.Quote(string(this.StructD.D2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC2606cd2b57d29245(lex *lexer.Lexer, this *StructA4) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v13, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.A = (string)(v13)

			case "d":
				v15, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D2 = (string)(v15)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA4) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC2606cd2b57d29245(w, this)
	return w.Bytes(), nil
}

func (this *StructA4) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC2606cd2b57d29245(lex, this)
}

func tinyjsonMarshalC1a02070f169c1121(w *bytes.Buffer, this *StructA5) {
	w.WriteString("{")
	w.WriteString("\"d\":")
	w.WriteString(strconv.Quote(string(this.StructD.D2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC1a02070f169c1121(lex *lexer.Lexer, this *StructA5) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "d":
				v42, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D2 = (string)(v42)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA5) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC1a02070f169c1121(w, this)
	return w.Bytes(), nil
}

func (this *StructA5) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC1a02070f169c1121(lex, this)
}

func tinyjsonMarshalC6054502fc5d6d268(w *bytes.Buffer, this *StructA6) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.StructB.B)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC6054502fc5d6d268(lex *lexer.Lexer, this *StructA6) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v124, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructB.B = (string)(v124)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA6) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC6054502fc5d6d268(w, this)
	return w.Bytes(), nil
}

func (this *StructA6) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC6054502fc5d6d268(lex, this)
}

func tinyjsonMarshalC4d65822107fcfd52(w *bytes.Buffer, this *StructA7) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.StructD.D)))
	w.WriteString(",")
	w.WriteString("\"b\":")
	tinyjsonMarshalC28b621587cb3ad0b(w, (*StructB)(&this.StructB))
	w.WriteString(",")
	w.WriteString("\"d\":")
	w.WriteString(strconv.Quote(string(this.StructD.D2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC4d65822107fcfd52(lex *lexer.Lexer, this *StructA7) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v18, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D = (string)(v18)

			case "b":
				tinyjsonUnmarshalC28b621587cb3ad0b(lex, (*StructB)(&this.StructB))
			case "d":
				v21, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.StructD.D2 = (string)(v21)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructA7) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC4d65822107fcfd52(w, this)
	return w.Bytes(), nil
}

func (this *StructA7) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC4d65822107fcfd52(lex, this)
}

func tinyjsonMarshalC28b621587cb3ad0b(w *bytes.Buffer, this *StructB) {
	w.WriteString("{")
	w.WriteString("\"a\":")
	w.WriteString(strconv.Quote(string(this.B)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC28b621587cb3ad0b(lex *lexer.Lexer, this *StructB) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "a":
				v24, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.B = (string)(v24)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func tinyjsonMarshalC6e661e92759805f5(w *bytes.Buffer, this *StructInStruct) {
	w.WriteString("{")
	w.WriteString("\"Key1\":")
	w.WriteString("{")
	w.WriteString("\"Key2\":")
	w.WriteString(strconv.Quote(string(this.Key1.Key2)))
	w.WriteString("}")
	w.WriteString("}")
}

func tinyjsonUnmarshalC6e661e92759805f5(lex *lexer.Lexer, this *StructInStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "Key1":
				data := lex.Data()
				if lex.Controls[0] == lexer.Nil {
					return nil
				} else if lex.Controls[0] != lexer.ObjectIn {
					lex.SkipValue()
					return lexer.ErrorUnexpectedType
				}
				lex.Controls = lex.Controls[1:]
				lex.Actions = lex.Actions[4:]
				for {
					switch lex.Controls[0] {
					case lexer.ObjectOut:
						lex.Controls = lex.Controls[1:]
						return nil
					case lexer.Key:
						key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
						lex.Controls = lex.Controls[1:]
						lex.Actions = lex.Actions[2:]
						switch key {
						case "Key2":
							v121, err := lex.ReadString()
							if err != nil {
								return err
							}
							this.Key1.Key2 = (string)(v121)

						default:
							lex.SkipValue()
						}
					}
				}
			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *StructInStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC6e661e92759805f5(w, this)
	return w.Bytes(), nil
}

func (this *StructInStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC6e661e92759805f5(lex, this)
}

func tinyjsonMarshalC430c8b35bb9457d8(w *bytes.Buffer, this *TaggedStruct) {
	w.WriteString("{")
	w.WriteString("\"key_1\":")
	w.WriteString(strconv.Quote(string(this.Key1)))
	w.WriteString(",")
	w.WriteString("\"key_2\":")
	w.WriteString(strconv.Quote(string(this.Key2)))
	w.WriteString("}")
}

func tinyjsonUnmarshalC430c8b35bb9457d8(lex *lexer.Lexer, this *TaggedStruct) error {
	data := lex.Data()
	if lex.Controls[0] == lexer.Nil {
		return nil
	} else if lex.Controls[0] != lexer.ObjectIn {
		lex.SkipValue()
		return lexer.ErrorUnexpectedType
	}
	lex.Controls = lex.Controls[1:]
	lex.Actions = lex.Actions[4:]
	for {
		switch lex.Controls[0] {
		case lexer.ObjectOut:
			lex.Controls = lex.Controls[1:]
			return nil
		case lexer.Key:
			key, _ := strconv.Unquote(string(data[lex.Actions[0]:lex.Actions[1]]))
			lex.Controls = lex.Controls[1:]
			lex.Actions = lex.Actions[2:]
			switch key {
			case "key_1":
				v101, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key1 = (string)(v101)

			case "key_2":
				v103, err := lex.ReadString()
				if err != nil {
					return err
				}
				this.Key2 = (string)(v103)

			default:
				lex.SkipValue()
			}
		}
	}
	return nil
}

func (this *TaggedStruct) MarshalJSON() ([]byte, error) {
	w := bytes.NewBuffer(nil)
	tinyjsonMarshalC430c8b35bb9457d8(w, this)
	return w.Bytes(), nil
}

func (this *TaggedStruct) UnmarshalJSON(data []byte) error {
	lex := lexer.NewLexer(data)
	lex.Parse()
	return tinyjsonUnmarshalC430c8b35bb9457d8(lex, this)
}
