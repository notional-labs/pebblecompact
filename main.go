package main

import (
	"github.com/cockroachdb/pebble"
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

	opts := &pebble.Options{}
	opts.EnsureDefaults()

	db, err := pebble.Open(dbPath, opts)
	if err != nil {
		panic(err)
	}

	defer func() {
		db.Close()
	}()

	iter := db.NewIter(nil)

	iter.First()
	start := cp(iter.Key())

	iter.Last()
	end := cp(iter.Key())

	if err := iter.Close(); err != nil {
		panic(err)
	}

	err = db.Compact(start, end, false)
	if err != nil {
		panic(err)
	}
}
