package dto

import (
	"github.com/zbwang163/ad_base_overpass"
	contentRpc "github.com/zbwang163/ad_content_overpass"
)

type SearchDTO struct {
	Docs       []*contentRpc.Case
	Pagination *ad_base_overpass.Pagination
}
