package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Moment struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	CatId       string             `bson:"catId,omitempty"`
	CommunityId string             `bson:"communityId,omitempty"`
	Photos      []string           `bson:"photos,omitempty"`
	Title       string             `bson:"title,omitempty"`
	Text        string             `bson:"text,omitempty"`
	UserId      string             `bson:"userId,omitempty"`
	UpdateAt    time.Time          `bson:"updateAt,omitempty" json:"updateAt,omitempty"`
	CreateAt    time.Time          `bson:"createAt,omitempty" json:"createAt,omitempty"`
}
