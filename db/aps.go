package db

import "gorm.io/gorm"

type Ansprechpartner struct {
	gorm.Model
	Name          string
	Telefon       *string
	Mobil         *string
	Mail          *string
	LieferantenId uint
}

func (d Database) CreateAnsprechpartner(Name string, Telefon, Mobil, Mail *string, LieferantenId uint) {
	d.db.Create(&Ansprechpartner{
		Name:          Name,
		Telefon:       Telefon,
		Mobil:         Mobil,
		Mail:          Mail,
		LieferantenId: LieferantenId,
	})
}

func (d Database) GetAnsprechpartner(id uint) Ansprechpartner {
	var ap Ansprechpartner
	d.db.First(&ap, id)
	return ap
}

func (d Database) GetAllAnsprechpartner() []Ansprechpartner {
	var aps []Ansprechpartner
	d.db.Order("Name asc").Find(&aps)
	return aps
}

func (d Database) UpdateAnsprechpartner(id uint, Name string, Telefon, Mobil, Mail *string) {
	var ap Ansprechpartner
	d.db.First(&ap, id)
	ap.Name = Name
	ap.Telefon = Telefon
	ap.Mobil = Mobil
	ap.Mail = Mail
	d.db.Save(&ap)
}

func (d Database) DeleteAnsprechpartner(id uint) {
	d.db.Delete(&Ansprechpartner{}, id)
}
