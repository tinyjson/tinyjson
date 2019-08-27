package generator

import (
	"fmt"
	"testing"
)

type xA struct {
	Name string
}

type xB *xA

type xC *xB

type yA int

type yB *yA

type yC *yB

type m map[string]string

type xD struct {
	m
	x xC
	y yC
}

func TestStruct(t *testing.T) {
	a := xA{"xxx"}
	ap := xB(&a)
	app := xC(&ap)
	b := yA(123)
	bp := yB(&b)
	bpp := yC(&bp)
	d := xD{
		m: map[string]string{},
		x: xC(app),
		y: bpp,
	}
	d.m["3"] = "3"
	fmt.Println(d)
}
