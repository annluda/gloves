package viewmodels

import (
	statusModel "gloves/app/models/status"
	"gloves/pkg/time"
)

// StatusViewModel 内容
type StatusViewModel struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt string
}

// NewStatusViewModelSerializer 数据展示
func NewStatusViewModelSerializer(s *statusModel.Status) *StatusViewModel {
	return &StatusViewModel{
		ID:        int(s.ID),
		Content:   s.Content,
		UserID:    int(s.UserID),
		CreatedAt: time.SinceForHuman(s.CreatedAt),
	}
}
