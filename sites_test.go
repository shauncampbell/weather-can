package weather

import (
	"encoding/xml"
	"fmt"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	data := `<siteList>
               <site code="s000001">
                   <nameEn>Test City</nameEn>
                   <nameFr>Ville de test</nameFr>
                   <provinceCode>ON</provinceCode>
               </site>
            </siteList>`

	var s = SiteList{}
	reader := strings.NewReader(data)
	err := xml.NewDecoder(reader).Decode(&s)
	fmt.Println(err)
	fmt.Println(s)
}

func TestSite(t *testing.T) {
	data := `<site code="s000001">
                   <nameEn>Test City</nameEn>
                   <nameFr>Ville de test</nameFr>
                   <provinceCode>ON</provinceCode>
               </site>`

	var s Site

	reader := strings.NewReader(data)
	err := xml.NewDecoder(reader).Decode(&s)
	fmt.Println(err)
	fmt.Println(s)
}
