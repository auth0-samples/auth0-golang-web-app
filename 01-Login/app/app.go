package app

import (
	"github.com/astaxie/beego/session"
)

var (
	GlobalSessions *session.Manager
)

func Init() {
	GlobalSessions, _ = session.NewManager("memory", &session.ManagerConfig{
		CookieName: "gosessionid",
		Gclifetime: 3600,
	})
	go GlobalSessions.GC()
}
