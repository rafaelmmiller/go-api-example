package models

import (
	z "github.com/Oudwins/zog"
)

type Product struct {
	ID          string
	Name        string
	Description string
	Price       float64
}

var ProductSchema = z.Struct(z.Schema{
	"id":          z.String().UUID().Required(),
	"name":        z.String().Min(1).Required(),
	"description": z.String().Min(1).Required(),
	"price":       z.Float().GT(0).Required(),
})
