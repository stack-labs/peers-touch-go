package bolt

import (
	"encoding/json"
	"fmt"

	"github.com/boltdb/bolt"
	log "github.com/stack-labs/peers-touch-go/logger"
	"github.com/stack-labs/peers-touch-go/store"
)

type boltStore struct {
	db      *bolt.DB
	name    string
	options store.Options
}

func (b *boltStore) Init(opts ...store.Option) (err error) {
	options := store.Options{}
	for _, opt := range opts {
		opt(&options)
	}

	b.options = options
	defer func() {
		if err != nil {
			log.Errorf("init bolt db err: %s ", err)
		}
	}()

	b.name = options.Database

	db, err := bolt.Open(options.Database, 0600, nil)
	if err != nil {
		err = fmt.Errorf("open bolt db file err: %s", err)
		return
	}

	if b.options.Context.Value(buckets{}) != nil {
		log.Infof("init boltdb buckets")
		buckets := b.options.Context.Value(buckets{}).([]string)
		err = db.Update(func(tx *bolt.Tx) error {
			for _, bucket := range buckets {
				log.Infof("init boltdb bucket: %s", bucket)
				if tx.Bucket([]byte(bucket)) != nil {
					log.Infof("init boltdb bucket: %s existed", bucket)
					continue
				}

				_, err := tx.CreateBucket([]byte(bucket))
				if err != nil {
					err = fmt.Errorf("gradually create bolt buckets err: %s", err)
					return err
				}

				log.Infof("init boltdb bucket: %s created", bucket)
			}

			return nil
		})

		if err != nil {
			err = fmt.Errorf("create bolt buckets err: %s", err)
			return
		}
	}

	b.db = db

	go func() {
		select {
		case <-options.Context.Done():
			err := db.Close()
			if err != nil {
				log.Errorf("bolt close err: %s", err)
			}
			return
		}
	}()

	return
}

func (b *boltStore) Options() store.Options {
	return b.options
}

func (b *boltStore) Read(key string, opts ...store.ReadOption) ([]*store.Record, error) {
	panic("implement me")
}

func (b *boltStore) Write(r *store.Record, opts ...store.WriteOption) error {
	options := store.WriteOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	return b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(options.Table))
		log.Infof("table: %s", options.Table)
		// todo codec
		v, _ := json.Marshal(r)
		err := b.Put([]byte(r.Key), v)
		if err != nil {
			log.Errorf("bolt store put data err: %s", err)
			return err
		}

		return nil
	})
}

func (b *boltStore) Tx(funcs ...func(tx store.Tx) error) error {
	return b.db.Update(func(tx *bolt.Tx) error {
		for _, f := range funcs {
			err := f(tx)
			if err != nil {
				err = tx.Rollback()
				return err
			}
		}

		return nil
	})
}

func (b *boltStore) Delete(key string, opts ...store.DeleteOption) error {
	return nil
}

func (b *boltStore) List(opts ...store.ListOption) (ret []*store.Record, err error) {
	options := store.ListOptions{}
	for _, opt := range opts {
		opt(&options)
	}

	err = b.db.View(func(tx *bolt.Tx) error {
		log.Debugf("bolt get cursor for %s %d", options.Table, options.Offset)
		c := tx.Bucket([]byte(options.Table)).Cursor()
		min := []byte(fmt.Sprintf("%08d", options.Offset))
		limit := options.Limit
		if limit == 0 {
			limit = 10
		}

		var count uint = 0
		for k, v := c.Seek(min); k != nil; k, v = c.Next() {
			if count == limit {
				break
			}
			re := &store.Record{}
			// todo codec
			_ = json.Unmarshal(v, re)
			re.Key = string(k)
			ret = append(ret, re)
			count++
		}

		return nil
	})

	return
}

func (b *boltStore) Close() error {
	return b.db.Close()
}

func (b *boltStore) String() string {
	return "bolt_" + b.options.Database
}

func NewStore(opts ...store.Option) (st store.Store, err error) {
	st = &boltStore{}
	if err = st.Init(opts...); err != nil {
		log.Errorf("create new bolt store error: %s", err)
	}

	return
}
