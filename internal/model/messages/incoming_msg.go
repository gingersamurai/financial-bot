package messages

import (
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
		memoryStorage = NewMemoryStorage()
		return s.tgClient.SendMessage("hello", msg.UserID)

	case strings.HasPrefix(msg.Text, "/addSpending"):
		addSpending(msg, &memoryStorage)
		return s.tgClient.SendMessage("трата добавлена", msg.UserID)
	default:
		return s.tgClient.SendMessage("я не знаю эту команду", msg.UserID)
	}
}
