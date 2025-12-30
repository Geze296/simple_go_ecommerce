package config

type MongoConfig struct{
	MongoURL string
	DatabaseName string
}

func LoadConfig() *MongoConfig{
	return &MongoConfig{
		MongoURL: "mongodb://127.0.0.1:27017",
		DatabaseName: "simple_ecommerce",
	}
}