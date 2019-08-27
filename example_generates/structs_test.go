package example_generates

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMyInt_MarshalJSON(t *testing.T) {
	a := MyInt(123)
	res, err := a.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `123`, string(res))
}

func TestMyInt_UnmarshalJSON(t *testing.T) {
	var a MyInt
	err := a.UnmarshalJSON([]byte(`123`))
	assert.NoError(t, err)
	assert.Equal(t, MyInt(123), a)
}

func TestMyString_MarshalJSON(t *testing.T) {
	a := MyString(`foo"bar`)
	res, err := a.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `"foo\"bar"`, string(res))
}

func TestMyString_UnmarshalJSON(t *testing.T) {
	var a MyString
	err := a.UnmarshalJSON([]byte(`"foo\"bar"`))
	assert.NoError(t, err)
	assert.Equal(t, MyString(`foo"bar`), a)
}

func TestMyFloat_MarshalJSON(t *testing.T) {
	a := MyFloat(12.3)
	res, err := a.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `12.3`, string(res))
}

func TestMyFloat_UnmarshalJSON(t *testing.T) {
	var a MyFloat
	err := a.UnmarshalJSON([]byte(`12.3`))
	assert.NoError(t, err)
	assert.Equal(t, MyFloat(12.3), a)
}

func TestMyExternal_MarshalJSON(t *testing.T) {
	a := MyExternal{
		Key: "k",
	}
	res, err := a.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"Key":"k"}`, string(res))
}

func TestMyExternal_UnmarshalJSON(t *testing.T) {
	raw := MyExternal{
		Key: "k",
	}
	var got MyExternal
	err := got.UnmarshalJSON([]byte(`{"Key":"k"}`))
	assert.NoError(t, err)
	assert.Equal(t, raw, got)
}

func TestMyExternalAlias_MarshalJSON(t *testing.T) {
	a := MyExternalAlias{
		Key: "k",
	}
	res, err := a.MarshalJSON()
	assert.NoError(t, err)
	assert.Equal(t, `{"Key":"k"}`, string(res))
}

func TestMyExternalAlias_UnmarshalJSON(t *testing.T) {
	raw := MyExternalAlias{
		Key: "k",
	}
	var got MyExternalAlias
	err := got.UnmarshalJSON([]byte(`{"Key":"k"}`))
	assert.NoError(t, err)
	assert.Equal(t, raw, got)
}

func TestMyArray_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		arr  []string
		want string
	}{
		{
			arr:  nil,
			want: `null`,
		},
		{
			arr:  []string{},
			want: `[]`,
		},
		{
			arr:  []string{"a"},
			want: `["a"]`,
		},
		{
			arr:  []string{"a", "b"},
			want: `["a","b"]`,
		},
	} {
		t.Run("", func(t *testing.T) {
			a := MyArray(tt.arr)
			res, err := a.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestMyArray_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want MyArray
		raw  string
	}{
		{
			want: nil,
			raw:  `null`,
		},
		{
			want: []string{},
			raw:  `[]`,
		},
		{
			want: []string{"a"},
			raw:  `["a"]`,
		},
		{
			want: []string{"a", "b"},
			raw:  `["a","b"]`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got MyArray
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			fmt.Printf("%#v %#v\n", tt.want, got)
		})
	}
}

func TestDoubleArray_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		arr  [][]string
		want string
	}{
		{
			arr:  nil,
			want: `null`,
		},
		{
			arr:  [][]string{},
			want: `[]`,
		},
		{
			arr:  [][]string{{}},
			want: `[[]]`,
		},
		{
			arr:  [][]string{nil},
			want: `[null]`,
		},
		{
			arr:  [][]string{{"a"}},
			want: `[["a"]]`,
		},
		{
			arr:  [][]string{{"a", "b"}},
			want: `[["a","b"]]`,
		},
		{
			arr:  [][]string{{"a"}, {"b"}},
			want: `[["a"],["b"]]`,
		},
		{
			arr:  [][]string{{"a"}, {"b", "c"}},
			want: `[["a"],["b","c"]]`,
		},
		{
			arr:  [][]string{{"a", "b"}, {"c"}},
			want: `[["a","b"],["c"]]`,
		},
	} {
		t.Run("", func(t *testing.T) {
			a := DoubleArray(tt.arr)
			res, err := a.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestDoubleArray_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want DoubleArray
		raw  string
	}{
		{
			want: nil,
			raw:  `null`,
		},
		{
			want: [][]string{},
			raw:  `[]`,
		},
		{
			want: [][]string{{}},
			raw:  `[[]]`,
		},
		{
			want: [][]string{nil},
			raw:  `[null]`,
		},
		{
			want: [][]string{{"a"}},
			raw:  `[["a"]]`,
		},
		{
			want: [][]string{{"a", "b"}},
			raw:  `[["a","b"]]`,
		},
		{
			want: [][]string{{"a"}, {"b"}},
			raw:  `[["a"],["b"]]`,
		},
		{
			want: [][]string{{"a"}, {"b", "c"}},
			raw:  `[["a"],["b","c"]]`,
		},
		{
			want: [][]string{{"a", "b"}, {"c"}},
			raw:  `[["a","b"],["c"]]`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got DoubleArray
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
			fmt.Printf("%#v %#v\n", tt.want, got)
		})
	}
}

func TestMapStringString_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    map[string]string
		want string
	}{
		{
			m:    nil,
			want: `null`,
		},
		{
			m:    make(MapStringString, 0),
			want: `{}`,
		},
		{
			m: map[string]string{
				"k": "v",
			},
			want: `{"k":"v"}`,
		},
		{
			m: map[string]string{
				"k":  "v",
				"k2": "v2",
			},
			want: `{"k":"v","k2":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			a := MapStringString(tt.m)
			res, err := a.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestMapStringString_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want MapStringString
		raw  string
	}{
		{
			want: nil,
			raw:  `null`,
		},
		{
			want: make(MapStringString, 0),
			raw:  `{}`,
		},
		{
			want: map[string]string{
				"k": "v",
			},
			raw: `{"k":"v"}`,
		},
		{
			want: map[string]string{
				"k":  "v",
				"k2": "v2",
			},
			raw: `{"k":"v","k2":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got MapStringString
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapIntString_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    map[int]string
		want string
	}{
		{
			m: map[int]string{
				1: "v",
			},
			want: `{"1":"v"}`,
		},
		{
			m: map[int]string{
				1: "v",
				2: "v2",
			},
			want: `{"1":"v","2":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			a := MapIntString(tt.m)
			res, err := a.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestMapIntString_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want MapIntString
		raw  string
	}{
		{
			want: map[int]string{
				1: "v",
			},
			raw: `{"1":"v"}`,
		},
		{
			want: map[int]string{
				1: "v",
				2: "v2",
			},
			raw: `{"1":"v","2":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got MapIntString
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapFloatString_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    map[float64]string
		want string
	}{
		{
			m: map[float64]string{
				1.2: "v",
			},
			want: `{"1.2":"v"}`,
		},
		{
			m: map[float64]string{
				1.2: "v",
				3.4: "v2",
			},
			want: `{"1.2":"v","3.4":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			a := MapFloatString(tt.m)
			res, err := a.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestMapFloatString_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want MapFloatString
		raw  string
	}{
		{
			want: map[float64]string{
				1.2: "v",
			},
			raw: `{"1.2":"v"}`,
		},
		{
			want: map[float64]string{
				1.2: "v",
				3.4: "v2",
			},
			raw: `{"1.2":"v","3.4":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got MapFloatString
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestMapMap_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    MapMap
		want string
	}{
		{
			m:    nil,
			want: `null`,
		},
		{
			m:    make(map[string]map[string]string, 0),
			want: `{}`,
		},
		{
			m: map[string]map[string]string{
				"k": nil,
			},
			want: `{"k":null}`,
		},
		{
			m: map[string]map[string]string{
				"k": make(map[string]string, 0),
			},
			want: `{"k":{}}`,
		},
		{
			m: map[string]map[string]string{
				"k1": {
					"k2": "v2",
				},
			},
			want: `{"k1":{"k2":"v2"}}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestMapMap_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want MapMap
		raw  string
	}{
		{
			want: nil,
			raw:  `null`,
		},
		{
			want: make(map[string]map[string]string, 0),
			raw:  `{}`,
		},
		{
			want: map[string]map[string]string{
				"k": nil,
			},
			raw: `{"k":null}`,
		},
		{
			want: map[string]map[string]string{
				"k": make(map[string]string, 0),
			},
			raw: `{"k":{}}`,
		},
		{
			want: map[string]map[string]string{
				"k1": {
					"k2": "v2",
				},
			},
			raw: `{"k1":{"k2":"v2"}}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got MapMap
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSimpleStruct_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    SimpleStruct
		want string
	}{
		{
			m: SimpleStruct{
				Key1: "v1",
				Key2: "v2",
			},
			want: `{"Key1":"v1","Key2":"v2"}`,
		},
		{
			m: SimpleStruct{
				Key1: "",
				Key2: "v2",
			},
			want: `{"Key1":"","Key2":"v2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestSimpleStruct_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		raw  string
		want SimpleStruct
	}{
		{
			raw: `{"Key1":"v1","Key2":"v2"}`,
			want: SimpleStruct{
				Key1: "v1",
				Key2: "v2",
			},
		},
		{
			raw: `{"Key1":"","Key2":"v2"}`,
			want: SimpleStruct{
				Key1: "",
				Key2: "v2",
			},
		},
	} {
		t.Run("", func(t *testing.T) {
			var got SimpleStruct
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAnonymousStruct_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    AnonymousStruct
		want string
	}{
		{
			m: AnonymousStruct{
				SimpleStruct: SimpleStruct{
					Key1: "v1",
					Key2: "v2",
				},
				Key3: "v3",
			},
			want: `{"Key1":"v1","Key2":"v2","Key3":"v3"}`,
		},
		{
			m: AnonymousStruct{
				SimpleStruct: SimpleStruct{
					Key1: "v1",
					Key2: "",
				},
				Key3: "v3",
			},
			want: `{"Key1":"v1","Key2":"","Key3":"v3"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestAnonymousStruct_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want AnonymousStruct
		raw  string
	}{
		{
			want: AnonymousStruct{
				SimpleStruct: SimpleStruct{
					Key1: "v1",
					Key2: "v2",
				},
				Key3: "v3",
			},
			raw: `{"Key1":"v1","Key2":"v2","Key3":"v3"}`,
		},
		{
			want: AnonymousStruct{
				SimpleStruct: SimpleStruct{
					Key1: "v1",
					Key2: "",
				},
				Key3: "v3",
			},
			raw: `{"Key1":"v1","Key2":"","Key3":"v3"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got AnonymousStruct
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestAnonymousAnonymousStruct_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    AnonymousAnonymousStruct
		want string
	}{
		{
			m: AnonymousAnonymousStruct{
				AnonymousStruct: AnonymousStruct{
					SimpleStruct: SimpleStruct{
						Key1: "v1",
						Key2: "v2",
					},
					Key3: "v3",
				},
				Key4: "v4",
			},
			want: `{"Key1":"v1","Key2":"v2","Key3":"v3","Key4":"v4"}`,
		},
		{
			m: AnonymousAnonymousStruct{
				AnonymousStruct: AnonymousStruct{
					SimpleStruct: SimpleStruct{
						Key1: "v1",
						Key2: "v2",
					},
					Key3: "v3",
				},
				Key4: "",
			},
			want: `{"Key1":"v1","Key2":"v2","Key3":"v3","Key4":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestAnonymousAnonymousStruct_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want AnonymousAnonymousStruct
		raw  string
	}{
		{
			want: AnonymousAnonymousStruct{
				AnonymousStruct: AnonymousStruct{
					SimpleStruct: SimpleStruct{
						Key1: "v1",
						Key2: "v2",
					},
					Key3: "v3",
				},
				Key4: "v4",
			},
			raw: `{"Key1":"v1","Key2":"v2","Key3":"v3","Key4":"v4"}`,
		},
		{
			want: AnonymousAnonymousStruct{
				AnonymousStruct: AnonymousStruct{
					SimpleStruct: SimpleStruct{
						Key1: "v1",
						Key2: "v2",
					},
					Key3: "v3",
				},
				Key4: "",
			},
			raw: `{"Key1":"v1","Key2":"v2","Key3":"v3","Key4":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got AnonymousAnonymousStruct
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructInStruct_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructInStruct
		want string
	}{
		{
			m: StructInStruct{
				Key1: struct {
					Key2 string
				}{
					Key2: "v3",
				},
			},
			want: `{"Key1":{"Key2":"v3"}}`,
		},
		{
			m: StructInStruct{
				Key1: struct {
					Key2 string
				}{
					Key2: "v0",
				},
			},
			want: `{"Key1":{"Key2":"v0"}}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructInStruct_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructInStruct
		raw  string
	}{
		{
			want: StructInStruct{
				Key1: struct {
					Key2 string
				}{
					Key2: "v3",
				},
			},
			raw: `{"Key1":{"Key2":"v3"}}`,
		},
		{
			want: StructInStruct{
				Key1: struct {
					Key2 string
				}{
					Key2: "v0",
				},
			},
			raw: `{"Key1":{"Key2":"v0"}}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructInStruct
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTaggedStruct_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    TaggedStruct
		want string
	}{
		{
			m: TaggedStruct{
				Key1: "v1",
				Key2: "v2",
			},
			want: `{"key_1":"v1","key_2":"v2"}`,
		},
		{
			m: TaggedStruct{
				Key1: "",
				Key2: "v2",
			},
			want: `{"key_1":"","key_2":"v2"}`,
		},
		{
			m: TaggedStruct{
				Key1: "v1",
				Key2: "",
			},
			want: `{"key_1":"v1","key_2":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestTaggedStruct_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want TaggedStruct
		raw  string
	}{
		{
			want: TaggedStruct{
				Key1: "v1",
				Key2: "v2",
			},
			raw: `{"key_1":"v1","key_2":"v2"}`,
		},
		{
			want: TaggedStruct{
				Key1: "",
				Key2: "v2",
			},
			raw: `{"key_1":"","key_2":"v2"}`,
		},
		{
			want: TaggedStruct{
				Key1: "v1",
				Key2: "",
			},
			raw: `{"key_1":"v1","key_2":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got TaggedStruct
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA1_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA1
		want string
	}{
		{
			m: StructA1{
				A: "a0",
				StructB: StructB{
					B: "b0",
				},
			},
			want: `{"a":"a0"}`,
		},
		{
			m: StructA1{
				A: "a0",
				StructB: StructB{
					B: "b1",
				},
			},
			want: `{"a":"a0"}`,
		},
		{
			m: StructA1{
				A: "a0",
				StructB: StructB{
					B: "",
				},
			},
			want: `{"a":"a0"}`,
		},
		{
			m: StructA1{
				A: "",
				StructB: StructB{
					B: "b0",
				},
			},
			want: `{"a":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA1_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA1
		raw  string
	}{
		{
			want: StructA1{
				A: "a0",
			},
			raw: `{"a":"a0","B":"b0"}`,
		},
		{
			want: StructA1{
				A: "a0",
			},
			raw: `{"a":"a0","B":"b1"}`,
		},
		{
			want: StructA1{
				A: "a0",
			},
			raw: `{"a":"a0"}`,
		},
		{
			want: StructA1{
				A: "",
			},
			raw: `{"a":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA1
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA2_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA2
		want string
	}{
		{
			m: StructA2{
				A: "a0",
				StructC: StructC{
					C: "c0",
				},
			},
			want: `{"a":"a0","c":"c0"}`,
		},
		{
			m: StructA2{
				A: "",
				StructC: StructC{
					C: "c0",
				},
			},
			want: `{"a":"","c":"c0"}`,
		},
		{
			m: StructA2{
				A: "a0",
				StructC: StructC{
					C: "",
				},
			},
			want: `{"a":"a0","c":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA2_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA2
		raw  string
	}{
		{
			want: StructA2{
				A: "a0",
				StructC: StructC{
					C: "c0",
				},
			},
			raw: `{"a":"a0","c":"c0"}`,
		},
		{
			want: StructA2{
				A: "",
				StructC: StructC{
					C: "c0",
				},
			},
			raw: `{"a":"","c":"c0"}`,
		},
		{
			want: StructA2{
				A: "a0",
				StructC: StructC{
					C: "",
				},
			},
			raw: `{"a":"a0","c":""}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA2
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA3_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA3
		want string
	}{
		{
			m: StructA3{
				A1: "1",
				A2: "2",
				A3: "3",
			},
			want: `{"a3":"3"}`,
		},
		{
			m: StructA3{
				A1: "",
				A2: "2",
				A3: "3",
			},
			want: `{"a3":"3"}`,
		},
		{
			m: StructA3{
				A1: "1",
				A2: "",
				A3: "3",
			},
			want: `{"a3":"3"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA3_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA3
		raw  string
	}{
		{
			want: StructA3{
				A3: "3",
			},
			raw: `{"A1":"1","A2":"2","a3":"3","A3":"3+"}`,
		},
		{
			want: StructA3{
				A3: "3",
			},
			raw: `{"A1":"1","A2":"2","A3":"3+","a3":"3"}`,
		},
		{
			want: StructA3{
				A3: "3",
			},
			raw: `{"a3":"3"}`,
		},
		{
			want: StructA3{
				A3: "3",
			},
			raw: `{"A2":"2","a3":"3","A3":"3+"}`,
		},
		{
			want: StructA3{
				A3: "3",
			},
			raw: `{"A1":"1","a3":"3","A3":"3+"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA3
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA4_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA4
		want string
	}{
		{
			m: StructA4{
				A: "a",
				StructB: StructB{
					B: "b",
				},
				StructD: StructD{
					D:  "d",
					D2: "d2",
				},
			},
			want: `{"a":"a","d":"d2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA4_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA4
		raw  string
	}{
		{
			want: StructA4{
				A: "a",
				StructD: StructD{
					D2: "d2",
				},
			},
			raw: `{"a":"a","B":"b","d":"d2","D":"d","D2":"d3"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA4
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA5_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA5
		want string
	}{
		{
			m: StructA5{
				StructB: StructB{
					B: "b",
				},
				StructD: StructD{
					D:  "d",
					D2: "d2",
				},
			},
			want: `{"d":"d2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA5_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA5
		raw  string
	}{
		{
			want: StructA5{
				StructD: StructD{
					D2: "d2",
				},
			},
			raw: `{"B":"b","D":"d","d":"d2","D2":"d2+"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA5
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA6_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA6
		want string
	}{
		{
			m: StructA6{
				StructB: StructB{
					B: "b1",
				},
				StructE: StructE{
					StructB: StructB{
						B: "b2",
					},
				},
			},
			want: `{"a":"b1"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA6_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA6
		raw  string
	}{
		{
			want: StructA6{
				StructB: StructB{
					B: "b1",
				},
			},
			raw: `{"a":"b1","B":"b2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA6
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStructA7_MarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		m    StructA7
		want string
	}{
		{
			m: StructA7{
				StructB: StructB{
					B: "b1",
				},
				StructD: StructD{
					D:  "d1",
					D2: "d2",
				},
			},
			want: `{"a":"d1","b":{"a":"b1"},"d":"d2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			res, err := tt.m.MarshalJSON()
			assert.NoError(t, err)
			assert.Equal(t, tt.want, string(res))
		})
	}
}

func TestStructA7_UnmarshalJSON(t *testing.T) {
	for _, tt := range []struct {
		want StructA7
		raw  string
	}{
		{
			want: StructA7{
				StructB: StructB{
					B: "b1",
				},
				StructD: StructD{
					D:  "d1",
					D2: "d2",
				},
			},
			raw: `{"a":"d1","b":{"a":"b1"},"d":"d2"}`,
		},
	} {
		t.Run("", func(t *testing.T) {
			var got StructA7
			err := got.UnmarshalJSON([]byte(tt.raw))
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
