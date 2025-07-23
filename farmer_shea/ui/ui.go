package ui

import (
	"github.com/rivo/tview"
)

// UI is the user interface for the application.	ype UI struct {
	app     *tview.Application
	logView *tview.TextView
}

// New creates a new UI.
func New() *UI {
	app := tview.NewApplication()
	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	app.SetRoot(logView, true)

	return &UI{
		app:     app,
		logView: logView,
	}
}

// Run runs the UI.
func (ui *UI) Run() error {
	return ui.app.Run()
}

// Log logs a message to the log view.
func (ui *UI) Log(message string) {
	ui.app.QueueUpdateDraw(func() {
		ui.logView.Write([]byte(message))
	})
}
