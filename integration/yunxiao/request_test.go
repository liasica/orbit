// Copyright (C) orbit. 2025-present.
//
// Created at 2025-11-04, by liasica

package yunxiao

import (
	"testing"

	"github.com/bytedance/sonic"

	"github.com/liasica/orbit/integration/yunxiao/entity"
)

var (
	testErrorResponseBytes   = []byte(`{"errorCode":"InvalidWorkitem.NotFound","errorMessage":"工作项id错误或不存在","errorMsg":"工作项id错误或不存在","success":false,"traceId":"0a032a1417622458234191409eb9a8"}`)
	testSuccessResponseBytes = []byte(`{"creator":{"name":"李耀辉","id":"644882a0daafbed6593d230c"},"modifier":{"name":"李耀辉","id":"644882a0daafbed6593d230c"},"gmtCreate":1762132293000,"gmtModified":1762223360000,"serialNumber":"DAUR-316","subject":"【勿删】测试任务","description":"","formatType":"RICHTEXT","assignedTo":{"name":"李耀辉","id":"644882a0daafbed6593d230c"},"status":{"name":"打开","nameEn":"","displayName":"打开","id":"fbddef925cc2b6f3317dbbd9e2"},"space":{"name":"极光出行","id":"3fcb02596d7ae5da2d15461da9"},"workitemType":{"name":"任务","id":"ba102e46bc6a8483d9b7f25c"},"logicalStatus":"NORMAL","customFieldValues":[{"fieldName":"代码仓库","fieldFormat":"list","values":[{"identifier":"aurservd","displayValue":"aurservd"}],"fieldId":"1e216b5770aac61d8778651c86"},{"fieldName":"审查人","fieldFormat":"multiUser","values":[{"identifier":"644882a0daafbed6593d230c","displayValue":"李耀辉"}],"fieldId":"909ce5954b576f348a9b5e9da4"},{"fieldName":"优先级","fieldFormat":"list","values":[{"identifier":"c53e54fa90560731c3cf31ddfd","displayValue":"中"}],"fieldId":"priority"}],"updateStatusAt":null,"trackers":null,"participants":null,"verifier":null,"sprint":null,"labels":null,"versions":null,"id":"6041e9a37b671d2724aec0ade2","idPath":null,"statusStageId":"1","categoryId":"Task","parentId":null}`)
)

func testUnmarshalNormalError() error {
	var resp entity.ErrorResponse
	return sonic.Unmarshal(testErrorResponseBytes, &resp)
}

func testUnmarshalNormalSuccess() error {
	var resp entity.Workitem
	return sonic.Unmarshal(testSuccessResponseBytes, &resp)
}

func testUnmarshalError() error {
	var resp entity.ErrorResponse
	return Unmarshal(testErrorResponseBytes, &resp)
}

func testUnmarshalSuccess() error {
	var resp entity.Workitem
	return Unmarshal(testSuccessResponseBytes, &resp)
}

func TestUnmarshalError(t *testing.T) {
	err := testUnmarshalError()
	t.Log(err)
}

func BenchmarkUnmarshalNormalError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testUnmarshalNormalError()
	}
}

func BenchmarkUnmarshalNormalSuccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testUnmarshalNormalSuccess()
	}
}

func BenchmarkUnmarshalError(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testUnmarshalError()
	}
}

func BenchmarkUnmarshalSuccess(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = testUnmarshalSuccess()
	}
}
