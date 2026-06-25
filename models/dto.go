package models

import "database/sql"
import "github.com/DecBat/DecCollectionManager/internal/store"
import "time"


type ItemResponse struct {
	ID int32 `json:"id`
	Title string `json:"title"`
	Price int32 `json:"Price"`
	UserID *int32 `json:"user_id"`
	Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`
}

type UserResponse struct {
	ID int32 `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
  Created *time.Time `json:"created"`
	Updated *time.Time `json:"updated"`	
}


func nullInt32ToPtr (n sql.NullInt32) *int32 {
	if !n.Valid {
		return nil
	}
	return &n.Int32
}

func nullTimeToPtr (n sql.NullTime) *time.Time {
	if !n.Valid {
		return nil
	}
	return &n.Time
}

func ToItemResponse (i store.Item) ItemResponse {
	return ItemResponse {
		ID: 		i.ID,
		Title: 	i.Title,
		Price: 	i.Price,
		UserID: nullInt32ToPtr(i.UserID).
		Created: nullTimeToPtr(i.Created),
		Updated: nullTimeToPtr(i.Updated),
	}
}


func ToItemResponse (items []store.item) []ItemResponse {
	out := make([]ItemResponse, 0, len(items))
	for _, i := range items {
		out = append(out, ToItemResponse(i))
	}
	return out
}

func ToUserResponse (u store.GetUserRow) UserResponse {
	return UserResponse {
		ID: 		i.ID,
		Username: 	i.Username,
		Email: 	i.Email,
		Created: nullTimeToPtr(i.Created),
		Updated: nullTimeToPtr(i.Updated),
	}
}





