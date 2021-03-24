package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"strconv"
	"time"
)

/*struct定义*/

// gorm.Model 的定义
type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

//用户
type User struct {
	UserOpenid        string    `json:"useropenid"  gorm:"column:USER_OPEN_ID"`
	UserName          string    `json:"username"  gorm:"column:USER_NAME"`
	UserAvatarUrl     string    `json:"useravatarurl"  gorm:"column:USER_AVATAR_URL"`
	UserCreatedTime   time.Time `gorm:"column:USER_CREATED_TIME"`
	UserState         int       `gorm:"column:USER_STATE"`
	UserBanTime       time.Time `gorm:"column:USER_BAN_TIME"`
	UserPrivilege     int       `gorm:"column:USER_PRIVILEGE"`
	UserZone          string    `gorm:"column:USER_ZONE"`
	UserZoneNumber    int       `gorm:"column:USER_ZONE_NUMBER"`
	UserFileNumber    int       `gorm:"column:USER_FILE_NUMBER"`
	UserContentNumber int       `gorm:"column:USER_CONTENT_NUMBER"`
	UserCommentNumber int       `gorm:"column:USER_COMMENT_NUMBER"`
	UserManager       string    `gorm:"column:USER_MANAGER"`
}

//动态内容
type Content struct {
	ContentID              string    `json:"contentid" gorm:"column:CONTENT_ID"`
	ContentCreatedBy       string    `json:"contentcreatedby" gorm:"column:CONTENT_CREATED_BY"`
	ContentCreatedTime     time.Time `json:"contentcreatedtime" gorm:"column:CONTENT_CREATED_TIME"`
	ContentCreatedTimeUnix int64     `json:"contentcreatedtimeunix" gorm:"column:CONTENT_CREATED_TIME_UNIX"`
	ContentText            string    `json:"contenttext" gorm:"column:CONTENT_TEXT"`
	ContentShare           string    `json:"contentshare" gorm:"column:CONTENT_SHARE"`
	ContentShareNumber     int       `json:"contentsharenumber" gorm:"column:CONTENT_SHARE_NUMBER"`
	ContentUpdatedTime     time.Time `json:"contentupdatedtime" gorm:"column:CONTENT_UPDATED_TIME"`
	ContentUpdatedTimeUnix int64     `json:"contentupdatedtimeunix" gorm:"column:CONTENT_UPDATED_TIME_UNIX"`
	ContentState           int       `json:"contentstate" gorm:"column:CONTENT_STATE"`
	ContentShareType       int       `json:"contentsharetype" gorm:"column:CONTENT_SHARE_TYPE"`
	ContentLikes           int       `json:"contentlikes" gorm:"column:CONTENT_LIKES"`
	ContentComments        int       `json:"contentcomments" gorm:"column:CONTENT_COMMENTS"`
	ContentIsHot           int       `json:"contentishot" gorm:"column:CONTENT_IS_HOT"`
}

//用于发送到前端的动态信息
type ContentWithUser struct {
	Content
	UserName      string `json:"username" gorm:"column:USER_NAME"`
	UserAvatarUrl string `json:"useravatarurl" gorm:"column:USER_AVATAR_URL"`
}

//小组
type Group struct {
	GroupID          string    `json:"groupid"  gorm:"column:GROUP_ID"`
	GroupName        string    `json:"groupname"  gorm:"column:GROUP_NAME"`
	GroupRemark      string    `json:"groupremark"  gorm:"column:GROUP_REMARK"`
	GroupNumber      int       `json:"groupnumber"  gorm:"column:GROUP_NUMBER"`
	GroupCreatedBy   string    `json:"groupcreatedby"  gorm:"column:GROUP_CREATED_BY"`
	GroupCreatedTime time.Time `json:"groupcreatedtime"  gorm:"column:GROUP_CREATED_TIME"`
	GroupState       int       `json:"groupstate"  gorm:"column:GROUP_STATE"`
}

//用户加入小组
type UserGroup struct {
	ID      int       `json:"id"  gorm:"column:ID"`
	GroupID string    `json:"groupid"  gorm:"column:GROUP_ID"`
	UserID  string    `json:"userid"  gorm:"column:USER_ID"`
	State   int       `json:"state"  gorm:"column:STATE"`
	InTime  time.Time `json:"intime"  gorm:"column:IN_TIME"`
	OutTime time.Time `json:"outtime"  gorm:"column:OUT_TIME"`
}

//返回的group列表
type GroupList struct {
	Group
	IsInGroup int `json:isingroup"`
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

/* user*/
// insert user
func (user *User) Insert() error {
	return DB.Create(user).Error
}

// update user
func (user *User) Update() error {
	return DB.Save(user).Error
}

//微信授权登录
func AppletsUserInfo(openid string, nickName string, avatarurl string) (*User, error) {

	var user = User{UserOpenid: openid, UserName: nickName, UserAvatarUrl: avatarurl, UserCreatedTime: time.Now()}

	err := DB.Where("USER_OPEN_ID=?", user.UserOpenid).Find(&user).Error

	if err != nil {
		log.Println(err)
		err := DB.Create(&user).Error
		if err != nil {
			log.Println(err)
		}
		return &user, err
	}

	if &user != nil {
		log.Println("用户已存在")
		return &user, err
	}

	return nil, err
}

func UserLogin(username string, useropenid string) (*User, error) {

	var user User
	err := DB.Where("USER_NAME = ? AND USER_OPEN_ID = ?", username, useropenid).Find(&user).Error

	data, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(data))
	return &user, err
}

func GetUser(openid string) error {
	var user User
	err := DB.Where("USER_OPEN_ID=?", openid).Find(&user).Error
	return err
}

func GetUserByUsername(username string) (*User, error) {
	var user User
	err := DB.First(&user, "EMAIL = ?", username).Error
	return &user, err
}

/*content*/

func CreateContent(openid string, text string, file string) (*Content, error) {
	timeUnix := time.Now().Unix()                                                           //获取时间戳
	randomnumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)         //获取六位随机数
	var contentid = "N" + strconv.FormatInt(timeUnix, 10) + strconv.Itoa(int(randomnumber)) //组成唯一ID

	var content = Content{ContentID: contentid, ContentCreatedBy: openid, ContentCreatedTime: time.Now(), ContentText: text, ContentShare: file, ContentShareNumber: 1, ContentComments: 0, ContentLikes: 0, ContentIsHot: 1, ContentShareType: 1, ContentState: 0, ContentCreatedTimeUnix: time.Now().Unix()}

	err := DB.Create(&content).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &content, err

}

func GetHotContentList() (contentwithuser []ContentWithUser) {
	var content []Content

	err := DB.Where("CONTENT_IS_HOT = ?", 1).Find(&content).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	slice := make([]ContentWithUser, len(content))
	DB.Table("CONTENT").Select("CONTENT_ID, CONTENT_CREATED_BY, CONTENT_CREATED_TIME, CONTENT_TEXT, CONTENT_SHARE, CONTENT_SHARE_NUMBER, "+
		"CONTENT_UPDATED_TIME, CONTENT_STATE, CONTENT_SHARE_TYPE, CONTENT_LIKES, CONTENT_COMMENTS, CONTENT_IS_HOT, CONTENT_CREATED_TIME_UNIX, CONTENT_UPDATED_TIME_UNIX").Where("CONTENT_IS_HOT = ?", 1).Scan(&slice)
	for i := 0; i < len(content); i++ {

		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM USER  WHERE USER_OPEN_ID = ?", content[i].ContentCreatedBy).First(&slice[i])
	}

	return slice
}

/*group*/
func CreateGroup(openid string, groupname string, groupremark string) (*Group, error) {
	timeUnix := time.Now().Unix()                                                         //获取时间戳
	randomnumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)       //获取六位随机数
	var groupid = "G" + strconv.FormatInt(timeUnix, 10) + strconv.Itoa(int(randomnumber)) //组成唯一ID

	var group = Group{GroupID: groupid, GroupCreatedBy: openid, GroupCreatedTime: time.Now(), GroupName: groupname, GroupNumber: 0, GroupRemark: groupremark, GroupState: 1}

	err := DB.Create(&group).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &group, err
}

func InGroup(openid string, groupid string) error {
	var usergroup UserGroup
	err := DB.Where("OPEN_ID = ? AND GROUP_ID = ?", openid, groupid).First(&usergroup).Error

	//查询结果为空，新建一条记录
	if err != nil {
		usergroup = UserGroup{UserID: openid, GroupID: groupid, InTime: time.Now(), State: 1}

		err = DB.Create(&usergroup).Error
		if err != nil {
			log.Println(err)
			return err
		}

	}

	if err == nil {
		DB.Model(&usergroup).Where("OPEN_ID = ? AND GROUP_ID = ? AND STATE = ?", openid, groupid, 0).Update("STATE", 1)
		return err
	}

	return err
}

func GetGroupList(openid string) (grouplist []GroupList) {
	var usergroup UserGroup
	var group []Group
	err := DB.Where("GROUP_STATE = ?", 1).Find(&group).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	slice := make([]GroupList, len(group))

	DB.Table("GROUP").Select("GROUP_ID, GROUP_NAME, GROUP_REMARK, GROUP_NUMBER, GROUP_CREATED_BY, GROUP_CREATED_TIME, GROUP_STATE").Where("GROUP_STATE = ?", 1).Scan(&slice)

	for i := 0; i < len(group); i++ {
		err = DB.Where("USER_ID = ? AND GROUP_ID = ? AND STATE = ?", openid, group[i].GroupID, 1).First(&usergroup).Error
		if err == nil {
			slice[i].IsInGroup = 1
		}

		if err != nil {
			slice[i].IsInGroup = 0
		}
	}

	return slice
}
