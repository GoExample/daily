package main

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserInfo struct {
    ID     uint
    Name   string `gorm:"unique;not null"`
    Age    int    `gorm:"default:0"`
    Gender string
    Hobby  string
}

func main() {
    db, err := gorm.Open("sqlite3", "file:gorm.db?cache=shared&mode=memory")
    if err != nil {
        panic(err)
    }
    defer db.Close()

    db.AutoMigrate(&UserInfo{})

    u1 := UserInfo{Name: "张三", Gender: "男", Hobby: "足球", Age: 28}
    u2 := UserInfo{Name: "李四", Gender: "男", Hobby: "篮球", Age: 35}
    u3 := UserInfo{Name: "王五", Gender: "男", Hobby: "双色球", Age: 38}

    // 创建记录
    db.Create(&u1)
    db.Create(&u2)
    db.Create(&u3)
    printAllRecords(db)

    var u = new(UserInfo)
    db.First(u)
    fmt.Printf("type is %T, value is %v\n", u, u)

    var uu UserInfo
    db.Find(&uu, "hobby=?", "足球")
    fmt.Printf("%#v\n", uu)

    // 更新
    db.Model(&u).Update("hobby", "双色球")
    printAllRecords(db)

    // 删除
    db.Delete(&u)
    printAllRecords(db)
}

func printAllRecords(db *gorm.DB) {
    var users []UserInfo
    result := db.Find(&users)
    fmt.Printf("Have %d users\n", result.RowsAffected)
    for _, user := range users {
        fmt.Printf("ID = %d, Name = %s, Age = %d, Hobby = %s\n", user.ID, user.Name, user.Age, user.Hobby)
    }
}
