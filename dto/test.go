package dto

type Test struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type TestGetReq struct {
	ID int64 `form:"id" json:"id" binding:"required"`
}

type TestCreateReq struct {
	Test
}

type TestUpdateReq struct {
	ID
	TestCreateReq
}

type TestDeleteReq struct {
	ID
}
