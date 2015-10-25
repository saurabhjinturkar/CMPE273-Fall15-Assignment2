package main

type Coordinate struct {
	Lat float64 `json:"lat" bson:"lat"`
	Lng float64 `json: "lng" bson:"lng"`
}

type Location struct {
	Id         int        `json:"id" bson:"_id,omitempty"`
	Name       string     `json:"name" bson:"name"`
	Address    string     `json:"address" bson:"address"`
	City       string     `json:"city" bson:"city"`
	State      string     `json:"state" bson:"state"`
	Zip        string     `json:"zip" bson:"zip"`
	Coordinate Coordinate `json:"coordinates" bson:"coordinates"`
}

func (l *Location) validate() bool {
	if len(l.Address) == 0 || len(l.City) == 0 || len(l.State) == 0 {
		return false
	}
	return true
}
