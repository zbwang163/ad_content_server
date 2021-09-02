package test

import (
	"context"
	"github.com/zbwang163/ad_common/rpc"
	contentRpc "github.com/zbwang163/ad_content_overpass"
	"github.com/zbwang163/ad_content_server/common/consts"
	"testing"
)

func TestSearch(t *testing.T) {
	rpc.InitRpc(consts.PSM)
	resp, err := rpc.ContentClient.Search(context.Background(), &contentRpc.SearchRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}
