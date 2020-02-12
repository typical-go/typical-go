package typbuild

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typenv"
)

func (b *Build) clean(ctx context.Context, c *Context) error {
	removeFile(typenv.AppBin)
	removeAllFile(typenv.Layout.Temp)
	removeFile(".env")
	removeFile(typenv.GeneratedConstructor)
	return nil
}

func removeFile(name string) {
	log.Infof("Remove: %s", name)
	if err := os.Remove(name); err != nil {
		log.Error(err.Error())
	}
}

func removeAllFile(path string) {
	log.Infof("Remove All: %s", path)
	if err := os.RemoveAll(path); err != nil {
		log.Error(err.Error())
	}
}
