package entity

type IncreasePostCommentCountMessage struct {
	PostID         string `json:"post_id"`
	IncreaseAmount int    `json:"increase_amount"`
}
type IncreaseUserPostCountMessage struct {
	UserID         string `json:"user_id"`
	IncreaseAmount int    `json:"increase_amount"`
}
