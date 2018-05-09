package network

import (
	"io"
	"os"
	"strings"
	"testing"
)

const (
	urlIndexpage   = "https://docs.freitagsrunde.org/Klausuren/"
	urlAnaIngIPage = "https://docs.freitagsrunde.org/Klausuren/Analysis_1_fuer_Ingenieure/"
)

func TestGetCourseList(t *testing.T) {
	lastElem := "Vorlage%20Pruefungsprotokoll.sxw"
	lastElem2 := "Vorlage Pruefungsprotokoll.sxw"

	notFirstElem := "/"
	notFirstElem2 := "Parent Directory"

	firstElem := "ABWL_I/"
	firstElem2 := "ABWL_I/"

	somewhere := [][]string{
		{"Semantic_Search/", "Semantic_Search/"},
		{"Telekommunikationsnetze/", "Telekommunikationsnetze/"},
	}

	got, err := GetCourseList(urlIndexpage)
	if err != nil {
		t.Error(err.Error())
	}

	if len(got) == 0 {
		t.Error("empty result")
	}
	if got[len(got)-1][0] != lastElem {
		t.Error("last elem missing")
	}
	if got[len(got)-1][1] != lastElem2 {
		t.Error("last elem wrong link")
	}
	if got[0][0] == notFirstElem {
		t.Error("first elem should be deleted")
	}
	if got[0][1] == notFirstElem2 {
		t.Error("first elem contains wrong link")
	}
	if got[0][0] != firstElem {
		t.Error("first elem missing")
	}
	if got[0][1] != firstElem2 {
		t.Error("first elem contains wrong link")
	}

	sw := 0
	for _, s := range got {
		for _, x := range somewhere {
			if s[0] == x[0] {
				sw++
			}
			if s[1] == x[1] {
				sw++
			}
		}
	}
	if sw != len(somewhere)*2 { // multiply with 2 (2-dim array)
		t.Error("some elems are missing")
	}
}

func TestGetExamList(t *testing.T) {
	// This test is for the "Analysis_I_fuer_Ingenieure" page

	lastElem := "2018.02_muendlich_AnaLinaKombi.pdf"
	lastElem2 := "2018.02_muendlich_AnaLinaKombi.pdf"

	notFirstElem := "/"
	notFirstElem2 := "Parent Directory"

	firstElem := "2001.01_probeklausur_loesung.pdf"
	firstElem2 := "2001.01_probeklausur_loesung.pdf"

	somewhere := [][]string{
		{"2003.07_B_RE.pdf", "2003.07_B_RE.pdf"},
		{"2012.03_klausur.pdf", "2012.03_klausur.pdf"},
	}

	got, err := GetCourseList(urlAnaIngIPage)
	if err != nil {
		t.Error(err.Error())
	}

	if len(got) == 0 {
		t.Error("empty result")
	}
	if got[len(got)-1][0] != lastElem {
		t.Error("last elem missing")
	}
	if got[len(got)-1][1] != lastElem2 {
		t.Error("last elem wrong link")
	}
	if got[0][0] == notFirstElem {
		t.Error("first elem should be deleted")
	}
	if got[0][1] == notFirstElem2 {
		t.Error("first elem contains wrong link")
	}
	if got[0][0] != firstElem {
		t.Error("first elem missing")
	}
	if got[0][1] != firstElem2 {
		t.Error("first elem contains wrong link")
	}

	sw := 0
	for _, s := range got {
		for _, x := range somewhere {
			if s[0] == x[0] {
				sw++
			}
			if s[1] == x[1] {
				sw++
			}
		}
	}
	if sw != len(somewhere)*2 { // multiply with 2 (2-dim array)
		t.Error("some elems are missing")
	}
}

func TestDownloadFile(t *testing.T) {
	// This test only works when checking the downloaded file by hand...

	pdfURL := strings.Join([]string{urlAnaIngIPage, "2001.01_probeklausur_loesung.pdf"}, "")
	txtURL := "https://docs.freitagsrunde.org/Klausuren/Advanced_Computer_Architectures/Ged%c3%a4chtnisprotokoll_WiSe_13_14.txt"

	dir := os.Getenv("HOME")
	if dir == "" {
		t.Error("error getting 'HOME'-variable")
	}

	// test pdf file
	f, err := DownloadFile(pdfURL)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	out, err := os.Create(dir + "/Desktop/test.pdf")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(out, f)
	if err != nil {
		t.Error(err)
	}

	// test txt file
	f, err = DownloadFile(txtURL)
	if err != nil {
		t.Error(err)
	}
	defer f.Close()

	out, err = os.Create(dir + "/Desktop/test.txt")
	if err != nil {
		t.Error(err)
	}

	_, err = io.Copy(out, f)
	if err != nil {
		t.Error(err)
	}
}
