package web

type AuthorRegisterRequest struct {
	Name     string `json:"Name"`
	Pen_Name string `json:"Pen_Name"`
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type AuthorLoginRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

type AuthorForgotPasswordRequest struct {
	Email string `json:"Email"`
}

type AuthorChangePassword struct {
	Old_Password string `json:"Old_Password"`
	New_Password string `json:"New_password"`
}

type RefreshTokenRequest struct {
	Refresh_Token string `json:"Refresh_Token"`
}

type AuthorUpdateProfileRequest struct {
	Author_ID int
	Name      string `json:"Name"`
	Pen_Name  string `json:"Pen_Name"`
}
