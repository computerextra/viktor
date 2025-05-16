package userdata

import (
	"encoding/json"
	"fmt"
	"os"
)

type UserData struct {
	Name *string
	Mail *string
	id   *string
}

func NewUserdata() *UserData {
	return &UserData{}
}

func (d UserData) GetId() string {
	return *d.id
}

func (d UserData) Login(name, mail, id string) (*UserData, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("no name given")
	}
	if len(mail) < 3 {
		return nil, fmt.Errorf("no mail given")
	}
	if len(id) < 3 {
		return nil, fmt.Errorf("no id given")
	}

	d.id = &id
	d.Mail = &mail
	d.Name = &name
	d.writeData()

	return &d, nil
}

func (d UserData) Logout() error {
	return d.deleteData()
}

func (d UserData) writeData() error {
	jsonData, err := json.Marshal(d)
	if err != nil {
		return err
	}
	err = os.WriteFile("userdata", jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (d UserData) deleteData() error {
	return os.Remove("userdata")
}
