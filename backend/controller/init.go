package controller

import (
	"context"
	"msgr/database"
	"msgr/reqres"
)

var ctx context.Context
var queries *database.Queries

func Initialize(furl string, c context.Context, q *database.Queries) {
	reqres.FrontendUrl = furl
	ctx = c
	queries = q
}
