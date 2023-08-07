package main

import (
	"time"
	"strings"
	"fmt"
	"io"
)

type DocWriter struct {
	W io.Writer

	// Depth is of stars for next header
	Depth uint
}

func (doc *DocWriter) emit(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	_, err := fmt.Fprint(doc.W, s)
	if err != nil {
		panic("Failure emitting org to file handle")
	}
}

func (doc *DocWriter) emitln(format string, a ...any) {
	doc.emit(format + "\n", a...)
}

func (doc *DocWriter) EmitPreamble() {
	doc.emitln("#+STARTUP: content logdone")
}

func (doc *DocWriter) EmitHeader(format string, a ...any) {
	stars := strings.Repeat("*", int(doc.Depth))
	if doc.Depth != 1 {
		stars = "\n" + stars
	}
	
	s := fmt.Sprintf(format, a...)
	doc.emitln(stars + " " + s)
}


func (doc *DocWriter) EmitContentLn(format string, a...any) {
	doc.emitln(format, a...)
}

func (doc *DocWriter) EmitListCheckboxLn(format string, a...any) {
	s := fmt.Sprintf(format, a...)
	doc.emitln(" - [ ] %s", s)
}

func FormatDate(t time.Time) string {
	return fmt.Sprintf("[%s]", t.Format("2006-01-02 Mon 15:04"))
}
