package db

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql" // init mysql driver
)

// Handler is struct which wraps database actions
type Handler struct {
	db    *sql.DB
	dbmap *gorp.DbMap
}

// Message is struct that holds one row in database
type Message struct {
	ID      int64  `db:"id"`
	Message string `db:"message"`
}

// Client is our concurrent global database handler
var Client = &Handler{}

func init() {
	db, err := sql.Open("mysql", "root:passwd@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		log.Fatal(err)
	}

	Client.db = db
	Client.dbmap = &gorp.DbMap{
		Db: db,
		Dialect: gorp.MySQLDialect{
			Engine:   "InnoDB",
			Encoding: "UTF8",
		},
	}

	Client.dbmap.AddTableWithName(Message{}, "messages").SetKeys(true, "ID")
	log.Println("Database client initialized")
}

// InsertMessage inserts new message to db
func (db *Handler) InsertMessage(msg *Message) error {
	return db.dbmap.Insert(msg)
}
