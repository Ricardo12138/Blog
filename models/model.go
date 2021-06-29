package models

import (
	//"database/sql"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Blog struct {
	Id      int `orm:"PK"`
	Title   string
	Content string
	Created time.Time
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(Blog))
	orm.RegisterDataBase("default", "mysql", "root:dxy971012@tcp(127.0.0.1:3306)/Blog?charset=utf8")
	orm.RunSyncdb("default", false, true)
}

func GetLink() orm.Ormer {
	o := orm.NewOrm()
	return o
}

func GetAll() (blogs []Blog) {
	db := GetLink()
	qs := db.QueryTable("blog")
	cnt, _ := qs.Count()
	var num int64
	for num = 0; num < cnt; num++ {
		var blog Blog
		blog.Id = int(num)
		err := db.Read(&blog)
		if err == orm.ErrNoRows {
			cnt++
			continue
		}
		blogs = append(blogs, blog)
	}
	return
}

func GetBlog(id int) (blog Blog) {
	db := GetLink()
	blog.Id = id
	db.Read(&blog)
	return
}

/*思路记录：
1.无id导致无输出
2.id递增
3.id加在中间？
4.id持续增加
*/
func GetId() (id int) {
	db := GetLink()
	qs := db.QueryTable("blog")
	cnt, _ := qs.Count()
	var num int64
	for num = 0; num < cnt; num++ {
		var blog Blog
		blog.Id = int(num)
		err := db.Read(&blog)
		if err == orm.ErrNoRows {
			return int(num)
		}
	}
	return int(cnt)
}

func SaveBlog(blog Blog) {
	db := GetLink()
	db.Insert(&blog)
	return
}

func UpdateBlog(blog Blog) {
	db := GetLink()
	db.Update(&blog)
	return
}

func DelBlog(blog Blog) {
	db := GetLink()
	tmp := blog.Id
	db.Delete(&blog)
	db.Raw("dbcc checkident(‘blog’,reseed,?);", tmp)
	return
}
