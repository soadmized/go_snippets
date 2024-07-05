package journal

import "time"

type Record struct {
	UserID    int32     `bson:"userId"`
	Timestamp time.Time `bson:"timestamp"`
}
