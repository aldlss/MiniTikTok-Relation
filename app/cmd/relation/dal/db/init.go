package db

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"os"
)

// Driver 这个 DriverWithContext 是 interface 类型，所以不用再声明为指针了
var Driver neo4j.DriverWithContext

func Init() {
	var err error
	uri := os.Getenv("NEO4J_URL")
	auth := neo4j.BasicAuth(os.Getenv("NEO4J_USERNAME"), os.Getenv("NEO4J_PASSWORD"), "")
	Driver, err = neo4j.NewDriverWithContext(uri, auth)
	if err != nil {
		panic(err)
	}

}
