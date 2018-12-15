package database

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"github.com/peterhellberg/swapi"
)

func GetInfor(DBPath string) {

	filmsCount := 7
	peopleCount := 87
	planetsCount := 61
	speciesCount := 37
	starshipsCount := 37
	vehiclesCount := 39

	//create Database
	db, _ := StartDB (DBPath)
	
	//Initial Database
	db.InitDB();

	//get film
	i := 1
	//open file
	// fd,_ := os.OpenFile(filmsPath, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0644)
	//get film information
	for ; i <= filmsCount; i++ {
		c := swapi.DefaultClient
		film := dump(c.Film(i))
		//transport json to data
		data, _ := json.Marshal(film)
		b := []byte(data)
		//add information to db
		db.AddObj("films", []byte(strconv.Itoa(i)), b)
		// //write in file
    	// fd.Write(b)
	}
    // fd.Close()
	fmt.Print("get file num: ")
	fmt.Println(i - 1)

	//get person
	i = 1
	// fd,_ = os.OpenFile(peoplePath, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0644)
	for ; i <= peopleCount; i++ {
		c := swapi.DefaultClient
		people := dump(c.Person(i))
		data, _ := json.Marshal(people)
		b := []byte(data)
		db.AddObj("people", []byte(strconv.Itoa(i)), b)
    	// fd.Write(b)
	}
    // fd.Close()
	fmt.Print("get people num: ")
	fmt.Println(i - 1)

	//get planet
	i = 1
	// fd,_ = os.OpenFile(planetPath, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0644)
	for ; i <= planetsCount; i++ {
		c := swapi.DefaultClient
		planet := dump(c.Planet(i))
		data, _ := json.Marshal(planet)
		b := []byte(data)
		db.AddObj("planets", []byte(strconv.Itoa(i)), b)
    	// fd.Write(b)
	}
    // fd.Close()
	fmt.Print("get planet num: ")
	fmt.Println(i - 1)

	//get species
	i = 1
	// fd,_ = os.OpenFile(speciesPath, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0644)
	for ; i <= speciesCount; i++ {
		c := swapi.DefaultClient
		species := dump(c.Species(i))
		data, _ := json.Marshal(species)
		b := []byte(data)
		db.AddObj("species", []byte(strconv.Itoa(i)), b)
    	// fd.Write(b)
	}
    // fd.Close()
	fmt.Print("get species num: ")
	fmt.Println(i - 1)

	//get starship
	i = 1
	// fd,_ = os.OpenFile(starshipsPath, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0644)
	for ; i <= starshipsCount; i++ {
		c := swapi.DefaultClient
		starships := dump(c.Starship(i))
		data, _ := json.Marshal(starships)
		b := []byte(data)
		db.AddObj("starships", []byte(strconv.Itoa(i)), b)
    	// fd.Write(b)
	}
    // fd.Close()
	fmt.Print("get starships num: ")
	fmt.Println(i - 1)

	//get vehicle
	i = 1
	// fd,_ = os.OpenFile(vehiclesPath, os.O_RDWR| os.O_CREATE| os.O_APPEND, 0644)
	for ; i <= vehiclesCount; i++ {
		c := swapi.DefaultClient
		vehicles := dump(c.Vehicle(i))
		data, _ := json.Marshal(vehicles)
		b := []byte(data)
		db.AddObj("vehicles", []byte(strconv.Itoa(i)), b)
    	// fd.Write(b)
	}
    // fd.Close()
	fmt.Print("get vehicles num: ")
	fmt.Println(i - 1)
	db.Close()
}

func analysis (src string) (inforType string, index string) {
	result := strings.Split(src, "/")
	return result[4], result[5]
}

func usage() {
	fmt.Println(strings.Join([]string{
		"Commands:",
		"film     [id]",
		"person   [id]",
		"planet   [id]",
		"species  [id]",
		"starship [id]",
		"vehicle  [id]",
	}, "\n\t"))
}

func dump(data interface{}, err error) (mes interface{}) {
	if err != nil {
		return
	}
	return data
}