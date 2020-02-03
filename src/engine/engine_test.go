package engine

import (
	"github.com/gin-gonic/gin"
	"testing"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"strings"
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
	assert.Equal(t,strings.Contains(middleWareNames[0], "Logger"), true)
	assert.Equal(t,strings.Contains(middleWareNames[1], "Recovery"), true)
}

func GetFunctionName(i interface{}) string {
    return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

