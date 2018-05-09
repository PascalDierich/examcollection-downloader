package ui

import "github.com/jroimartin/gocui"

func keybindings(g *gocui.Gui) error {
	// Any:
	// Quit Programm with <CTRL_C>
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		return err
	}

	// View: C - courses:
	// Select course with <ENTER>
	if err := g.SetKeybinding(C, gocui.KeyEnter, gocui.ModNone, selectCourse); err != nil {
		return err
	}
	// Go down with <ARROW_DOWN>
	if err := g.SetKeybinding(C, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	// Go up with <ARROW_UP>
	if err := g.SetKeybinding(C, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	// Change focus to exams with <ARROW_RIGHT>
	if err := g.SetKeybinding(C, gocui.KeyArrowRight, gocui.ModNone, changeFocus); err != nil {
		return err
	}

	// View: E - exams:
	// Download selected exams with <ENTER>
	if err := g.SetKeybinding(E, gocui.KeyEnter, gocui.ModNone, downloadExams); err != nil {
		return err
	}
	// Select exams with <SPACE>
	if err := g.SetKeybinding(E, gocui.KeySpace, gocui.ModNone, selectExam); err != nil {
		return err
	}
	// Download all exams with <CTRL_D>
	if err := g.SetKeybinding(E, gocui.KeyCtrlD, gocui.ModNone, downloadAll); err != nil {
		return err
	}
	// Go down with <ARROW_DOWN>
	if err := g.SetKeybinding(E, gocui.KeyArrowDown, gocui.ModNone, cursorDown); err != nil {
		return err
	}
	// Go up with <ARROW_UP>
	if err := g.SetKeybinding(E, gocui.KeyArrowUp, gocui.ModNone, cursorUp); err != nil {
		return err
	}
	// Change focus to courses with <ARROW_LEFT>
	if err := g.SetKeybinding(E, gocui.KeyArrowLeft, gocui.ModNone, changeFocus); err != nil {
		return err
	}

	return nil
}
