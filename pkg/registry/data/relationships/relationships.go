package relationships

type ResourceLink struct {
	ID   interface{} `json:"id" bson:"_id"`
	Link string      `json:"link" bson:"link"`
}
