package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Photo struct {
	ID          primitive.Binary `bson:"_id"`
	UserID      uint64           `bson:"user_id"`
	Kind        PhotoKind        `bson:"kind"`
	FileName    string           `bson:"file_name"`
	FileSize    int64            `bson:"file_size"`
	ContentType string           `bson:"content_type"`
	UploadedAt  time.Time        `bson:"uploaded_at"`
	Metadata    Metadata         `bson:"metadata,inline"`
}

type PhotoKind string

const (
	Media      PhotoKind = "media"
	Screenshot PhotoKind = "screenshot"
	Meme       PhotoKind = "meme"
)
