package weather

import (
	"encoding/xml"
	"fmt"
	"golang.org/x/net/html/charset"
	"net/http"
	"strings"
)

type SiteList struct {
	Sites []Site `xml:"site"`
}

type Site struct {
	Code        string   `xml:"code,attr"`
	EnglishName string   `xml:"nameEn"`
	FrenchName  string   `xml:"nameFr"`
	Province    Province `xml:"provinceCode"`
}

const (
	NewfoundLandAndLabrador Province = "NL"
	PrinceEdwardIsland      Province = "PEI"
	NovaScotia              Province = "NS"
	NewBrunswick            Province = "NB"
	Quebec                  Province = "QC"
	Ontario                 Province = "ON"
	Manitoba                Province = "MB"
	Saskatchewan            Province = "SK"
	Alberta                 Province = "AB"
	BritishColumbia         Province = "BC"
	Yukon                   Province = "YT"
	NorthwestTerritories    Province = "NT"
	Nunavut                 Province = "NU"
)
const weatherUrlFormat = "https://dd.weather.gc.ca/citypage_weather/xml/{province}/{site}_{language}.xml"

func (s *Site) GetSiteData(language Language) (*SiteData, error) {
	url := strings.NewReplacer("{province}", string(s.Province), "{site}", s.Code, "{language}", string(language)).Replace(weatherUrlFormat)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to load city: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to load city: server returned non-ok status code")
	}
	var data SiteData
	decoder := xml.NewDecoder(resp.Body)
	decoder.CharsetReader = charset.NewReaderLabel
	err = decoder.Decode(&data)

	if err != nil {
		return nil, fmt.Errorf("failed to load city: %w", err)
	}

	return &data, nil
}
