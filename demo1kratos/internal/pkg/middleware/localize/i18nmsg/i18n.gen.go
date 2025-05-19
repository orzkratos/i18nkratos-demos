package i18nmsg

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

type ErrormsgParam struct {
	ErrorMessage any
}

func (p *ErrormsgParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.ErrorMessage != nil {
		res["ErrorMessage"] = p.ErrorMessage
	}
	return res
}

func NewErrormsg(data *ErrormsgParam) (string, map[string]any) {
	return "errormsg", data.GetTemplateValues()
}

func I18nErrormsg(data *ErrormsgParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "errormsg",
		TemplateData: data.GetTemplateValues(),
	}
}

type GreetingParam struct {
	Name any
}

func (p *GreetingParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.Name != nil {
		res["Name"] = p.Name
	}
	return res
}

func NewGreeting(data *GreetingParam) (string, map[string]any) {
	return "greeting", data.GetTemplateValues()
}

func I18nGreeting(data *GreetingParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "greeting",
		TemplateData: data.GetTemplateValues(),
	}
}

type WelcomeParam struct {
	User any
}

func (p *WelcomeParam) GetTemplateValues() map[string]any {
	res := make(map[string]any)
	if p.User != nil {
		res["User"] = p.User
	}
	return res
}

func NewWelcome(data *WelcomeParam) (string, map[string]any) {
	return "welcome", data.GetTemplateValues()
}

func I18nWelcome(data *WelcomeParam) *i18n.LocalizeConfig {
	return &i18n.LocalizeConfig{
		MessageID:    "welcome",
		TemplateData: data.GetTemplateValues(),
	}
}
