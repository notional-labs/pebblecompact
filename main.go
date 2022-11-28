package main

import (
	"encoding/hex"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
	"os"
)

func cp(bz []byte) (ret []byte) {
	ret = make([]byte, len(bz))
	copy(ret, bz)
	return ret
}

func main() {
	if len(os.Args) != 2 {
		panic("Usage: pebble <dbPath>")
	}

	dbPath := os.Args[1]

	//opts := &pebble.Options{
	//	MaxOpenFiles: 100,
	//}
	//opts.EnsureDefaults()

	//db, err := pebble.Open(dbPath, opts)
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
	}()

	//iter := db.NewIter(nil)
	//var start, end []byte
	//
	//if iter.First() {
	//	start = cp(iter.Key())
	//}
	//
	//if iter.Last() {
	//	end = cp(iter.Key())
	//}
	//
	//if err := iter.Close(); err != nil {
	//	panic(err)
	//}

	//err = db.Compact(start, end, false)
	start_key, errDecode := hex.DecodeString("ffd1e16a90b7b05050324904fa3c05c996da4833d3b4d128bfb95d7b658e0584")
	if errDecode != nil {
		panic(errDecode)
	}
	err = db.CompactRange(util.Range{Start: start_key, Limit: nil})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Done!")
}
