package dao

import (
	"log"
	"testing"
)

func TestSession(t *testing.T) {
	// t.Run("开始测试删除Session", testDeleteSession)
}

func testDeleteSession(t *testing.T) {
	err := DeleteSession("4440a445-a74c-4cf3-7e68-98006f0c8488")
	if err != nil {
		log.Fatal(err)
	}
}
