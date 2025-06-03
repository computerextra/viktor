package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"sync"
	"time"

	"slices"

	"github.com/dgraph-io/badger/v4"
	"github.com/dgraph-io/badger/v4/options"
)

var (
	ErrNotFound = errors.New("storage: key not found")
	ErrBadValue = errors.New("storage: bad value")

	storageSync sync.Once
	instance    *Database
)

type Database struct {
	DB *badger.DB
}

type Ansprechpartner struct {
	Id            string
	Name          string
	Telefon       *string
	Mobil         *string
	Mail          *string
	LieferantenId string
}

type Kanban struct {
	Id        string
	Name      string
	UserId    string
	Posts     []Post
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Post struct {
	Id          string
	KanbanId    string
	Name        string
	Description *string
	Status      string
	Importance  string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Lieferant struct {
	Id              string
	Firma           string
	Kundennummer    *string
	Webseite        *string
	Ansprechpartner []Ansprechpartner
}

type Mitarbeiter struct {
	Id               string
	Name             string
	Short            *string
	Gruppenwahl      *string
	InternTelefon1   *string
	InternTelefon2   *string
	FestnetzPrivat   *string
	FestnetzBusiness *string
	HomeOffice       *string
	MobilBusiness    *string
	MobilPrivat      *string
	Email            *string
	Azubi            bool
	Geburtstag       *time.Time
	Paypal           bool
	Abonniert        bool
	Geld             *string
	Pfand            *string
	Dinge            *string
	Abgeschickt      *time.Time
	Bild1            *string
	Bild2            *string
	Bild3            *string
	Bild1Date        *time.Time
	Bild2Date        *time.Time
	Bild3Date        *time.Time
}

type User struct {
	Id            string
	Password      string
	Mail          string
	MitarbeiterId string
	Mitarbeiter   Mitarbeiter
	Boards        []Kanban
}

func NewDatabase(connString string) *Database {
	return initDatabase(connString)
}

func initDatabase(connString string) *Database {
	storageSync.Do(func() {
		opts := badger.DefaultOptions(connString).WithLogger(nil).WithCompression(options.ZSTD).WithIndexCacheSize(100 << 20)
		db, err := badger.Open(opts)
		if err != nil {
			panic(fmt.Errorf("failed to connect database: %v", err))
		}
		instance = &Database{DB: db}
		go instance.runGC()
	})

	if instance == nil {
		panic("Failed to initialize storage")
	}

	return instance
}

func (d *Database) Set(key string, data any) error {
	if data == nil {
		return ErrBadValue
	}

	var buf []byte
	buf, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return d.DB.Update(func(txn *badger.Txn) error {
		e := badger.NewEntry([]byte(key), buf)
		err := txn.SetEntry(e)
		return err
	})
}

func (d *Database) Get(key string) (any, error) {
	var data any
	err := d.DB.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err == badger.ErrKeyNotFound {
			return ErrNotFound
		}
		if err != nil {
			return err
		}
		var valCopy []byte
		err = item.Value(func(val []byte) error {
			// Copying or parsing val is valid.
			valCopy = slices.Clone(val)

			return nil
		})
		if err != nil {
			return err
		}

		err = json.Unmarshal(valCopy, &data)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (d *Database) Update(key string, data any) error {
	if data == nil {
		return ErrBadValue
	}
	values, err := d.Get(key)
	if err != nil {
		return err
	}
	if reflect.TypeOf(values).Kind() != reflect.Slice {
		return errors.New("malformed data")
	}

	tmp := values.([]any)
	values = append(tmp, data)

	return d.Set(key, values)
}

func (d *Database) Delete(key string, id string) error {
	values, err := d.Get(key)
	if err != nil {
		return err
	}
	if reflect.TypeOf(values).Kind() != reflect.Slice {
		return errors.New("malformed data")
	}
	var newData []any
	switch key {
	case "Ansprechpartner":
		for _, x := range values.([]Ansprechpartner) {
			if x.Id != id {
				newData = append(newData, x)
			}
		}
	case "User":
		for _, x := range values.([]User) {
			if x.Id != id {
				newData = append(newData, x)
			}
		}
	case "Kanban":
		for _, x := range values.([]Kanban) {
			if x.Id != id {
				newData = append(newData, x)
			}
		}
	case "Post":
		for _, x := range values.([]Post) {
			if x.Id != id {
				newData = append(newData, x)
			}
		}
	case "Mitarbeiter":
		for _, x := range values.([]Mitarbeiter) {
			if x.Id != id {
				newData = append(newData, x)
			}
		}
	case "Lieferant":
		for _, x := range values.([]Lieferant) {
			if x.Id != id {
				newData = append(newData, x)
			}
		}
	}

	return d.Update("ap", newData)
}

func (d *Database) Close() error {
	if d.DB != nil {
		return d.DB.Close()
	}
	return nil
}

func (d *Database) IsHealthy() bool {
	if d.DB == nil {
		return false
	}

	err := d.DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte("health_check"))
		if err == badger.ErrKeyNotFound {
			return nil
		}
		return err
	})
	return err == nil
}

func (d *Database) runGC() {
	ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("[Storage] GC panic recovered: %v", r)
				}
			}()

			err := d.DB.RunValueLogGC(0.5)
			if err != nil && err != badger.ErrNoRewrite {
				log.Printf("[Storage] GC error: %v", err)
			}
		}()
	}
}

func (d *Database) HasKey(key string) bool {
	err := d.DB.View(func(txn *badger.Txn) error {
		_, err := txn.Get([]byte(key))
		return err
	})
	return err == nil
}
