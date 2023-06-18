package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

const Guild_Slug = "https://thatsmybis.com/9044/reset/"

type config struct {
	DiscordUsername string `env:"DISCORD_USERNAME"`
	DiscordPassword string `env:"DISCORD_PASSWORD"`
}

func main() {
	// load env
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	// // create go-rod launcher
	// l := launcher.New().
	// 	Headless(true).
	// 	Devtools(false)
	// defer l.Cleanup()

	// // create browser
	// url := l.MustLaunch()
	browser := rod.New().
		SlowMotion(time.Millisecond * 45).
		MustConnect()
	defer browser.MustClose()

	// login
	page := browser.MustPage("https://thatsmybis.com/auth/discord")
	page.MustElement("input[name=\"email\"]").MustInput(cfg.DiscordUsername)
	page.MustElement("input[name=\"password\"]").MustInput(cfg.DiscordPassword)
	page.MustElement("button[type=\"submit\"]").Click(proto.InputMouseButtonLeft, 1)
	time.Sleep(time.Second * 5)

	// create temp dir and set as download location
	wd, _ := os.Getwd()
	tempDir, _ := os.MkdirTemp(wd, "tmp")
	defer func() {
		os.RemoveAll(tempDir)
	}()
	// only using this to set the download location, the wait wont work
	browser.WaitDownload(tempDir)

	// download big json blob
	page.Navigate(Guild_Slug + "export")
	downloadUrl := Guild_Slug + "export/characters-with-items/json"
	selector := fmt.Sprintf("a[href=\"%s\"]", downloadUrl)
	page.MustElement(selector).MustClick()
	time.Sleep(time.Second * 5)

	// move to the right place
	filepath.WalkDir(tempDir, func(s string, d fs.DirEntry, e error) error {
		if !d.IsDir() {
			os.Rename(s, "../app/data/character-json.json")
		}
		return nil
	})
}
