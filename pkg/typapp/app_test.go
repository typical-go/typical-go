package typapp_test

import (
	"testing"

	"github.com/urfave/cli/v2"

	"github.com/stretchr/testify/require"

	"github.com/typical-go/typical-go/app"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func TestTypicalApp(t *testing.T) {
	t.Run("SHOULD implement App", func(t *testing.T) {
		var _ typcore.App = &typapp.TypicalApp{}
	})
	t.Run("SHOULD implement Preconditioner", func(t *testing.T) {
		var _ typbuildtool.Preconditioner = &typapp.TypicalApp{}
	})
	t.Run("SHOULD implement Provider", func(t *testing.T) {
		var _ typapp.Provider = &typapp.TypicalApp{}
	})
	t.Run("SHOULD implement Destroyer", func(t *testing.T) {
		var _ typapp.Destroyer = &typapp.TypicalApp{}
	})
	t.Run("SHOULD implement Preparer", func(t *testing.T) {
		var _ typapp.Preparer = &typapp.TypicalApp{}
	})
	t.Run("SHOULD implement EntryPointer", func(t *testing.T) {
		var _ typapp.EntryPointer = &typapp.TypicalApp{}
	})
	t.Run("SHOULD implement Commander", func(t *testing.T) {
		var _ typapp.Commander = &typapp.TypicalApp{}
	})
}

func TestAppModule(t *testing.T) {
	t.Run("GIVEN empty app sources", func(t *testing.T) {
		require.Equal(t,
			[]string{"app"},
			typapp.AppModule(app.New()).AppSources(),
		)
	})
	t.Run("GIVEN some app sources", func(t *testing.T) {
		require.Equal(t,
			[]string{"some-app-sources"},
			typapp.AppModule(app.New(), "some-app-sources").AppSources(),
		)
	})
}

func TestProvide(t *testing.T) {
	c1 := typapp.NewConstructor(nil)
	c2 := typapp.NewConstructor(nil)
	c3 := typapp.NewConstructor(nil)
	app := typapp.AppModule(dummyProvider(c1)).WithModules(dummyProvider(c2))
	typapp.AppendConstructor(c3)

	require.EqualValues(t,
		[]*typapp.Constructor{c1, c2, c3},
		app.Provide(),
	)
}

func TestDestoy(t *testing.T) {
	i1 := typapp.NewDestruction(nil)
	i2 := typapp.NewDestruction(nil)
	i3 := typapp.NewDestruction(nil)
	app := typapp.AppModule(dummyDestroyers(i1)).WithModules(dummyDestroyers(i2, i3))

	require.EqualValues(t,
		[]*typapp.Destruction{i1, i2, i3},
		app.Destroy(),
	)
}

func TestPrepare(t *testing.T) {
	i1 := typapp.NewPreparation(nil)
	i2 := typapp.NewPreparation(nil)
	i3 := typapp.NewPreparation(nil)
	app := typapp.AppModule(dummyPreparer(i1)).WithModules(dummyPreparer(i2, i3))

	require.EqualValues(t,
		[]*typapp.Preparation{i1, i2, i3},
		app.Prepare(),
	)
}

func TestEntryPoint(t *testing.T) {
	i1 := typapp.NewMainInvocation(nil)
	i2 := typapp.NewMainInvocation(nil)
	app := typapp.AppModule(dummyEntryPointer(i1)).WithModules(dummyEntryPointer(i2))

	require.EqualValues(t, i1, app.EntryPoint())
}

func TestApp(t *testing.T) {
	c1 := &cli.Command{}
	c2 := &cli.Command{}
	c3 := &cli.Command{}
	fn := struct{}{}

	app := typapp.
		AppModule(struct {
			typapp.EntryPointer
			typapp.Commander
		}{
			EntryPointer: dummyEntryPointer(typapp.NewMainInvocation(fn)),
			Commander:    dummyCommander(c1, c2),
		}).
		WithModules(dummyCommander(c3))

	cliApp := app.App(&typcore.Descriptor{
		Name:    "some-name",
		Version: "some-version",
	})

	require.EqualValues(t, []*cli.Command{c1, c2, c3}, cliApp.Commands)
	require.Equal(t, "some-name", cliApp.Name)
	require.Equal(t, "some-version", cliApp.Version)
}
