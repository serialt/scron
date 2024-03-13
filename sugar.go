package main

import (
	"os"

	"log/slog"

	"github.com/robfig/cron/v3"
	"github.com/serialt/crab"
)

func service() {
	slog.Debug("debug msg")
	slog.Info("info msg")
	slog.Error("error msg")

	// 定时任务
	c := cron.New()

	for _, v := range config.Cron {
		// sender := &smail.Mailer{
		// 	User:    config.Mailer.User,
		// 	Pass:    config.Mailer.Pass,
		// 	Smtp:    config.Mailer.Smtp,
		// 	Port:    config.Mailer.Port,
		// 	MailTo:  config.Mailer.MailTo,
		// 	Subject: v.Subject,
		// 	Body:    v.Msg,
		// }
		c.AddFunc(v.Time, func() {
			result, err := crab.RunCMD(v.Action, "/tmp")
			if err != nil {
				slog.Info("run cmd succeed", "cmd", v.Action, "result", result)
			}
			slog.Debug("run cmd", "cmd", v.Action)
		})

	}

	c.Start()
}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}
