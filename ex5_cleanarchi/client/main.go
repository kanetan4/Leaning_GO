package main

import (
	"context"
	"fmt"
	user "ex5_cleanarchi/kitex_gen/user/userservice"
	gen "ex5_cleanarchi/kitex_gen/user"
	"github.com/cloudwego/kitex/client"
)

func main() {
	cli, err := user.NewClient("user-service", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	// req := &gen.CreateUserRequest{
	// 	Name:  "Kanyuuuue",
	// 	Email: "kane@kitex.dev",
	// }

	// resp, err := cli.CreateUser(context.Background(), req)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("✅ User created:", resp)
	
	req := &gen.GetUserRequest{
		Id: 5,
	}

	resp, err := cli.Getuser(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ User gotten:", resp)
}
