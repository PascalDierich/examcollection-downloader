package main

import (
	"fmt"
	"testing"

	"github.com/pascaldierich/examcollection-downloader/network"
)

func TestViewController(t *testing.T) {
	el, err := network.GetExamList(url + "TheGI_3_Bachelor/")
	if err != nil {
		t.Error(err)
	}

	// Return to view only the exam names, not the links.
	var names []string
	for _, n := range el {
		names = append(names, n[network.NameField])
	}

	for _, n := range names {
		fmt.Println(n)
	}

	el, err = network.GetExamList(url + "TheGI_4/")
	if err != nil {
		t.Error(err)
	}

	// Return to view only the exam names, not the links.
	names = nil
	for _, n := range el {
		names = append(names, n[network.NameField])
	}

	for _, n := range names {
		fmt.Println(n)
	}
}

func TestCreateDir(t *testing.T) {
	dir, err := getHomeDir()
	if err != nil {
		t.Error("could not read home directory: " + err.Error())
	}

	dir = dir + "/Desktop/test"
	err = createDir(dir)
	if err != nil {
		t.Error(err)
	}
}

func TestInitDirectory(t *testing.T) {
	dir, err := getHomeDir()
	if err != nil {
		t.Error("could not read home directory: " + err.Error())
	}

	_, err = initDirectory(dir)
	if err != nil {
		t.Error(err)
	}
}

func TestGetFileExtension(t *testing.T) {
	from := []string{
		"hallo.txt",
		"hal.lo.txt",
		"hallo..pdf",
		".hallo.txt",
		"hallo",
	}
	want := []string{
		".txt",
		".txt",
		".pdf",
		".txt",
		"",
	}
	var got []string

	for _, s := range from {
		g := getFileExtension(s)
		got = append(got, g)
	}

	for i, s := range got {
		if want[i] != s {
			t.Errorf("got=%s, want=%s", s, want[i])
		}
	}
}
