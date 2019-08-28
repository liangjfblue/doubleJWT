package service

// CommentReq commit comment
type CommentReq struct {
    TopicId     uint    `json:"topic_id"`
    TopicType   string  `json:"topic_type"`
    Content     string  `json:"content"`
    FromUid     uint    `json:"from_uid"`
}

// CommentRes respond
type CommentRes struct {

}
