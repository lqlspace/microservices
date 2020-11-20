package logic

import (
	"fmt"
)

type (
	BasicListRequest struct {
		Type string `form:"type"`
	}

	BasicListResponse struct {
		Data string `json:"data"`
	}

	BasicListLogic struct {

	}
)

func NewBasicListLogic() *BasicListLogic {
	return &BasicListLogic{}
}

func (bl *BasicListLogic) GetBasicListLogic(req *BasicListRequest) (*BasicListResponse, error) {
	return &BasicListResponse{Data:fmt.Sprintf("allen: %s\n", req.Type)}, nil
}
