package typgen

type (
	// Interface interface declaration
	Interface struct {
		TypeDecl
	}
)

func CreateInterfaceDecl(typeDecl TypeDecl) *Interface {
	return &Interface{TypeDecl: typeDecl}
}
