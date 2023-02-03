package main

import (
	"fmt"
	"log"

	"github.com/dgraph-io/badger/v3"
)

func main() {
	opts := badger.DefaultOptions("./cmd/badger/data")
	db, err := badger.Open(opts)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	// set
	err = db.Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte("answer"), []byte("42"))
		return err
	})
	if err != nil {
		panic(err)
	}
	// get
	var valNot, valCopy []byte
	err = db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("answer"))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			// This func with val would only be called if item.Value encounters no error.

			// Accessing val here is valid.
			fmt.Printf("[test]The answer is: %s\n", val)

			// Copying or parsing val is valid.
			valCopy = append([]byte{}, val...)

			// Assigning val slice to another variable is NOT OK.
			valNot = val // Do not do this.
			return nil
		})
		if err != nil {
			return err
		}
		fmt.Printf("[test]Outside The answer is: %s %s\n", valNot, valCopy)
		return nil
	})
	if err != nil {
		panic(err)
	}
	// iterate
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchSize = 10
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("[test]iterate key=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// Prefix scans
	db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		prefix := []byte("ans")
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			err := item.Value(func(v []byte) error {
				fmt.Printf("[test]Prefix scans bkey=%s, value=%s\n", k, v)
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// iterate keys
	err = db.View(func(txn *badger.Txn) error {
		opts := badger.DefaultIteratorOptions
		opts.PrefetchValues = false
		it := txn.NewIterator(opts)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			k := item.Key()
			fmt.Printf("[test]iterate key=%s\n", k)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	// delete
	err = db.Update(func(txn *badger.Txn) error {
		return txn.Delete([]byte("answer"))
	})
	if err != nil {
		panic(err)
	}
}
