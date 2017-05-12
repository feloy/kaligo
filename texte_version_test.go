package kaligo

import (
	//	"fmt"
	"testing"
)

func TestDecodeTexteVersionUnknown(t *testing.T) {
	f := "specs/notfound.xml"
	_, err := DecodeTexteVersion(f)
	if err == nil {
		t.Errorf("Error unhandled not found file")
	}
}

func TestDecodeTexteVersionInvalid(t *testing.T) {
	f := "specs/conteneur_KALICONT000017577652.xml"
	_, err := DecodeTexteVersion(f)
	if err == nil {
		t.Errorf("Error unhandled invalid file")
	}
}

func TestDecodeTexteVersion1(t *testing.T) {
	f := "specs/texte_version_KALITEXT000006934491.xml"
	v, err := DecodeTexteVersion(f)
	if err != nil {
		t.Errorf("Error decoding file")
		return
	}

	assertStringEquals(t, "K4U92016", v.AncienId)
	assertStringEquals(t, "KALITEXT000006934491", v.Id)
	assertStringEquals(t, "ACCORD", v.Nature)
	assertStringEquals(t, "KALI", v.Origine)
	assertStringEquals(t, "texte/version/KALI/TEXT/00/00/06/93/44/KALITEXT000006934491.xml", v.Url)
	assertStringEquals(t, "KALITEXT000006934491", v.Cid)
	assertStringEquals(t, "", v.Num)
	assertStringEquals(t, "", v.NumSequence)
	assertStringEquals(t, "ASET0050882M", v.Nor)

	assertStringEquals(t, "2999-01-01", v.DatePubli)
	assertStringEquals(t, "2000-10-24", v.DateTexte)
	assertStringEquals(t, "2006-12-11", v.DerniereModification)
	assertStringEquals(t, "", v.OriginePubli)
	assertStringEquals(t, "", v.PageDebPubli)
	assertStringEquals(t, "", v.PageFinPubli)

	assertStringEquals(t, "Salaires (Nord - Pas-De-Calais)", v.Titre)
	assertStringEquals(t, "Nord - Pas-De-Calais Accord du 24 octobre 2000 relatif aux salaires", v.TitreFull)
	assertStringEquals(t, "REGIONAL NORD PAS-DE-CALAIS", v.AppliGeo)
	assertStringEquals(t, "VIGUEUR_ETEN", v.Etat)
	assertStringEquals(t, "2000-10-24", v.DateDebut)
	assertStringEquals(t, "2999-01-01", v.DateFin)

	assertLenEquals(t, 2, len(v.McsTxt))
	assertStringEquals(t, "cabinets d'architecte", v.McsTxt[0])
	assertStringEquals(t, "Maîtrise d'oeuvre", v.McsTxt[1])

	assertStringEquals(t, "Un texte de visa", v.Visas.Contenu)
	assertStringEquals(t, "Un texte de nota", v.Nota.Contenu)
	assertStringEquals(t, "Le présent accord est applicable à compter de la publication de son arrêté d'extension.", v.ConditionDiffere)

	assertStringEquals(t, "", v.Signataires.Execution.Contenu)
	assertStringEquals(t, "", v.Signataires.OrgAdhere.Contenu)
	assertStringEquals(t, "", v.Signataires.OrgDenonce.Contenu)
	assertStringEquals(t, "", v.Signataires.SignExt.Contenu)
	assertStringEquals(t, "<br/>Le syndicat des architectes de la région Nord - Pas-de-Calais (membre de l'UNSFA),<br/>", v.Signataires.SignPatron.Contenu)
	assertStringEquals(t, "<br/>Le syndicat CFDT,<br/>", v.Signataires.SignSyndic.Contenu)
}
