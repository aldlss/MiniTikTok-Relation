package dal

import "github.com/aldlss/MiniTikTok-Relation/app/cmd/relation/dal/db"

// Init init dal
func Init() {
	db.Init() // neo4j init
}
