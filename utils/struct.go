package utils

type Product struct {
	ID       string `bson:"_id" json:"_id"`
	Name     string `bson:"name" json:"name"`
	Price    int    `bson:"price" json:"price"`
	Quantity int    `bson:"quantity" json:"quantity"`
}
