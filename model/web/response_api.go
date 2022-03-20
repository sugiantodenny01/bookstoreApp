package web

type LoginResponse struct {
	Message string
	Data    map[string]string
}

type FailResponse struct {
	Message       string
	Error_key     interface{}
	Error_message string
	Error_data    interface{}
}

type AuthorProfileResponse struct {
	Author_ID int
	Name      string
	Pen_Name  string
	Email     string
}

func SuccessLoginResponse(data map[string]string) LoginResponse {
	var res LoginResponse
	res.Message = "Success"
	res.Data = data
	return res
}

func ToFailResponse(err error, information string) FailResponse {
	var res FailResponse
	res.Message = "Failed"
	res.Error_key = err.Error()
	res.Error_message = information

	return res
}
