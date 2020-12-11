package dao

import (
	"fmt"
	"log"
	"testing"
)

func TestCar(t *testing.T) {
	// t.Run("开始测试GetCarByUserID", testGetCarByUserID)
}

func testGetCarByUserID(t *testing.T) {
	car, err := GetCarByUserID(2)
	fmt.Println(car)
	log.Fatal(err)
}
