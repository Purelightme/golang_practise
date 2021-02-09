package model

type UserProfile struct {
	ID int64
	Coin int64
	IsVip int8
	UserID int64
	User User
}
