package printer

import "github.com/pterm/pterm"

func Header(message string) {

	println()
	pterm.DefaultHeader.
		WithBackgroundStyle(pterm.NewStyle(pterm.BgLightMagenta)).
		WithTextStyle(pterm.NewStyle(pterm.FgBlack)).
		Println(message)
	println()
}

func Info(message string) {

	pterm.Info.Println(message)
}
