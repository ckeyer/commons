package mongo

import (
	"fmt"
	"strings"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	Database string
}

func NewMdb(host, port, database, userName, password string) (*mgo.Database, error) {
	if database == "" {
		return nil, fmt.Errorf("invalid database name.")
	}
	conf := &Config{
		Host:     host,
		Port:     port,
		UserName: userName,
		Password: password,
		Database: database,
	}
	ss, err := connect(conf)

	if err != nil {
		return nil, err
	}
	return ss.DB(database), nil
}

func NewMdbWithURL(url string) (*mgo.Database, error) {
	ss, err := mgo.Dial(url)
	if err != nil {
		return nil, err
	}

	dbName := getDbNameFromURL(url)
	if dbName == "" {
		return nil, fmt.Errorf("invalid database name.")
	}

	return ss.DB(dbName), nil
}

func getDbNameFromURL(url string) string {
	url = strings.TrimPrefix(url, "mongodb://")

	i := strings.LastIndex(url, "/")
	if i <= 0 || i >= len(url)-1 {
		return ""
	}

	dbopt := string(url[i+1:])
	optI := strings.Index(dbopt, "?")
	if optI == 0 {
		return ""
	}
	if optI > 0 {
		return string(dbopt[:optI])
	}
	return dbopt
}

func connect(c *Config) (*mgo.Session, error) {
	//连接url ： [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	url := c.Host
	if c.UserName != "" && c.Password != "" {
		url = c.UserName + ":" + c.Password + "@" + url
	}
	if c.Port != "" {
		url = url + ":" + c.Port
	}
	if c.Database != "" {
		url = url + "/" + c.Database
	}
	return mgo.Dial(url)
}

func WithC(db *mgo.Database, collection string, job func(*mgo.Collection) error) error {
	return job(db.C(collection))
}

func Upsert(db *mgo.Database, collection string, selector interface{}, change interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		_, err := c.Upsert(selector, change)
		return err
	})
}

func UpdateId(db *mgo.Database, collection string, id interface{}, change interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.UpdateId(id, change)
	})
}
func Update(db *mgo.Database, collection string, selector, change interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Update(selector, change)
	})
}
func UpdateAll(db *mgo.Database, collection string, selector, change interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		_, err := c.UpdateAll(selector, change)
		return err
	})
}

func Insert(db *mgo.Database, collection string, data ...interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Insert(data...)
	})
}

func All(db *mgo.Database, collection string, query interface{}, result interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Find(query).All(result)
	})
}

// 返回所有复合 query 条件的item， 并且被 projection 限制返回的fields
func AllSelect(db *mgo.Database, collection string, query interface{}, projection interface{}, result interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Find(query).Select(projection).All(result)
	})
}

func One(db *mgo.Database, collection string, query interface{}, result interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Find(query).One(result)
	})
}

func OneSelect(db *mgo.Database, collection string, query interface{}, projection interface{}, result interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Find(query).Select(projection).One(result)
	})
}

func FindId(db *mgo.Database, collection string, id interface{}, result interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Find(bson.M{"_id": id}).One(result)
	})
}

func RemoveId(db *mgo.Database, collection string, id interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		err := c.RemoveId(id)
		return err
	})
}
func Remove(db *mgo.Database, collection string, selector interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		err := c.Remove(selector)
		return err
	})
}
func RemoveAll(db *mgo.Database, collection string, selector interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		_, err := c.RemoveAll(selector)
		return err
	})
}

func CountId(db *mgo.Database, collection string, id interface{}) (n int) {
	WithC(db, collection, func(c *mgo.Collection) error {
		var err error
		n, err = c.FindId(id).Count()
		return err
	})
	return n
}
func Count(db *mgo.Database, collection string, query interface{}) (n int) {
	WithC(db, collection, func(c *mgo.Collection) error {
		var err error
		n, err = c.Find(query).Count()
		return err
	})
	return n
}
func Exist(db *mgo.Database, collection string, query interface{}) bool {
	return Count(db, collection, query) != 0
}
func ExistId(db *mgo.Database, collection string, id interface{}) bool {
	return CountId(db, collection, id) != 0
}

func Page(db *mgo.Database, collection string, query bson.M, offset int, limit int, result interface{}) error {
	return WithC(db, collection, func(c *mgo.Collection) error {
		return c.Find(query).Skip(offset).Limit(limit).All(result)
	})
}

//获取页面数据和“所有”符合条件的记录“总共”的条数
func PageAndCount(db *mgo.Database, collection string, query bson.M, offset int, limit int, result interface{}) (total int, err error) {
	err = WithC(db, collection, func(c *mgo.Collection) error {
		total, err = c.Find(query).Count()
		if err != nil {
			return err
		}
		return c.Find(query).Skip(offset).Limit(limit).All(result)
	})
	return total, err
}

//等同与UpdateId(db *mgo.Database, collection,id,bson.M{"$set":change})
func SetId(db *mgo.Database, collection string, id interface{}, change interface{}) error {
	return UpdateId(db, collection, id, bson.M{"$set": change})
}

func Exec(db *mgo.Database, callback func(*mgo.Database) error) error {
	var (
		ss  = db.Session.Clone()
		ndb = ss.DB(db.Name)
	)
	defer ss.Close()
	return callback(ndb)
}
