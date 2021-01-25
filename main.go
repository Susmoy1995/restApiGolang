package main

import (
  "encoding/json"
  "net/http"
  "github.com/gorilla/mux"
  "math/rand"
  "strconv"
  "log"
)

// Shopping Item
type Item struct {
  ID string `json:"id"`
  Name string `json:"name"`
  Quantity int `json:"quantity"`
  Price float64 `json:"price"`
}

// initalize a shopping struct
var item []Item

// get all list from shopping
func getShoppingList(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(item)
}

// get specific shopping item
func getShoppingItem(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r) // get params

  // loop through items to get id specific data
  for _, i := range item {
    if i.ID == params["id"] {
      json.NewEncoder(w).Encode(i)
      return
    }
  }

  json.NewEncoder(w).Encode(nil)
}

// create a new item
func createShoppingItem(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  var newProduct Item
  _ = json.NewDecoder(r.Body).Decode(&newProduct)

  // generate id and assign it into new product
  newProduct.ID = strconv.Itoa(rand.Intn(1000))
  // fetch value from postman input
  item = append(item, newProduct)
  json.NewEncoder(w).Encode(newProduct)
}

// update an item
func updateShoppingItem(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for index, i := range item {
    if i.ID == params["id"] {
      item = append(item[:index], item[index+1:]...)
      var newProduct Item
      _ = json.NewDecoder(r.Body).Decode(&newProduct)

      newProduct.ID = params["id"]
      item = append(item, newProduct)
      json.NewEncoder(w).Encode(newProduct)
      return
    }
  }

  json.NewEncoder(w).Encode(item)
}

// delete an item
func deleteShoppingItem(w http.ResponseWriter, r *http.Request)  {
  w.Header().Set("Content-Type", "application/json")
  params := mux.Vars(r)

  for index, i := range item {
    if i.ID == params["id"] {
      item = append(item[:index], item[index+1:]...)
      break
    }
  }

  json.NewEncoder(w).Encode(item)
}

func main()  {
  // Initialize router
  router := mux.NewRouter()

  // demo data
  item = append(item, Item{ID: "1", Name: "Hair Brush", Quantity: 2, Price: 300.00})
  item = append(item, Item{ID: "2", Name: "Chocolate Bar", Quantity: 2, Price: 200.00})

  // Routes
  router.HandleFunc("/api/items", getShoppingList).Methods("GET")
  router.HandleFunc("/api/items/{id}", getShoppingItem).Methods("GET")
  router.HandleFunc("/api/items", createShoppingItem).Methods("POST")
  router.HandleFunc("/api/items/{id}", updateShoppingItem).Methods("PUT")
  router.HandleFunc("/api/items/{id}", deleteShoppingItem).Methods("DELETE")

  // listening port
  log.Fatal(http.ListenAndServe(":8080", router))
}
