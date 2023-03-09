/*
 * @Author: Bin
 * @Date: 2023-03-05
 * @FilePath: /gpt-zmide-server/helper/default.go
 */
package helper

import (
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"time"
)

var AppName = "server"

func init() {
	type obj struct{}
	AppName = reflect.TypeOf(obj{}).PkgPath()
	AppName = AppName[:len(AppName)-7]
}

func IsRelease() bool {
	arg1 := strings.ToLower(os.Args[0])
	name := filepath.Base(arg1)
	return (strings.Index(name, AppName) == 0 &&
		strings.Index(arg1, "go-build") < 0 &&
		os.Getenv("DEBUG") != "")
}

// 生成随机字符串
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandStr(n int) string {
	seededRand := rand.New(
		rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}
