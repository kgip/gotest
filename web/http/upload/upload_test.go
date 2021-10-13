package main

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/smartystreets/goconvey/convey"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSaveFile(t *testing.T) {
	convey.Convey("文件上传测试", t, func() {
		recorder := httptest.NewRecorder()
		request := httptest.NewRequest(http.MethodPost, "/test/upload", nil)
		SaveFile(recorder, request)
	})
}

func TestName(t *testing.T) {
	convey.Convey("Name测试", t, func() {
		t.Log("hello")
		assert.PanicMatches(t, func() {
			panic(errors.New("panic出错测试"))
		}, "panic出错测试")
	})
}
