package adapter

import (
	"context"
	"github.com/go-playground/validator"
	accountRpc "github.com/zbwang163/ad_account_overpass"
	"github.com/zbwang163/ad_base_overpass"
	"github.com/zbwang163/ad_common/biz_error"
	logs "github.com/zbwang163/ad_common/logger"
	"github.com/zbwang163/ad_common/rpc"
	"github.com/zbwang163/ad_content_server/biz/service/dto"
	"github.com/zbwang163/ad_content_server/biz/service/query"
)

type ContentAdapter interface {
	GetQueryAdapter() QueryAdapter
	GetDtoAdapter() DtoAdapter

	Search(context.Context, *query.SearchQuery) (*dto.SearchDTO, *biz_error.BizError)
}

type AdContentAdapter struct {
	QueryAdapter QueryAdapter
	DtoAdapter   DtoAdapter
}

func NewContentAdapter() ContentAdapter {
	return &AdContentAdapter{
		QueryAdapter: newQueryAdapterV1(),
		DtoAdapter:   newDtoAdapterV1(),
	}
}

func (a AdContentAdapter) GetQueryAdapter() QueryAdapter {
	return a.QueryAdapter
}

func (a AdContentAdapter) GetDtoAdapter() DtoAdapter {
	return a.DtoAdapter
}

func (a AdContentAdapter) Search(ctx context.Context, query *query.SearchQuery) (*dto.SearchDTO, *biz_error.BizError) {
	err := validator.New().Struct(query)
	if err != nil {
		return nil, biz_error.NewParamError(err)
	}

	resp, err := rpc.AccountClient.Login(ctx, &accountRpc.LoginRequest{})
	if err != nil {
		logs.CtxError(ctx, "call account service Login err:%v", err.Error())
		return nil, biz_error.NewResourceError(err)
	}
	return &dto.SearchDTO{
		Docs:       nil,
		Pagination: &ad_base_overpass.Pagination{QueryIds: []string{resp.String()}},
	}, nil
}
