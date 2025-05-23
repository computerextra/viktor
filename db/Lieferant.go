package db

import (
	"gorm.io/gorm"
)

type Lieferant struct {
	gorm.Model
	Firma           string
	Kundennummer    *string
	Webseite        *string
	Ansprechpartner []Ansprechpartner `gorm:"foreignKey:LieferantenId;constraint:OnDelete:CASCADE"`
}

func (d Database) CreateLieferant(Firma string, Kundennummer, Webseite *string) {
	d.db.Omit("Ansprechpartner.*").Create(&Lieferant{
		Firma:        Firma,
		Kundennummer: Kundennummer,
		Webseite:     Webseite,
	})
}

func (d Database) GetLieferant(id uint) Lieferant {
	var l Lieferant
	d.db.Model(&Lieferant{}).Preload("Ansprechpartner").First(&l, id)
	return l
}

func (d Database) GetLieferanten() []Lieferant {
	var l []Lieferant
	d.db.Preload("Ansprechpartner").Find(&l).Order("Firma asc")
	return l
}

func (d Database) UpdateLieferant(id uint, Firma string, Kundennummer, Webseite *string) {
	var l Lieferant
	d.db.First(&l, id)
	l.Firma = Firma
	l.Kundennummer = Kundennummer
	l.Webseite = Webseite
	d.db.Save(&l)
}

func (d Database) DeleteLieferant(id uint) {
	d.db.Delete(&Lieferant{}, id)
}
