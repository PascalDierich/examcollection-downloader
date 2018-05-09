# ExamCollection Downloader

`go get github.com/pascaldierich/examcollection-downloader`

The _examcollection-downloader_ is a terminal UI based application to support a __faster__ and __easier__ way to download old exams used by the TU Berlin.

## Internals
- All content gets downloaded from the [freitagsrunde-website](https://docs.freitagsrunde.org/Klausuren/).
- It creates a folder `old_exams`.
- Inside the `old_exams` folder it stores the exams in a corresponding course folder.

## Usage
Per default the `old_exams` folder is placed inside the home-directory, but you can change this through the
`-p` flag (in doubt run the `-h` flag).

#### Keybindings

- Move around the courses/exams or windows with the `ARROW` keys.
- To open the exams of one course hit `ENTER`.
- If you want to select a course for downloading use the `SPACE` key, and after that `ENTER` to download.
- In case you want to download __all__ exams of one course, run `CTRL-D`.
- Finally, use `CTRL-C` to __exit__.

## Screenshot

![screenshot](https://i.imgur.com/TuoEZQ6.png?raw=true)
