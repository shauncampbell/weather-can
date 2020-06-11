package weather

import (
	"fmt"
	"testing"
)

func TestWeatherSites(t *testing.T) {
	w, err := New()
	if err != nil {
		// do something.
		return
	}

	f, err := w.GetSite("Burlington", Ontario, English)
	if err != nil {
		// do something.
		return
	}

	c, err := f.GetSiteData(English)
	if err != nil {
		// do something.
		return
	}
	fmt.Printf("in %s, %s it is currently %.2f %s\n", f.EnglishName, f.Province, c.CurrentConditions.Temperature.Value, c.CurrentConditions.Temperature.Units)

}
