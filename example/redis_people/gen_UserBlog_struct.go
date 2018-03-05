package test

import "time"

var _ time.Time

type UserBlog struct {
	Key   string `db:"key" json:"key"`
	Value int32  `db:"value" json:"value"`
	isNew bool
}

const ()

func (p *UserBlog) GetNameSpace() string {
	return "people"
}

func (p *UserBlog) GetClassName() string {
	return "UserBlog"
}
func (p *UserBlog) GetStoreType() string {
	return "set"
}

func (p *UserBlog) GetPrimaryKey() string {
	return ""
}

func (p *UserBlog) GetIndexes() []string {
	idx := []string{}
	return idx
}

type _UserBlogMgr struct {
}

var UserBlogMgr *_UserBlogMgr

// Get_UserBlogMgr returns the orm manager in case of its name starts with lower letter
func Get_UserBlogMgr() *_UserBlogMgr { return UserBlogMgr }

func (m *_UserBlogMgr) NewUserBlog() *UserBlog {
	rval := new(UserBlog)
	return rval
}
