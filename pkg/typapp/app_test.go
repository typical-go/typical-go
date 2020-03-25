package typapp_test

import (
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/stretchr/testify/require"
	"github.com/typical-go/typical-go/pkg/typdep"

	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestNewApp(t *testing.T) {
	t.Run("SHOULD implement App", func(t *testing.T) {
		var _ typcore.App = typapp.Create(nil)
	})
	t.Run("SHOULD implement Preconditioner", func(t *testing.T) {
		var _ typbuildtool.Preconditioner = typapp.Create(nil)
	})
	t.Run("SHOULD implement Provider", func(t *testing.T) {
		var _ typapp.Provider = typapp.Create(nil)
	})
	t.Run("SHOULD implement Destroyer", func(t *testing.T) {
		var _ typapp.Destroyer = typapp.Create(nil)
	})
	t.Run("SHOULD implement Preparer", func(t *testing.T) {
		var _ typapp.Preparer = typapp.Create(nil)
	})
	t.Run("SHOULD implement EntryPointer", func(t *testing.T) {
		var _ typapp.EntryPointer = typapp.Create(nil)
	})
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typapp.Commander = typapp.Create(nil)
	})
	t.Run("SHOULD implement Sourceable", func(t *testing.T) {
		var _ typcore.AppSourceable = typapp.Create(nil)
	})
}

func TestProvide(t *testing.T) {
	c1 := typdep.NewConstructor(nil)
	c2 := typdep.NewConstructor(nil)
	c3 := typdep.NewConstructor(nil)
	app := typapp.Create(dummyProvider(c1)).WithModules(dummyProvider(c2))
	typapp.AppendConstructor(c3)

	require.EqualValues(t,
		[]*typdep.Constructor{c1, c2, c3},
		app.Provide(),
	)
}

func TestDestoy(t *testing.T) {
	i1 := typdep.NewInvocation(nil)
	i2 := typdep.NewInvocation(nil)
	i3 := typdep.NewInvocation(nil)
	app := typapp.Create(dummyDestroyers(i1)).WithModules(dummyDestroyers(i2, i3))

	require.EqualValues(t,
		[]*typdep.Invocation{i1, i2, i3},
		app.Destroy(),
	)
}

func TestPrepare(t *testing.T) {
	i1 := typdep.NewInvocation(nil)
	i2 := typdep.NewInvocation(nil)
	i3 := typdep.NewInvocation(nil)
	app := typapp.Create(dummyPreparer(i1)).WithModules(dummyPreparer(i2, i3))

	require.EqualValues(t,
		[]*typdep.Invocation{i1, i2, i3},
		app.Prepare(),
	)
}

func TestEntryPoint(t *testing.T) {
	i1 := typdep.NewInvocation(nil)
	i2 := typdep.NewInvocation(nil)
	app := typapp.Create(dummyEntryPointer(i1)).WithModules(dummyEntryPointer(i2))

	require.EqualValues(t, i1, app.EntryPoint())
}

func TestApp(t *testing.T) {
	c1 := &cli.Command{}
	c2 := &cli.Command{}
	c3 := &cli.Command{}
	fn := struct{}{}

	app := typapp.
		Create(struct {
			typapp.EntryPointer
			typapp.Commander
		}{
			EntryPointer: dummyEntryPointer(typdep.NewInvocation(fn)),
			Commander:    dummyCommander(c1, c2),
		}).
		WithModules(dummyCommander(c3))

	cliApp := app.App(&typcore.Descriptor{
		Name:    "some-name",
		Version: "some-version",
	})

	require.EqualValues(t, []*cli.Command{c1, c2}, cliApp.Commands)
	require.Equal(t, "some-name", cliApp.Name)
	require.Equal(t, "some-version", cliApp.Version)
}
