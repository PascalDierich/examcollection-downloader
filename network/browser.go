// Package network implements the network functionality.
package network

import (
	"errors"
	"io"
	"net/http"

	"golang.org/x/net/html"
)

// NOTE: the functions GetCourseList and GetExamList are identical,
// 		 but it's better to have them seperated for future updates.

const (
	indexListTableID = "indexlist"
	indexHeadID      = "indexhead"
)

const (
	// LinkField is the spot to the Link field of the returned list.
	LinkField = iota
	// NameField is the equivalent for the Name field.
	NameField
)

// GetCourseList returns a list with the courses and the corresponding links to the exams.
func GetCourseList(url string) (list [][2]string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("error appeared while downloading the Indexpage")
	}
	defer resp.Body.Close()

	list, err = parseSite(resp.Body)
	if err != nil {
		err = errors.New("unexpected error appeared while parsing the Indexpage. Structure of Indexpage changed?")
	} else if list == nil {
		err = errors.New("error appeared while parsing the Indexpage")
	}

	// We need to throw out the "Parent Directory" field
	list = list[1:]
	return
}

// GetExamList returns a list with the exam name and the corresponding links to the files.
func GetExamList(url string) (list [][2]string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = errors.New("error appeared while downloading the Coursepage")
	}
	defer resp.Body.Close()

	list, err = parseSite(resp.Body)
	if err != nil {
		err = errors.New("unexpected error appeared while parsing the Coursepage. Structure of Coursepage changed?")
	} else if list == nil {
		err = errors.New("error appeared while parsing the Coursepage")
	}

	// We need to throw out the "Parent Directory" field
	list = list[1:]
	return
}

func parseSite(site io.Reader) (list [][2]string, err error) {
	t := html.NewTokenizer(site)

	// Find <table id="indexlist">
indexList:
	for {
		tokenType := t.Next()
		switch tokenType {
		case html.StartTagToken:
			token := t.Token()
			for _, a := range token.Attr {
				if a.Key == "id" && a.Val == indexListTableID {
					break indexList // found table
				}
			}
		}
	}
	t.Next() // skip <tbody> tag

	// Add all courses and links from the "even" and "odd" classes to the list
	for {
		tokenType := t.Next()
		switch tokenType {
		case html.StartTagToken:
			token := t.Token()
			for _, a := range token.Attr {
				if a.Key == "class" && a.Val == indexHeadID {
					skipIndexHead(t)
					continue // jump over element with id="index head"
				}

				if a.Key == "class" {
					if a.Val == "even" || a.Val == "odd" {
						t.Next()
						rs, err := getCourseAndLink(t)
						if err != nil {
							return nil, errors.New("unexpected design of tr-elements")
						}
						list = append(list, rs)
						break
					}
				}
			}
		case html.EndTagToken:
			token := t.Token()
			if token.Data == "table" {
				return // finished
			}
		}
	}
}

func getCourseAndLink(t *html.Tokenizer) (rs [2]string, err error) {
	defer func() {
		// We need to skip a few tags afterwards:
		t.Next() // </a>
		t.Next() // </td>
		t.Next() // td-tag class="indexcollastmod"
		t.Next() // inside of above tag
		t.Next() // end of above tag
		t.Next() // td-tag class="indexcolsize"
		t.Next() // inside of above tag
		t.Next() // end of above tag
		t.Next() // </tr>
	}()
	t.Next() // skip <td class="indexcolname">

	token := t.Token()
	switch token.Type {
	case html.StartTagToken:
		for _, a := range token.Attr {
			if a.Key == "href" {
				rs[LinkField] = a.Val // Link

				if t.Next() != html.TextToken {
					err = errors.New("missing Text Token")
				}

				rs[NameField] = t.Token().Data // Name
				return                         // finished
			}
			err = errors.New("missing href-key")
			break
		}
	default:
		err = errors.New("missing StartTag")
	}
	return
}

func skipIndexHead(t *html.Tokenizer) {
	for {
		t.Next()
		token := t.Token()
		switch token.Type {
		case html.EndTagToken:
			if token.Data == "tr" {
				return
			}
		}
	}
}
