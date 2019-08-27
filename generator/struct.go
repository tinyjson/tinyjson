package generator

type Object struct {
	PointerLvl int
	Package    string
	Filepath   string
	Name       string
}

type BaseObject struct {
	Object
	TypeName string
}

type Map struct {
	Object
	Key   interface{}
	Value interface{}
}

type Array struct {
	Object
	Value interface{}
}

type Struct struct {
	Object
	EmbededStructs []Struct
	Fields         map[string]Struct
}

type (
	Bool   BaseObject
	Int    BaseObject
	Uint   BaseObject
	Float  BaseObject
	String BaseObject
)
