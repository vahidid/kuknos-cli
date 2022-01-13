package helpers

import (
	"github.com/pterm/pterm"
)

func IntroScreen() {
	ptermLogo, _ := pterm.DefaultBigText.WithLetters(
		pterm.NewLettersFromStringWithStyle("Kuknos ", pterm.NewStyle(pterm.FgBlue)),
		pterm.NewLettersFromStringWithStyle("Cli", pterm.NewStyle(pterm.FgLightMagenta))).
		Srender()

	pterm.DefaultCenter.Print(ptermLogo)

	pterm.DefaultCenter.Print(pterm.DefaultHeader.WithFullWidth().WithBackgroundStyle(pterm.NewStyle(pterm.BgLightBlue)).WithMargin(10).Sprint("Kuknos cli is a command line tool for managing Kuknos projects."))
	pterm.Println()

}
