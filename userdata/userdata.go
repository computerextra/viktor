package userdata

import (
	"encoding/json"
	"fmt"
	"os"
)

const userdata = "userdata.tmp"

type UserData struct {
	Name   *string
	Mail   *string
	Id     *uint
	UserId *uint
}

func NewUserdata() *UserData {
	return &UserData{}
}

func (d UserData) GetId() uint {
	return *d.Id
}

func (d UserData) Login(name, mail string, id, userId uint) (*UserData, error) {
	if len(name) < 3 {
		return nil, fmt.Errorf("no name given")
	}
	if len(mail) < 3 {
		return nil, fmt.Errorf("no mail given")
	}

	d.Id = &id
	d.Mail = &mail
	d.Name = &name
	d.UserId = &userId
	err := d.writeData()
	if err != nil {
		return nil, err
	}
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
	err = os.WriteFile(userdata, jsonData, 0o644)
	if err != nil {
		return err
	}

	return nil
}

func ReadFile() (*UserData, error) {
	data, err := os.ReadFile(userdata)
	if err != nil {
		return nil, err
	}
	var jsonData UserData
	err = json.Unmarshal(data, &jsonData)
	if err != nil {
		return nil, err
	}
	return &jsonData, nil
}

func (d UserData) deleteData() error {
	return os.Remove(userdata)
}
