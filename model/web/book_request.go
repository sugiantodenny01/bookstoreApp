package web

type AddBookRequest struct {
	Title           string `json:"Title"`
	Summary         string `json:"Summary"`
	Price           int    `json:"Price"`
	Stock           int    `json:"Stock"`
	Cover_Image     string `json:"Cover_Image_Base64"`
	Image_Extension string `json:"Image_Extension"`
}

type UpdateCoverBookRequest struct {
	Book_ID         int    `json:"id"`
	Cover_Image     string `json:"Cover_Image_Base64"`
	Image_Extension string `json:"Image_Extension"`
}

type GetBookByIdRequest struct {
	Book_ID int `json:"id"`
}
