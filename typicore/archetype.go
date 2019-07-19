package typicore

type ArcheType interface {
	Name() string
	Statements() []Statement
}
