package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LocationController struct{}

var id int = 0

var url string = "test:test@ds059888.mongolab.com:59888/cmpe273"// "mongodb://localhost"


func generateID() int {
	id = id + 1
	return id
}

func (l *LocationController) CreateLocation(location Location) (loc Location, err error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return loc, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cmpe273").C("location")

	location.Id = generateID()
	g := fetchGeocode(location)
	location.Coordinate = g
	err = c.Insert(location)
	if err != nil {
		return loc, err
	}
	return location, nil
}

func (l *LocationController) GetLocation(id int) (loc Location, err error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return loc, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cmpe273").C("location")
	result := Location{}
	err = c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return loc, err
	}
	fmt.Println(result)
	return result, nil
}

func (l *LocationController) DeleteLocation(id int) error {
	session, err := mgo.Dial(url)
	if err != nil {
		return err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cmpe273").C("location")
	err = c.Remove(bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}

func (l *LocationController) UpdateLocation(id int, location Location) (loc Location, err error) {
	session, err := mgo.Dial(url)
	if err != nil {
		return loc, err
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("cmpe273").C("location")
	colQuerier := bson.M{"_id": id}
	location.Coordinate = fetchGeocode(location)
	change := bson.M{"$set": bson.M{"address": location.Address, "city": location.City, "state": location.State, "zip": location.Zip, "coordinates": location.Coordinate}}
	err = c.Update(colQuerier, change)
	if err != nil {
		return loc, err
	}
	result := Location{}
	err = c.Find(bson.M{"_id": id}).One(&result)
	if err != nil {
		return loc, err
	}
	return result, nil
}
