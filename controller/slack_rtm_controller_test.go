package controller

import (
	"testing"

	"github.com/ShotaKitazawa/gh-assigner/domain"
)

var (
	flagPong                             = false
	flagShowDefault                      = false
	flagShowHelp                         = false
	flagSendImageWithReviewWaitTimeGraph = false
)

type SlackRTMControllerInteractorMock struct{}
type SlackRTMControllerLoggerMock struct{}

func newSlackRTMController() *SlackRTMController {
	return &SlackRTMController{
		Interactor: &SlackRTMControllerInteractorMock{},
		Logger:     &SlackRTMControllerLoggerMock{},
	}
}
func (i SlackRTMControllerInteractorMock) Pong(d domain.SlackMessage) error {
	flagPong = true
	return nil
}
func (i SlackRTMControllerInteractorMock) ShowDefault(d domain.SlackMessage) error {
	flagShowDefault = true
	return nil
}
func (i SlackRTMControllerInteractorMock) ShowHelp(d domain.SlackMessage) error {
	flagShowHelp = true
	return nil
}
func (i SlackRTMControllerInteractorMock) SendImageWithReviewWaitTimeGraph(d domain.SlackMessage) error {
	flagSendImageWithReviewWaitTimeGraph = true
	return nil
}
func (l SlackRTMControllerLoggerMock) Debug(args ...interface{}) {
}
func (l SlackRTMControllerLoggerMock) Info(args ...interface{}) {
}
func (l SlackRTMControllerLoggerMock) Warn(args ...interface{}) {
}
func (l SlackRTMControllerLoggerMock) Error(args ...interface{}) {
}

func TestSlackRTMController(t *testing.T) {
	// Initialize
	t.Parallel()

	t.Run("MessageEvent()", func(t *testing.T) {
		// Initialize
		t.Parallel()

		t.Run("", func(t *testing.T) {
			//TODO
		})
	})
}
