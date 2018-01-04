package main

import (
	"fmt"
	"log"
	"strconv"

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
	defer db.Close()
	parser := db.AsParser(500000)
	parser.PreserveTmp = true
	parser.LoadTempKV()
	/*
		f, err := os.Open("/d/dev/datasets/PH.tar.bz2")
		if nil != err {
			log.Fatal(err)
		}

		parser.SkipNodes = 23368917
		parser.SkipWays = 3767187
		parser.SkipRelations = 9327

		if err = parser.ParseFromTarBz2(f); nil != err {
			log.Fatal(err)
		}
	*/
	log.Println("dbpath=", db.DbPath())
	results, err := db.Search("cebu city")
	for _, d := range results.Hits {
		id, _ := strconv.ParseInt(d.ID, 10, 64)
		v, err := db.GetInfo(id)
		fmt.Println("err=", err, ", v=", v)
	}

	fmt.Printf("count n=%d, w=%d, r=%d\n", parser.NodeCnt, parser.WayCnt, parser.RelationCnt)
	fmt.Printf("missed w=%d, r=%d\n", parser.MissedWays, parser.MissedRelations)
	fmt.Printf("%% missed w=%.2f%%, r=%.2f%%\n", parser.MissedWaysPercent(), parser.MissedRelationsPercent())

	fmt.Println("dbpath=", parser.DbPath())
	fmt.Println("nodes count=", parser.CountNodesInTmpKV())
	fmt.Println("ways count=", parser.CountWaysInTmpKV())
}
