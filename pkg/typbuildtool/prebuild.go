package typbuildtool

import (
	"context"

	"github.com/typical-go/typical-go/pkg/typcore"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcore/walker"
	"github.com/typical-go/typical-go/pkg/typenv"
)

func prebuild(ctx context.Context, d *typcore.ProjectDescriptor) (err error) {
	var (
		stdPrebuilder typcore.StandardPrebuilder
		projInfo      typcore.ProjectInfo
		events        walker.Declarations
	)
	if projInfo, err = typcore.ReadProject(typenv.Layout.App); err != nil {
		log.Fatal(err.Error())
	}
	log.Info("Walk the project")
	if events, err = walker.Walk(projInfo.Files); err != nil {
		return
	}
	pc := &typcore.PrebuildContext{
		Context:           ctx,
		ProjectDescriptor: d,
		ProjectInfo:       projInfo,
		Declarations:      events,
	}
	if err = stdPrebuilder.Prebuild(pc); err != nil {
		return
	}
	return
}
