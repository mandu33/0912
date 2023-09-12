package mysql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB //连接池对象
func CloseDB() {
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = dbSQL.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库连接已关闭")
}

// Go连接Mysql
func InitDB() {
	//用户名&密码mandu 123456数据库名称db
	//用户名:密码啊@tcp(ip:端口)/数据库的名字
	dsn := "root:12345678@tcp(127.0.0.1:3306)/tiktok?charset=utf8&parseTime=true"
	//连接数据集
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}) //open不会检验用户名和密码
	if err != nil {
		fmt.Printf("dsn:%s invalid,err:%v\n", dsn, err)
		return
	}

	//err = db.AutoMigrate(&Follow{}, &Video{}, &Comment{}, &User{}, &Favorite{})
	if err != nil {
		panic(err)
	}
	dbSQL, err := db.DB()
	if err != nil {
		panic(err)
	}
	err = dbSQL.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库已成功连接")

	// DB,err := db.DB()
	// if err != nil {
	//     panic("Failed to get underlying *sql.DB")
	// 	fmt.Println("连接数据库失败")
	// }
	// fmt.Println("连接数据库成功")
	//设置数据库连接池的最大连接数
	// db.SetMaxOpenConns(100)
	// db.SetMaxIdleConns(20)
	// db.SetConnMaxLifetime(60 * time.Minute)
	// DB.Close()
	// fmt.Println("关闭数据库成功")
	// 创建user表
	// _, err = db.Exec(`CREATE TABLE user (
	//     id INT AUTO_INCREMENT PRIMARY KEY,
	//     username VARCHAR(255) NOT NULL UNIQUE,
	//     password VARCHAR(255) NOT NULL,
	//     profession VARCHAR(255),
	//     age INT
	// 	following INT
	// 	followed INT
	// )`)
	// _, err = db.Exec(`DROP TABLE IF EXISTS user`)
	//likes表
	// _, err = db.Exec(`CREATE TABLE likes (
	//     id INT AUTO_INCREMENT PRIMARY KEY,
	//     user_id VARCHAR(255) NOT NULL,
	//     video_id VARCHAR(255) NOT NULL,
	//     like_time DATETIME NOT NULL
	// )`)
	//comments表
	// _, err = db.Exec(`CREATE TABLE comments (
	//     id INT AUTO_INCREMENT PRIMARY KEY,
	//     user_id VARCHAR(255) NOT NULL,
	//     video_id VARCHAR(255) NOT NULL,
	//     message VARCHAR(255) NOT NULL,
	//     comment_time DATETIME NOT NULL
	// )`)
	//videos表
	// _, err = db.Exec(`CREATE TABLE videos (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	user_id VARCHAR(255) NOT NULL,
	// 	video_url VARCHAR(255) NOT NULL,
	// 	pic_url VARCHAR(255) NOT NULL,
	// 	theme VARCHAR(255) NOT NULL,
	// 	comment_time DATETIME NOT NULL );`)
	//follows表
	// _, err = db.Exec(`CREATE TABLE follows (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	user_id VARCHAR(255) NOT NULL,
	// 	dst_id VARCHAR(255) NOT NULL,
	// 	follow_time DATETIME NOT NULL
	// );`)
	//chats表
	// _, err = db.Exec(`CREATE TABLE chats (
	// 	id INT AUTO_INCREMENT PRIMARY KEY,
	// 	user_id VARCHAR(255) NOT NULL,
	// 	dst_id VARCHAR(255) NOT NULL,
	// 	message VARCHAR(255) NOT NULL,
	// 	send_time DATETIME NOT NULL
	// );`)
	// if err != nil {
	//     panic(err.Error())
	// }
	// fmt.Println("新建表单成功")

}

//curd增删改查
// func insertUser(){
// 	sqlStr := `insert into user(username,password,profession,age) values("mandu","123456","student",22)`//sql语句
// 	ret, err := db.Exec(sqlStr)//执行sql语句
// 	if err != nil {
// 		fmt.Printf("insert failed,err:%v\n", err)
// 		return
// 	}
// 	//如果是插入数据的操作，能够拿到插入数据的id
// 	id, err := ret.LastInsertId()
// 	if err != nil {
// 		fmt.Printf("get id failed,err:%v\n", err)
// 		return
// 	}
// 	fmt.Println("id", id)
// }
