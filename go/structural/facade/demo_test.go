package facade

import (
	"testing"
)

func TestUserFacade(t *testing.T) {
	facade := &UserFacade{}
	result := facade.GetUserProfile("xujiajun", "group1")

	if (result != "username:xujiajun@groupname:group1") {
		t.Fatal("error")
	}
}
