/*
 * SWAPI
 *
 * This is a RESTful API SWAPI written in GO. You can find out more about SWAPI at [https://www.swapi.co](https://www.swapi.co). And you can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/). 
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
	"fmt"
	"net/http"

	db "github.com/blesswxl/MySWAPI/StarWars/database"
)

/*
 * Add a people
 * return:
 *   200: Successful operation
 *   401: Unauthorized
 *   404: Invalid input
 */
func AddPeople(w http.ResponseWriter, r *http.Request) {
	myDb, err := db.StartDB("mydb.db")
	if err != nil {
		fmt.Printf("Fail in open database: %v\n", err)
		return
	}

	// Verify token
	token := r.Header.Get("AuthToken")
	if (!myDb.IsLogIn([]byte(token))) {
		fmt.Printf("Unauthorized: %v\n", err)
		// 401: Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Add a people
	vars := mux.Vars(r)
	peopleId, err := strconv.Atoi(vars["peopleId"])

	body, err := ioutil.ReadAll(r.Body)
    if err != nil {
		fmt.Printf("Read body error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	if err := myDb.AddObj("people", []byte(strconv.Itoa(peopleId)),[]byte(body)); err != nil {
		fmt.Printf("Read body error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}
	
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
 * Delete a people
 * return:
 *   200: Successful operation
 *   401: Unauthorized
 *   404: Invalid ID supplied
 */
func DeletePeople(w http.ResponseWriter, r *http.Request) {
	myDb, err := db.StartDB("mydb.db")
	if err != nil {
		fmt.Printf("Fail in open database: %v\n", err)
		return
	}

	// Verify token
	token := r.Header.Get("AuthToken")
	if (!myDb.IsLogIn([]byte(token))) {
		fmt.Printf("Unauthorized: %v\n", err)
		// 401: Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}	

	// Delete a people
	vars := mux.Vars(r)
	peopleId, err := strconv.Atoi(vars["peopleId"])

	if err := myDb.DeleteObj("people", []byte(strconv.Itoa(peopleId))); err != nil {
		fmt.Printf("Read body error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

/*
 * Get people by ID
 * return:
 *   200: Successful operation
 *   401: Unauthorized
 *   404: Invalid ID supplied
 */
func GetPeopleById(w http.ResponseWriter, r *http.Request) {
	myDb, err := db.StartDB("mydb.db")
	if err != nil {
		fmt.Printf("Fail in open database: %v\n", err)
		return
	}

	// Verify token
	token := r.Header.Get("AuthToken")
	if (!myDb.IsLogIn([]byte(token))) {
		fmt.Printf("Unauthorized: %v\n", err)
		// 401: Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get people by ID
	vars := mux.Vars(r)
	peopleId, err := strconv.Atoi(vars["peopleId"])

	data := myDb.SearchByID("people", []byte(strconv.Itoa(peopleId)))
	if data != nil {
		fmt.Printf("Read body error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	// Write information to response
	w.Write(data)
}

/*
 * Get people by page
 * return:
 *   200: Successful operation
 *   401: Unauthorized
 *   404: Invalid ID supplied
 */
func GetPeopleByPage(w http.ResponseWriter, r *http.Request) {
	myDb, err := db.StartDB("mydb.db")
	if err != nil {
		fmt.Printf("Fail in open database: %v\n", err)
		return
	}

	// Verify token
	token := r.Header.Get("AuthToken")
	if (!myDb.IsLogIn([]byte(token))) {
		fmt.Printf("Unauthorized: %v\n", err)
		// 401: Unauthorized
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	
	// Get people by page
	r.ParseForm()
	page, err := strconv.Atoi(r.Form["page"][0])

	data := myDb.SearchByPage("people", page)
	if data != nil {
		fmt.Printf("Read body error: %v\n", err)
		w.WriteHeader(http.StatusNotFound)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
