// Package attendance code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package larkattendance

import (
	"bytes"
	"context"
	"net/http"

	"github.com/larksuite/oapi-sdk-go/v3/core"
)

func NewService(config *larkcore.Config) *AttendanceService {
	a := &AttendanceService{config: config}
	a.ApprovalInfo = &approvalInfo{service: a}
	a.File = &file{service: a}
	a.Group = &group{service: a}
	a.Shift = &shift{service: a}
	a.UserApproval = &userApproval{service: a}
	a.UserDailyShift = &userDailyShift{service: a}
	a.UserFlow = &userFlow{service: a}
	a.UserSetting = &userSetting{service: a}
	a.UserStatsData = &userStatsData{service: a}
	a.UserStatsField = &userStatsField{service: a}
	a.UserStatsView = &userStatsView{service: a}
	a.UserTask = &userTask{service: a}
	a.UserTaskRemedy = &userTaskRemedy{service: a}
	return a
}

type AttendanceService struct {
	config         *larkcore.Config
	ApprovalInfo   *approvalInfo   // 假勤审批
	File           *file           // 文件
	Group          *group          // 考勤组管理
	Shift          *shift          // 考勤班次
	UserApproval   *userApproval   // 假勤审批
	UserDailyShift *userDailyShift // 考勤排班
	UserFlow       *userFlow       // 考勤记录
	UserSetting    *userSetting    // 用户设置
	UserStatsData  *userStatsData  // 考勤统计
	UserStatsField *userStatsField // 考勤统计
	UserStatsView  *userStatsView  // 考勤统计
	UserTask       *userTask       // 考勤记录
	UserTaskRemedy *userTaskRemedy // 考勤补卡
}

type approvalInfo struct {
	service *AttendanceService
}
type file struct {
	service *AttendanceService
}
type group struct {
	service *AttendanceService
}
type shift struct {
	service *AttendanceService
}
type userApproval struct {
	service *AttendanceService
}
type userDailyShift struct {
	service *AttendanceService
}
type userFlow struct {
	service *AttendanceService
}
type userSetting struct {
	service *AttendanceService
}
type userStatsData struct {
	service *AttendanceService
}
type userStatsField struct {
	service *AttendanceService
}
type userStatsView struct {
	service *AttendanceService
}
type userTask struct {
	service *AttendanceService
}
type userTaskRemedy struct {
	service *AttendanceService
}

// 通知审批状态更新
//
// - 对于只使用飞书考勤系统而未使用飞书审批系统的企业，可以通过该接口更新写入飞书考勤系统中的三方系统审批状态，例如请假、加班、外出、出差、补卡等审批，状态包括通过、不通过、撤销等。
//
// - 发起状态的审批才可以被更新为通过、不通过，已经通过的审批才可以被更新为撤销。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/approval_info/process
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/process_approvalInfo.go
func (a *approvalInfo) Process(ctx context.Context, req *ProcessApprovalInfoReq, options ...larkcore.RequestOptionFunc) (*ProcessApprovalInfoResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/approval_infos/process"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, a.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ProcessApprovalInfoResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 下载文件
//
// - 通过文件 ID 下载指定的文件。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/download
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/download_file.go
func (f *file) Download(ctx context.Context, req *DownloadFileReq, options ...larkcore.RequestOptionFunc) (*DownloadFileResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/files/:file_id/download"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, f.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DownloadFileResp{ApiResp: apiResp}
	// 如果是下载，则设置响应结果
	if apiResp.StatusCode == http.StatusOK {
		resp.File = bytes.NewBuffer(apiResp.RawBody)
		resp.FileName = larkcore.FileNameByHeader(apiResp.Header)
		return resp, err
	}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 上传文件
//
// - 上传文件并获取文件 ID，可用于“修改用户设置”接口中的 face_key 参数。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/file/upload
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/upload_file.go
func (f *file) Upload(ctx context.Context, req *UploadFileReq, options ...larkcore.RequestOptionFunc) (*UploadFileResp, error) {
	options = append(options, larkcore.WithFileUpload())
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/files/upload"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, f.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UploadFileResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建或修改考勤组
//
// - 考勤组，是对部门或者员工在某个特定场所及特定时间段内的出勤情况（包括上下班、迟到、早退、病假、婚假、丧假、公休、工作时间、加班情况等）的一种规则设定。;;通过设置考勤组，可以从部门、员工两个维度，来设定考勤方式、考勤时间、考勤地点等考勤规则。
//
// - 出于安全考虑，目前通过该接口只允许修改自己创建的考勤组。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/create_group.go
func (g *group) Create(ctx context.Context, req *CreateGroupReq, options ...larkcore.RequestOptionFunc) (*CreateGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/groups"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, g.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateGroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除考勤组
//
// - 通过班次 ID 删除班次。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/delete_group.go
func (g *group) Delete(ctx context.Context, req *DeleteGroupReq, options ...larkcore.RequestOptionFunc) (*DeleteGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/groups/:group_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, g.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteGroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取考勤组详情
//
// - 通过考勤组 ID 获取考勤组详情。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/get_group.go
func (g *group) Get(ctx context.Context, req *GetGroupReq, options ...larkcore.RequestOptionFunc) (*GetGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/groups/:group_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, g.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetGroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取考勤组列表
//
// - 翻页获取所有考勤组列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/list_group.go
func (g *group) List(ctx context.Context, req *ListGroupReq, options ...larkcore.RequestOptionFunc) (*ListGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/groups"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, g.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListGroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (g *group) ListByIterator(ctx context.Context, req *ListGroupReq, options ...larkcore.RequestOptionFunc) (*ListGroupIterator, error) {
	return &ListGroupIterator{
		ctx:      ctx,
		req:      req,
		listFunc: g.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 按名称查询考勤组
//
// - 按考勤组名称查询考勤组摘要信息。查询条件支持名称精确匹配和模糊匹配两种方式。查询结果按考勤组修改时间 desc 排序，且最大记录数为 10 条。
//
// - 该接口依赖的数据和考勤组主数据间存在数据同步延时（正常数据同步 2 秒以内），因此在使用该接口时需注意评估数据延迟潜在风险。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/group/search
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/search_group.go
func (g *group) Search(ctx context.Context, req *SearchGroupReq, options ...larkcore.RequestOptionFunc) (*SearchGroupResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/groups/search"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, g.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &SearchGroupResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建班次
//
// - 班次是描述一次考勤任务时间规则的统称，比如一天打多少次卡，每次卡的上下班时间，晚到多长时间算迟到，晚到多长时间算缺卡等。
//
// - - 创建一个考勤组前，必须先创建一个或者多个班次。;- 一个公司内的班次是共享的，你可以直接引用他人创建的班次，但是需要注意的是，若他人修改了班次，会影响到你的考勤组及其考勤结果。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/create_shift.go
func (s *shift) Create(ctx context.Context, req *CreateShiftReq, options ...larkcore.RequestOptionFunc) (*CreateShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/shifts"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 删除班次
//
// - 通过班次 ID 删除班次。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/delete
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/delete_shift.go
func (s *shift) Delete(ctx context.Context, req *DeleteShiftReq, options ...larkcore.RequestOptionFunc) (*DeleteShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/shifts/:shift_id"
	apiReq.HttpMethod = http.MethodDelete
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &DeleteShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取班次详情
//
// - 通过班次 ID 获取班次详情。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/get_shift.go
func (s *shift) Get(ctx context.Context, req *GetShiftReq, options ...larkcore.RequestOptionFunc) (*GetShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/shifts/:shift_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取班次列表
//
// - 翻页获取所有班次列表。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/list
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/list_shift.go
func (s *shift) List(ctx context.Context, req *ListShiftReq, options ...larkcore.RequestOptionFunc) (*ListShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/shifts"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ListShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
func (s *shift) ListByIterator(ctx context.Context, req *ListShiftReq, options ...larkcore.RequestOptionFunc) (*ListShiftIterator, error) {
	return &ListShiftIterator{
		ctx:      ctx,
		req:      req,
		listFunc: s.List,
		options:  options,
		limit:    req.Limit}, nil
}

// 按名称查询班次
//
// - 通过班次的名称查询班次信息。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/shift/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_shift.go
func (s *shift) Query(ctx context.Context, req *QueryShiftReq, options ...larkcore.RequestOptionFunc) (*QueryShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/shifts/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, s.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 写入审批结果
//
// - 由于部分企业使用的是自己的审批系统，而不是飞书审批系统，因此员工的请假、加班等数据无法流入到飞书考勤系统中，导致员工在请假时间段内依然收到打卡提醒，并且被记为缺卡。;;对于这些只使用飞书考勤系统，而未使用飞书审批系统的企业，可以通过考勤开放接口的形式，将三方审批结果数据回写到飞书考勤系统中。
//
// - 目前支持写入加班、请假、出差和外出这四种审批结果，写入只会追加(insert)，不会覆盖(update)
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/create_userApproval.go
func (u *userApproval) Create(ctx context.Context, req *CreateUserApprovalReq, options ...larkcore.RequestOptionFunc) (*CreateUserApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_approvals"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取审批通过数据
//
// - 获取员工在某段时间内的请假、加班、外出和出差四种审批的通过数据。
//
// - 请假的假期时长字段，暂未开放提供，待后续提供。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_approval/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userApproval.go
func (u *userApproval) Query(ctx context.Context, req *QueryUserApprovalReq, options ...larkcore.RequestOptionFunc) (*QueryUserApprovalResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_approvals/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserApprovalResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 创建或修改班表
//
// - 班表是用来描述考勤组内人员每天按哪个班次进行上班。目前班表支持按一个整月对一位或多位人员进行排班。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/batch_create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/batchCreate_userDailyShift.go
func (u *userDailyShift) BatchCreate(ctx context.Context, req *BatchCreateUserDailyShiftReq, options ...larkcore.RequestOptionFunc) (*BatchCreateUserDailyShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_daily_shifts/batch_create"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchCreateUserDailyShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询班表信息
//
// - 支持查询多个用户的排班情况，查询的时间跨度不能超过 30 天。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_daily_shift/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userDailyShift.go
func (u *userDailyShift) Query(ctx context.Context, req *QueryUserDailyShiftReq, options ...larkcore.RequestOptionFunc) (*QueryUserDailyShiftResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_daily_shifts/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserDailyShiftResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 导入打卡流水记录
//
// - 导入授权内员工的打卡流水记录。导入后，会根据员工所在的考勤组班次规则，计算最终的打卡状态与结果。
//
// - 适用于考勤机数据导入等场景。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/batch_create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/batchCreate_userFlow.go
func (u *userFlow) BatchCreate(ctx context.Context, req *BatchCreateUserFlowReq, options ...larkcore.RequestOptionFunc) (*BatchCreateUserFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_flows/batch_create"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &BatchCreateUserFlowResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取打卡流水记录
//
// - 通过打卡记录 ID 获取用户的打卡流水记录。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/get
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/get_userFlow.go
func (u *userFlow) Get(ctx context.Context, req *GetUserFlowReq, options ...larkcore.RequestOptionFunc) (*GetUserFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_flows/:user_flow_id"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &GetUserFlowResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 批量查询打卡流水记录
//
// - 批量查询授权内员工的实际打卡流水记录。例如，企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，但是该员工在这期间打了多次卡，该接口会把所有的打卡记录都返回。
//
// - 如果只需获取打卡结果，而不需要详细的打卡数据，可使用“获取打卡结果”的接口。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_flow/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userFlow.go
func (u *userFlow) Query(ctx context.Context, req *QueryUserFlowReq, options ...larkcore.RequestOptionFunc) (*QueryUserFlowResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_flows/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserFlowResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 修改用户设置
//
// - 修改授权内员工的用户设置信息，包括人脸照片文件 ID。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/modify
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/modify_userSetting.go
func (u *userSetting) Modify(ctx context.Context, req *ModifyUserSettingReq, options ...larkcore.RequestOptionFunc) (*ModifyUserSettingResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_settings/modify"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &ModifyUserSettingResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 批量查询用户设置
//
// - 批量查询授权内员工的用户设置信息，包括人脸照片文件 ID、人脸照片更新时间。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_setting/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userSetting.go
func (u *userSetting) Query(ctx context.Context, req *QueryUserSettingReq, options ...larkcore.RequestOptionFunc) (*QueryUserSettingResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_settings/query"
	apiReq.HttpMethod = http.MethodGet
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserSettingResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询统计数据
//
// - 查询日度统计或月度统计的统计数据。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_data/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userStatsData.go
func (u *userStatsData) Query(ctx context.Context, req *QueryUserStatsDataReq, options ...larkcore.RequestOptionFunc) (*QueryUserStatsDataResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_stats_datas/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserStatsDataResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询统计表头
//
// - 查询考勤统计支持的日度统计或月度统计的统计表头。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_field/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userStatsField.go
func (u *userStatsField) Query(ctx context.Context, req *QueryUserStatsFieldReq, options ...larkcore.RequestOptionFunc) (*QueryUserStatsFieldResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_stats_fields/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserStatsFieldResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 查询统计设置
//
// - 查询开发者定制的日度统计或月度统计的统计报表表头设置信息。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userStatsView.go
func (u *userStatsView) Query(ctx context.Context, req *QueryUserStatsViewReq, options ...larkcore.RequestOptionFunc) (*QueryUserStatsViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_stats_views/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserStatsViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 更新统计设置
//
// - 更新开发者定制的日度统计或月度统计的统计报表表头设置信息。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_stats_view/update
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/update_userStatsView.go
func (u *userStatsView) Update(ctx context.Context, req *UpdateUserStatsViewReq, options ...larkcore.RequestOptionFunc) (*UpdateUserStatsViewResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_stats_views/:user_stats_view_id"
	apiReq.HttpMethod = http.MethodPut
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &UpdateUserStatsViewResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取打卡结果
//
// - 获取企业内员工的实际打卡结果，包括上班打卡结果和下班打卡结果。
//
// - - 如果企业给一个员工设定的班次是上午 9 点和下午 6 点各打一次上下班卡，即使员工在这期间打了多次卡，该接口也只会返回 1 条记录。;- 如果要获取打卡的详细数据，如打卡位置等信息，可使用“获取打卡流水记录”或“批量查询打卡流水记录”的接口。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userTask.go
func (u *userTask) Query(ctx context.Context, req *QueryUserTaskReq, options ...larkcore.RequestOptionFunc) (*QueryUserTaskResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_tasks/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserTaskResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 通知补卡审批发起
//
// - 对于只使用飞书考勤系统而未使用飞书审批系统的企业，可以通过该接口，将在三方审批系统中发起的补卡审批数据，写入到飞书考勤系统中，状态为审批中。写入后可以由[通知审批状态更新](https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/approval_info/process) 进行状态更新
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/create
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/create_userTaskRemedy.go
func (u *userTaskRemedy) Create(ctx context.Context, req *CreateUserTaskRemedyReq, options ...larkcore.RequestOptionFunc) (*CreateUserTaskRemedyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_task_remedys"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &CreateUserTaskRemedyResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取补卡记录
//
// - 获取授权内员工的补卡记录。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/query_userTaskRemedy.go
func (u *userTaskRemedy) Query(ctx context.Context, req *QueryUserTaskRemedyReq, options ...larkcore.RequestOptionFunc) (*QueryUserTaskRemedyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_task_remedys/query"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserTaskRemedyResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}

// 获取用户可补卡时间
//
// - 获取用户某天可以补的第几次上 / 下班卡的时间。
//
// - 官网API文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/attendance-v1/user_task_remedy/query_user_allowed_remedys
//
// - 使用Demo链接:https://github.com/larksuite/oapi-sdk-go/tree/v3_main/sample/apiall/attendancev1/queryUserAllowedRemedys_userTaskRemedy.go
func (u *userTaskRemedy) QueryUserAllowedRemedys(ctx context.Context, req *QueryUserAllowedRemedysUserTaskRemedyReq, options ...larkcore.RequestOptionFunc) (*QueryUserAllowedRemedysUserTaskRemedyResp, error) {
	// 发起请求
	apiReq := req.apiReq
	apiReq.ApiPath = "/open-apis/attendance/v1/user_task_remedys/query_user_allowed_remedys"
	apiReq.HttpMethod = http.MethodPost
	apiReq.SupportedAccessTokenTypes = []larkcore.AccessTokenType{larkcore.AccessTokenTypeTenant}
	apiResp, err := larkcore.Request(ctx, apiReq, u.service.config, options...)
	if err != nil {
		return nil, err
	}
	// 反序列响应结果
	resp := &QueryUserAllowedRemedysUserTaskRemedyResp{ApiResp: apiResp}
	err = apiResp.JSONUnmarshalBody(resp)
	if err != nil {
		return nil, err
	}
	return resp, err
}
