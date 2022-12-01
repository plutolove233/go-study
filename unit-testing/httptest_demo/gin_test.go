/*
@Coding : utf-8
@Time : 2022/4/10 15:45
@Author : 刘浩宇
@Software: GoLand
*/
package httptest_demo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	tests := []struct{
		name string
		param string
		expect string
	}{
		{"base case", `{"Name": "shyhao"}`,"Hello shyhao"},
		{"wrong case", "","获取请求信息失败"},
	}

	r := SetupRouter()
	for _,tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(
				"POST",
				"/hello",
				strings.NewReader(tt.param),
			)

			w := httptest.NewRecorder()
			r.ServeHTTP(w,req)
			assert.Equal(t, 200, w.Code)
			// 解析并检验响应内容是否复合预期
			var resp map[string]string
			err := json.Unmarshal([]byte(w.Body.String()), &resp)
			assert.Nil(t, err)
			assert.Equal(t, tt.expect, resp["message"])
		})
	}
}