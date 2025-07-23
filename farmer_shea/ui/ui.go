package ui

import (
	"fmt"

	"github.com/rivo/tview"
)

// UI is the user interface for the application.	ype UI struct {
	app         *tview.Application
	logView     *tview.TextView
	portfolioView *tview.Table
	pnlView     *tview.TextView
}

// New creates a new UI.
func New() *UI {
	app := tview.NewApplication()

	logView := tview.NewTextView().
		SetDynamicColors(true).
		SetRegions(true).
		SetWordWrap(true)

	portfolioView := tview.NewTable().
		SetBorders(true)

	pnlView := tview.NewTextView().
		SetBorders(true)

	flex := tview.NewFlex().
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(portfolioView, 0, 1, false).
			AddItem(pnlView, 0, 1, false), 0, 1, false).
		AddItem(logView, 0, 2, false)

	app.SetRoot(flex, true)

	return &UI{
		app:         app,
		logView:     logView,
		portfolioView: portfolioView,
		pnlView:     pnlView,
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

// UpdatePortfolio updates the portfolio view.
func (ui *UI) UpdatePortfolio(portfolio map[string]float64) {
	ui.app.QueueUpdateDraw(func() {
		ui.portfolioView.Clear()
		ui.portfolioView.SetCell(0, 0, tview.NewTableCell("Asset"))
		ui.portfolioView.SetCell(0, 1, tview.NewTableCell("Amount"))

		i := 1
		for asset, amount := range portfolio {
			ui.portfolioView.SetCell(i, 0, tview.NewTableCell(asset))
			ui.portfolioView.SetCell(i, 1, tview.NewTableCell(fmt.Sprintf("%.2f", amount)))
			i++
		}
	})
}

// UpdatePnL updates the PnL view.
func (ui *UI) UpdatePnL(pnl float64) {
	ui.app.QueueUpdateDraw(func() {
		ui.pnlView.SetText(fmt.Sprintf("PnL: %.2f", pnl))
	})
}
