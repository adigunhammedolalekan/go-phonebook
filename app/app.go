package app

import (
	"github.com/gorilla/mux"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
	"log"
)


type App struct {

	Db *gorm.DB;
	Router *mux.Router;

}

func (a *App) Init (dbConnectionString string)  {

	var db, err = gorm.Open("mysql", dbConnectionString)
	if err != nil {
		panic(err)
	}

	a.Db = db;
	a.Router = mux.NewRouter()
}

func (a *App) Run()  {

	err := http.ListenAndServe("127.0.0.1:9000", a.Router)
	if err != nil {
		log.Fatal(err)
	}
}
