package models
import "time"
type User struct {
	// ID bson.ObjectId `bson:"_id"`
	Name      string    `bson:"name"`
	Age       int       `bson:"age"`
	JonedAt   time.Time `bson:"joned_at"`
	Interests []string  `bson:"interests"`
}