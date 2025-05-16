package db

// TODO: Datenbank abfragen gehen alle nicht!!!!!!!

import (
	"fmt"
)

type Database struct {
	ConnectionString string
}

const (
	Ansprechpartner = "Ansprechpartner"
	Lieferant       = "Lieferant"
	Mitarbeiter     = "Mitarbeiter"
	User            = "User"
	Version         = "Version"
)

type Model string

var Models = map[Model]string{
	Ansprechpartner: "Ansprechpartner",
	Lieferant:       "Lieferante",
	Mitarbeiter:     "Mitarbeiter",
	User:            "User",
	Version:         "Version",
}

func NewDatabase(uploadFolder string) *Database {
	return &Database{
		ConnectionString: fmt.Sprintf("file:%s\\viktor.db?cache=shared&mode=rwc&_fk=1", uploadFolder),
	}
}

func (d Database) Create(model Model, params any) error {
	switch model {
	case Ansprechpartner:
		p, ok := params.(AnsprechpartnerParams)
		if ok {
			return d.createAnsprechpartner(p)
		} else {
			return fmt.Errorf("wrong params")
		}

	case Lieferant:
		p, ok := params.(LieferantParams)
		if ok {
			return d.createLieferant(p)
		} else {
			return fmt.Errorf("wrong params")
		}

	case Mitarbeiter:
		p, ok := params.(MitarbeiterParams)
		if ok {
			return d.createMitarbeiter(p)
		} else {
			return fmt.Errorf("wrong params")
		}

	case User:
		p, ok := params.(UserParams)
		if ok {
			return d.createUser(p)
		} else {
			return fmt.Errorf("wrong params")
		}

	case Version:
		p, ok := params.(VersionParams)
		if ok {
			return d.createVersion(p)
		} else {
			return fmt.Errorf("wrong params")
		}
	}

	return fmt.Errorf("wrong model")
}

func (d Database) Update(model Model, params any, idString *string, idInt *int) error {
	switch model {
	case Ansprechpartner:
		p, ok := params.(AnsprechpartnerParams)
		if ok {
			if len(*idString) > 0 {
				return d.updateAnsprechpartner(*idString, p)
			} else {
				return fmt.Errorf("wrong id given")
			}
		} else {
			return fmt.Errorf("wrong params")
		}

	case Lieferant:
		p, ok := params.(LieferantParams)
		if ok {
			if len(*idString) > 0 {
				return d.updateLieferant(*idString, p)
			} else {
				return fmt.Errorf("wrong id given")
			}
		} else {
			return fmt.Errorf("wrong params")
		}

	case Mitarbeiter:
		p, ok := params.(MitarbeiterParams)
		if ok {
			if len(*idString) > 0 {
				return d.updateMitarbeiter(*idString, p)
			} else {
				return fmt.Errorf("wrong id given")
			}
		} else {
			return fmt.Errorf("wrong params")
		}

	case User:
		p, ok := params.(UserParams)
		if ok {
			if len(*idString) > 0 {
				return d.updateUser(*idString, p)
			} else {
				return fmt.Errorf("wrong id given")
			}
		} else {
			return fmt.Errorf("wrong params")
		}

	case Version:
		p, ok := params.(VersionParams)
		if ok {
			if *idInt > 0 {
				return d.updateVersion(*idInt, p)
			} else {
				return fmt.Errorf("wrong id given")
			}
		} else {
			return fmt.Errorf("wrong params")
		}
	}

	return fmt.Errorf("wrong model")
}

func (d Database) Read(model Model, idString *string, idInt *int) (any, error) {
	switch model {
	case Ansprechpartner:
		return d.readAnsprechpartner(idString)

	case Lieferant:
		return d.readLieferant(idString)

	case Mitarbeiter:
		return d.readMitarbeiter(idString)

	case User:
		return d.readUser(idString)

	case Version:
		return d.readVersion(idInt)
	}

	return nil, fmt.Errorf("wrong model")
}

func (d Database) Delete(model Model, idString *string, idInt *int) error {
	switch model {
	case Ansprechpartner:
		if len(*idString) == 0 {
			return fmt.Errorf("wrong id given")
		} else {
			return d.deleteAnsprechpartner(*idString)
		}

	case Lieferant:
		if len(*idString) == 0 {
			return fmt.Errorf("wrong id given")
		} else {
			return d.deleteLieferant(*idString)
		}

	case Mitarbeiter:
		if len(*idString) == 0 {
			return fmt.Errorf("wrong id given")
		} else {
			return d.deleteMitarbeiter(*idString)
		}

	case User:
		if len(*idString) == 0 {
			return fmt.Errorf("wrong id given")
		} else {
			return d.deleteUser(*idString)
		}

	case Version:
		if *idInt == 0 {
			return fmt.Errorf("wrong id given")
		} else {
			return d.deleteVersion(*idInt)
		}
	}

	return fmt.Errorf("wrong model")
}
