package controller

import (
	"context"
	"msgr/database"

	"github.com/unrolled/render"
)

var ctx context.Context
var queries *database.Queries
var rndr *render.Render
var frontendUrl string

func Initialize(furl string, c context.Context, q *database.Queries) {
	frontendUrl = furl
	rndr = render.New()
	ctx = c
	queries = q
}
