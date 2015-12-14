package mongo

import (
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

//mongo的辅助类
type Mdb struct {
	*Config
	baseSession *mgo.Session
}

func NewMdbWithHost(Host string) *Mdb {
	return NewMdb(Host, "", "", "", "")
}

func NewMdb(Host string, Port string, Database string, UserName string, Password string) *Mdb {
	mdb := &Mdb{
		Config: &Config{
			Host:     Host,
			Port:     Port,
			UserName: UserName,
			Password: Password,
			Database: Database,
		},
	}
	mdb.connect()
	return mdb
}

func NewMdbWithConf(c *Config) (mdb *Mdb) {
	mdb = &Mdb{
		Config: c,
	}
	mdb.connect()
	return
}

func (self *Mdb) connect() {
	//连接url ： [mongodb://][user:pass@]host1[:port1][,host2[:port2],...][/database][?options]
	url := self.Host
	if self.UserName != "" && self.Password != "" {
		url = self.UserName + ":" + self.Password + "@" + url
	}
	if self.Port != "" {
		url = url + ":" + self.Port
	}
	if self.Database != "" {
		url = url + "/" + self.Database
	}
	var err error
	self.baseSession, err = mgo.Dial(url)
	if err != nil {
		panic(err)
	}
}

func (self *Mdb) Session() *mgo.Session {
	return self.baseSession.New()
}

func (self *Mdb) DB(s *mgo.Session) *mgo.Database {
	return s.DB(self.Config.Database)
}

func (self *Mdb) WithC(collection string, job func(*mgo.Collection) error) error {
	s := self.baseSession.New()
	defer s.Close()
	return job(s.DB(self.Config.Database).C(collection))
}

func (self *Mdb) Upsert(collection string, selector interface{}, change interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		_, err := c.Upsert(selector, change)
		return err
	})
}

func (self *Mdb) UpdateId(collection string, id interface{}, change interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.UpdateId(id, change)
	})
}
func (self *Mdb) Update(collection string, selector, change interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Update(selector, change)
	})
}
func (self *Mdb) UpdateAll(collection string, selector, change interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		_, err := c.UpdateAll(selector, change)
		return err
	})
}

func (self *Mdb) Insert(collection string, data ...interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Insert(data...)
	})
}

func (self *Mdb) All(collection string, query interface{}, result interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Find(query).All(result)
	})
}

// 返回所有复合 query 条件的item， 并且被 projection 限制返回的fields
func (self *Mdb) AllSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Find(query).Select(projection).All(result)
	})
}

func (self *Mdb) One(collection string, query interface{}, result interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Find(query).One(result)
	})
}

func (self *Mdb) OneSelect(collection string, query interface{}, projection interface{}, result interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Find(query).Select(projection).One(result)
	})
}

//等效于: self.One(collection,bson.M{"_id":id},result)
func (self *Mdb) FindId(collection string, id interface{}, result interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Find(bson.M{"_id": id}).One(result)
	})
}

func (self *Mdb) RemoveId(collection string, id interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		err := c.RemoveId(id)
		return err
	})
}
func (self *Mdb) Remove(collection string, selector interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		err := c.Remove(selector)
		return err
	})
}
func (self *Mdb) RemoveAll(collection string, selector interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		_, err := c.RemoveAll(selector)
		return err
	})
}

func (self *Mdb) CountId(collection string, id interface{}) (n int) {
	self.WithC(collection, func(c *mgo.Collection) error {
		var err error
		n, err = c.FindId(id).Count()
		return err
	})
	return n
}
func (self *Mdb) Count(collection string, query interface{}) (n int) {
	self.WithC(collection, func(c *mgo.Collection) error {
		var err error
		n, err = c.Find(query).Count()
		return err
	})
	return n
}
func (self *Mdb) Exist(collection string, query interface{}) bool {
	return self.Count(collection, query) != 0
}
func (self *Mdb) ExistId(collection string, id interface{}) bool {
	return self.CountId(collection, id) != 0
}

func (self *Mdb) Page(collection string, query bson.M, offset int, limit int, result interface{}) error {
	return self.WithC(collection, func(c *mgo.Collection) error {
		return c.Find(query).Skip(offset).Limit(limit).All(result)
	})
}

//获取页面数据和“所有”符合条件的记录“总共”的条数
func (self *Mdb) PageAndCount(collection string, query bson.M, offset int, limit int, result interface{}) (total int, err error) {
	err = self.WithC(collection, func(c *mgo.Collection) error {
		total, err = c.Find(query).Count()
		if err != nil {
			return err
		}
		return c.Find(query).Skip(offset).Limit(limit).All(result)
	})
	return total, err
}

//等同与UpdateId(collection,id,bson.M{"$set":change})
func (self *Mdb) SetId(collection string, id interface{}, change interface{}) error {
	return self.UpdateId(collection, id, bson.M{"$set": change})
}
