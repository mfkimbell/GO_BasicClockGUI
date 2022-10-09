package main

import (
	"time"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func updateClock(clock *widget.Label) {
	var format = time.Now().Format("Time: 03:04:05")
	clock.SetText(format)
}

func main() {
	var a = app.New()
	var w = a.NewWindow("Clock")

	clock := widget.NewLabel("")
	updateClock(clock)

	w.SetContent(clock)
	go func() {

		for range time.Tick(time.Second) {

			updateClock(clock)
		}
	}()
	w.ShowAndRun()
}
