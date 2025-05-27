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

func (d Database) CreateLieferant(Firma string, Kundennummer, Webseite *string) error {
	return d.db.Omit("Ansprechpartner.*").Create(&Lieferant{
		Firma:        Firma,
		Kundennummer: Kundennummer,
		Webseite:     Webseite,
	}).Error
}

func (d Database) GetLieferant(id uint) (*Lieferant, error) {
	var l Lieferant
	err := d.db.Model(&Lieferant{}).Preload("Ansprechpartner").First(&l, id).Error
	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (d Database) GetLieferanten() ([]Lieferant, error) {
	var l []Lieferant
	err := d.db.Preload("Ansprechpartner").Find(&l).Order("Firma asc").Error
	if err != nil {
		return nil, err
	}
	return l, nil
}

func (d Database) UpdateLieferant(id uint, Firma string, Kundennummer, Webseite *string) error {
	var l Lieferant
	err := d.db.First(&l, id).Error
	if err != nil {
		return err
	}
	l.Firma = Firma
	l.Kundennummer = Kundennummer
	l.Webseite = Webseite
	return d.db.Save(&l).Error
}

func (d Database) DeleteLieferant(id uint) error {
	return d.db.Delete(&Lieferant{}, id).Error
}
