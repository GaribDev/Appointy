package models

type Participant struct {
	Name  string `bson:"name" json:"name"`
	Email string `bson:"email" json:"email"`
	Rsvp  string `bson:"rsvp" json:"rsvp"`
}
