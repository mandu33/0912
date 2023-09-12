package mysql

import (
	"context"
	"fmt"
	"log"
	// "pro2/user/kitex_gen/user"
	// "pro2/user/kitex_gen/user/userservice"
)

func main() {
	InitDB()
	res := new(User)
	// usr := db.User{
	// 	UserName: "admin",
	// 	PassWord: "12345",
	// }
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
	err := db.WithContext(context.Background()).Where("user_name = ?", "admin").Limit(1).Find(&res).Error
	if err != nil {
		log.Fatal(err)
	}
	// log.Println(resp)
	fmt.Printf("res:%s", res)

}
