package test

import (
	"search-trace-server/model"
	"search-trace-server/server"
	"search-trace-server/utils"
	"testing"
)

var encryption = "search-trace-server"

func passwordEncryption(pass string) string {
	return utils.Sha256(utils.Sha256(pass) + encryption)
}
func TestPass(t *testing.T) {
	println(passwordEncryption("123456"))
}
func TestInitAdmin(t *testing.T) {
	err := server.UserCreate(&model.User{
		Name:  "admin",
		Email: "admin@gmail.com",
		Pass:  "123456",
		Role:  1,
	})
	if err != nil {
		t.Error(err)
	}
}
