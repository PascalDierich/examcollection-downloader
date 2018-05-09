// Package ui provides the views, controls and functions for the terminal UI.
package ui

import (
	"fmt"

	"github.com/jroimartin/gocui"
)

const (
	// C is the left list, showing the courses.
	C = "courses"
	// E is the right list, shwowing the corresponding exams.
	E = "exams"
)

// This functions are set in the view-controllers init process
// and working as boundary between the UI and the use-cases.
var (
	// GetExams returns the list of exam names for the course hold by view-controller.
	GetExams func() (list []string, err error)

	// DownloadExams downloads and saves the exam.
	DownloadExams func() error

	// DownloadAllExams downloads and saves all exams for the selected course.
	DownloadAllExams func() error

	// CourseSelected transmits the current selected course to the view-controller.
	CourseSelected func(courseName string) error

	// ExamSelected add a new exam to view-controllers "to-download"-exam list.
	ExamSelected func(examName string) error
)

/*
 * View C
 */

func selectCourse(g *gocui.Gui, v *gocui.View) (err error) {
	_, y := v.Cursor()
	c, err := v.Line(y)
	if err != nil {
		return
	}

	CourseSelected(c)
	v, err = g.SetCurrentView(E)
	if err != nil {
		return
	}
	showExams(g, v)

	return
}

/*
 * View E
 */

// Mark exam as selected
var markExam = func(v *gocui.View, y int) {
	v.SetCursor(0, y)
	v.EditWrite('*')
	v.EditWrite(' ')
}

func showExams(g *gocui.Gui, v *gocui.View) (err error) {
	exs, err := GetExams()
	if err != nil {
		return
	}

	v.Clear()
	for _, s := range exs {
		fmt.Fprintln(v, s)
	}
	v.Line(0)

	return
}

func downloadExams(g *gocui.Gui, v *gocui.View) (err error) {
	err = DownloadExams()
	if err != nil {
		return
	}

	v.Clear()
	v, err = g.SetCurrentView(C)
	if err != nil {
		return
	}

	return
}

func selectExam(g *gocui.Gui, v *gocui.View) (err error) {
	_, y := v.Cursor()
	e, err := v.Line(y)
	if err != nil {
		return
	}

	err = ExamSelected(e)
	if err != nil {
		return
	}
	markExam(v, y)

	return
}

func downloadAll(g *gocui.Gui, v *gocui.View) (err error) {
	err = DownloadAllExams()
	if err != nil {
		return
	}

	v.Clear()
	v, err = g.SetCurrentView(C)
	if err != nil {
		return
	}
	return
}

/*
 * Any
 */

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func cursorUp(g *gocui.Gui, v *gocui.View) (err error) {
	v.MoveCursor(0, -1, false)
	return
}

func cursorDown(g *gocui.Gui, v *gocui.View) (err error) {
	v.MoveCursor(0, 1, false)
	return
}

func changeFocus(g *gocui.Gui, v *gocui.View) (err error) {
	switch v.Name() {
	case E:
		_, err = g.SetCurrentView(C)
	case C:
		_, err = g.SetCurrentView(E)
	}
	return
}
