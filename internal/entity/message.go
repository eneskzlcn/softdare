package entity

type IncreasePostCommentCountMessage struct {
	PostID         string `json:"post_id"`
	IncreaseAmount int    `json:"increase_amount"`
}
