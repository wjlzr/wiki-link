package sidekiqs

import (
	"bytes"
	"encoding/json"
	"html/template"
	"reflect"
	"time"

	"github.com/oldfritter/sidekiq-go"

	"wiki-link/baseServices/email"
)

func CreateSendEmail(w *sidekiq.Worker) sidekiq.WorkerI {
	return &SendEmail{*w, email.MailerPayload{}}
}

type SendEmail struct {
	sidekiq.Worker
	Payload email.MailerPayload
}

func (worker *SendEmail) Work() (err error) {
	start := time.Now().UnixNano()
	reflect.ValueOf(&worker.Payload).MethodByName(worker.Payload.I18nFuncName()).Call([]reflect.Value{})
	t, err := template.New(worker.Payload.Method+".html").Funcs(worker.Payload.FuncMap).ParseFiles(
		"public/workers/email/"+worker.Payload.Mailer+"/"+worker.Payload.Method+".html",
		"public/workers/email/head.html",
		"public/workers/email/footer.html",
		"public/workers/email/content.html",
	)
	if err != nil {
		return
	}
	var tpl bytes.Buffer
	if err = t.Execute(&tpl, worker.Payload); err != nil {
		return
	}
	if err = email.SendMail(worker.Payload.Email, worker.Payload.Subject, tpl.String()); err != nil {
		return
	}
	worker.LogInfo("payload: ", worker.Payload, ", time:", (time.Now().UnixNano()-start)/1000000, " ms")
	return
}

func (worker *SendEmail) SetPayload(payload string) {
	json.Unmarshal([]byte(payload), &worker.Payload)
}
