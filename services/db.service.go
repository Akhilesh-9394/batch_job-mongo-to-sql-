package services

import (
	"batch_job/db"
	"batch_job/model"
	"context"
	"time"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	database    = db.Newdatabase()
	client, _   = database.DBinstance()
	dbtoconnect = database.OpenDatabase(client, "test")
	collection  = database.OpenCollection(client, dbtoconnect, "users")
	users       []*model.Users
	sqldb             = db.ConnectSql()
	skip        int64 = 0
)

func Getdata(limit int64) {

	options := options.Find().SetLimit(limit)
	options.SetSkip(skip)
	cursor, err := collection.Find(context.TODO(), bson.D{}, options)
	if err != nil {
		log.Info(err)
	}
	if err = cursor.All(context.TODO(), &users); err != nil {
		log.Info(err)
	}
	for _, user := range users {
		insert(user)
	}
	skip = skip + limit
	log.Debugln(skip)
}

func insert(user *model.Users) {
	query := "INSERT INTO userinfo(userId,firstname,lastname,email,gender) VALUES (?,?,?,?,?)"
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	stmt, err := sqldb.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error %s when preparing SQL statement", err)
	}
	defer stmt.Close()
	res, err := stmt.ExecContext(ctx, user.Id, user.FirstName, user.LastName, user.Email, user.Gender)
	if err != nil {
		log.Printf("Error %s when inserting row into products table", err)
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when finding rows affected", err)
	}
	log.Printf("%d products created ", rows)
}
