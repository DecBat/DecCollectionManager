package models

import (
	"database/sql"
	"time"

	"github.com/DecBat/DecCollectionManager/internal/store"
)

type ItemResponse struct {
	ID      int32      `json:"id"`
	Title   string     `json:"title"`
	Price   int32      `json:"price"`
	UserID  *int32     `json:"user_id"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

type UserResponse struct {
	ID       int32      `json:"id"`
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Created  *time.Time `json:"created"`
	Updated  *time.Time `json:"updated"`
}

type CreateItemRequest struct {
	Title  string `json:"title"`
	Price  int32  `json:"price"`
	UserID *int32 `json:"user_id"`
}

func nullInt32ToPtr(n sql.NullInt32) *int32 {
	if !n.Valid {
		return nil
	}
	return &n.Int32
}

func nullTimeToPtr(n sql.NullTime) *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Time
}

func ToItemResponse(i store.Item) ItemResponse {
	return ItemResponse{
		ID:      i.ID,
		Title:   i.Title,
		Price:   i.Price,
		UserID:  nullInt32ToPtr(i.UserID),
		Created: nullTimeToPtr(i.Created),
		Updated: nullTimeToPtr(i.Updated),
	}
}

func ToItemResponses(items []store.Item) []ItemResponse {
	out := make([]ItemResponse, 0, len(items))
	for _, i := range items {
		out = append(out, ToItemResponse(i))
	}
	return out
}

func ToUserResponse(u store.GetUserRow) UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Created:  nullTimeToPtr(u.Created),
		Updated:  nullTimeToPtr(u.Updated),
	}
}
