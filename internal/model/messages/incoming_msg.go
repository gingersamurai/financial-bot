package messages

import (
	"log"
	"strings"

	"github.com/gingersamurai/financial-bot/internal/model/storage"
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

var myMemoryStorage storage.MemoryStorage

func (s *Model) IncomingMessage(msg Message) error {
	switch {
	case msg.Text == "/start":
		myMemoryStorage = storage.NewMemoryStorage()
		return s.tgClient.SendMessage("hello", msg.UserID)

	case strings.HasPrefix(msg.Text, "/addSpending"):
		err := addSpending(msg, &myMemoryStorage)
		if err != nil {
			log.Println(err)
			return s.tgClient.SendMessage(err.Error(), msg.UserID)
		}
		return s.tgClient.SendMessage("трата добавлена", msg.UserID)
	case strings.HasPrefix(msg.Text, "/getReport"):
		result, err := getReport(msg, &myMemoryStorage)
		if err != nil {
			log.Println(err)
			return s.tgClient.SendMessage(err.Error(), msg.UserID)
		}

		return s.tgClient.SendMessage(result, msg.UserID)
	default:
		return s.tgClient.SendMessage("я не знаю эту команду", msg.UserID)
	}
}
