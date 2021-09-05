package typgen

type (
	// InterfaceDecl interface declaration
	InterfaceDecl struct {
		TypeDecl
	}
)

func CreateInterfaceDecl(typeDecl TypeDecl) *InterfaceDecl {
	return &InterfaceDecl{TypeDecl: typeDecl}
}
