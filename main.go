package main

import (
	"fmt"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var version = "dev"

func main() {
	fyneApp := app.New()
	window := fyneApp.NewWindow("Arduino Manager " + version)
	window.Resize(fyne.NewSize(620, 470))

	status := widget.NewMultiLineEntry()
	status.Wrapping = fyne.TextWrapWord
	status.SetText("Select an action to query Arduino CLI. Arduino CLI must be installed separately.")

	run := func(args ...string) {
		output, err := RunArduinoCLI(args...)
		commandLine := "$ arduino-cli " + strings.Join(args, " ")
		if err != nil {
			status.SetText(commandLine + "\n\n" + err.Error() + "\n\n" + output)
			return
		}

		if len(args) >= 2 && args[0] == "board" && args[1] == "list" {
			boards := ParseBoardList(output)
			if len(boards) == 0 {
				status.SetText(commandLine + "\n\nNo boards detected. Check the USB cable and permissions.\n\n" + output)
				return
			}

			var report strings.Builder
			fmt.Fprintf(&report, "%s\n\nDetected boards: %d\n", commandLine, len(boards))
			for index, board := range boards {
				fmt.Fprintf(&report, "\n%d. Port: %s\n   FQBN: %s\n", index+1, board.Port, board.FQBN)
			}
			status.SetText(report.String())
			return
		}

		status.SetText(commandLine + "\n\n" + output)
	}

	window.SetContent(container.NewVBox(
		widget.NewLabelWithStyle("Arduino Manager", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		widget.NewLabel("A lightweight GUI for common Arduino CLI inspection commands."),
		widget.NewLabel("Version: "+version),
		container.NewGridWithColumns(2,
			widget.NewButton("Detect connected boards", func() { run("board", "list") }),
			widget.NewButton("List installed cores", func() { run("core", "list") }),
			widget.NewButton("List installed libraries", func() { run("lib", "list") }),
			widget.NewButton("Show Arduino CLI version", func() { run("version") }),
		),
		widget.NewSeparator(),
		status,
		widget.NewButton("Exit", func() { fyneApp.Quit() }),
	))

	window.ShowAndRun()
}
