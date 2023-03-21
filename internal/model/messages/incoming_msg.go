package messages

import (
	"fmt"
	"strings"
)

type MessageSender interface {
	SendMessage(text string, userID int64) error
}

type Model struct {
	tgClient MessageSender
}

func New(tgClient MessageSender) *Model {
	return &Model{
		tgClient: tgClient,
	}
}

type Message struct {
	Text   string
	UserID int64
}

func (s *Model) IncomingMessage(msg Message) error {
	switch {
	case msg.Text == "/start":
		myStorage = NewMemoryStorage()
		// if myStorage.storage == nil {
		// 	log.Fatal("it is nil even in start")
		// } else {
		// 	log.Println("in is not nil in start")
		// }
		return s.tgClient.SendMessage("hello", msg.UserID)

	case strings.HasPrefix(msg.Text, "/addSpending"):
		// if myStorage.storage == nil {
		// 	log.Fatal("it is nil even in /addSpending handlenr")
		// } else {
		// 	log.Println("it is not nil even in /addSpending handlenr")
		// }
		addSpending(msg, &myStorage)
		return s.tgClient.SendMessage(fmt.Sprintln(myStorage.storage), msg.UserID)
	default:
		return s.tgClient.SendMessage("я не знаю эту команду", msg.UserID)
	}
}
