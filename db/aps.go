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

func (d Database) GetAnsprechpartner(id uint) (*Ansprechpartner, error) {
	var ap Ansprechpartner
	err := d.db.First(&ap, id).Error
	if err != nil {
		return nil, err
	}
	return &ap, nil
}

func (d Database) GetAllAnsprechpartner() ([]Ansprechpartner, error) {
	var aps []Ansprechpartner
	err := d.db.Order("Name asc").Find(&aps).Error
	if err != nil {
		return nil, err
	}
	return aps, nil
}

func (d Database) UpdateAnsprechpartner(id uint, Name string, Telefon, Mobil, Mail *string) error {
	var ap Ansprechpartner
	err := d.db.First(&ap, id).Error
	if err != nil {
		return err
	}
	ap.Name = Name
	ap.Telefon = Telefon
	ap.Mobil = Mobil
	ap.Mail = Mail
	return d.db.Save(&ap).Error
}

func (d Database) DeleteAnsprechpartner(id uint) error {
	return d.db.Delete(&Ansprechpartner{}, id).Error
}
