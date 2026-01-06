package internal

type Counts struct {
	Abteilungen int
	Angebote    int
	Jobs        int
	Mitarbeiter int
	Partner     int
}

type KundenResponse struct {
	SG_Adressen_PK int
	Suchbegriff    *string
	KundNr         *string
	LiefNr         *string
	Homepage       *string
	Telefon1       *string
	Telefon2       *string
	Mobiltelefon1  *string
	Mobiltelefon2  *string
	EMail1         *string
	EMail2         *string
	KundUmsatz     *float64
	LiefUmsatz     *float64
}

type AccessArtikel struct {
	Id            int
	Artikelnummer string
	Artikeltext   string
	Preis         float64
}

type AusstellerArtikel struct {
	Id            int
	Artikelnummer string
	Artikelname   string
	Specs         string
	Preis         float64
}

type User struct {
	Kundennummer string
	Name         string
	Vorname      string
}

type Sepa struct {
	Kundennummer string `json:"kundennummer,omitempty"`
	Firma        string `json:"firma,omitempty"`
	Nachname     string `json:"nachname,omitempty"`
	Vorname      string `json:"vorname,omitempty"`
	Angelegt     string `json:"angelegt,omitempty"`
	Aktiviert    string `json:"aktiviert,omitempty"`
}
