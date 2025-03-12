package dto

import "time"

type VocabularyRequest struct {
	Vocabulary  string    `json:"vocabulary"  bson:"vocabulary"`
	Translation string    `json:"translation" bson:"translation"`
	CreatedAt   time.Time `json:"created_at"  bson:"createdAt"`
	UpdatedAt   time.Time `json:"updated_at"  bson:"updatedAt"`
}
