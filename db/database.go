package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type Database struct {
	DB *sql.DB
}

func NewDatabase(connString string) *Database {
	conn, err := initDatabase(connString)
	if err != nil {
		panic(err)
	}
	return &Database{
		DB: conn,
	}
}

func initDatabase(connString string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", connString)
	if err != nil {
		return nil, err
	}

	// Create Tables
	Ansprechpartner := `CREATE TABLE IF NOT EXISTS Ansprechpartner (Id TEXT not null primary key, Name TEXT not null, Telefon TEXT, Mobil TEXT,  Mail TEXT, LieferantenId TEXT not null);`

	Kanban := `CREATE TABLE IF NOT EXISTS Kanban (Id TEXT not null primary key, Name TEXT not null, UserId TEXT not null, CreatedAt DATETIME not null default (DATETIME('now')), UpdatedAt DATETIME not null default (DATETIME('now')));`

	Post := `CREATE TABLE IF NOT EXISTS Post (Id TEXT not null primary key, KanbanId TEXT not null, Name TEXT not null, Description TEXT,  Status TEXT not null, Importance TEXT not null, CreatedAt DATETIME not null default (DATETIME('now')), UpdatedAt DATETIME not null default (DATETIME('now')));`

	Lieferant := `CREATE TABLE IF NOT EXISTS Lieferant ( Id text not null primary key, Firma text not null unique, Kundennummer text, Webseite text);`

	Mitarbeiter := `CREATE TABLE IF NOT EXISTS Mitarbeiter (Id text not null primary key, Name text not null unique, Short text default null, Gruppenwahl text default null, InternTelefon1 text default null, InternTelefon2 text default null, FestnetzPrivat text default null, FestnetzBusiness text default null, HomeOffice text default null, MobilBusiness text default null, MobilPrivat text default null, Email text default null, Azubi BOOLEAN default false, Geburtstag DATETIME default null, Paypal BOOLEAN default false, Abonniert BOOLEAN default false, Geld text default null, Pfand text default null, Dinge text default null, Abgeschickt DATETIME default null, Bild1 BLOB default null, Bild2 BLOB default null, Bild3 BLOB default null, Bild1Date DATETIME default null, Bild2Date DATETIME default null, Bild3Date DATETIME default null);`

	User := `CREATE TABLE IF NOT EXISTS User (Id string not null primary key, Passwort string not null, Mail string not null unique, MitarbeiterId string not null unique);`

	_, err = conn.Exec(Ansprechpartner)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(Kanban)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(Post)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(Lieferant)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(Mitarbeiter)
	if err != nil {
		return nil, err
	}
	_, err = conn.Exec(User)
	if err != nil {
		return nil, err
	}

	return conn, err
}
