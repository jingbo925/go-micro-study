package main

import (
	"fmt"
	"go_micro_service/domain/repository"
	serivce2 "go_micro_service/domain/service"
	"go_micro_service/handler"
	"go_micro_service/proto/user"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/micro/go-micro/v2"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name("go.micro.server.user"),
		micro.Version("latest"),
	)
	// 初始化服务
	srv.Init()

	// 创建数据库连接
	db, err := gorm.Open("mysql", "root:go-micro@tcp(127.0.0.1:3306)/go-micro?charset=utf8&parseTime=True&loc=Local")
	db.SingularTable(true)
	if err != nil {
		fmt.Println(err)
	}

	defer db.Close()
	// 只执行一次
	rp := repository.NewUserRepository(db)
	rp.InitTable()
	// 创建服务实例
	userDataService := serivce2.NewUserDataService(repository.NewUserRepository(db))
	// 注册handle
	err = user.RegisterUserHandler(srv.Server(),
		&handler.User{UserDateService: userDataService})
	if err != nil {
		fmt.Println(err)
	}

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
