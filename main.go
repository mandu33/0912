package main

import (
	"context"
	"fmt"
	db "pro2/dal/mysql/db"
	//"log"
	// "pro2/user/kitex_gen/user"
	// "pro2/user/kitex_gen/user/userservice"
)

func main() {
	// dsn := "root:12345678@tcp(127.0.0.1:3306)/tiktok?charset=utf8&parseTime=true"
	// //连接数据集
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //open不会检验用户名和密码
	// if err != nil {
	// 	fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
	// 	return
	// }
	db.InitDB()
	res := new(db.User)
	usr := db.User{
		UserName: "admin",
		PassWord: "12345",
	}
	//db.Create(&usr)
	//err = db.WithContext(context.Background()).Where("id = ?", 1).Limit(1).Find(&res).Error
	// if err != nil {
	// 	fmt.Printf("err:%v\n", err)
	// 	return
	// }
	// req := &user.DouyinUserRegisterRequest{
	// 	Username: "admin1",
	// 	Password: "123456",
	// }
	//db.InitDB()
	//res, err := db.QueryUser(context.Background(), "admin")
	err = db.WithContext(context.Background()).Where("user_name = ?", userName).Limit(1).Find(&res).Error
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println(resp)
	fmt.Printf("res:%s", res)

}
