package query

import (
	contentRpc "github.com/zbwang163/ad_content_overpass"
)

type SearchQuery struct {
	Page       string                              `protobuf:"bytes,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit      string                              `protobuf:"bytes,2,opt,name=limit,proto3" json:"limit,omitempty"`
	SearchType contentRpc.SearchRequest_SearchType `protobuf:"varint,3,opt,name=search_type,json=searchType,proto3,enum=SearchRequest_SearchType" json:"search_type,omitempty"`
	QueryIds   []string                            `protobuf:"bytes,4,rep,name=query_ids,json=queryIds,proto3" json:"query_ids,omitempty"`
}
