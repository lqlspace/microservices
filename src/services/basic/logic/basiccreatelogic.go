package logic

import (
	"fmt"
)

type (
	BasicCreateRequest struct {
		Name  string   `json:"name"`
		Age   int64    `json:"age"`
		Man   bool     `json:"man"`
		Hobby []string `json:"hobby"`
	}

	BasicCreateLogic struct {
	}
)

func NewBasicCreateLogic() *BasicCreateLogic {
	return &BasicCreateLogic{}
}

func (bl *BasicCreateLogic) CreateBasicElem(req *BasicCreateRequest) error {
	fmt.Println(*req)
	return nil
}
