package kaligo

import (
	"encoding/xml"
	"io/ioutil"
)

type Conteneur struct {
	XMLName xml.Name `xml:"IDCC"`
	// META
	// - META_COMMUN
	Id       string `xml:"META>META_COMMUN>ID"`
	AncienId string `xml:"META>META_COMMUN>ANCIEN_ID"`
	Origine  string `xml:"META>META_COMMUN>ORIGINE"`
	Url      string `xml:"META>META_COMMUN>URL"`
	Nature   string `xml:"META>META_COMMUN>NATURE"`

	// - META_SPEC
	//  - META_CONTENEUR
	Titre     string `xml:"META>META_SPEC>META_CONTENEUR>TITRE"`
	Etat      string `xml:"META>META_SPEC>META_CONTENEUR>ETAT"`
	Num       string `xml:"META>META_SPEC>META_CONTENEUR>NUM"`
	DatePubli string `xml:"META>META_SPEC>META_CONTENEUR>DATE_PUBLI"`

	// STRUCTURE_TXT
	LiensTxt []LienTxt `xml:"STRUCTURE_TXT>LIEN_TXT"`
	Tms      []Tm      `xml:"STRUCTURE_TXT>TM"`

	// TXTS_ISO

	// ACTS_PRO
	ActsPro []string `xml:"ACTS_PRO>ACT_PRO"`

	// NUMS_BROCH
	NumsBroch []string `xml:"NUMS_BROCH>NUM_BROCH"`
}

type LienTxt struct {
	Idtxt    string `xml:"idtxt,attr"`
	Titretxt string `xml:"titretxt,attr"`
}

type Tm struct {
	Niv      string    `xml:"niv,attr"`
	LiensTxt []LienTxt `xml:"LIEN_TXT"`
	TitreTm  string    `xml:"TITRE_TM"`
	Tms      []Tm      `xml:"TM"`
}

func DecodeConteneur(filename string) (*Conteneur, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var v Conteneur
	err = xml.Unmarshal(content, &v)
	if err != nil {
		return nil, err
	}
	return &v, nil
}
