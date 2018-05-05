package authentication

import (
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/boltdb/bolt"
	"github.com/golang/protobuf/proto"
)

var (
	DB *bolt.DB
)

func Authenticate(pubK string, privK string) (*Credential, error) {
	cred := Credential{}
	x := fmt.Sprintf("%s:%s", pubK, privK)
	encoded := base64.StdEncoding.EncodeToString([]byte(x))
	err := DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("CREDENTIALS"))
		v := b.Get([]byte(encoded))
		if v == nil {
			return errors.New("invalid creds")
		}
		if err := proto.Unmarshal(v, &cred); err != nil {
			return errors.New("failed to unmarshal data")
		}
		return nil
	})
	if err != nil {
		return &Credential{}, errors.New("invalid creds")
	}

	return &cred, nil
}

func init() {
	setupDB()
}

func setupDB() (*bolt.DB, error) {
	DB, err := bolt.Open("creds.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("could not open db, %v", err)
	}
	err = DB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("CREDENTIALS"))
		if err != nil {
			return fmt.Errorf("could not create bucket: %v", err)
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("could not set up buckets, %v", err)
	}
	fmt.Println("DB Setup Done")
	return DB, nil
}
