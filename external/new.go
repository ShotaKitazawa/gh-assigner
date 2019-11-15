package external

import (
	"github.com/ShotaKitazawa/gh-assigner/controller"
	"github.com/ShotaKitazawa/gh-assigner/infrastructure/github"
	"github.com/ShotaKitazawa/gh-assigner/usecase"
)

func NewGitHubController() *controller.GitHubWebhookController {
	return &controller.GitHubWebhookController{
		Interactor: usecase.GitHubInteractor{
			GitRepository: &github.GitRepository{},
		},
	}
}
