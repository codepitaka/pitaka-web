package engine

import (
	"github.com/gin-gonic/gin"
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"strings"
	"os"
)


func Test_New(t *testing.T) {
	engineWrapper := New()
	expectedEngineType := gin.New()
	actualEngineType := engineWrapper.engine
	assert.Equal(t, reflect.TypeOf(expectedEngineType), reflect.TypeOf(actualEngineType))
}

func Test_AddMiddlewares(t *testing.T) {
	engineWrapper := New()
	engineWrapper.AddMiddlewares()
	
	middleWareNames := []string{
		GetFunctionName(engineWrapper.engine.RouterGroup.Handlers[0]),
		GetFunctionName(engineWrapper.engine.RouterGroup.Handlers[1]),
	}
	
	assert.Equal(t,len(middleWareNames), 2)
	assert.Equal(t,strings.Contains(middleWareNames[0], "Logger"), true)
	assert.Equal(t,strings.Contains(middleWareNames[1], "Recovery"), true)
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func Test_LoadStaticFiles(t *testing.T){
	_ = os.Mkdir("./testDir", os.FileMode(0777))
	_, _ = os.Create("./testDir/testFile1.tmpl")
	defer os.RemoveAll("./testDir")
	
	expectedFile := "testDir/testFile1.tmpl"
	
	engineWrapper := New()
	engineWrapper.LoadStaticFiles("./testDir")
	r := reflect.ValueOf(engineWrapper.engine.HTMLRender)
	f := reflect.Indirect(r).FieldByName("Files")
	actualFile :=f.Index(0).String()
	
	assert.Equal(t, expectedFile, actualFile)
}
