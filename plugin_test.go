package main
import (
	"testing"
    "os"
)

func TestFunction(t *testing.T){
    var creds = map[string] string {"user_key": os.Getenv("PUSHOVER_USER_KEY"), "api_token": os.Getenv("PUSHOVER_API_TOKEN")}
    NotificationPlugin.Init(creds)
    _, err := NotificationPlugin.SendMessage("test message")
    if err != nil {
        t.Error(err)
    }
}