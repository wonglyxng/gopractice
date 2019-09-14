// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package v1 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

package 命名使用 {discovery_id}.{version} 的方式, version 形如 v1, v2, v1beta ..
NOTE: 不知道的 discovery_id 请询问大佬, 新项目找大佬申请 discovery_id，先到先得抢注
e.g. account.service.v1

It is generated from these files:
	api.proto
*/
package v1

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathZeusMatch = "/live.zeus.v1.Zeus/Match"

// ==============
// Zeus Interface
// ==============

type ZeusBMServer interface {
	// `method:"POST"`
	Match(ctx context.Context, req *MatchRequest) (resp *MatchResponse, err error)
}

var v1ZeusSvc ZeusBMServer

func zeusMatch(c *bm.Context) {
	p := new(MatchRequest)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v1ZeusSvc.Match(c, p)
	c.JSON(resp, err)
}

// RegisterZeusBMServer Register the blademaster route
func RegisterZeusBMServer(e *bm.Engine, server ZeusBMServer) {
	v1ZeusSvc = server
	e.POST("/live.zeus.v1.Zeus/Match", zeusMatch)
}
