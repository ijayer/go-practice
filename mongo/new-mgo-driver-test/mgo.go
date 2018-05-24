/*
* 说明：新mgo驱动测试
* 作者：zhe
* 时间：2018-05-17 1:16 PM
* 更新：
 */

package main

import (
	"fmt"
	"os"
	"time"

	"qx-api/src/utils"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

const DB = "hello"

// 初始化 logrus 对象
var l = logrus.NewEntry(&logrus.Logger{
	Out: os.Stdout,
	Formatter: &logrus.TextFormatter{
		TimestampFormat: utils.TimeLayout,
		FullTimestamp:   true,
		ForceColors:     true,
	},
	Level: logrus.InfoLevel,
})

// InitDB 初始化mongodb, 并返回session
func InitDB() *mgo.Session {
	var err error
	var session *mgo.Session

	info := new(mgo.DialInfo)
	info.Database = DB
	info.Addrs = []string{"127.0.0.1:27017"}
	info.Timeout = 60 * time.Second

	session, err = mgo.DialWithInfo(info)
	if err != nil {
		l.Panicf("dial with %v failed.", info.Addrs)
	}
	// Optional. Monotonic 模式下session 的读操作开始是向其他服务器发起（且通过一个唯一的连接），
	// 只要出现了一次写操作，session 的连接就会切换至主服务器。由此可见此模式下，能够分散一些读操作
	// 到其他服务器，但是读操作不一定能够获得最新的数据。
	session.SetMode(mgo.Monotonic, true)

	return session
}

var num = 1000000

func main() {
	session := InitDB()
	bookDao := NewBookDao(session)

	start := time.Now()
	for i := 0; i < num; i++ {
		book := Book{
			Id:       bson.NewObjectId(),
			CreateAt: utils.Now(),
			DeleteAt: utils.Date(),
			Date:     utils.Date(),
			Name:     fmt.Sprintf("mgo-%v", i),
			Price:    fmt.Sprintf("%v", i+1),
			Author:   fmt.Sprintf("mgo-%v", i),
			Serial:   uuid.New().String(),
		}
		err := bookDao.Create(book)
		if err != nil {
			l.Errorf("create book doc failed. Error: %v\n", err)
		}
	}
	l.Infof("insert %d docs time cost: %v", num, time.Since(start))
}

// Book & Books 表示数据库集合
type Book struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	CreateAt string        `json:"create_at" bson:"create_at"`
	DeleteAt string        `json:"delete_at" bson:"delete_at"`
	Date     string        `json:"date"`
	Name     string        `json:"name"`
	Price    string        `json:"price" `
	Author   string        `json:"author"`
	Serial   string        `json:"serial"`
}
type Books []Book

// BookDao 表示数据库集合 book 的数据库访问对象
type BookDao struct {
	session *mgo.Session
	Name    string
	Keys    []string
}

func NewBookDao(session *mgo.Session) *BookDao {
	return &BookDao{
		session: session,
		Name:    "movies",
		Keys:    []string{"name", "serial"},
	}
}

// Create
func (b *BookDao) Create(docs interface{}) error {
	session := b.session.Copy()
	defer session.Close()
	coll := session.DB(DB).C(b.Name)

	var index mgo.Index
	if len(b.Keys) != 0 {
		index = mgo.Index{
			Key:        b.Keys, // 索引键
			Unique:     true,   // 创建唯一索引
			DropDups:   true,   // 删除重复索引
			Background: true,   // 在后台创建
			Sparse:     true,   // 不存在字段不启用索引
		}
		if err := coll.EnsureIndex(index); err != nil {
			return err
		}
	}

	return coll.Insert(docs)
}

// Delete
func (b *BookDao) Delete() {

}

// Update
func (b *BookDao) Update() {

}

// Find
func (b *BookDao) Find(m bson.M) {
	// session := b.session.Copy()
	// defer session.Close()
	// coll := session.DB(DB).C(b.Name)
}

// FindOne
func (b *BookDao) FindOne() {

}
