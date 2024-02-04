package china_area

import (
	"encoding/json"
	"github.com/ice-cream-heaven/log"
	"os"
	"strings"
)

type Area struct {
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`

	Children map[string]*Area `json:"children,omitempty"`
}

func (p *Area) SearchSubArea(subArea string) (*Area, bool) {
	for _, area := range p.Children {
		if area.Name == subArea {
			return area, true
		}

		if strings.HasPrefix(area.Name, subArea) {
			return area, true
		}
	}

	return nil, false
}

type AreaManage struct {
	areas map[string]*Area
}

func NewAreaManage() *AreaManage {
	return &AreaManage{
		areas: make(map[string]*Area),
	}
}

func (p *AreaManage) Load() error {
	buf, err := os.ReadFile("./resources/china_area.json")
	if err != nil {
		log.Errorf("err:%v", err)

		log.Warnf("try to download china_area.json from remote")

		err = download()
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}

		buf, err = os.ReadFile("./resources/china_area.json")
		if err != nil {
			log.Errorf("err:%v", err)
			return err
		}
	}

	err = json.Unmarshal(buf, &p.areas)
	if err != nil {
		log.Errorf("err:%v", err)
		return err
	}

	return nil
}

func (p *AreaManage) SearchProvince(province string) (*Area, bool) {
	for _, area := range p.areas {
		if area.Name == province {
			return area, true
		}

		if strings.HasPrefix(area.Name, province) {
			return area, true
		}
	}

	return nil, false
}
