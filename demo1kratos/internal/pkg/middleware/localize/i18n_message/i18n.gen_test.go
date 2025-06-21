package i18n_message_test

import (
	"testing"

	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/middleware/localize/i18n_message"
	"github.com/yyle88/goi18n"
	"github.com/yyle88/neatjson/neatjsons"
	"github.com/yyle88/osexistpath/osmustexist"
	"github.com/yyle88/runpath/runtestpath"
	"github.com/yyle88/zaplog"
)

//go:generate go test -v -run ^TestGenerate$
func TestGenerate(t *testing.T) {
	bundle, messageFiles := i18n_message.LoadI18nFiles(true)
	zaplog.SUG.Debugln(neatjsons.S(bundle.LanguageTags()))

	outputPath := osmustexist.FILE(runtestpath.SrcPath(t))
	options := goi18n.NewOptions().WithOutputPathWithPkgName(outputPath)
	t.Log(neatjsons.S(options))
	goi18n.Generate(messageFiles, options)
}
