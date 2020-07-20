package utils

import(	
	"os"
	"log"
	"sync"
	"database/sql"
	_ "github.com/lib/pq"
)

type dbUtils struct{
	db *sql.DB
}

var dbInstance *dbUtils
var dbOnce sync.Once

func GetDBConnection() *sql.DB {
	dbOnce.Do(func() {
		log.Println("Initialize db connection...")
		// create connection to postgresql
		log.Println("host=" + os.Getenv("DATABASE_HOST") + " port=" + os.Getenv("DATABASE_PORT"))
		connection := "host=" + os.Getenv("DATABASE_HOST") + " port=" + os.Getenv("DATABASE_PORT") + " user=" + os.Getenv("USERNAME_DB") + " dbname=" + os.Getenv("DATABASE_NAME") +
			" password=" + os.Getenv("PASSWORD_DB") + " sslmode=" + os.Getenv("DATABASE_SSL")
		db, err := sql.Open(os.Getenv("DATABASE_TYPE"), connection)

		if err != nil {
			log.Println(err)
			return
		}
		if err != nil {
			log.Println(err)
		}
		err = db.Ping()
		if err != nil {		 	
		 	panic(err) 
		}

		dbInstance = &dbUtils{
			db: db,
		}
	})

	return dbInstance.db
}
