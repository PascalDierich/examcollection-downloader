// Examcollection-Downloader is a terminal UI based application to download
// old exams from the TU Berlin.
//
// The old exams and other material is provided by the freitagsrunde,
// https://wiki.freitagsrunde.org/Hauptseite.
//
// The CLI creates a folder ``old_exams`` with path specified by the
// ``-p`` flag (defaults to home directory).
//
// Keybindings:
//
//		<arrow-keys> - move vertically and horizontally
//		<enter>      - open exams OR start dowload
//		<space>      - mark exam to download it
//		<ctrl-d>     - download all exams for given course
//		<ctrl-c>     - exit
//
package main

import (
	"errors"
	"flag"
	"io"
	"os"

	"github.com/jroimartin/gocui"
	"github.com/pascaldierich/examcollection-downloader/network"
	"github.com/pascaldierich/examcollection-downloader/ui"
)

const url = "https://docs.freitagsrunde.org/Klausuren/"

var p = flag.String("p", "", "Path to download the exams.")

func main() {
	flag.Parse()

	// Get working file-path for application-folder.
	path, err := initDirectory(*p)
	if err != nil {
		panic(err)
	}

	// Create new UI.
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	// Start view controller.
	err = viewController(g, path)
	if err != nil {
		panic(err)
	}

	// Start main loop.
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		panic(err)
	}
}

func viewController(g *gocui.Gui, path string) (err error) {
	var cl [][2]string         // Course list
	var el [][2]string         // Corresponding exams list
	lu := make(map[string]int) // Lookup map for course list
	le := make(map[string]int) // Lookup map for exams list

	var courseName string // Name of current selected course
	var courseLink string // Relative link to current selected course

	var selectedExams [][2]string // Selected Exams

	cl, err = network.GetCourseList(url)
	if err != nil {
		return
	}

	// Fill courses lookup map
	for i, c := range cl {
		lu[c[network.NameField]] = i
	}

	// View only needs the course-names, not the links.
	var names []string
	for _, n := range cl {
		names = append(names, n[network.NameField])
	}

	// Set view's functions:

	// ui.getExams returns the list of exam names
	// for the currently selected course.
	ui.GetExams = func() ([]string, error) {
		if courseLink == "" || courseName == "" {
			return nil, errors.New("no course selected")
		}

		el, err = network.GetExamList(url + courseLink)
		if err != nil {
			return nil, err
		}

		// Fill exams lookup map
		for i, e := range el {
			le[e[network.NameField]] = i
		}

		// Empty selected exams
		selectedExams = selectedExams[:0]

		// Return to view only the exam names, not the links.
		var names []string
		for _, n := range el {
			names = append(names, n[network.NameField])
		}

		return names, nil
	}

	// ui.DownloadExam downloads and saves the exam in corresponding course folder.
	ui.DownloadExams = func() (err error) {
		for _, e := range selectedExams {
			en := e[network.LinkField]
			f, err := network.DownloadFile(url + courseLink + en)
			if err != nil {
				return err
			}
			defer f.Close()

			err = createDir(path + courseName)
			if err != nil {
				return err
			}

			err = saveFile(f, path+courseName+en)
		}
		return
	}

	// ui.DownloadAllExams downloads and saves all exams in corresponding course folder.
	ui.DownloadAllExams = func() (err error) {
		// Fill selected exams list
		for _, s := range el {
			selectedExams = append(selectedExams, s)
		}
		return ui.DownloadExams()
	}

	// ui.CourseSelected set's the courseName and courseLink variables.
	ui.CourseSelected = func(cn string) error {
		i, ok := lu[cn]
		if !ok {
			return errors.New("could not find courseName in dataset. View edited?")
		}

		courseLink = cl[i][network.LinkField]
		courseName = cl[i][network.NameField]
		return nil
	}

	// ui.ExamSelected add a new exam to the "to-download-list".
	ui.ExamSelected = func(en string) error {
		i, ok := le[en]
		if !ok {
			return errors.New(en)
		}

		selectedExams = append(selectedExams, el[i])
		return nil
	}

	// Init UI.
	err = ui.Init(g, names)
	if err != nil {
		panic(err)
	}

	return
}

// saveFile write `f` to path `p`
func saveFile(f io.Reader, p string) (err error) {
	out, err := os.Create(p)
	if err != nil {
		return
	}

	_, err = io.Copy(out, f)
	return
}

// initDirectory returns the application path.
func initDirectory(home string) (path string, err error) {
	if home == "" {
		home, err = getHomeDir()
		if err != nil {
			return
		}
	}

	path = home + "/old_exams/"
	err = createDir(path)
	return
}

// getHomeDir return the path to the user's home directory.
func getHomeDir() (dir string, err error) {
	dir = os.Getenv("HOME")
	if dir == "" {
		err = errors.New("No HomeDir")
	}
	return
}

// createDir checks if directory exists, creates otherwise.
func createDir(p string) (err error) {
	_, err = os.Stat(p)
	if err == nil {
		return
	}
	if os.IsNotExist(err) {
		err = os.Mkdir(p, 0744)
	}
	return

}
