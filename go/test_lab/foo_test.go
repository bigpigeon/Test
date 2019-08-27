package test_lab

import (
	"bytes"
	"crypto/rand"
	"encoding/gob"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

type Data []byte

func TestFoo(t *testing.T) {
	// 一些你执行测试的代码
	var data Data = make([]byte, 128)
	// defer 在Test也是有用的，并且优先级比t.Skip高
	defer func() {
		t.Log("defer func run !!")
	}()

	n, err := rand.Read(data)

	// 如果错误不为空,则标记失败并跳过该函数剩下的流程,也可以用t.FailNow()来代替
	if err != nil {
		// 标记失败
		t.Fail()
		// 跳过流程
		t.Skip(err)
	}
	// 其他打印在 === RUN   TestFoo/sub_test1 之后
	log.Print("run to here, rand.Read's n is", n)
	// t.Log的打印会在PASS: TestFoo之后，并且只有在 go test -v 才会打印
	t.Log("run to here, rand.Read's n is", n)

	var buff bytes.Buffer
	err = gob.NewEncoder(&buff).Encode(data)
	// 也可以使用一个便捷库 "github.com/stretchr/testify/assert" 来处理错误
	// 如果你想错误的同时跳过流程可以使用"github.com/stretchr/testify/require".NoError
	// "github.com/stretchr/testify/assert".NoError = if err != nil { t.Fail(); }
	// "github.com/stretchr/testify/require".NoError = if err != nil {t.Fail(); t.Skip(err); }
	assert.NoError(t, err)
	// 并发运行子测试函数,
	t.Run("sub test1", func(t *testing.T) {
		// 只要加上这个就表示并发运行
		t.Parallel()
		time.Sleep(100 * time.Millisecond)
		t.Log("t log test 1")
	})
	t.Run("sub test2", func(t *testing.T) {
		t.Parallel()
		t.Log("t log test 2")
	})
}
