package response

import (
	"time"

	"github.com/google/uuid"
	ListModel "github.com/saegus/test-technique-romain-chenard/internal/modules/list/models"
)

type ListResponse struct {
	ID       uuid.UUID `json:"id"`
	Name string `json:"name"`
	CreatedAt    time.Time    `json:"createdAt"`
	UpdatedAt    time.Time    `json:"updatedAt"`

	UserId	uuid.UUID `json:"userId"`
}

func ToListResponse (list ListModel.List)ListResponse{
	return ListResponse{
		ID: list.ID,
		Name: list.Name,
		CreatedAt: list.CreatedAt,
		UpdatedAt: list.UpdatedAt,
		UserId: list.UserId,
	}
}

func ToListArrayResponse (listArray []ListModel.List) []ListResponse{
	listRes := make([]ListResponse, 0)
	for _, l := range listArray{
		listRes = append(listRes, ToListResponse(l))
	}
	return listRes
}