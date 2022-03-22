package web

type LoginResponse struct {
	Message string
	Data    interface{}
}

type BookResponsePage struct {
	Message         string
	Data            interface{}
	Pagination_Data interface{}
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

type BookResponse struct {
	Book_ID         int
	Author_ID       int
	Title           string
	Summary         string
	Stock           int
	Price           int
	Cover_URL       string
	Author_Pen_name string
}

type PageInfo struct {
	Current_Page      int
	Max_Data_Per_Page int
	Max_Page          int
	Total_All_Data    int
}

func SuccessResponse(data interface{}) LoginResponse {
	var res LoginResponse
	res.Message = "Success"
	res.Data = data
	return res
}

func SuccessBookByPage(data interface{}, pageInfo interface{}) BookResponsePage {
	var res BookResponsePage
	res.Message = "Success"
	res.Data = data
	res.Pagination_Data = pageInfo
	return res

}

func ToFailResponse(err error, information string) FailResponse {
	var res FailResponse
	res.Message = "Failed"
	res.Error_key = err.Error()
	res.Error_message = information

	return res
}
