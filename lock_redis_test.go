package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redismock/v8"
	"log"
	"testing"
)

// 加锁成功，并执行业务代码
func TestLockSuccess(t *testing.T) {
	ctx := context.Background()
	key := "test_key_TestLockSuccess"
	token := "some_token"

	// 创建 redismock 客户端
	db, mock := redismock.NewClientMock()

	// 创建 RedisLock 实例
	lock := New(ctx, db, key, WithToken(token))

	// 设置模拟锁获取成功的行为
	mock.ExpectEval(lockScript, []string{key}, token, lockTime.Seconds()).SetVal("OK")

	err := lock.Lock()
	if err != nil {
		t.Errorf("Lock() returned unexpected error: %v", err)
		return
	}

	defer lock.UnLock()

	// 在这里执行锁定期间的任务，确保任务可以在锁定期间正常执行
	// ...
	log.Println("Execute Business Code: start")
	fmt.Println("任务执行")
	log.Println("Execute Business Code: end")
}
