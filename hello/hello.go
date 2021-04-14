package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"

	"github.com/marcusolsson/tui-go"
)

type itens struct {
	website, url string
}

var personalItens = []itens{
	{website: "Linkedin", url: "https://www.linkedin.com/in/ramonlimaramos"},
	{website: "Github", url: "https://github.com/ramonlimaramos"},
	{website: "Twitter", url: "https://twitter.com/ramonlimaramos"},
	{website: "Instagram", url: "https://instagram.com/ramonlimaramos"},
}

var logoAscii = `
.----.---.-.--------.-----.-----.
|   _|  _  |        |  _  |     |
|__| |___._|__|__|__|_____|__|__|
.----.---.-.--------.-----.-----.
|   _|  _  |        |  _  |__ --|
|__| |___._|__|__|__|_____|_____|
`

func main() {
	theme := tui.NewTheme()

	// LOGO
	logo := tui.NewLabel(logoAscii)
	logo.SetStyleName("logo")
	theme.SetStyle("label.logo", tui.Style{Bg: tui.ColorDefault, Fg: tui.ColorWhite})

	// ITENS LIST
	itens := tui.NewList()
	itens.SetFocused(true)
	for _, p := range personalItens {
		itens.AddItems(fmt.Sprintf("%s: %s", p.website, p.url))
	}
	itens.SetSelected(0)
	theme.SetStyle("list.item.selected", tui.Style{Bg: tui.ColorGreen, Fg: tui.ColorWhite})
	itens.OnItemActivated(func(t *tui.List) {
		openBrowser(personalItens[t.Selected()])
	})

	// WINDOW COMPOSITION
	window := tui.NewVBox(
		tui.NewPadder(10, 1, logo),
		tui.NewPadder(18, 0, tui.NewLabel("Software Engineer")),
		tui.NewPadder(1, 1, itens),
	)
	window.SetBorder(true)

	wrapper := tui.NewVBox(
		tui.NewSpacer(),
		window,
		tui.NewSpacer(),
	)

	// CONTENT ADJUSTMENT
	content := tui.NewHBox(tui.NewSpacer(), wrapper, tui.NewSpacer())

	status := tui.NewStatusBar("[ESC] Quit")

	root := tui.NewVBox(
		content,
		status,
	)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}

	ui.SetTheme(theme)
	ui.SetKeybinding("Esc", func() { ui.Quit() })

	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}

func openBrowser(iten itens) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", iten.url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", iten.url).Start()
	case "darwin":
		err = exec.Command("open", iten.url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}

	if err != nil {
		log.Fatal(err)
	}

}
