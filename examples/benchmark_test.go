package main

import (
	"github.com/jinzhu/copier"
	"testing"
	"time"
)

type User struct {
	ID int
	Mobile       string
	Level       int
	Name           string
	RegisterTime time.Time

}

var fromUser User

func init() {
	fromUser.ID = 1
	fromUser.Mobile = "130000000"
	fromUser.Level = 99
	fromUser.Name = "Tom"
	fromUser.RegisterTime = time.Now()
}


func BenchmarkMapper(b *testing.B) {

	var toUser User
	for n := 0; n < b.N; n++ {
		toUser.ID = fromUser.ID
		toUser.Name = fromUser.Name
		toUser.RegisterTime = fromUser.RegisterTime
	}
}

func BenchmarkCopier(b *testing.B) {
	var toUser User
	for n := 0; n < b.N; n++ {
		copier.Copy(&toUser, fromUser)
	}
}
