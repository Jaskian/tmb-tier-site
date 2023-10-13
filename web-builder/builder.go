package webbuilder

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"
	"time"

	"github.com/jaskian/tmb-tier-site/shared"
)

var (
	//go:embed "templates/*"
	pageTemplates embed.FS
)

type SiteRenderer struct {
	pageTempl *template.Template
}

func NewSiteRenderer() (*SiteRenderer, error) {

	t, err := template.New("pageTemplate.gohtml").Funcs(template.FuncMap{
		"getPhaseData":       getPhaseData,
		"getPhaseSlotData":   getPhaseSlotData,
		"getSlotImage":       getSlotImage,
		"getClassColor":      getClassColor,
		"getLootStatusClass": getLootStatusClass,
		"subtract":           subtract,
	}).ParseFS(pageTemplates, "templates/*.gohtml")

	return &SiteRenderer{t}, err
}

func (s *SiteRenderer) BuildWebsite(data shared.TMBData, currentPhase int) (map[string]string, error) {
	pages := map[string]string{}

	for _, phase := range shared.PHASES {
		renderData := struct {
			Phase        int
			Data         shared.TMBData
			Slots        []shared.Slot
			ShowWishlist bool
			UpdateDate   string
		}{phase, data, shared.SLOTS, phase == currentPhase, time.Now().Format("2006-01-02 15:04")}

		b := bytes.Buffer{}
		pageName := fmt.Sprintf("p%d.html", phase)
		err := s.pageTempl.Execute(&b, renderData)
		if err != nil {
			return pages, err
		}
		pages[pageName] = b.String()
	}

	return pages, nil
}

func getPhaseData(c shared.Character, phase int) shared.PhaseData {
	return c.Phases[phase]
}

func getPhaseSlotData(pd shared.PhaseData, slot shared.Slot) shared.SlotData {
	return pd.Slots[int(slot)]
}

func getSlotImage(slot shared.Slot) string {
	return shared.SLOT_IMAGE_URLS[slot]
}

func getClassColor(class string) string {
	color, ok := shared.ClassColors[class]
	if ok {
		return color
	} else {
		return "white"
	}
}

func getLootStatusClass(item shared.WishlistLoot) string {
	if item.Received {
		return "receivedLoot"
	}
	return ""
}

func subtract(x int, y int) int {
	return x - y
}
