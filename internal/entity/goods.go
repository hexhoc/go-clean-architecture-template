package entity

type Goods struct {
	Id    int    `json:"id" example:"1"`
	Name  string `json:"name" example:"mac book air"`
	Count int    `json:"count" example:"10"`
}
