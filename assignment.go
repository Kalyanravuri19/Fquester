package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

type vg struct {
	Name string `json:"name"`
	Qty  int    `json:"qty"`
}

func displayDetails(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	minQty, _ := strconv.Atoi(vars["quantity"])
	flag := false
	fruit, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/fruits")
	vegetable, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/vegetable")
	grain, _ := http.Get("https://f8776af4-e760-4c93-97b8-70015f0e00b3.mock.pstmn.io/grains")
	response, _ := ioutil.ReadAll(fruit.Body)
	var allItems []vg
	var fruits []vg
	json.Unmarshal(response, &fruits)
	allItems = append(allItems, fruits...)
	response2, _ := ioutil.ReadAll(vegetable.Body)
	var vegetables []vg
	json.Unmarshal(response2, &vegetables)
	allItems = append(allItems, vegetables...)
	response3, _ := ioutil.ReadAll(grain.Body)
	var grains []vg
	json.Unmarshal(response3, &grains)
	allItems = append(allItems, grains...)
	sort.Slice(allItems, func(a, b int) bool {
		return allItems[a].Name < allItems[b].Name
	})
	for _, item := range allItems {

		if item.Qty <= minQty {

			flag = true

			fmt.Println(w, item)

		}

	}
	if flag == false {
		fmt.Fprintln(w, "Not found")
	}
}
func handleRequests() {

	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/quest/{quantity}", displayDetails).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8086", myRouter))

}

func main() {

	handleRequests()

}
