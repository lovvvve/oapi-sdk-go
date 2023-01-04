//  code generated by oapi sdk gen
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

package lark

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/acs/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/admin/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/application/v6"
	"github.com/larksuite/oapi-sdk-go/v3/service/approval/v4"
	"github.com/larksuite/oapi-sdk-go/v3/service/attendance/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/authen/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/baike/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/bitable/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/block/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/calendar/v4"
	"github.com/larksuite/oapi-sdk-go/v3/service/contact/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/docx/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/ehr/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/event/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/ext"
	"github.com/larksuite/oapi-sdk-go/v3/service/face_detection/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/gray_test_open_sg/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/helpdesk/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/hire/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/human_authentication/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/im/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/mail/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/okr/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/optical_char_recognition/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/passport/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/search/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/sheets/v3"
	"github.com/larksuite/oapi-sdk-go/v3/service/speech_to_text/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/task/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/tenant/v2"
	"github.com/larksuite/oapi-sdk-go/v3/service/translation/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/vc/v1"
	"github.com/larksuite/oapi-sdk-go/v3/service/wiki/v2"
)

type Client struct {
	config                 *larkcore.Config
	Acs                    *larkacs.AcsService                                         // 智能门禁
	Admin                  *larkadmin.AdminService                                     // 管理后台-企业勋章
	Application            *larkapplication.ApplicationService                         // 应用信息
	Approval               *larkapproval.ApprovalService                               // 审批
	Attendance             *larkattendance.AttendanceService                           // 打卡
	Authen                 *larkauthen.AuthenService                                   //
	Baike                  *larkbaike.BaikeService                                     // 企业百科
	Bitable                *larkbitable.BitableService                                 // 云文档-多维表格
	Block                  *larkblock.BlockService                                     // 小组件
	Calendar               *larkcalendar.CalendarService                               // 日历
	Contact                *larkcontact.ContactService                                 // 搜索
	Corehr                 *larkcorehr.CorehrService                                   // CoreHR
	Docx                   *larkdocx.DocxService                                       // 云文档-文档
	Drive                  *larkdrive.DriveService                                     // 云文档-文件管理
	Ehr                    *larkehr.EhrService                                         // 智能人事
	Event                  *larkevent.EventService                                     // 事件订阅
	FaceDetection          *larkface_detection.FaceDetectionService                    // AI能力
	GrayTestOpenSg         *larkgray_test_open_sg.GrayTestOpenSgService                //
	Helpdesk               *larkhelpdesk.HelpdeskService                               // 服务台
	Hire                   *larkhire.HireService                                       // 招聘
	HumanAuthentication    *larkhuman_authentication.HumanAuthenticationService        // 实名认证
	Im                     *larkim.ImService                                           // 消息与群组
	Mail                   *larkmail.MailService                                       // 邮箱
	Okr                    *larkokr.OkrService                                         // OKR
	OpticalCharRecognition *larkoptical_char_recognition.OpticalCharRecognitionService // AI能力
	Passport               *larkpassport.PassportService                               // 帐号
	Search                 *larksearch.SearchService                                   // 搜索
	Sheets                 *larksheets.SheetsService                                   // 云文档-电子表格
	SpeechToText           *larkspeech_to_text.SpeechToTextService                     // AI能力
	Task                   *larktask.TaskService                                       // 任务
	Tenant                 *larktenant.TenantService                                   // 企业信息
	Translation            *larktranslation.TranslationService                         // AI能力
	Vc                     *larkvc.VcService                                           // 视频会议
	Wiki                   *larkwiki.WikiService                                       // 云文档-知识库
	Ext                    *larkext.ExtService
}

type ClientOptionFunc func(config *larkcore.Config)

func WithAppType(appType larkcore.AppType) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.AppType = appType
	}
}

func WithMarketplaceApp() ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.AppType = larkcore.AppTypeMarketplace
	}
}

func WithEnableTokenCache(enableTokenCache bool) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.EnableTokenCache = enableTokenCache
	}
}

func WithLogger(logger larkcore.Logger) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Logger = logger
	}
}

func WithOpenBaseUrl(baseUrl string) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.BaseUrl = baseUrl
	}
}

func WithLogLevel(logLevel larkcore.LogLevel) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.LogLevel = logLevel
	}
}

func WithTokenCache(cache larkcore.Cache) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.TokenCache = cache
	}
}

func WithLogReqAtDebug(printReqRespLog bool) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.LogReqAtDebug = printReqRespLog
	}
}

func WithHttpClient(httpClient larkcore.HttpClient) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.HttpClient = httpClient
	}
}

func WithHelpdeskCredential(helpdeskID, helpdeskToken string) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.HelpDeskId = helpdeskID
		config.HelpDeskToken = helpdeskToken
		if helpdeskID != "" && helpdeskToken != "" {
			config.HelpdeskAuthToken = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", helpdeskID, helpdeskToken)))
		}
	}
}

func WithSerialization(serializable larkcore.Serializable) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Serializable = serializable
	}
}

func WithReqTimeout(reqTimeout time.Duration) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.ReqTimeout = reqTimeout
	}
}

// 设置每次请求都会携带的 header
func WithHeaders(header http.Header) ClientOptionFunc {
	return func(config *larkcore.Config) {
		config.Header = header
	}
}

func NewClient(appId, appSecret string, options ...ClientOptionFunc) *Client {
	// 构建配置
	config := &larkcore.Config{
		BaseUrl:          FeishuBaseUrl,
		AppId:            appId,
		AppSecret:        appSecret,
		EnableTokenCache: true,
		AppType:          larkcore.AppTypeSelfBuilt,
	}
	for _, option := range options {
		option(config)
	}

	// 构建日志器
	larkcore.NewLogger(config)

	// 构建缓存
	larkcore.NewCache(config)

	// 创建序列化器
	larkcore.NewSerialization(config)

	// 创建httpclient
	larkcore.NewHttpClient(config)

	// 创建sdk-client，并初始化服务
	client := &Client{config: config}
	initService(client, config)

	// 触发重推 appTicket,如果是 ISV 的话
	resendAppTicketIfNeed(client)
	return client
}

func initService(client *Client, config *larkcore.Config) {
	client.Acs = larkacs.NewService(config)
	client.Admin = larkadmin.NewService(config)
	client.Application = larkapplication.NewService(config)
	client.Approval = larkapproval.NewService(config)
	client.Attendance = larkattendance.NewService(config)
	client.Authen = larkauthen.NewService(config)
	client.Baike = larkbaike.NewService(config)
	client.Bitable = larkbitable.NewService(config)
	client.Block = larkblock.NewService(config)
	client.Calendar = larkcalendar.NewService(config)
	client.Contact = larkcontact.NewService(config)
	client.Corehr = larkcorehr.NewService(config)
	client.Docx = larkdocx.NewService(config)
	client.Drive = larkdrive.NewService(config)
	client.Ehr = larkehr.NewService(config)
	client.Event = larkevent.NewService(config)
	client.FaceDetection = larkface_detection.NewService(config)
	client.GrayTestOpenSg = larkgray_test_open_sg.NewService(config)
	client.Helpdesk = larkhelpdesk.NewService(config)
	client.Hire = larkhire.NewService(config)
	client.HumanAuthentication = larkhuman_authentication.NewService(config)
	client.Im = larkim.NewService(config)
	client.Mail = larkmail.NewService(config)
	client.Okr = larkokr.NewService(config)
	client.OpticalCharRecognition = larkoptical_char_recognition.NewService(config)
	client.Passport = larkpassport.NewService(config)
	client.Search = larksearch.NewService(config)
	client.Sheets = larksheets.NewService(config)
	client.SpeechToText = larkspeech_to_text.NewService(config)
	client.Task = larktask.NewService(config)
	client.Tenant = larktenant.NewService(config)
	client.Translation = larktranslation.NewService(config)
	client.Vc = larkvc.NewService(config)
	client.Wiki = larkwiki.NewService(config)
	client.Ext = larkext.NewService(config)
}

func resendAppTicketIfNeed(client *Client) {
	defer func() {
		if err := recover(); err != nil {
			client.config.Logger.Error(context.Background(), fmt.Sprintf("resendAppTicketIfNeed, error:%v", err))
		}
	}()

	if client.config.AppType == larkcore.AppTypeMarketplace {
		ctx := context.Background()
		resp, err := client.ResendAppTicket(ctx, &larkcore.ResendAppTicketReq{
			AppID:     client.config.AppId,
			AppSecret: client.config.AppSecret,
		})
		if err != nil {
			client.config.Logger.Info(ctx, fmt.Sprintf("resend appTicket error:%v", err))
			return
		}
		client.config.Logger.Info(ctx, fmt.Sprintf("resend appTicket resp:%v", resp))

	}
}

func (cli *Client) Post(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Do(ctx context.Context, apiReq *larkcore.ApiReq, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return larkcore.Request(ctx, apiReq, cli.config, options...)
}
func (cli *Client) Get(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodGet,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Delete(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodDelete,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Put(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPut,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) Patch(ctx context.Context, httpPath string, body interface{}, accessTokeType larkcore.AccessTokenType, options ...larkcore.RequestOptionFunc) (*larkcore.ApiResp, error) {
	return cli.Do(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPatch,
		ApiPath:                   httpPath,
		Body:                      body,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{accessTokeType},
	}, options...)
}

func (cli *Client) GetAppAccessTokenBySelfBuiltApp(ctx context.Context, req *larkcore.SelfBuiltAppAccessTokenReq) (*larkcore.AppAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.AppAccessTokenInternalUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.AppAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) GetAppAccessTokenByMarketplaceApp(ctx context.Context, req *larkcore.MarketplaceAppAccessTokenReq) (*larkcore.AppAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.AppAccessTokenUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.AppAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) GetTenantAccessTokenBySelfBuiltApp(ctx context.Context, req *larkcore.SelfBuiltTenantAccessTokenReq) (*larkcore.TenantAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.TenantAccessTokenInternalUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.TenantAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) GetTenantAccessTokenByMarketplaceApp(ctx context.Context, req *larkcore.MarketplaceTenantAccessTokenReq) (*larkcore.TenantAccessTokenResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.TenantAccessTokenUrlPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.TenantAccessTokenResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

func (cli *Client) ResendAppTicket(ctx context.Context, req *larkcore.ResendAppTicketReq) (*larkcore.ResendAppTicketResp, error) {
	rawResp, err := larkcore.Request(ctx, &larkcore.ApiReq{
		HttpMethod:                http.MethodPost,
		ApiPath:                   larkcore.ApplyAppTicketPath,
		Body:                      req,
		SupportedAccessTokenTypes: []larkcore.AccessTokenType{larkcore.AccessTokenTypeNone},
	}, cli.config)

	if err != nil {
		return nil, err
	}
	resp := &larkcore.ResendAppTicketResp{}
	err = json.Unmarshal(rawResp.RawBody, resp)
	if err != nil {
		return nil, err
	}
	resp.ApiResp = rawResp

	return resp, nil
}

var FeishuBaseUrl = "https://open.feishu.cn"
var LarkBaseUrl = "https://open.larksuite.com"
