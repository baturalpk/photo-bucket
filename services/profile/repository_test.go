package profile

import (
	"context"
	"testing"

	"github.com/baturalpk/photo-bucket/clients/entclient"
	"github.com/baturalpk/photo-bucket/pkg/crypto"
	"github.com/baturalpk/photo-bucket/services/profile/contracts"
	"github.com/baturalpk/photo-bucket/test/data"
)

func beforeTest(t *testing.T) {
	crypto.LoadES256KeysIntoMemory()
	if err := entclient.InitConnection(); err != nil {
		t.Error(err)
		t.FailNow()
	}
}

func TestRepository_SignUp(t *testing.T) {
	beforeTest(t)
	repo := NewRepository(entclient.Client())

	for _, u := range data.TestUsers() {
		if u.Email != "" { // Case: Email
			if err := repo.SignUp(context.Background(), contracts.SignUpFormWithEmail{
				Email: u.Email,
				SignUpForm: contracts.SignUpForm{
					Username: u.Username,
					Password: u.Password,
				},
			}); err != nil {
				t.Error(err)
				t.Fail()
			}
		} else { // Case: Phone
			if err := repo.SignUp(context.Background(), contracts.SignUpFormWithPhone{
				Phone: u.Phone,
				SignUpForm: contracts.SignUpForm{
					Username: u.Username,
					Password: u.Password,
				},
			}); err != nil {
				t.Error(err)
				t.Fail()
			}
		}
	}
}

func TestRepository_SignIn(t *testing.T) {
	beforeTest(t)
	repo := NewRepository(entclient.Client())

	for i, u := range data.TestUsers() {
		acc := u.Email
		if acc == "" {
			acc = u.Username
		}
		p, token, err := repo.SignIn(context.Background(), contracts.SignInForm{
			Account:  acc,
			Password: u.Password,
		})
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		t.Log(i, p, token)
	}
}

func TestRepository_GetByUsername(t *testing.T) {
	beforeTest(t)
	repo := NewRepository(entclient.Client())

	for i, u := range data.TestUsers() {
		p, err := repo.GetByUsername(context.Background(), u.Username)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		t.Log(i, *p)
	}
}
