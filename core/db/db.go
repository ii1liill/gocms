package db

// postgres
import (
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"

	"github.com/ii1liill/gocms/core/config"
)

var db *gorm.DB
var err error
var once sync.Once

func New() *gorm.DB{
	once.Do(func() {
		db, err = gorm.Open(config.GetString("db.driver"), config.GetString("db.dsn"))
		if err != nil {
			panic(err)
		}
	})
	return db
}