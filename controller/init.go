package controller

import (
	"context"
	"msgr/database"

	"github.com/unrolled/render"
)

var ctx context.Context
var queries *database.Queries
var rndr *render.Render

func Initialize(c context.Context, q *database.Queries) {
	rndr = render.New()
	ctx = c
	queries = q
}
