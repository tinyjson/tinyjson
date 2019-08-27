package example_generates

import (
	"github.com/tinyjson/tinyjson/example_generates/external_lib"
	al "github.com/tinyjson/tinyjson/example_generates/external_lib"
)

type (
	//tinyjson:json
	MyInt int

	//tinyjson:json
	MyUInt16 uint16

	//tinyjson:json
	IntPtr *int

	//tinyjson:json
	IntPtrPtr **int

	//tinyjson:json
	MyString string

	//tinyjson:json
	MyStringAlias MyString

	//tinyjson:json
	MyFloat float64

	//tinyjson:json
	MyExternal external_lib.ExternalClass

	//tinyjson:json
	MyExternalAlias al.ExternalClass

	//tinyjson:json
	MyArray []string

	//tinyjson:json
	DoubleArray [][]string

	//tinyjson:json
	MapStringString map[string]string

	//tinyjson:json
	MapMap map[string]map[string]string

	//tinyjson:json
	MapIntString map[int]string

	//tinyjson:json
	MapFloatString map[float64]string

	//tinyjson:json
	SimpleStruct struct {
		Key1 string
		Key2 string
	}

	//tinyjson:json
	TaggedStruct struct {
		Key1 string `json:"key_1"`
		Key2 string `json:"key_2"`
	}

	//tinyjson:json
	AnonymousStruct struct {
		SimpleStruct
		Key3 string
	}

	//tinyjson:json
	AnonymousAnonymousStruct struct {
		AnonymousStruct
		Key4 string
	}

	//tinyjson:json
	StructInStruct struct {
		Key1 struct {
			Key2 string
		}
	}

	//tinyjson:json
	StructA1 struct {
		A string `json:"a"`
		StructB
	}

	//tinyjson:json
	StructA2 struct {
		A string `json:"a"`
		StructC
	}

	//tinyjson:json
	StructA3 struct {
		A1 string `json:"a"`
		A2 string `json:"a"`
		A3 string `json:"a3"`
	}

	//tinyjson:json
	StructA4 struct {
		A string `json:"a"`
		StructB
		StructD
	}

	//tinyjson:json
	StructA5 struct {
		StructB
		StructD
	}

	//tinyjson:json
	StructA6 struct {
		StructB
		StructE
	}

	//tinyjson:json
	StructA7 struct {
		StructB `json:"b"`
		StructD
	}

	StructB struct {
		B string `json:"a"`
	}

	StructC struct {
		C string `json:"c"`
	}

	StructD struct {
		D  string `json:"a"`
		D2 string `json:"d"`
	}

	StructE struct {
		StructB
	}
)
