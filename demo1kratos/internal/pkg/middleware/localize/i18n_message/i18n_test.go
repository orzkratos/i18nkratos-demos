package i18n_message_test

import (
	"testing"

	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/middleware/localize/i18n_message"
)

func TestLoadI18nFiles(t *testing.T) {
	bundle, messageFiles := i18n_message.LoadI18nFiles(true)
	t.Log(len(messageFiles))
	t.Log(len(bundle.LanguageTags()))
}
