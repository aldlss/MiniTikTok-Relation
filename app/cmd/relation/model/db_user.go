package model

type User struct {
	Id              int64
	Name            string
	FollowCount     int64
	FollowerCount   int64
	Avatar          string
	BackgroundImage string
	Signature       string
	TotalFavorited  int64
	WorkCount       int64
	FavoriteCount   int64
}
