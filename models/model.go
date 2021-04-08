package models

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"math/rand"
	"os"
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
	ContentFrom            string    `json:"contentfrom" gorm:"column:CONTENT_FROM"`
}

//用于发送到前端的动态信息
type ContentWithUser struct {
	Content
	UserName      string `json:"username" gorm:"column:USER_NAME"`
	UserAvatarUrl string `json:"useravatarurl" gorm:"column:USER_AVATAR_URL"`
	IsLiked       int    `json:"isliked"`
	IsCollected   int    `json:"iscollected"`
	GroupName     string `json:"groupname"  gorm:"column:GROUP_NAME"`
	FileType      string `json:"filetype" gorm:"column:FILE_TYPE"`
	FileName      string `json:"filename" gorm:"column:FILE_NAME"`
}

//用户点赞表
type UserLikeContent struct {
	ID          int       `json:"id"  gorm:"column:ID"`
	ContentID   string    `json:"contentid"  gorm:"column:CONTENT_ID"`
	UserID      string    `json:"userid"  gorm:"column:USER_ID"`
	LikedUserID string    `json:"likeduserid"  gorm:"column:LIKED_USER_ID"`
	State       int       `json:"state"  gorm:"column:STATE"`
	InTime      time.Time `json:"intime"  gorm:"column:IN_TIME"`
	InTimeUnix  int64     `json:"intimeunix"  gorm:"column:IN_TIME_UNIX"`
	OutTime     time.Time `json:"outime"  gorm:"column:OUT_TIME"`
}

type LikeList struct {
	UserLikeContent
	UserName      string `json:"username" gorm:"column:USER_NAME"`
	UserAvatarUrl string `json:"useravatarurl" gorm:"column:USER_AVATAR_URL"`
	ContentText   string `json:"contenttext" gorm:"column:CONTENT_TEXT"`
}

//用户收藏表
type UserCollectContent struct {
	ID              int       `json:"id"  gorm:"column:ID"`
	ContentID       string    `json:"contentid"  gorm:"column:CONTENT_ID"`
	UserID          string    `json:"userid"  gorm:"column:USER_ID"`
	CollectedUserID string    `json:"collecteduserid"  gorm:"column:COLLECTED_USER_ID"`
	State           int       `json:"state"  gorm:"column:STATE"`
	InTime          time.Time `json:"intime"  gorm:"column:IN_TIME"`
	InTimeUnix      int64     `json:"intimeunix"  gorm:"column:IN_TIME_UNIX"`
	OutTime         time.Time `json:"outime"  gorm:"column:OUT_TIME"`
}

type CollectList struct {
	UserCollectContent
	UserName      string `json:"username" gorm:"column:USER_NAME"`
	UserAvatarUrl string `json:"useravatarurl" gorm:"column:USER_AVATAR_URL"`
	ContentText   string `json:"contenttext" gorm:"column:CONTENT_TEXT"`
}

//小组
type Group struct {
	GroupID          string    `json:"groupid"  gorm:"column:GROUP_ID"`
	GroupName        string    `json:"groupname"  gorm:"column:GROUP_NAME"`
	GroupAvatarUrl   string    `json:"groupavatarurl"  gorm:"column:GROUP_AVATAR_URL"`
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

//评论
type Comment struct {
	CommentID              string    `json:"commentid"  gorm:"column:COMMENT_ID"`
	CommentCreatedBy       string    `json:"commentcreatedby"  gorm:"column:COMMENT_CREATED_BY"`
	CommentCreatedTime     time.Time `json:"commentcreatedtime"  gorm:"column:COMMENT_CREATED_TIME"`
	CommentCreatedTimeUnix int64     `json:"commentcreatedtimeunix"  gorm:"column:COMMENT_CREATED_TIME_UNIX"`
	CommentContentID       string    `json:"commentcontentid"  gorm:"column:COMMENT_CONTENT_ID"`
	CommentUserID          string    `json:"commentuserid"  gorm:"column:COMMENT_USER_ID"`
	CommentText            string    `json:"commenttext"  gorm:"column:COMMENT_TEXT"`
	CommentBackCommentID   string    `json:"commentbackcommentid"  gorm:"column:COMMENT_BACK_COMMENT_ID"`
	CommentState           int       `json:"commentstate"  gorm:"column:COMMENT_STATE"`
	CommentBackUserID      string    `json:"commentbackuserid"  gorm:"column:COMMENT_BACK_USER_ID"`
	CommentIsBack          int       `json:"commentisback"  gorm:"column:COMMENT_IS_BACK"`
}

type CommentList struct {
	Comment
	UserName      string `json:"username" gorm:"column:USER_NAME"`
	UserAvatarUrl string `json:"useravatarurl" gorm:"column:USER_AVATAR_URL"`
	ContentText   string `json:"contenttext" gorm:"column:CONTENT_TEXT"`
}

//folder
type Folder struct {
	FolderID              string    `json:"folderid" gorm:"column:FOLDER_ID"`
	FolderName            string    `json:"foldername" gorm:"column:FOLDER_NAME"`
	FolderCreatedBy       string    `json:"foldercreatedby" gorm:"column:FOLDER_CREATED_BY"`
	FolderCreatedTime     time.Time `json:"foldercreatedtime" gorm:"column:FOLDER_CREATED_TIME"`
	FolderCreatedTimeUnix int64     `json:"foldercreatedtimeunix"  gorm:"column:FOLDER_CREATED_TIME_UNIX"`
	FolderParentID        string    `json:"folderparentid" gorm:"column:FOLDER_PARENT_ID"`
	FolderState           int       `json:"folderstate" gorm:"column:FOLDER_STATE"`
}

//file
type File struct {
	FileID              string    `json:"fileid" gorm:"column:FILE_ID"`
	FileCreatedBy       string    `json:"filecreatedby" gorm:"column:FILE_CREATED_BY"`
	FileCreatedTime     time.Time `json:"filecreatedtime" gorm:"column:FILE_CREATED_TIME"`
	FileCreatedTimeUnix int64     `json:"filecreatedtimeunix" gorm:"column:FILE_CREATED_TIME_UNIX"`
	FileName            string    `json:"filename" gorm:"column:FILE_NAME"`
	FileSize            int64     `json:"filesize" gorm:"column:FILE_SIZE"`
	FileContent         string    `json:"filecontent" gorm:"column:FILE_CONTENT"`
	FileType            string    `json:"filetype" gorm:"column:FILE_TYPE"`
	FileFolderID        string    `json:"filefolderid" gorm:"column:FILE_FOLDER_ID"`
	FileState           int       `json:"filestate" gorm:"column:FILE_STATE"`
	FileShare           int       `json:"fileshare" gorm:"column:FILE_SHARE"`
}

/*数据库*/
var DB *gorm.DB

func InitDB() (*gorm.DB, error) {

	//db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/shuwo?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", "root:Zzq19990404.@(sh-cynosdbmysql-grp-8t1kxa1e.sql.tencentcdb.com:26707)/shuwo?charset=utf8&parseTime=True&loc=Local")

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

	var user = User{UserOpenid: openid, UserName: nickName, UserAvatarUrl: avatarurl, UserCreatedTime: time.Now(), UserBanTime: time.Date(1999, 4, 4, 0, 0, 0, 0, time.Local)}

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

func CreateContent(openid string, text string, file string, groupid string) (*Content, error) {
	timeUnix := time.Now().Unix()                                                           //获取时间戳
	randomnumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)         //获取六位随机数
	var contentid = "N" + strconv.FormatInt(timeUnix, 10) + strconv.Itoa(int(randomnumber)) //组成唯一ID

	var content = Content{ContentID: contentid, ContentCreatedBy: openid, ContentCreatedTime: time.Now(), ContentText: text, ContentShare: file, ContentShareNumber: 1,
		ContentComments: 0, ContentLikes: 0, ContentIsHot: 1, ContentShareType: 1, ContentState: 0, ContentCreatedTimeUnix: time.Now().Unix(), ContentUpdatedTime: time.Date(1999, 4, 4, 0, 0, 0, 0, time.Local), ContentFrom: groupid}

	err := DB.Create(&content).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &content, err

}

func LikeContent(openid string, contentid string) error {
	var (
		userlikecontent UserLikeContent
		content         Content
	)
	err := DB.Where("CONTENT_ID = ? ", contentid).First(&content).Error
	if err != nil {
		log.Println(err)
		return err
	}
	err = DB.Where("USER_ID = ? AND CONTENT_ID = ? ", openid, contentid).First(&userlikecontent).Error
	timeUnix := time.Now().Unix()
	//查询结果为空，新建一条记录
	if err != nil {
		userlikecontent = UserLikeContent{UserID: openid, LikedUserID: content.ContentCreatedBy, ContentID: contentid, InTime: time.Now(), InTimeUnix: timeUnix, State: 1, OutTime: time.Date(1999, 4, 4, 0, 0, 0, 0, time.Local)}

		err = DB.Create(&userlikecontent).Error
		if err != nil {
			log.Println(err)
			return err
		}

	} else {
		//如果查询结果存在，且用户已不在改小组，则更新状态

		err = DB.Model(&userlikecontent).Where("USER_ID = ? AND CONTENT_ID = ? AND STATE = ?", openid, contentid, 0).Update("STATE", 1).Error
		if err != nil {
			fmt.Println(err)
			return err

		}
	}

	DB.Where("CONTENT_ID = ?", contentid).First(&content)
	err = DB.Model(&content).Where("CONTENT_ID = ?", contentid).Update("CONTENT_LIKES", content.ContentLikes+1).Error
	if err != nil {
		fmt.Println(err)
		return err

	}
	return err
}

func UnlikeContent(openid string, contentid string) error {
	var (
		userlikecontent UserLikeContent
		content         Content
	)

	err := DB.Model(&userlikecontent).Where("USER_ID = ? AND CONTENT_ID = ? AND STATE = ?", openid, contentid, 1).Update("STATE", 0).Error
	if err != nil {
		fmt.Println(err)
		return err

	}
	DB.Where("CONTENT_ID = ?", contentid).First(&content)
	err = DB.Model(&content).Where("CONTENT_ID = ?", contentid).Update("CONTENT_LIKES", content.ContentLikes-1).Error
	if err != nil {
		fmt.Println(err)
		return err

	}
	return err
}

func GetLikeList(contentid string) (likelist []LikeList) {
	var like []UserLikeContent

	err := DB.Where("CONTENT_ID = ?", contentid).Order("IN_TIME DESC").Find(&like).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	slice := make([]LikeList, len(like))
	DB.Table("user_like_content").Select("ID, CONTENT_ID, USER_ID, STATE, IN_TIME, OUT_TIME").Where("CONTENT_ID = ?", contentid).Scan(&slice)
	for i := 0; i < len(like); i++ {
		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM user WHERE USER_OPEN_ID = ?", like[i].UserID).First(&slice[i])
	}

	return slice
}

func CollectContent(openid string, contentid string) error {
	var (
		usercollectcontent UserCollectContent
		content            Content
	)
	timeUnix := time.Now().Unix()

	err := DB.Where("CONTENT_ID = ? ", contentid).First(&content).Error
	if err != nil {
		log.Println(err)
		return err
	}

	err = DB.Where("USER_ID = ? AND CONTENT_ID = ? ", openid, contentid).First(&usercollectcontent).Error
	//查询结果为空，新建一条记录
	if err != nil {
		usercollectcontent = UserCollectContent{UserID: openid, CollectedUserID: content.ContentCreatedBy, ContentID: contentid, InTime: time.Now(), InTimeUnix: timeUnix, State: 1, OutTime: time.Date(1999, 4, 4, 0, 0, 0, 0, time.Local)}

		err = DB.Create(&usercollectcontent).Error
		if err != nil {
			log.Println(err)
			return err
		}

	} else {
		//如果查询结果存在，且用户已不在改小组，则更新状态

		err = DB.Model(&usercollectcontent).Where("USER_ID = ? AND CONTENT_ID = ? AND STATE = ?", openid, contentid, 0).Update("STATE", 1).Error
		if err != nil {
			fmt.Println(err)
			return err

		}
	}

	return err
}

func UncollectContent(openid string, contentid string) error {
	var (
		usercollectcontent UserCollectContent
	)

	err := DB.Model(&usercollectcontent).Where("USER_ID = ? AND CONTENT_ID = ? AND STATE = ?", openid, contentid, 1).Update("STATE", 0).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	return err
}

func GetContentList(openid string, groupid string) (contentwithuser []ContentWithUser) {
	var content []Content
	var userlikecontent UserLikeContent
	var usercollectcontent UserCollectContent
	if groupid != "" {
		err := DB.Where(" CONTENT_FROM = ?", groupid).Order("CONTENT_CREATED_TIME DESC").Find(&content).Error
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	if groupid == "" {
		err := DB.Find(&content).Order("CONTENT_CREATED_TIME DESC").Error
		if err != nil {
			fmt.Println(err)
			return nil
		}
	}

	slice := make([]ContentWithUser, len(content))
	if groupid != "" {
		DB.Table("content").Select("CONTENT_ID, CONTENT_CREATED_BY, CONTENT_CREATED_TIME, CONTENT_TEXT, CONTENT_SHARE, CONTENT_SHARE_NUMBER, "+
			"CONTENT_UPDATED_TIME, CONTENT_STATE, CONTENT_SHARE_TYPE, CONTENT_LIKES, CONTENT_COMMENTS, CONTENT_IS_HOT, CONTENT_CREATED_TIME_UNIX, CONTENT_UPDATED_TIME_UNIX, CONTENT_FROM").Where("CONTENT_FROM = ?", groupid).Order("CONTENT_CREATED_TIME DESC").Scan(&slice)
	}

	if groupid == "" {
		DB.Table("content").Select("CONTENT_ID, CONTENT_CREATED_BY, CONTENT_CREATED_TIME, CONTENT_TEXT, CONTENT_SHARE, CONTENT_SHARE_NUMBER, " +
			"CONTENT_UPDATED_TIME, CONTENT_STATE, CONTENT_SHARE_TYPE, CONTENT_LIKES, CONTENT_COMMENTS, CONTENT_IS_HOT, CONTENT_CREATED_TIME_UNIX, CONTENT_UPDATED_TIME_UNIX, CONTENT_FROM").Order("CONTENT_CREATED_TIME DESC").Scan(&slice)
	}

	for i := 0; i < len(content); i++ {
		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM user WHERE USER_OPEN_ID = ?", slice[i].ContentCreatedBy).First(&slice[i])
		DB.Table("group").Select("GROUP_NAME").Where("GROUP_ID = ?", slice[i].ContentFrom).Scan(&slice[i])
		DB.Table("file").Select("FILE_NAME, FILE_TYPE").Where("FILE_ID = ?", slice[i].ContentShare).Scan(&slice[i])
		err := DB.Where("USER_ID = ? AND CONTENT_ID = ? AND STATE = ?", openid, slice[i].ContentID, 1).First(&userlikecontent).Error
		fmt.Println("1")
		if err == nil {
			slice[i].IsLiked = 1
			fmt.Println("2")
		}
		if err != nil {
			slice[i].IsLiked = 0
			fmt.Println("3")
		}

		err = DB.Where("USER_ID = ? AND CONTENT_ID = ? AND STATE = ?", openid, slice[i].ContentID, 1).First(&usercollectcontent).Error
		if err == nil {
			slice[i].IsCollected = 1
		}
		if err != nil {
			slice[i].IsCollected = 0
		}

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
	var group Group
	err := DB.Where("USER_ID = ? AND GROUP_ID = ? ", openid, groupid).First(&usergroup).Error

	//查询结果为空，新建一条记录
	if err != nil {
		usergroup = UserGroup{UserID: openid, GroupID: groupid, InTime: time.Now(), State: 1, OutTime: time.Date(1999, 4, 4, 0, 0, 0, 0, time.Local)}

		err = DB.Create(&usergroup).Error
		if err != nil {
			log.Println(err)
			return err
		}
		return err
	}

	//如果查询结果存在，且用户已不在改小组，则更新状态
	usergroup.State = 1
	err = DB.Model(&usergroup).Where("USER_ID = ? AND GROUP_ID = ? AND STATE = ?", openid, groupid, 0).Update("STATE", 1).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	DB.Where("GROUP_ID = ?", groupid).First(&group)
	err = DB.Model(&group).Where("GROUP_ID = ?", groupid).Update("GROUP_NUMBER", group.GroupNumber+1).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func OutGroup(openid string, groupid string) error {
	var usergroup UserGroup
	var group Group

	err := DB.Model(&usergroup).Where("USER_ID = ? AND GROUP_ID = ? AND STATE = ?", openid, groupid, 1).Update("STATE", 0).Error
	if err != nil {
		fmt.Println(err)
		return err

	}

	DB.Where("GROUP_ID = ?", groupid).First(&group)
	err = DB.Model(&group).Where("GROUP_ID = ?", groupid).Update("GROUP_NUMBER", group.GroupNumber-1).Error
	if err != nil {
		fmt.Println(err)
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

	DB.Table("group").Select("GROUP_ID, GROUP_NAME, GROUP_REMARK, GROUP_NUMBER, GROUP_CREATED_BY, GROUP_CREATED_TIME, GROUP_STATE").Where("GROUP_STATE = ?", 1).Scan(&slice)

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

/*comment*/

func CreateComment(openid string, text string, contentid string, commentbackuserid string) (*Comment, error) {
	timeUnix := time.Now().Unix()                                                           //获取时间戳
	randomnumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)         //获取六位随机数
	var commentid = "C" + strconv.FormatInt(timeUnix, 10) + strconv.Itoa(int(randomnumber)) //组成唯一ID
	var content Content
	err := DB.Where("CONTENT_ID = ?", contentid).First(&content).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var comment = Comment{CommentID: commentid, CommentCreatedBy: openid, CommentCreatedTime: time.Now(), CommentCreatedTimeUnix: time.Now().Unix(),
		CommentText: text, CommentContentID: contentid, CommentUserID: content.ContentCreatedBy, CommentBackUserID: commentbackuserid, CommentState: 1, CommentIsBack: 0}

	err = DB.Create(&comment).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}

	DB.Where("CONTENT_ID = ?", contentid).First(&content)
	err = DB.Model(&content).Where("CONTENT_ID = ?", contentid).Update("CONTENT_COMMENTS", content.ContentComments+1).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &comment, err

}

func GetCommentList(contentid string) (commentlist []CommentList) {
	var comment []Comment
	err := DB.Where("COMMENT_CONTENT_ID = ?", contentid).Find(&comment).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	slice := make([]CommentList, len(comment))
	DB.Table("comment").Select("COMMENT_ID, COMMENT_CREATED_BY, COMMENT_CREATED_TIME, COMMENT_CREATED_TIME_UNIX, COMMENT_CONTENT_ID, "+
		"COMMENT_TEXT, COMMENT_BACK_COMMENT_ID, COMMENT_STATE, COMMENT_BACK_USER_ID, COMMENT_IS_BACK").Where("COMMENT_CONTENT_ID = ?", contentid).Order("COMMENT_CREATED_TIME DESC").Scan(&slice)
	for i := 0; i < len(comment); i++ {
		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM user WHERE USER_OPEN_ID = ?", comment[i].CommentCreatedBy).First(&slice[i])
	}

	return slice
}

func GetMyGroupList(openid string) (group []Group) {
	var usergroup []UserGroup

	err := DB.Where("USER_ID = ? AND STATE = ?", openid, 1).Find(&usergroup).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}

	slice := make([]Group, len(usergroup))

	for i := 0; i < len(usergroup); i++ {
		DB.Table("group").Select("GROUP_ID, GROUP_NAME, GROUP_REMARK, GROUP_NUMBER, GROUP_CREATED_BY, GROUP_CREATED_TIME, GROUP_STATE, GROUP_AVATAR_URL").Where("GROUP_ID = ?", usergroup[i].GroupID).First(&slice[i])
	}

	return slice
}

//folder
func CreateFolder(openid string, foldername string, parentid string) (*Folder, error) {
	timeUnix := time.Now().Unix()                                                           //获取时间戳
	randomnumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)         //获取六位随机数
	var folderid = "FD" + strconv.FormatInt(timeUnix, 10) + strconv.Itoa(int(randomnumber)) //组成唯一ID

	var folder = Folder{FolderCreatedBy: openid, FolderCreatedTime: time.Now(), FolderCreatedTimeUnix: timeUnix, FolderID: folderid, FolderName: foldername, FolderParentID: parentid, FolderState: 1}

	err := DB.Create(&folder).Error
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &folder, err
}

func GetFolderList(openid string, parentid string) (folder []Folder) {

	err := DB.Where("FOLDER_CREATED_BY = ? AND FOLDER_STATE = ? AND FOLDER_PARENT_ID = ?", openid, 1, parentid).Order("FOLDER_CREATED_TIME DESC").Find(&folder).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return folder
}

func DeleteFolder(openid string, folderid string) error {
	var folder Folder

	err := DB.Model(&folder).Where("FOLDER_CREATED_BY = ? AND FOLDER_STATE = ? AND FOLDER_ID = ?", openid, 1, folderid).Update("FOLDER_STATE", 0).Error
	if err != nil {
		fmt.Println(err)
		return err

	}
	return err
}

//file
func CreateFile(openid string, filename string, filesize int64, filetype string, filepath string, foldid string) (string, error) {
	timeUnix := time.Now().Unix()                                                        //获取时间戳
	randomnumber := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)      //获取六位随机数
	var fileid = "F" + strconv.FormatInt(timeUnix, 10) + strconv.Itoa(int(randomnumber)) //组成唯一ID

	path := filepath + fileid + "." + filetype

	var file = File{FileCreatedBy: openid, FileContent: path, FileCreatedTime: time.Now(), FileCreatedTimeUnix: timeUnix, FileID: fileid, FileName: filename, FileFolderID: foldid, FileShare: 0, FileSize: filesize, FileState: 1, FileType: filetype}

	err := DB.Create(&file).Error
	if err != nil {
		log.Println(err)
		return "", err
	}
	return fileid, err
}

func GetFileList(openid string, folderid string) (file []File) {

	err := DB.Where("FILE_CREATED_BY = ? AND FILE_STATE = ? AND FILE_FOLDER_ID = ?", openid, 1, folderid).Order("FILE_CREATED_TIME DESC").Find(&file).Error

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return file
}

func DeleteFile(openid string, fileid string) error {
	var file File

	err := DB.Model(&file).Where("FILE_CREATED_BY = ? AND FILE_STATE = ? AND FILE_ID = ?", openid, 1, fileid).Update("FILE_STATE", 0).Error
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = DB.Where("FILE_CREATED_BY = ? AND FILE_STATE = ? AND FILE_ID = ?", openid, 0, fileid).First(&file).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	err = os.Remove(file.FileContent)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

//message
func GetLikeMessageList(openid string) (likelist []LikeList) {
	var like []UserLikeContent
	err := DB.Where("LIKED_USER_ID = ? AND STATE = ?", openid, 1).Find(&like).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	slice := make([]LikeList, len(like))
	DB.Table("user_like_content").Select("ID, CONTENT_ID, USER_ID, STATE, IN_TIME, OUT_TIME, LIKED_USER_ID, IN_TIME_UNIX").Where("USER_ID = ? AND STATE = ?", openid, 1).Order("IN_TIME DESC").Scan(&slice)
	for i := 0; i < len(like); i++ {
		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM user WHERE USER_OPEN_ID = ?", slice[i].UserID).First(&slice[i])
		DB.Raw("SELECT CONTENT_TEXT FROM content WHERE CONTENT_ID = ?", slice[i].ContentID).First(&slice[i])

	}

	return slice
}

func GetCommentMessageList(openid string) (commentlist []CommentList) {
	var comment []Comment
	err := DB.Where("COMMENT_USER_ID = ?", openid).Find(&comment).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	slice := make([]CommentList, len(comment))
	DB.Table("comment").Select("COMMENT_ID, COMMENT_CREATED_BY, COMMENT_CREATED_TIME, COMMENT_CREATED_TIME_UNIX, COMMENT_CONTENT_ID, COMMENT_USER_ID, COMMENT_TEXT, COMMENT_BACK_COMMENT_ID, COMMENT_STATE, COMMENT_BACK_USER_ID, COMMENT_IS_BACK").Order("COMMENT_CREATED_TIME DESC").Where("COMMENT_USER_ID = ? AND COMMENT_STATE = ?", openid, 1).Scan(&slice)
	for i := 0; i < len(comment); i++ {
		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM user WHERE USER_OPEN_ID = ?", slice[i].CommentCreatedBy).First(&slice[i])
		DB.Raw("SELECT CONTENT_TEXT FROM content WHERE CONTENT_ID = ?", slice[i].CommentContentID).First(&slice[i])

	}

	return slice
}

func GetCollectMessageList(openid string) (collectlist []CollectList) {
	var collect []UserCollectContent
	err := DB.Where("COLLECTED_USER_ID = ?", openid).Find(&collect).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	slice := make([]CollectList, len(collect))
	DB.Table("user_collect_content").Select("ID, CONTENT_ID, USER_ID, STATE, IN_TIME, OUT_TIME, COLLECTED_USER_ID, IN_TIME_UNIX").Order("IN_TIME DESC").Where("USER_ID = ? AND STATE = ?", openid, 1).Scan(&slice)
	for i := 0; i < len(collect); i++ {
		DB.Raw("SELECT USER_NAME, USER_AVATAR_URL FROM user WHERE USER_OPEN_ID = ?", slice[i].UserID).First(&slice[i])
		DB.Raw("SELECT CONTENT_TEXT FROM content WHERE CONTENT_ID = ?", slice[i].ContentID).First(&slice[i])

	}

	return slice
}
