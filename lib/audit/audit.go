package audit

import (
	"fmt"
	"github.com/ckeyer/commons/email"
)

func AuditError(msg ...interface{}) {
	var body string
	for _, v := range msg {
		body += fmt.Sprintln(v)
	}
	email.SendMail(email.DefaultAccount, "fxgc: [Error]", body, "dev@ckeyer.com")
}
