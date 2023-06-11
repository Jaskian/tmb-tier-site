package webbuilder

import (
	"bytes"
	"embed"
	"fmt"
	"text/template"

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
		"getPhaseSlotData": getPhaseSlotData,
		"getSlotImage":     getSlotImage,
		"getClassColor":    getClassColor,
	}).ParseFS(pageTemplates, "templates/*.gohtml")

	return &SiteRenderer{t}, err
}

func (s *SiteRenderer) BuildWebsite(data shared.TMBData) (map[string]string, error) {
	pages := map[string]string{}

	for _, phase := range shared.PHASES {
		renderData := struct {
			Phase int
			Data  shared.TMBData
			Slots []shared.Slot
		}{phase, data, shared.SLOTS}

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

func getPhaseSlotData(c shared.Character, phase int, slot shared.Slot) shared.SlotData {
	return c.Phases[phase][int(slot)]
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
