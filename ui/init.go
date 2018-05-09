package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

var half = func(w int) int {
	return w / 2
}

// Init function for package UI.
func Init(g *gocui.Gui, cl []string) (err error) {
	g.Highlight = true
	g.SelFgColor = gocui.ColorBlue
	g.BgColor = gocui.ColorBlack
	g.FgColor = gocui.ColorWhite

	g.SetManagerFunc(layout)

	terminalWidth, terminalHeight := g.Size()

	// Course view.
	courseView, err := g.SetView(C, 0, 0,
		half(terminalWidth), terminalHeight-1)
	if err != nil && err != gocui.ErrUnknownView {
		return err
	}

	courseView.Title = "Courses"
	courseView.FgColor = gocui.ColorWhite
	courseView.Highlight = true
	courseView.SelBgColor = gocui.ColorRed
	courseView.SelFgColor = gocui.ColorWhite

	// Exam view.
	examView, err := g.SetView(E, half(terminalWidth)+1, 0,
		terminalWidth, terminalHeight-1)
	if err != nil && err != gocui.ErrUnknownView {
		return
	}

	examView.Title = "Exams"
	examView.FgColor = gocui.ColorWhite
	examView.Highlight = true
	examView.SelBgColor = gocui.ColorRed
	examView.SelFgColor = gocui.ColorWhite

	// Set keybindings.
	err = keybindings(g)
	if err != nil {
		return
	}

	// Set inital view.
	_, err = g.SetCurrentView(C)

	// Update view with course-list
	g.Update(
		func(g *gocui.Gui) error {
			v, err := g.View(C)
			if err != nil {
				return err
			}
			v.Clear()

			for _, n := range cl {
				fmt.Fprintln(v, n)
			}
			return nil
		})

	return
}

// layout handler calculates all sizes depending on the current terminal size.
func layout(g *gocui.Gui) (err error) {
	tw, th := g.Size()

	// Course view.
	_, err = g.SetView(C, 0, 0, half(tw), th-1)
	if err != nil {
		return
	}

	// Exam view.
	_, err = g.SetView(E, half(tw)+1, 0, tw, th-1)
	if err != nil {
		return
	}

	return
}
