package profile

import (
	"context"
	"errors"
	"time"

	"github.com/baturalpk/photo-bucket/ent"
	"github.com/baturalpk/photo-bucket/ent/profile"
	"github.com/baturalpk/photo-bucket/services/profile/contracts"
	"github.com/baturalpk/photo-bucket/utils/crypto"
	"github.com/golang-jwt/jwt/v4"
)

type repository struct {
	main *ent.Client
}

func NewRepository(mainDB *ent.Client) *repository {
	return &repository{
		main: mainDB,
	}
}

func (r repository) SignUp(ctx context.Context, form interface{}) error {
	p := r.main.Profile.Create()
	var innerForm contracts.SignUpForm

	switch form.(type) {
	case contracts.SignUpFormWithEmail:
		p.SetEmail(form.(contracts.SignUpFormWithEmail).Email)
		innerForm = form.(contracts.SignUpFormWithEmail).SignUpForm
	case contracts.SignUpFormWithPhone:
		p.SetPhone(form.(contracts.SignUpFormWithPhone).Phone)
		innerForm = form.(contracts.SignUpFormWithPhone).SignUpForm
	default:
		return errors.New("invalid form type")
	}

	hash, err := crypto.GenerateHashFromPassword(innerForm.Password)
	if err != nil {
		return err
	}
	p.SetPasswordHash(hash)
	p.SetUsername(innerForm.Username)

	if _, err := p.Save(ctx); err != nil {
		return err
	}
	return nil
}

func (r repository) SignIn(ctx context.Context, form interface{}) (*ent.Profile, string, error) {
	f, ok := form.(contracts.SignInForm)
	if !ok {
		return nil, "", errors.New("invalid form type")
	}

	p, err := r.main.Profile.Query().Where(profile.Or(
		profile.Email(f.Account),
		profile.Username(f.Account),
	)).Only(ctx)
	if err != nil {
		return nil, "", err
	}

	if !crypto.DoHashAndPasswordMatch(p.PasswordHash, f.Password) {
		return nil, "", errors.New("wrong username, email, and/or password")
	}

	claims := jwt.RegisteredClaims{
		Subject: p.ID.String(),
		ExpiresAt: &jwt.NumericDate{
			Time: time.Now().Add(time.Minute * time.Duration(15)),
		},
		IssuedAt: &jwt.NumericDate{
			Time: time.Now(),
		},
	}

	token, err := crypto.SignNewJWTWithClaims(claims)
	return nil, token, err
}

func (r repository) GetByUsername(ctx context.Context, username string) (*ent.Profile, error) {
	return r.main.Profile.Query().Where(profile.Username(username)).Select(
		profile.FieldID,
		profile.FieldUsername,
		profile.FieldPictureURL,
		profile.FieldName,
		profile.FieldBiography,
	).Only(ctx)
}
