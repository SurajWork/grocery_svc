package grocery_svc

import (
	"github.com/gorilla/mux"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var db *gorm.DB

func main() {

	// gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	db, err := gorm.Open(sqlite.Open("./test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("unable to connect to DB", err.Error())
	}

	if err = db.AutoMigrate(&Grocery{}); err != nil {
		log.Fatal("migration failed", err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/allgroceries", AllGroceries)                        // ----> To request all groceries
	r.HandleFunc("/groceries/{name}", SingleGrocery)                   // ----> To request a specific grocery
	r.HandleFunc("/groceries", GroceriesToBuy).Methods("POST")         // ----> To add  new grocery to buy
	r.HandleFunc("/groceries/{name}", UpdateGrocery).Methods("PUT")    // ----> To update a grocery
	r.HandleFunc("/groceries/{name}", DeleteGrocery).Methods("DELETE") // ----> Delete a grocery

	log.Fatal(http.ListenAndServe(":8000", r))

}
