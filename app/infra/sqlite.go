package infra

import (
	"context"
	"log"
	"strconv"
	"tamaribacms/ent"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite3Connection() (*ent.Client, error) {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Printf("failed opening connection to sqlite: %v", err)
	}
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		log.Printf("failed creating schema resources: %v", err)
	}

	for i := 1; i <= 5; i++ {
		CreateSample(ctx, client, i)
	}

	return client, err
}

func CreateSample(ctx context.Context, client *ent.Client, id int) (*ent.User, error) {
	user, err := client.User.
		Create().
		SetUsername("user" + strconv.Itoa(id)).
		SetAge(10 * id).
		SetCreatedAt(time.Now()).
		SetUpdatedAt(time.Now()).
		Save(ctx)

	if err != nil {
		log.Print("failed creating sample: ", err)
	}
	log.Printf("created: %v", user)

	return user, err
}
