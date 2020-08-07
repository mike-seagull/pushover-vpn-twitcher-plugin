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

func (pp *pushoverplugin) SendMessage(msg string) (string, error) {
    app := pushover.New(pp.api_token)
    recipient := pushover.NewRecipient(pp.user_key)
    message := pushover.NewMessageWithTitle(msg, "VPN-Twitcher")
    resp, err := app.SendMessage(message, recipient) // returns response and err
    if err != nil {
        fmt.Println(err)
    }
    return resp.String(), err
}

var NotificationPlugin pushoverplugin