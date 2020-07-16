package typapp

import "log"

// Start the entry points
func Start(entryPoint interface{}) {
	app := &App{
		EntryPoint: entryPoint,
		Ctors:      GetCtors(),
		Dtors:      GetDtors(),
	}
	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}
