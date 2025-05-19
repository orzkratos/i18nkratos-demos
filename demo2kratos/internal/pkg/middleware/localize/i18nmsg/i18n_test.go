package i18nmsg_test

import (
	"testing"

	"github.com/orzkratos/demokratos/demo2kratos/internal/pkg/middleware/localize/i18nmsg"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := i18nmsg.LoadI18nFiles(true)
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}
