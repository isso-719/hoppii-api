package user

import (
	"os"
	"testing"
)

// TestUserService_CreateUser : CreateUser の単純なテスト, login() を扱い Hoppii にログインを行う
func TestUserService_CreateUser(t *testing.T) {
	username := os.Getenv("STUDENT_ID")
	if username == "" {
		username = "99z9999"
	}
	password := os.Getenv("STUDENT_PASSWORD")
	if password == "" {
		password = "password"
	}

	userInput := &UserInput{
		StudentId: username,
		Password:  password,
	}

	userService := NewUserService()
	userOutput, err := userService.CreateUser(userInput)
	if err != nil {
		t.Error(err)
	}

	if userOutput.StudentId != username {
		t.Errorf("StudentId is not %s", username)
	}

	if userOutput.Password != password {
		t.Errorf("Password is not %s", password)
	}

	if userOutput.Cookie == nil {
		t.Errorf("Cookie is nil")
	}
}
