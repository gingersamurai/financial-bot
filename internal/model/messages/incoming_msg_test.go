package messages

import (
	"testing"

	mocks "github.com/gingersamurai/financial-bot/internal/mocks/messages"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_OnStartCommand_ShouldAnswerWithIntroMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl)
	model := New(sender)

	sender.EXPECT().SendMessage("hello", int64(123))

	err := model.IncomingMessage(
		Message{
			Text:   "/start",
			UserID: 123,
		},
	)

	assert.NoError(t, err)
}

func Test_OnUnknownCommand_ShouldAnswerWithHelpMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	sender := mocks.NewMockMessageSender(ctrl) //моком сделали клиент
	model := New(sender)

	sender.EXPECT().SendMessage("я не знаю эту команду", int64(123))

	err := model.IncomingMessage(
		Message{
			Text:   "abacaba",
			UserID: 123,
		},
	)

	assert.NoError(t, err)
}
