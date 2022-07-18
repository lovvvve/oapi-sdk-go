// Package tenant code generated by oapi sdk gen
package larktenant

import (
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

// 构建业务域服务实例
func NewService(config *larkcore.Config) *TenantService {
	t := &TenantService{config: config}
	t.Tenant = &tenant{service: t}
	return t
}

// 业务域服务定义
type TenantService struct {
	config *larkcore.Config
	Tenant *tenant
}

// 资源服务定义
type tenant struct {
	service *TenantService
}

// 资源服务方法定义
func (t *tenant) Query(ctx context.Context, options ...larkcore.RequestOptionFunc) (*QueryTenantResp, error) {
	// 发起请求
	httpReq := &larkcore.HttpReq{
		PathParams:  larkcore.PathParams{},
		QueryParams: larkcore.QueryParams{},
	}
	httpReq.ApiPath = "/open-apis/tenant/v2/tenant/query"
	httpReq.HttpMethod = http.MethodGet
	httpReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	rawResp, err := larkcore.Request(ctx, httpReq, t.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryTenantResp{RawResponse: rawResp}
	err = rawResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
