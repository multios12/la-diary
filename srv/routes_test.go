package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	name, err := ioutil.TempDir("data", "unit")
	dataDir = name
	if err != nil {
		return
	}
	var lines []lineModel
	lines = append(lines, lineModel{Day: "2022-02-01", Memo: "test"})
	lines = append(lines, lineModel{Day: "2022-02-02", Memo: "test2"})
	err = writeMonthFile("202202", lines)
	if err != nil {
		return
	}

	exitVal := m.Run()

	os.RemoveAll(dataDir)
	os.Exit(exitVal)
}

// ----------------------------------------------------------------------------
func TestRouterInitial(t *testing.T) {
	routerInitial(static)
}

func TestStatic(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest(http.MethodGet, "/dummy.txt", nil)
	getStatic(c)
}

func TestGetDay(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/api/2022/02/01", nil)
	c.Params = gin.Params{gin.Param{Key: "year", Value: "2022"}, gin.Param{Key: "month", Value: "02"}, gin.Param{Key: "day", Value: "02"}}
	getDay(c)
	if c.Writer.Status() != http.StatusOK {
		t.Error()
	}

	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/api/2022/12/01", nil)
	c.Params = gin.Params{gin.Param{Key: "year", Value: "2022"}, gin.Param{Key: "month", Value: "12"}, gin.Param{Key: "day", Value: "01"}}
	getDay(c)
	if c.Writer.Status() != http.StatusNotFound {
		t.Error()
	}
}

func TestGetMonth(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/api/2022/02/02", nil)
	c.Params = gin.Params{gin.Param{Key: "year", Value: "2022"}, gin.Param{Key: "month", Value: "02"}}
	getMonth(c)
	if c.Writer.Status() != http.StatusOK {
		t.Error()
	}
}

func TestPostDay(t *testing.T) {
	c, _ := gin.CreateTestContext(httptest.NewRecorder())
	b := bytes.NewBufferString(`{"Day":"2021-02-03","Memo":"test"}`)
	c.Request = httptest.NewRequest("GET", "/api/2022/02/03", b)
	c.Params = gin.Params{gin.Param{Key: "year", Value: "2022"}, gin.Param{Key: "month", Value: "02"}, gin.Param{Key: "day", Value: "03"}}
	postDay(c)
	if c.Writer.Status() != http.StatusOK {
		t.Errorf("status:%d", c.Writer.Status())
	}

	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	b = bytes.NewBufferString(`{"Day":"2021-02-03","Memo":""}`)
	c.Request = httptest.NewRequest("GET", "/api/2022/02/03", b)
	c.Params = gin.Params{gin.Param{Key: "year", Value: "2022"}, gin.Param{Key: "month", Value: "02"}, gin.Param{Key: "day", Value: "03"}}
	postDay(c)
	if c.Writer.Status() != http.StatusBadRequest {
		t.Errorf("status:%d", c.Writer.Status())
	}

	c, _ = gin.CreateTestContext(httptest.NewRecorder())
	c.Request = httptest.NewRequest("GET", "/api/2022/02/03", nil)
	c.Params = gin.Params{gin.Param{Key: "year", Value: "2022"}, gin.Param{Key: "month", Value: "02"}, gin.Param{Key: "day", Value: "03"}}
	postDay(c)
	if c.Writer.Status() != http.StatusBadRequest {
		t.Errorf("status:%d", c.Writer.Status())
	}
}

// ----------------------------------------------------------------------------

func TestReadLine(t *testing.T) {
	l := readLine("2022-02-02")
	if l.Day != "2022-02-02" || l.Memo != "test2" {
		t.Error(l)
	}
}

func TestWriteLine(t *testing.T) {
	err := writeLine("2022-03-01", "test3")
	if err != nil {
		t.Error(err)
	}
}

func TestReadMonthFile(t *testing.T) {
	lines := readMonthFile("202202")
	if len(lines) != 2 {
		t.Error(lines)
	}
}
func TestWriteMonthFile(t *testing.T) {

	var lines []lineModel
	lines = append(lines, lineModel{Day: "2022-01-01", Memo: "test"})
	lines = append(lines, lineModel{Day: "2022-01-02", Memo: "test2"})
	err := writeMonthFile("202201", lines)
	if err != nil {
		t.Error(err)
	}
}
