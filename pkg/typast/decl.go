package typast

type (
	// Decl stand of declaration
	Decl struct {
		Name    string
		Path    string
		Package string
		Type    DeclType
	}
	// DeclType is declaration type
	DeclType int
)

//
// DeclType
//

const (
	// FuncType type
	FuncType DeclType = iota + 1
	// InterfaceType type
	InterfaceType
	// StructType type
	StructType
	// GenericType type
	GenericType
)

func (d DeclType) String() string {
	return [...]string{"", "Function", "Interface", "Struct", "Generic"}[d]
}
