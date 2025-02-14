package model

type CreateTagsRequest struct {
	Tag string `json:"tag" validate:"required"`
}

type CreateTagsResponse struct {
	Tag string `json:"tag"`
}
