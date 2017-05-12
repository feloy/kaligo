package kaligo

import (
	//	"fmt"
	"testing"
)

func TestDecodeConteneurUnknown(t *testing.T) {
	f := "specs/notfound.xml"
	_, err := DecodeConteneur(f)
	if err == nil {
		t.Errorf("Error unhandled not found file")
	}
}
/*
func TestDecodeConteneurInvalid(t *testing.T) {
	f := "specs/texte_version_LEGITEXT000005627819.xml"
	_, err := DecodeConteneur(f)
	if err == nil {
		t.Errorf("Error unhandled invalid file")
	}
}*/

func TestDecodeConteneur(t *testing.T) {
	f := "specs/conteneur_KALICONT000017577652.xml"
	v, err := DecodeConteneur(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "KALICONT000017577652", v.Id)
	assertStringEquals(t, "UN_ANCIEN_ID", v.AncienId)
	assertStringEquals(t, "KALI", v.Origine)
	assertStringEquals(t, "conteneur/KALI/CONT/00/00/17/57/76/KALICONT000017577652.xml", v.Url)
	assertStringEquals(t, "IDCC", v.Nature)

	assertStringEquals(t, "Convention collective nationale du sport du 7 juillet 2005 étendue par arrêté du 21 novembre 2006", v.Titre)
	assertStringEquals(t, "VIGUEUR_ETEN", v.Etat)
	assertStringEquals(t, "2511", v.Num)
	assertStringEquals(t, "2005-07-07", v.DatePubli)

	assertLenEquals(t, 0, len(v.LiensTxt))
	assertLenEquals(t, 4, len(v.Tms))

	assertStringEquals(t, "1", v.Tms[0].Niv)

	assertStringEquals(t, "Texte de base", v.Tms[0].TitreTm)
	assertLenEquals(t, 1, len(v.Tms[0].LiensTxt))
	assertStringEquals(t, "KALITEXT000017577657", v.Tms[0].LiensTxt[0].Idtxt)
	assertStringEquals(t, "Convention collective nationale du 7 juillet 2005", v.Tms[0].LiensTxt[0].Titretxt)
	assertLenEquals(t, 109, len(v.Tms[1].LiensTxt))

	assertLenEquals(t, 1, len(v.ActsPro))
	assertStringEquals(t, "Sport", v.ActsPro[0])

	assertLenEquals(t, 1, len(v.NumsBroch))
	assertStringEquals(t, "3328", v.NumsBroch[0])
}

