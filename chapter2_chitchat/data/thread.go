package data

import "time"

type Thread struct{
	Id int
	Uuid string
	Topic string
	UserId int
	CreatedAt time.Time
}


func Threads()(threads []Thread,err error){
	conv:=  Thread{}
	threads=append(threads,conv)

	return threads,err
}


