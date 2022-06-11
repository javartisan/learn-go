package user_test

import (
	user "learn-go/test"
	"testing"
)

func TestLogin(t *testing.T) {
	login := user.Login("zhangsan", "zhangsan")
	if !login {
		t.Error("login test failed")
	}
	login = user.Login("zhangsan", "lisi")
	if login {
		t.Error("login test failed")
	}
}

func TestLogOut(t *testing.T) {
	login := user.LogOut("zhangsan", "zhangsan")
	if !login {
		t.Error("login test failed")
	}
	login = user.LogOut("zhangsan", "lisi")
	if login {
		t.Error("login test failed")
	}
}
