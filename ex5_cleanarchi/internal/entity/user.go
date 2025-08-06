package entity

type User struct {
	Id    int64  `thrift:"id,1" frugal:"1,default,i64" json:"id"`
	Name  string `thrift:"name,2" frugal:"2,default,string" json:"name"`
	Email string `thrift:"email,3" frugal:"3,default,string" json:"email"`
}