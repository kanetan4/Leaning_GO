package entity

type Video struct {
	Id     int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Title  string `thrift:"title,2" frugal:"2,default,string" json:"title"`
	Length int64  `thrift:"length,3" frugal:"3,default,i64" json:"length"`
	Date   int64  `thrift:"date,4" frugal:"4,default,i64" json:"date"`
	Owner  string `thrift:"owner,5" frugal:"5,default,string" json:"owner"`
	Status string `thrift:"status,6" frugal:"6,default,string" json:"status"`
	Format string `thrift:"format,7" frugal:"7,default,string" json:"format"`
}

type VideoList struct {
	Videos []*Video `thrift:"videos,1" frugal:"1,default,list<Video>" json:"videos"`
	Total  int64    `thrift:"total,2" frugal:"2,default,i64" json:"total"`
}