package controller

import (
	"context"
	"msgr/database"
	"msgr/reqres"

	"github.com/unrolled/render"
)

var ctx context.Context
var queries *database.Queries

func Initialize(furl string, c context.Context, q *database.Queries) {
	reqres.FrontendUrl = furl
	reqres.Rndr = render.New()
	ctx = c
	queries = q
}
