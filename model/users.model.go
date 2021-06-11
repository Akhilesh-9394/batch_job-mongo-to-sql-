package model

type Users struct {
	Id        float64 `bson:"id" json:"id"`
	FirstName string  `bson:"first_name" json:"first_name"`
	LastName  string  `bson:"last_name" json:"last_name"`
	Email     string  `bson:"email" json:"email"`
	Gender    string  `bson:"gender" json:"gender"`
}
