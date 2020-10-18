package dao

import (
	"log"

	. "github.com/GaribDev/Appointy/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type MeetingDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "movies"
)

// Establish a connection to database
func (m *MeetingDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of meetings
func (m *MeetingDAO) FindAll() ([]Movie, error) {
	var meetings []Meeting
	err := db.C(COLLECTION).Find(bson.M{}).All(&meetings)
	return meetings, err
}

// Find a meeting by its id
func (m *MeetingDAO) FindById(id string) (Meeting, error) {
	var meet Meeting
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&meet)
	return meet, err
}

// Insert a meeting into database
func (m *MeetingDAO) Insert(meeting Meeting) error {
	err := db.C(COLLECTION).Insert(&meeting)
	return err
}
