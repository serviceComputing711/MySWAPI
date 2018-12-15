package main

import (
	// "encoding/json"
	"fmt"
	// "strconv"
	"os"
	// "strings"
	// "github.com/boltdb/bolt"
	// "github.com/peterhellberg/swapi"
	"github.com/golang/StarWars/database"
)

func main () {
	dir, _ := os.Getwd()
	fmt.Println("Creat database")
	db, _ := database.StartDB ("mydb.db")
	database.GetInfor(dir + "mydb.db")

	//test func HasObj
	name := "Greedo"
	flag := db.HasObj("people", []byte(name))
	fmt.Println(flag)

	//test func SearchByID
	ID := "1"
	fmt.Printf("%s\n", db.SearchByID("people", []byte(ID)))

	//test func SearchFilmByName
	filmName := "A New Hope"
	fmt.Printf("%s\n", db.SearchFilmByName([]byte(filmName)))

	//test func SearchPersonByName
	personName := "Luke Skywalker"
	fmt.Printf("%s\n", db.SearchPersonByName([]byte(personName)))

	//test func SearchPlanetByName
	planetName := "climate"
	fmt.Printf("%s\n", db.SearchPlanetByName([]byte(planetName)))

	//test func SearchSpeciesByName
	speciesName := "Human"
	fmt.Printf("%s\n", db.SearchSpeciesByName([]byte(speciesName)))

	//test func SearchStarshipByName
	starshipName := "model"
	fmt.Printf("%s\n", db.SearchStarshipByName([]byte(starshipName)))

	//test func SearchVehicleByName
	vehicleName := "model"
	fmt.Printf("%s\n", db.SearchVehicleByName([]byte(vehicleName)))

	//test func SearchByPage
	fmt.Printf("%s\n", db.SearchByPage("people", 1))
}