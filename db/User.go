package db

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Password      string
	Mail          string
	MitarbeiterId uint
	Mitarbeiter   Mitarbeiter
}

func (d Database) CreateUser(Mail, Password string) string {
	var m Mitarbeiter
	res := d.db.Where(&Mitarbeiter{Email: &Mail}).First(&m)
	if res.Error != nil {
		return res.Error.Error()
	}
	if len(*m.Email) < 3 {
		return "Kein Mitarbeiter mit dieser E-Mail Adresse gefunden"
	}

	res = d.db.Create(&User{
		Password:      Password,
		Mail:          Mail,
		Mitarbeiter:   m,
		MitarbeiterId: m.ID,
	})
	if res.Error != nil {
		return res.Error.Error()
	}
	return "OK"
}

func (d Database) GetUser(id uint) User {
	var u User
	d.db.First(&u, id)
	return u
}

func (d Database) GetUserByMail(Mail string) *User {
	var u User
	res := d.db.Where(&User{Mail: Mail}).Joins("Mitarbeiter").First(&u)
	if res.Error != nil {
		return nil
	}

	return &u
}

func (d Database) CheckUser(Mail, Password string) bool {
	u := d.GetUserByMail(Mail)
	return u.Password == Password
}

func (d Database) ChangePassword(id uint, old, new string) {
	u := d.GetUser(id)
	if u.Password == old {
		u.Password = new
	}
	d.db.Save(&u)
}

func (d Database) DeleteUser(id uint) {
	d.db.Delete(&User{}, id)
}
