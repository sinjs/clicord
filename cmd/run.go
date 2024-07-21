package cmd

import (
	"log"

	"github.com/rivo/tview"
	"github.com/sinjs/clicord/internal/config"
	"github.com/sinjs/clicord/internal/logger"
	"github.com/sinjs/clicord/internal/ui"
)

var (
	discordState *State

	cfg      *config.Config
	app      = tview.NewApplication()
	mainFlex *MainFlex
)

func init() {
	var err error
	cfg, err = config.Load()
	if err != nil {
		panic(err)
	}
}

func Run(token string) error {
	if err := logger.Load(); err != nil {
		return err
	}

	if token == "" {
		lf := ui.NewLoginForm(cfg)

		go func() {
			// mainFlex must be initialized before opening a new state.
			mainFlex = newMainFlex()

			token := <-lf.Token
			if token.Error != nil {
				app.Stop()
				log.Fatal(token.Error)
			}

			if err := openState(token.Value); err != nil {
				app.Stop()
				log.Fatal(err)
			}

			app.QueueUpdateDraw(func() {
				app.SetRoot(mainFlex, true)
			})
		}()

		app.SetRoot(lf, true)
	} else {
		mainFlex = newMainFlex()
		if err := openState(token); err != nil {
			return err
		}

		app.SetRoot(mainFlex, true)
	}

	app.EnableMouse(cfg.Mouse)
	return app.Run()
}
