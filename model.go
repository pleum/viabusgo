package viabusgo

type RegisterAnonymousRequest struct {
	NativeFirstName   string `json:"native_first_name"`
	EnglishLastName   string `json:"english_last_name"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	NativeMiddleName  string `json:"native_middle_name"`
	NativeLastName    string `json:"native_last_name"`
	Gender            string `json:"gender"`
	EnglishMiddleName string `json:"english_middle_name"`
	EnglishFirstName  string `json:"english_first_name"`
}

type RegisterAnonymousResponse struct {
	Auth struct {
		Status int    `json:"status"`
		Title  string `json:"title"`
		Detail string `json:"detail"`
	} `json:"auth"`
}
