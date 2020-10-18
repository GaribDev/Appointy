package models

import "gopkg.in/mgo.v2/bson"

type Participant struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Rsvp  string `bson:"rsvp" json:"rsvp"`
}

type Meeting struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Title     string        `bson:"title" json:"title"`
	StartTime string        `bson:"stime" json:"stime"`
	EndTime   string        `bson:"etime" json:"etime"`
	Timestamp string        `bson:"timestamp" json:"timestamp"`
	Participant
}
