package main

import (
	"context"
	contentRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_content_server/biz/adapter"
	"log"
)

type server struct {
	contentRpc.UnimplementedContentServiceServer
	contentAdapter adapter.ContentAdapter
}

func NewServer() *server {
	return &server{
		contentAdapter: adapter.NewContentAdapter(),
	}
}

func (s *server) Search(ctx context.Context, req *contentRpc.SearchRequest) (*contentRpc.SearchResponse, error) {
	log.Println("Search called")
	query := s.contentAdapter.GetQueryAdapter().GetSearchQuery(ctx, req)
	dto, bizError := s.contentAdapter.Search(ctx, query)
	log.Printf("search dto:%v", dto)
	return s.contentAdapter.GetDtoAdapter().GetSearchResponse(ctx, dto, bizError), nil
}
