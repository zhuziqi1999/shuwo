package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"time"
)

/*struct定义*/

// gorm.Model 的定义
type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	UserOpenid        string    `json:"useropenid" binding:"required" gorm:"column:USER_OPEN_ID"`
	UserName          string    `json:"username" binding:"required" gorm:"column:USER_NAME"`
	UserPhone         string    `gorm:"column:user_phone"`
	UserCreatedTime   time.Time `gorm:"column:user_created_time"`
	UserState         int       `gorm:"column:user_state"`
	UserBanTime       time.Time `gorm:"column:user_ban_time"`
	UserPrivilege     int       `gorm:"column:user_privilege"`
	UserZone          string    `gorm:"column:user_zone"`
	UserZoneNumber    int       `gorm:"column:user_zone_number"`
	UserFileNumber    int       `gorm:"column:user_file_number"`
	UserContentNumber int       `gorm:"column:user_content_number"`
	UserCommentNumber int       `gorm:"column:user_comment_number"`
	UserManager       string    `gorm:"column:user_manager"`
}

/*数据库*/
var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/shuwo?charset=utf8&parseTime=True&loc=Local")

	if err == nil {
		DB = db
		db.LogMode(true)
		db.SingularTable(true)
		//db.AutoMigrate( &User{})
		//db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
		return db, err
	}
	return nil, err
}

// user
// insert user
func (user *User) Insert() error {
	return DB.Create(user).Error
}

// update user
func (user *User) Update() error {
	return DB.Save(user).Error
}

//微信授权登录
func AppletsUserInfo(openid string, nickName string) (*User, error) {

	var user = User{UserOpenid: openid, UserName: nickName, UserCreatedTime: time.Now()}

	err := DB.Where("user_open_id=?", user.UserOpenid).Find(&user).Error

	if err != nil {
		log.Println(err)
		return nil, err
	}

	if &user != nil {
		log.Println("用户已存在")
		return &user, err
	}

	if &user == nil {
		err := DB.Create(&user).Error
		if err != nil {
			log.Println(err)
		}
		return &user, err
	}

	return nil, err
}

func UserLogin(username string, useropenid string) (*User, error) {

	var user User
	err := DB.Where("user_name = ? AND user_open_id = ?", username, useropenid).Find(&user).Error

	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	return &user, err
}

func GetUser(openid string) error {
	var user User
	err := DB.Where("user_open_id=?", openid).Find(&user).Error
	return err
}

//
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.First(&user, "email = ?", username).Error
	return &user, err
}
