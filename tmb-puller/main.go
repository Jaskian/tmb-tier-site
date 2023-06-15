package main

import (
	"fmt"
	"os"
	"time"

	"github.com/caarlos0/env"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
)

const Guild_Slug = "https://thatsmybis.com/9044/reset/"
const tmbAuthLink = "https://thatsmybis.com/auth/discord"

type config struct {
	DiscordUsername string `env:"DISCORD_USERNAME"`
	DiscordPassword string `env:"DISCORD_PASSWORD"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	l := launcher.New().
		Headless(false).
		Devtools(false)

	defer l.Cleanup() // remove launcher.FlagUserDataDir

	url := l.MustLaunch()

	// Trace shows verbose debug information for each action executed
	// SlowMotion is a debug related function that waits 2 seconds between
	// each action, making it easier to inspect what your code is doing.
	browser := rod.New().
		ControlURL(url).
		SlowMotion(time.Millisecond * 45).
		MustConnect()

	defer browser.MustClose()
	// ServeMonitor plays screenshots of each tab. This feature is extremely
	// useful when debugging with headless mode.
	// You can also enable it with flag "-rod=monitor"
	//launcher.Open(browser.ServeMonitor(""))

	page := browser.MustPage("https://thatsmybis.com/auth/discord")

	page.MustElement("input[name=\"email\"]").MustInput(cfg.DiscordUsername)
	page.MustElement("input[name=\"password\"]").MustInput(cfg.DiscordPassword)

	page.MustElement("button[type=\"submit\"]").Click(proto.InputMouseButtonLeft, 1)

	selector := fmt.Sprintf("a[href=\"%s\"]", tmbAuthLink)
	if has, element, err := page.HasX(selector); has {
		element.Click(proto.InputMouseButtonLeft, 1)
		time.Sleep(time.Second * 5)
	} else if err != nil {
		panic(err)
	}

	page = browser.MustPage(Guild_Slug + "export")

	wait := page.Browser().MustWaitDownload()

	downloadUrl := Guild_Slug + "export/characters-with-items/html"
	page.MustElement(fmt.Sprintf("a[href=\"%s\"]", downloadUrl)).Click(proto.InputMouseButtonLeft, 1)

	res := wait()
	err := os.WriteFile("../app/data/character-json1.json", res, 0644)

	if err != nil {
		panic(err)
	}
}
