package db

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Password      string
	Mail          string
	MitarbeiterId uint
	Mitarbeiter   Mitarbeiter
}

func (d Database) CreateUser(Mail, Password string) error {
	var m Mitarbeiter
	res := d.db.Where(&Mitarbeiter{Email: &Mail}).First(&m)
	if res.Error != nil {
		return res.Error
	}
	if len(*m.Email) < 3 {
		return fmt.Errorf("kein Mitarbeiter mit dieser E-Mail Adresse gefunden")
	}

	res = d.db.Create(&User{
		Password:      Password,
		Mail:          Mail,
		Mitarbeiter:   m,
		MitarbeiterId: m.ID,
	})
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (d Database) GetUser(id uint) (*User, error) {
	var u User
	err := d.db.First(&u, id).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (d Database) GetUserByMail(Mail string) (*User, error) {
	var u User
	res := d.db.Where(&User{Mail: Mail}).Joins("Mitarbeiter").First(&u)
	if res.Error != nil {
		return nil, res.Error
	}

	return &u, nil
}

func (d Database) CheckUser(Mail, Password string) (bool, error) {
	u, err := d.GetUserByMail(Mail)
	if err != nil {
		return false, err
	}
	return u.Password == Password, nil
}

func (d Database) ChangePassword(id uint, old, new string) error {
	u, err := d.GetUser(id)
	if err != nil {
		return err
	}
	if u.Password == old {
		u.Password = new
	}
	return d.db.Save(&u).Error
}

func (d Database) DeleteUser(id uint) error {
	return d.db.Delete(&User{}, id).Error
}
