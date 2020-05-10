package mypackage

import "fmt"

type MyType struct {
	EmbeddedType
}

type EmbeddedType string

func (e EmbeddedType) ExportedMethod() {
	fmt.Println("Exported")
}
func (e EmbeddedType) unExportedMethod() {
	fmt.Println("Exported")
}
