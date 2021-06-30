package Module

type Comment struct {
	CommentId string `json:"comment_id"`
	CommentType int `json:"comment_type"`
	Reader string `json:"reader"`
	Sender string `json:"sender"`
	InnerData string `json:"inner_data"`
	SentTime string `json:"sent_time"`
}