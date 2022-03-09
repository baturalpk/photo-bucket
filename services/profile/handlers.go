package profile

import (
	"github.com/baturalpk/photo-bucket/ent"
	"github.com/baturalpk/photo-bucket/services/profile/contracts"
	"github.com/baturalpk/photo-bucket/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type handlers struct {
	repo *repository
}

type authProfile struct {
	Profile     ent.Profile `json:"profile"`
	TokenString string      `json:"token_string"`
}

func NewHandlers(profileRepo *repository) *handlers {
	return &handlers{
		repo: profileRepo,
	}
}

func (h handlers) signUp(c *fiber.Ctx, form interface{}) error {
	if err := c.BodyParser(form); err != nil {
		return err
	}
	if err := validator.GetValidatorInstance().Struct(form); err != nil {
		return err
	}
	return h.repo.SignUp(c.UserContext(), form)
}

func (h handlers) SignUpViaEmail(c *fiber.Ctx) error {
	form := new(contracts.SignUpFormWithEmail)
	return h.signUp(c, form)
}

func (h handlers) SignUpViaPhone(c *fiber.Ctx) error {
	form := new(contracts.SignUpFormWithPhone)
	return h.signUp(c, form)
}

func (h handlers) SignIn(c *fiber.Ctx) error {
	f := new(contracts.SignInForm)
	if err := c.BodyParser(f); err != nil {
		return err
	}
	if err := validator.GetValidatorInstance().Struct(f); err != nil {
		return err
	}
	profile, token, err := h.repo.SignIn(c.UserContext(), f)
	if err != nil {
		return err
	}
	return c.JSON(authProfile{
		Profile:     *profile,
		TokenString: token,
	})
}

func (h handlers) GetByUsername(c *fiber.Ctx) error {
	profile, err := h.repo.GetByUsername(c.UserContext(), c.Params("username"))
	if err != nil {
		return err
	}
	return c.JSON(*profile)
}
