package email

import (
	"fmt"
	"html/template"
	"strings"

	. "wiki-link/i18n"
)

type MailerPayload struct {
	Mailer  string `json:"mailer"`
	Method  string `json:"method"`
	Locale  string `json:"locale"`
	Subject string `json:"subject"`
	Email   string `json:"email"`
	Content string `json:"content"`
	Times   int    `json:"times"`
	FuncMap template.FuncMap
}

func (mp *MailerPayload) I18nFuncName() string {
	return strings.Join([]string{strings.Title(mp.Mailer), strings.Title(mp.Method)}, "")
}

func (mp *MailerPayload) NotifyOKLink500() {
	mp.Subject = fmt.Sprint(I18n.T(mp.Locale, "mailer.notify.OKLink500.subject"))
	mp.FuncMap = template.FuncMap{
		"title": func() string {
			return fmt.Sprint(I18n.T(mp.Locale, "mailer.notify.OKLink500.title", map[string]interface{}{"email": mp.Email}))
		},
		"content": func() string {
			return fmt.Sprint(I18n.T(mp.Locale, "mailer.notify.OKLink500.content", map[string]interface{}{"content": mp.Content}))
		},
		"foot": func() template.HTML {
			return template.HTML(fmt.Sprint(I18n.T(mp.Locale, "mailer.footer.contact", map[string]interface{}{"contact": SmtpConfig.Username})))
		},
	}
}

func (mp *MailerPayload) NotifyTokenViewKeyExpired() {
	mp.Subject = fmt.Sprint(I18n.T(mp.Locale, "mailer.notify.TokenViewKeyExpired.subject"))
	mp.FuncMap = template.FuncMap{
		"title": func() string {
			return fmt.Sprint(I18n.T(mp.Locale, "mailer.notify.TokenViewKeyExpired.title", map[string]interface{}{"email": mp.Email}))
		},
		"content": func() string {
			return fmt.Sprint(I18n.T(mp.Locale, "mailer.notify.TokenViewKeyExpired.content", map[string]interface{}{"content": mp.Content}))
		},
		"foot": func() template.HTML {
			return template.HTML(fmt.Sprint(I18n.T(mp.Locale, "mailer.footer.contact", map[string]interface{}{"contact": SmtpConfig.Username})))
		},
	}
}
