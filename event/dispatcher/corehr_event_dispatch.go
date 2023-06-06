// Package dispatcher code generated by oapi sdk gen
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

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v1"
)

//
//
// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2ContractCreatedV1(handler func(ctx context.Context, event *larkcorehr.P2ContractCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.contract.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.contract.created_v1")
	}
	dispatcher.eventType2EventHandler["corehr.contract.created_v1"] = larkcorehr.NewP2ContractCreatedV1Handler(handler)
	return dispatcher
}

// 部门创建
//
// - 飞书人事中「部门被创建」时将触发此事件。触发时间为部门实际生效时间，如在 2022-01-01 创建部门，部门生效时间设置为 2022-05-01，事件将在 2022-05-01 进行推送。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/events/created
func (dispatcher *EventDispatcher) OnP2DepartmentCreatedV1(handler func(ctx context.Context, event *larkcorehr.P2DepartmentCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.department.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.department.created_v1")
	}
	dispatcher.eventType2EventHandler["corehr.department.created_v1"] = larkcorehr.NewP2DepartmentCreatedV1Handler(handler)
	return dispatcher
}

// 部门删除
//
// - 飞书人事中「部门被删除」时将触发此事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/events/deleted
func (dispatcher *EventDispatcher) OnP2DepartmentDeletedV1(handler func(ctx context.Context, event *larkcorehr.P2DepartmentDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.department.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.department.deleted_v1")
	}
	dispatcher.eventType2EventHandler["corehr.department.deleted_v1"] = larkcorehr.NewP2DepartmentDeletedV1Handler(handler)
	return dispatcher
}

// 部门更新
//
// - 飞书人事中「部门信息被更新」时将触发此事件。触发时间为部门更新实际生效时间，如在 2022-01-01 更新部门，部门更新生效时间设置为 2022-05-01，事件将在 2022-05-01 进行推送。
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/department/events/updated
func (dispatcher *EventDispatcher) OnP2DepartmentUpdatedV1(handler func(ctx context.Context, event *larkcorehr.P2DepartmentUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.department.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.department.updated_v1")
	}
	dispatcher.eventType2EventHandler["corehr.department.updated_v1"] = larkcorehr.NewP2DepartmentUpdatedV1Handler(handler)
	return dispatcher
}

// 员工转正
//
// - 员工在飞书人事转正完成后将触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/events/converted
func (dispatcher *EventDispatcher) OnP2EmploymentConvertedV1(handler func(ctx context.Context, event *larkcorehr.P2EmploymentConvertedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.employment.converted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.employment.converted_v1")
	}
	dispatcher.eventType2EventHandler["corehr.employment.converted_v1"] = larkcorehr.NewP2EmploymentConvertedV1Handler(handler)
	return dispatcher
}

// 雇佣信息创建
//
// - 员工在飞书人事的「雇佣信息被创建」时将触发此事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/events/created
func (dispatcher *EventDispatcher) OnP2EmploymentCreatedV1(handler func(ctx context.Context, event *larkcorehr.P2EmploymentCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.employment.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.employment.created_v1")
	}
	dispatcher.eventType2EventHandler["corehr.employment.created_v1"] = larkcorehr.NewP2EmploymentCreatedV1Handler(handler)
	return dispatcher
}

// 雇佣信息删除
//
// - 员工在飞书人事的「雇佣信息被删除」时将触发此事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/events/deleted
func (dispatcher *EventDispatcher) OnP2EmploymentDeletedV1(handler func(ctx context.Context, event *larkcorehr.P2EmploymentDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.employment.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.employment.deleted_v1")
	}
	dispatcher.eventType2EventHandler["corehr.employment.deleted_v1"] = larkcorehr.NewP2EmploymentDeletedV1Handler(handler)
	return dispatcher
}

// 员工完成离职
//
// - 员工完成离职，即离职日期的次日凌晨时，员工雇佣状态更改为“离职”后触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/events/resigned
func (dispatcher *EventDispatcher) OnP2EmploymentResignedV1(handler func(ctx context.Context, event *larkcorehr.P2EmploymentResignedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.employment.resigned_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.employment.resigned_v1")
	}
	dispatcher.eventType2EventHandler["corehr.employment.resigned_v1"] = larkcorehr.NewP2EmploymentResignedV1Handler(handler)
	return dispatcher
}

// 雇佣信息更新
//
// - 员工在飞书人事的「雇佣信息被更新」时将触发此事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/employment/events/updated
func (dispatcher *EventDispatcher) OnP2EmploymentUpdatedV1(handler func(ctx context.Context, event *larkcorehr.P2EmploymentUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.employment.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.employment.updated_v1")
	}
	dispatcher.eventType2EventHandler["corehr.employment.updated_v1"] = larkcorehr.NewP2EmploymentUpdatedV1Handler(handler)
	return dispatcher
}

// 异动状态变更事件
//
// - 在异动发起审批和产生审批结果时触发该事件，审批结果产生的场景包括撤销、审批通过、审批拒绝
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_change/events/updated
func (dispatcher *EventDispatcher) OnP2JobChangeUpdatedV1(handler func(ctx context.Context, event *larkcorehr.P2JobChangeUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.job_change.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.job_change.updated_v1")
	}
	dispatcher.eventType2EventHandler["corehr.job_change.updated_v1"] = larkcorehr.NewP2JobChangeUpdatedV1Handler(handler)
	return dispatcher
}

// 员工异动
//
// - 员工在飞书人事异动完成后将触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_data/events/changed
func (dispatcher *EventDispatcher) OnP2JobDataChangedV1(handler func(ctx context.Context, event *larkcorehr.P2JobDataChangedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.job_data.changed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.job_data.changed_v1")
	}
	dispatcher.eventType2EventHandler["corehr.job_data.changed_v1"] = larkcorehr.NewP2JobDataChangedV1Handler(handler)
	return dispatcher
}

// 员工完成入职
//
// - 在「飞书人事」将待入职员工手动操作“完成入职”后，触发该事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/job_data/events/employed
func (dispatcher *EventDispatcher) OnP2JobDataEmployedV1(handler func(ctx context.Context, event *larkcorehr.P2JobDataEmployedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.job_data.employed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.job_data.employed_v1")
	}
	dispatcher.eventType2EventHandler["corehr.job_data.employed_v1"] = larkcorehr.NewP2JobDataEmployedV1Handler(handler)
	return dispatcher
}

// 离职状态变更事件
//
// - 在离职发起审批和产生审批结果时触发该事件，审批结果产生的场景包括撤销、审批通过、审批拒绝
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/offboarding/events/updated
func (dispatcher *EventDispatcher) OnP2OffboardingUpdatedV1(handler func(ctx context.Context, event *larkcorehr.P2OffboardingUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.offboarding.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.offboarding.updated_v1")
	}
	dispatcher.eventType2EventHandler["corehr.offboarding.updated_v1"] = larkcorehr.NewP2OffboardingUpdatedV1Handler(handler)
	return dispatcher
}

//
//
// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2OrgRoleAuthorizationUpdatedV1(handler func(ctx context.Context, event *larkcorehr.P2OrgRoleAuthorizationUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.org_role_authorization.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.org_role_authorization.updated_v1")
	}
	dispatcher.eventType2EventHandler["corehr.org_role_authorization.updated_v1"] = larkcorehr.NewP2OrgRoleAuthorizationUpdatedV1Handler(handler)
	return dispatcher
}

//
//
// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2PersonCreatedV1(handler func(ctx context.Context, event *larkcorehr.P2PersonCreatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.person.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.person.created_v1")
	}
	dispatcher.eventType2EventHandler["corehr.person.created_v1"] = larkcorehr.NewP2PersonCreatedV1Handler(handler)
	return dispatcher
}

//
//
// -
//
// - 事件描述文档链接:
func (dispatcher *EventDispatcher) OnP2PersonDeletedV1(handler func(ctx context.Context, event *larkcorehr.P2PersonDeletedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.person.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.person.deleted_v1")
	}
	dispatcher.eventType2EventHandler["corehr.person.deleted_v1"] = larkcorehr.NewP2PersonDeletedV1Handler(handler)
	return dispatcher
}

// 个人信息更新
//
// - 员工在飞书人事的「个人信息被更新」时将触发此事件
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/corehr-v1/person/events/updated
func (dispatcher *EventDispatcher) OnP2PersonUpdatedV1(handler func(ctx context.Context, event *larkcorehr.P2PersonUpdatedV1) error) *EventDispatcher {
	_, existed := dispatcher.eventType2EventHandler["corehr.person.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "corehr.person.updated_v1")
	}
	dispatcher.eventType2EventHandler["corehr.person.updated_v1"] = larkcorehr.NewP2PersonUpdatedV1Handler(handler)
	return dispatcher
}
