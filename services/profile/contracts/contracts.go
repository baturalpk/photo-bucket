package contracts

type SignInForm struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpForm struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type SignUpFormWithEmail struct {
	Email string `json:"email" validate:"required"`
	SignUpForm
}

type SignUpFormWithPhone struct {
	Phone string `json:"phone" validate:"required"`
	SignUpForm
}

//type PublicProfileData struct {
//	ID         string         `json:"id"`
//	Username   string         `json:"username"`
//	PictureURL sql.NullString `json:"picture_url,omitempty"`
//	Name       string         `json:"name,omitempty"`
//	Biography  string         `json:"biography,omitempty"`
//}
//
//type SelfProfile struct {
//	Email string `json:"email,omitempty"`
//	Phone string `json:"phone,omitempty"`
//	PublicProfileData
//}
