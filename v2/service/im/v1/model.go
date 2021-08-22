// Package imv1 code generated by lark suite oapi sdk gen
package imv1

import (
	"github.com/larksuite/oapi-sdk-go/v2"
	"io"
)

type Mention struct {
	Key       string `json:"key,omitempty"`
	Id        string `json:"id,omitempty"`
	IdType    string `json:"id_type,omitempty"`
	Name      string `json:"name,omitempty"`
	TenantKey string `json:"tenant_key,omitempty"`
}

type MessageBody struct {
	Content string `json:"content,omitempty"`
}

type Sender struct {
	Id         string `json:"id,omitempty"`
	IdType     string `json:"id_type,omitempty"`
	SenderType string `json:"sender_type,omitempty"`
	TenantKey  string `json:"tenant_key,omitempty"`
}

type Message struct {
	MessageId      string       `json:"message_id,omitempty"`
	RootId         string       `json:"root_id,omitempty"`
	ParentId       string       `json:"parent_id,omitempty"`
	MsgType        string       `json:"msg_type,omitempty"`
	CreateTime     int64        `json:"create_time,omitempty,string"`
	UpdateTime     int64        `json:"update_time,omitempty,string"`
	Deleted        bool         `json:"deleted,omitempty"`
	Updated        bool         `json:"updated,omitempty"`
	ChatId         string       `json:"chat_id,omitempty"`
	Sender         *Sender      `json:"sender,omitempty"`
	Body           *MessageBody `json:"body,omitempty"`
	Mentions       []*Mention   `json:"mentions,omitempty"`
	UpperMessageId string       `json:"upper_message_id,omitempty"`
}

type ReceiveIdType string

func (receiveIdType ReceiveIdType) Ptr() *ReceiveIdType {
	return &receiveIdType
}

const (
	ReceiveIdTypeOpenId  ReceiveIdType = "open_id"
	ReceiveIdTypeUserId  ReceiveIdType = "user_id"
	ReceiveIdTypeUnionId ReceiveIdType = "union_id"
	ReceiveIdTypeEmail   ReceiveIdType = "email"
	ReceiveIdTypeChatId  ReceiveIdType = "chat_id"
)

type MessageCreateReq struct {
	ReceiveIdType *ReceiveIdType       `query:"receive_id_type"`
	Body          MessageCreateReqBody `body:""`
}

type MessageCreateReqBody struct {
	ReceiveId *string `json:"receive_id,omitempty"`
	Content   *string `json:"content,omitempty"`
	MsgType   *string `json:"msg_type,omitempty"`
}

type MessageCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *Message `json:"data"`
}

type FileCreateReqBody struct {
	FileType *string   `json:"file_type"`
	FileName *string   `json:"file_name"`
	Duration *int      `json:"duration"`
	File     io.Reader `json:"file"`
}

type FileCreateReq struct {
	Body *FileCreateReqBody `body:""`
}

type UserId struct {
	UserId  string `json:"user_id,omitempty"`
	OpenId  string `json:"open_id,omitempty"`
	UnionId string `json:"union_id,omitempty"`
}

type EventSender struct {
	SenderId   *UserId `json:"sender_id,omitempty"`
	SenderType string  `json:"sender_type,omitempty"`
	TenantKey  string  `json:"tenant_key,omitempty"`
}

type MentionEvent struct {
	Key       string  `json:"key,omitempty"`
	Id        *UserId `json:"id,omitempty"`
	Name      string  `json:"name,omitempty"`
	TenantKey string  `json:"tenant_key,omitempty"`
}

type EventMessage struct {
	MessageId   string          `json:"message_id,omitempty"`
	RootId      string          `json:"root_id,omitempty"`
	ParentId    string          `json:"parent_id,omitempty"`
	CreateTime  int64           `json:"create_time,omitempty,string"`
	ChatId      string          `json:"chat_id,omitempty"`
	ChatType    string          `json:"chat_type,omitempty"`
	MessageType string          `json:"message_type,omitempty"`
	Content     string          `json:"content,omitempty"`
	Mentions    []*MentionEvent `json:"mentions,omitempty"`
}

type MessageReceiveEventData struct {
	Sender  *EventSender  `json:"sender,omitempty"`
	Message *EventMessage `json:"message,omitempty"`
}

type MessageReceiveEvent struct {
	*lark.EventV2Base
	Event *MessageReceiveEventData `json:"event"`
}

type FileCreateRespData struct {
	FileKey string `json:"file_key"`
}

type FileCreateResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	Data *FileCreateRespData `json:"data"`
}

type FileGetReq struct {
	FileKey string `path:"file_key"`
}

type FileGetResp struct {
	*lark.RawResponse `json:"-"`
	lark.CodeError
	File     io.Reader `json:"-"`
	FileName string    `json:"-"`
}
