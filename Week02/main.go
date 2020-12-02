package main

import (
	"github.com/EndlessIdea/Go-000/Week02/handler"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			glog.Fatalf("service went wrong %+v", r)
		}
	}()

	route := gin.Default()
	route.POST("/user/:id", handler.GetUserInfo)

	if err := route.Run(":8088"); err != nil {
		glog.Fatalf("service failed to start %+v", err)
	}
}
