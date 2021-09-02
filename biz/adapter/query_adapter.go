package adapter

import (
	"context"
	contentRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_content_server/biz/service/query"
)

type QueryAdapter interface {
	GetSearchQuery(context.Context, *contentRpc.SearchRequest) *query.SearchQuery
}

type AdQueryAdapter struct{}

func newQueryAdapterV1() QueryAdapter {
	return &AdQueryAdapter{}
}

func (q AdQueryAdapter) GetSearchQuery(ctx context.Context, req *contentRpc.SearchRequest) *query.SearchQuery {
	if req == nil {
		return nil
	}
	return &query.SearchQuery{
		Page:       req.Page,
		Limit:      req.Limit,
		SearchType: req.SearchType,
		QueryIds:   req.QueryIds,
	}
}
