package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Tables []interface{}
	dsn    string
	db     *gorm.DB
}

var err error

var DatabaseConn Database

func (DB Database) Start() {
	DatabaseConn = DB
	DatabaseConn.addTables()
	DatabaseConn.dsn = "root:@tcp(127.0.0.1:3306)/sojectmanagement?charset=utf8mb4&parseTime=True&loc=Local"
	DatabaseConn.connect()
	DatabaseConn.dropTables()
	DatabaseConn.createTables()
	addBaseData()
}

func (DB *Database) connect() {

	DB.db, err = gorm.Open(mysql.Open(DB.dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		println("Database connected! ")
	}

}

func (DB *Database) addTables() {
	DB.Tables = append(DB.Tables,
		&User{},
		&Role{},
		&Access{},
		&Company{},
		&Tag{},
	)
}

func (DB *Database) dropTables() {
	DB.db.Migrator().DropTable(DB.Tables...)
}
func (DB *Database) createTables() {
	DB.db.AutoMigrate(DB.Tables...)
}
