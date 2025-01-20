// Needs a copy of .env file in folder

package tests

import (
	"context"
	"fmt"
	"log"
	"msgr/database"
	"msgr/models"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var queries *database.Queries
var ctx context.Context
var test *testing.T

func TestQueries(t *testing.T) {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env file was not found!")
	}

	var dbUrl string = os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("variable DB_URL was not found in .env file")
	}

	test = t

	cotx, conn, err := database.GetConnection(dbUrl)
	if err != nil {
		test.Fatalf(err.Error())
	}
	ctx = cotx
	defer conn.Close(ctx)

	queries = database.New(conn)

	id := insertQuery()

	getQuery(id)

	deleteQuery(id)
}

func insertQuery() uuid.UUID {
	id := uuid.New()
	name := "Juan"
	password := "123"
	email := "email@example.com"

	insert := database.InsertUserParams{
		ID:       models.ToPgtypeUUID(id),
		Name:     name,
		Password: password,
		Email:    email,
	}

	pgid, err := queries.InsertUser(ctx, insert)
	if err != nil {
		test.Fatalf(err.Error())
	}

	if id.String() != pgid.String() {
		test.Fatalf("%s != %s", id.String(), pgid.String())
	}

	return models.ToGoogleUUID(pgid)
}

func getQuery(id uuid.UUID) {
	user, err := queries.GetUser(ctx, models.ToPgtypeUUID(id))
	if err != nil {
		test.Fatalf(err.Error())
	}
	fmt.Println(user)
	fmt.Println()
}

func deleteQuery(id uuid.UUID) {
	pgid, err := queries.DeleteUser(ctx, models.ToPgtypeUUID(id))
	if err != nil {
		test.Fatalf("Did not delete user")
	} else {
		fmt.Printf("Deleted user: %s", pgid.String())
	}
}
