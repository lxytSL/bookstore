package dao

import (
	"fmt"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	// t.Run("开始测试登录", testLogin)
	// t.Run("开始测试注册", testRegister)
	// t.Run("开始测试保存", testSave)

}

func testLogin(t *testing.T) {
	u, err := CheckUser("admin", "123")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}

func testRegister(t *testing.T) {
	u, err := CheckUserName("admin")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u)
}

func testSave(t *testing.T) {
	err := SaveUser("admin5", "123", "2@qq.com")
	if err != nil {
		log.Fatal(err)
	}
}
