package appx

// DBInfra database infrastructure
type DBInfra interface {
	Create() (err error)
	Drop() (err error)
	Migrate(source string) error
	Rollback(source string) error
}
