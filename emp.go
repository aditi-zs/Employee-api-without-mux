package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Emp struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

type allEmployees []Emp

var employees = allEmployees{{"1", "Aditi", 22, "UP"}}

func Employee(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		empID := r.URL.Query().Get("id") //fetching id from request
		//getting one employee's data by id
		if empID != "" {
			for _, val := range employees {
				if (val.ID) == empID {
					//json.NewEncoder(w).Encode(val)
					w.WriteHeader(http.StatusOK)
					respBody, err := json.Marshal(val)
					w.Write(respBody)
					if err != nil {
						log.Println(err)
					}
				} else {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, "ID doesn't exist")
				}
			}
			return
		}
		//get all records
		w.WriteHeader(http.StatusOK)
		respBody, err := json.Marshal(employees)
		w.Write(respBody)
		if err != nil {
			log.Println(err)
		}
		//json.NewEncoder(w).Encode(employees)

	case "POST":
		var emp Emp
		w.Header().Set("Content-Type", "application/json")
		req, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Fprintf(w, "enter data")
		}
		json.Unmarshal(req, &emp)
		employees = append(employees, emp)
		w.WriteHeader(http.StatusCreated)
		respBody, err := json.Marshal(emp)
		w.Write(respBody)
		if err != nil {
			log.Println(err)
		}
	//	json.NewEncoder(w).Encode(emp)
	default:
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

func main() {
	http.HandleFunc("/emp", Employee)
	fmt.Println(("server at port 8000"))

	log.Fatal(http.ListenAndServe(":8000", nil))
}
