package localize

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/orzkratos/demokratos/demo1kratos/internal/pkg/middleware/localize/i18nmsg"
	"github.com/yyle88/tern/zerotern"
)

type localizerKey struct{}

// I18N 组件用于翻译
// cp from: https://github.com/go-kratos/examples/blob/3a46aa32f7dbecbb01f2e3ecb28af187b2d9b53c/i18n/internal/pkg/middleware/localize/localize.go#L16
func I18N() middleware.Middleware {
	bundle, _ := i18nmsg.LoadI18nFiles(false)

	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var accept string
			// parse accept-language from transport context
			if tr, ok := transport.FromServerContext(ctx); ok {
				accept = tr.RequestHeader().Get("accept-language") //这里有可能得到个空字符串
			}
			// when accept == "" set accept = defaultValue
			accept = zerotern.VV(accept, i18nmsg.DefaultLanguage.String()) //在这里设置默认语言

			localizer := i18n.NewLocalizer(bundle, accept)
			ctx = context.WithValue(ctx, localizerKey{}, localizer)
			return handler(ctx, req)
		}
	}
}

func FromContext(ctx context.Context) *i18n.Localizer {
	return ctx.Value(localizerKey{}).(*i18n.Localizer)
}
