package image

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

const (
	titleReviewWaitTime = "Review Wait Time"
	xReviewWaitTime     = "Reviewed PR ID"
	yReviewWaitTime     = "Minutes"
)

// ImageInfrastructure is Infrastructure
type ImageInfrastructure struct {
	Path string
}

func (r ImageInfrastructure) CreateGraphWithReviewWaitTime(times map[uint]time.Duration) (filepath string, err error) {
	dirname, err := r.Initialize()
	if err != nil {
		return
	}

	p, err := plot.New()
	if err != nil {
		return
	}

	p.Title.Text = titleReviewWaitTime
	p.X.Label.Text = xReviewWaitTime
	p.Y.Label.Text = yReviewWaitTime

	var nominal []string
	var pts plotter.Values
	for idx, val := range times {
		nominal = append(nominal, strconv.Itoa(int(idx)))
		pts = append(pts, float64(int64(val/time.Minute)))
	}
	if len(nominal) != 0 {
		p.NominalX(nominal...)
	}

	bars, err := plotter.NewBarChart(pts, vg.Points(20))
	if err != nil {
		return
	}
	p.Add(bars)

	imagePath := joinPath(dirname, randString(8)+".png")

	// Save the plot to a PNG file.
	if err = p.Save(4*vg.Inch, 4*vg.Inch, imagePath); err != nil {
		return
	}

	filepath = imagePath
	return
}

func (r ImageInfrastructure) DeleteFile(filepath string) (err error) {
	if err != nil {
		return
	}
	return os.Remove(filepath)
}

func (r ImageInfrastructure) Initialize() (dirname string, err error) {
	if r.Path == "" {
		r.Path, err = os.Getwd()
		if err != nil {
			return
		}
	}

	err = initDir(r.Path)
	if err != nil {
		return
	}

	dirname = r.Path
	return
}
