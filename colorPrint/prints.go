package colorPrint

import "github.com/fatih/color"

type Writer struct {
	// Headers
	warningHeader string
	errorHeader   string

	//Colors
	messageColor *color.Color
	warningColor *color.Color
	errorColor   *color.Color
}

func (w *Writer) SetDefaultMessages() {
	w.warningHeader = "--> Warning!"
	w.warningColor = color.New(color.FgYellow)

	w.errorHeader = "--> Error!"
	w.errorColor = color.New(color.FgRed)

	w.messageColor = color.New(color.FgCyan)
}

func (w *Writer) PrintWarning(s string) {
	w.warningColor.Println(w.warningHeader)
	w.warningColor.Println(s)
}

func (w *Writer) PrintError(s string) {
	w.errorColor.Println(w.errorHeader)
	w.errorColor.Println(s)
}

func (w *Writer) PrintMessage(s string) {
	w.messageColor.Println(s)
}
