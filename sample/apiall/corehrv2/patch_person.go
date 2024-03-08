// Package corehr code generated by oapi sdk gen
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

package main

import (
	"context"
	"fmt"

	"github.com/larksuite/oapi-sdk-go/v3"
	"github.com/larksuite/oapi-sdk-go/v3/core"
	"github.com/larksuite/oapi-sdk-go/v3/service/corehr/v2"
)

// PATCH /open-apis/corehr/v2/persons/:person_id
func main() {
	// 创建 Client
	client := lark.NewClient("appID", "appSecret")
	// 创建请求对象
	req := larkcorehr.NewPatchPersonReqBuilder().
		PersonId("12454646").
		ClientToken("12454646").
		NoNeedQuery(false).
		PersonInfo(larkcorehr.NewPersonInfoBuilder().
			NameList([]*larkcorehr.PersonName{larkcorehr.NewPersonNameBuilder().Build()}).
			Gender(larkcorehr.NewEnumBuilder().Build()).
			DateOfBirth("2020-01-01").
			NationalityIdV2("6862995757234914821").
			Race(larkcorehr.NewEnumBuilder().Build()).
			MaritalStatus(larkcorehr.NewEnumBuilder().Build()).
			PhoneList([]*larkcorehr.Phone{larkcorehr.NewPhoneBuilder().Build()}).
			AddressList([]*larkcorehr.Address{larkcorehr.NewAddressBuilder().Build()}).
			EmailList([]*larkcorehr.Email{larkcorehr.NewEmailBuilder().Build()}).
			WorkExperienceList([]*larkcorehr.WorkExperienceInfo{larkcorehr.NewWorkExperienceInfoBuilder().Build()}).
			EducationList([]*larkcorehr.Education{larkcorehr.NewEducationBuilder().Build()}).
			BankAccountList([]*larkcorehr.BankAccount{larkcorehr.NewBankAccountBuilder().Build()}).
			NationalIdList([]*larkcorehr.NationalId{larkcorehr.NewNationalIdBuilder().Build()}).
			DependentList([]*larkcorehr.Dependent{larkcorehr.NewDependentBuilder().Build()}).
			EmergencyContactList([]*larkcorehr.EmergencyContact{larkcorehr.NewEmergencyContactBuilder().Build()}).
			DateEnteredWorkforce("2020-10-01").
			ProfileImageId("dfysuc8x76dsfsw").
			PersonalProfile([]*larkcorehr.PersonalProfile{larkcorehr.NewPersonalProfileBuilder().Build()}).
			NativeRegion("6863326262618752111").
			HukouType(larkcorehr.NewEnumBuilder().Build()).
			HukouLocation("山东省平阴县").
			TalentId("6863326262618752123").
			CustomFields([]*larkcorehr.CustomFieldData{larkcorehr.NewCustomFieldDataBuilder().Build()}).
			BornCountryRegion("中国").
			IsDisabled(true).
			DisableCardNumber("1110000").
			IsMartyrFamily(true).
			MartyrCardNumber("1110000").
			IsOldAlone(true).
			ResidentTaxes([]*larkcorehr.ResidentTax{larkcorehr.NewResidentTaxBuilder().Build()}).
			FirstEntryTime("2021-01-02").
			LeaveTime("2022-01-02").
			Build()).
		Build()
	// 发起请求
	resp, err := client.Corehr.V2.Person.Patch(context.Background(), req)

	// 处理错误
	if err != nil {
		fmt.Println(err)
		return
	}

	// 服务端错误处理
	if !resp.Success() {
		fmt.Println(resp.Code, resp.Msg, resp.RequestId())
		return
	}

	// 业务处理
	fmt.Println(larkcore.Prettify(resp))
}
