package user

type User struct {
	ID      int32  `bson:"_id"`
	Name    string `bson:"name"`
	Comment string `bson:"comment"`
}
