package main

import (
	"log"
	"strconv"

	"github.com/kr/pretty"
	"github.com/noypi/kv/leveldb"
	"github.com/noypi/spatial/db"
	"github.com/noypi/spatial/osm"
)

func main() {

	db, err := osm.New(spatial.OptKVName{leveldb.Name},
		spatial.OptKVDir{"/d/dev/datasets/osmtmp"})
	if nil != err {
		log.Fatal(err)
	}
	/*
		f, err := os.Open("/d/dev/datasets/PH.tar.bz2")
		if nil != err {
			log.Fatal(err)
		}

			parser := db.AsParser(500000)
			parser.SkipNodes = 23300000
			parser.SkipWays = 3700000
			parser.PreserveTmp = true

			if err = parser.ParseFromTarBz2(f); nil != err {
				log.Fatal(err)
			}
	*/

	log.Println("dbpath=", db.DbPath())
	results, err := db.Search("cebu city")
	for _, d := range results.Hits {
		id, _ := strconv.ParseInt(d.ID, 10, 64)
		v, err := db.GetInfo(id)
		pretty.Println("err=", err, ", v=", v)
	}

}
