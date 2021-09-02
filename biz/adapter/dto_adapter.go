package adapter

import (
	"context"
	"github.com/zbwang163/ad_common/biz_error"
	contentRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_content_server/biz/service/dto"
	"github.com/zbwang163/ad_content_server/common/utils"
)

type DtoAdapter interface {
	GetSearchResponse(context.Context, *dto.SearchDTO, *biz_error.BizError) *contentRpc.SearchResponse
}

type AdDtoAdapter struct{}

func newDtoAdapterV1() DtoAdapter {
	return &AdDtoAdapter{}
}

func (d AdDtoAdapter) GetSearchResponse(ctx context.Context, searchDTO *dto.SearchDTO, bizError *biz_error.BizError) *contentRpc.SearchResponse {
	if searchDTO == nil {
		return nil
	}
	data := &contentRpc.SearchResponse_SearchData{
		Docs:       searchDTO.Docs,
		Pagination: searchDTO.Pagination,
	}
	// 固定写法
	resp := &contentRpc.SearchResponse{}
	resp.Data = data
	resp.Base = utils.GetBaseResp(bizError)
	return resp
}
