package mongo

import (
	"gopkg.in/mgo.v2"
)

type MgoIndexs map[string][]mgo.Index

func (m MgoIndexs) Setup(db *mgo.Database) error {
	return Exec(db, func(idb *mgo.Database) error {
		for col, indexes := range m {
			for _, idx := range indexes {
				if err := idb.C(col).EnsureIndex(idx); err != nil {
					return err
				}
				idb.C(col).EnsureIndexKey("_id")
			}
		}
		return nil
	})
}
