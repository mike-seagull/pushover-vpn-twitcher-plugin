package main
import (
	"github.com/gregdel/pushover"
    "fmt"
)

type pushoverplugin struct {
    user_key string
    api_token string
}
func (pp *pushoverplugin) Init(options map[string]string){
    pp.user_key = options["user_key"]
    pp.api_token = options["api_token"]
}

func (pp *pushoverplugin) SendMessage(msg string) {
    app := pushover.New(pp.api_token)
    recipient := pushover.NewRecipient(pp.user_key)
    message := pushover.NewMessage(msg)
    _, err := app.SendMessage(message, recipient)
    if err != nil {
        fmt.Println(err)
    }
}

var NotificationPlugin pushoverplugin