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
	UserOpenid	string	 		`json:"useropenid" binding:"required" gorm:"column:USER_OPEN_ID"`
	UserName	string  		`json:"username" binding:"required" gorm:"column:USER_NAME"`
	UserPhone   string			`gorm:"column:user_phone"`
	UserCreatedTime time.Time	`gorm:"column:user_created_time"`
	UserState int				`gorm:"column:user_state"`
	UserBanTime time.Time		`gorm:"column:user_ban_time"`
	UserPrivilege int			`gorm:"column:user_privilege"`
	UserZone string				`gorm:"column:user_zone"`
	UserZoneNumber int			`gorm:"column:user_zone_number"`
	UserFileNumber int			`gorm:"column:user_file_number"`
	UserContentNumber int		`gorm:"column:user_content_number"`
	UserCommentNumber int		`gorm:"column:user_comment-number"`
	UserManager string			`gorm:"column:user_manager"`
}

/*数据库*/
var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/zuccshare?charset=utf8&parseTime=True&loc=Local")

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

func UserLogin(username string,useropenid string) (*User, error){

	var user User
	err := DB.Where("user_name = ? AND user_open_id = ?", username,useropenid).Find(&user).Error

	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	return &user, err
}

func GetUser() (*User, error) {
	var user User
	err := DB.Where("user_name=?","zzq").Find(&user).Error
	return &user, err
}

//
func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.First(&user, "email = ?", username).Error
	return &user, err
}


