package weather

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
)

type Weather struct {
	Sites *SiteList
}

type Province string
type Language string

const (
	English                 Language = "e"
	French                  Language = "f"
)

const cityListUrl = "https://dd.weather.gc.ca/citypage_weather/xml/siteList.xml"

func New() (*Weather, error) {
	resp, err := http.Get(cityListUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to load list of cities: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to load list of cities: server returned non-ok status code")
	}
	var sites SiteList
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&sites)

	if err != nil {
		return nil, fmt.Errorf("failed to load list of cities: %w", err)
	}

	w := Weather{Sites: &sites}
	return &w, nil
}

func (w *Weather) GetSite(name string, province Province, language Language) (*Site, error) {
	for _, s := range w.Sites.Sites {
		if language == English && strings.ToLower(s.EnglishName) == strings.ToLower(name) && s.Province == province {
			return &s, nil
		} else if language == French && strings.ToLower(s.FrenchName) == strings.ToLower(name) && s.Province == province {
			return &s, nil
		}
	}
	return nil, fmt.Errorf("site not found")
}
