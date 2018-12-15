package database

import (
	"fmt"
	"github.com/boltdb/bolt"
	"encoding/json"
    "bytes"
)

type SWDB struct {
	db *bolt.DB
}

/*
**	close the database
*/
func (m *SWDB) Close() error {
	return m.db.Close()
}

/*
**	open the database with Path
**	if data not exsist, create a database
**	return a database in SWDB
**	if wrong, return wrong message
*/
func StartDB(Path string) (*SWDB, error) {
	//open database
	db, err := bolt.Open(Path, 0644, nil)
	fmt.Println("Open the database: " + Path)
	if err != nil {
		return nil, err
	}
	return &SWDB{db}, nil
}

/*
**	initial database
**	initial with 7 buckets, buckets "user" record the information of the assign user
**	other buckets record the information of films, people, etc
**	buckets is empty
*/
func (m *SWDB) InitDB() (error) {
	//use operation bussiness
	err := m.db.Update(func(tx *bolt.Tx) error {
		//create buckets
		_, err := tx.CreateBucketIfNotExists([]byte("films"))
		_, err = tx.CreateBucketIfNotExists([]byte("people"))
		_, err = tx.CreateBucketIfNotExists([]byte("planets"))
		_, err = tx.CreateBucketIfNotExists([]byte("species"))
		_, err = tx.CreateBucketIfNotExists([]byte("starships"))
		_, err = tx.CreateBucketIfNotExists([]byte("vehicles"))
		_, err = tx.CreateBucketIfNotExists([]byte("user"))
		return err
	})
	return err
}

/*
**	add information to bucket
**	has 3 arguement:
**	buc: buckets type, as films, people, etc
**	key: object id
**	value: object information as Json
**	return wrong message
*/
func (m *SWDB) AddObj(buc string, key []byte, value []byte) (error) {
	//use operation bussiness
	err := m.db.Update(func(tx *bolt.Tx) error {
		//find the buckets
		b := tx.Bucket([]byte(buc))
		//add information; has error, return wrong message
		return b.Put(key, value)
	})
	// fmt.Println("add a new information")
	return err
}

/*
**	delete a object from database
**	has 2 arguement
**	buc: buckets type, as films, people, etc
**	key: object id
**	if has error, return wrong message
*/
func (m *SWDB) DeleteObj(buc string, key []byte) (error) {
	//use operation bussiness
	err := m.db.Update(func(tx *bolt.Tx) error {
		//find the buckets
		b := tx.Bucket([]byte(buc))
		//delete element whose id == key
		return b.Delete(key)
	})
	return err
}

/*
**	check object exsist or not by name
**	has 2 arguement
**	buc: buckets type, as films, people, etc
**	key: object name
**	return bool
*/
func (m *SWDB) HasObj (buc string, name []byte) (bool) {
	flag := false
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(buc))
        c := b.Cursor()
        var p Person
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				flag = true
				break
			}
		}
        return nil
    })
    if (flag) {
    	return true
    } else {
    	return false
    }
}

/*
**	search object by ID
**	has 2 arguement
**	buc: buckets type, as films, people, etc
**	key: object ID
**	return object by Json
*/
func (m *SWDB) SearchByID (buc string, key []byte) ([]byte) {
	var v []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(buc))
        v = b.Get(key)
        return nil
    })
    return v
}

/*
**	search film by Name
**	has 1 arguement
**	key: object Name
**	return object by Json
*/
func (m *SWDB) SearchFilmByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("films"))
        c := b.Cursor()
        var p Film
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}

/*
**	search person by Name
**	has 1 arguement
**	key: object Name
**	return object by Json
*/
func (m *SWDB) SearchPersonByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("people"))
        c := b.Cursor()
        var p Person
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}

/*
**	search planet by Name
**	has 1 arguement
**	key: object Name
**	return planet by Json
*/
func (m *SWDB) SearchPlanetByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("planets"))
        c := b.Cursor()
        var p Planet
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}

/*
**	search species by Name
**	has 1 arguement
**	key: object Name
**	return species by Json
*/
func (m *SWDB) SearchSpeciesByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("species"))
        c := b.Cursor()
        var p Species
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}

/*
**	search starship by Name
**	has 1 arguement
**	key: object Name
**	return starship by Json
*/
func (m *SWDB) SearchStarshipByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("starships"))
        c := b.Cursor()
        var p Starship
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}

/*
**	search vehicle by Name
**	has 1 arguement
**	key: object Name
**	return Vehicle by Json
*/
func (m *SWDB) SearchVehicleByName (name []byte) ([]byte) {
	var res []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("vehicles"))
        c := b.Cursor()
        var p Vehicle
	    for k, v := c.First(); k != nil; k, v = c.Next() {
			json.Unmarshal(v, &p)
			if p.getName() == string(name) {
				res = v
				break
			}
		}
        return nil
    })
    return res
}

/*
**	check log in or not
**	has 1 arguement
**	token: token
**	return bool
*/
func  (m *SWDB) IsLogIn (token []byte) (bool) {
	var v []byte
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte("user"))
        v = b.Get(token)
        return nil
    })
    if (v == nil) {
    	return false
    } else {
    	return true
    }
}

/*
**	user log in
**	add a key-value to bucket("user")
**	key: token
**	value: name
**	if has error, return wrong message
*/
func (m *SWDB) LogIn(name []byte, token []byte) (error) {
	err := m.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("user"))
		return b.Put(token, name)
	})
	return err
}

/*
**	search information by page
**	has 2 arguement
**	buc: buckets, like people, films, etc
**	page: page number
*/
func (m *SWDB) SearchByPage (buc string, page int) ([]byte) {
	var buffer bytes.Buffer
	endNum := 5 * page
	startNum := endNum - 5
	currentNum := 0
	m.db.View(func(tx *bolt.Tx) error {
        b := tx.Bucket([]byte(buc))
        c := b.Cursor()
	    for k, v := c.First(); k != nil; k, v = c.Next() {
	    	if currentNum < endNum && currentNum >= startNum {
	    		buffer.Write(v)
	    	}
			currentNum++
			if currentNum == endNum {
				break
			}
		}
        return nil
    })
    return buffer.Bytes()
}