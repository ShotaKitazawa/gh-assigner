package controller

import (
	"context"

	"github.com/ShotaKitazawa/gh-assigner/controller/interfaces"
)

// CronController is Controller
type CronController struct {
	Interactor interfaces.CronInteractor
	Logger     interfaces.Logger
}

// Event is called by Cron
func (c CronController) Event(ctx context.Context) (err error) {
	/*
	organization := ctx.Get("organization")
	repository := ctx.Get("repository")
	period := ctx.Get("period")

	err = c.Interactor.SendImageWithReviewWaitTimeGraph(organization, repository, period)
	*/

	return nil
}
