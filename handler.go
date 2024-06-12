package grocery_svc

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

var groceries = []Grocery{
	{Name: "Almod Milk", Quantity: 2},
	{Name: "Apple", Quantity: 6},
}

func AllGroceries(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: returnAllGroceries")
	json.NewEncoder(w).Encode(groceries)
}

func SingleGrocery(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params["name"] == "" {
		json.NewEncoder(w).Encode("grocery name cannot be empty")
	}
	for _, grocery := range groceries {
		if grocery.Name == params["name"] {
			json.NewEncoder(w).Encode(grocery)
		}

	}
}

func GroceriesToBuy(w http.ResponseWriter, r *http.Request) {
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		json.NewEncoder(w).Encode("unmarshalling failed")
	}
	var grocery Grocery
	err = json.Unmarshal(reqBody, &grocery)
	if err != nil {
		json.NewEncoder(w).Encode("unmarshalling failed")
	}
	groceries = append(groceries, grocery)
	json.NewEncoder(w).Encode(groceries)
}

func DeleteGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)
		}
	}
}

func UpdateGrocery(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	name := vars["name"]

	for index, grocery := range groceries {
		if grocery.Name == name {
			groceries = append(groceries[:index], groceries[index+1:]...)

			var updateGrocery Grocery

			json.NewDecoder(r.Body).Decode(&updateGrocery)
			groceries = append(groceries, updateGrocery)
			fmt.Println("Endpoint hit: UpdateGroceries")
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}

}
