package test

import (
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ezbuy/ezorm/db"

	"gopkg.in/mgo.v2"
)

var (
	mgoInstances     []*mgo.Session
	mgoConfig        *db.MongoConfig
	mgoInstanceOnce  sync.Once
	mgoInstanceIndex uint32
)

var ErrOperaBeforeInit = errors.New("please set db.SetOnFinishInit when needed operating db in init()")

const mgoMaxSessions = 8

const defaultRefresheDuration = 3 * time.Minute

func MgoSetup(config *db.MongoConfig) {
	mgoConfig = config
	mgoInstances = db.MustNewMgoSessions(config)
	db.SetupIdleSessionRefresher(config, mgoInstances, defaultRefresheDuration)
}

func Col(col string) (*mgo.Session, *mgo.Collection) {
	return getCol("test", col)
}

func DbName() string {
	return "test"
}

func getCol(dbName, col string) (*mgo.Session, *mgo.Collection) {
	i := atomic.AddUint32(&mgoInstanceIndex, 1)
	i = i % uint32(len(mgoInstances))
	session := mgoInstances[int(i)].Clone()
	var mgoCol *mgo.Collection
	if dbName == "" {
		mgoCol = session.DB(mgoConfig.DBName).C(col)
	} else {
		mgoCol = session.DB(dbName).C(col)
	}
	return session, mgoCol
}
