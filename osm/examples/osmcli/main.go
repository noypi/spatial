package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/kr/pretty"
	"github.com/noypi/kv/leveldb"
	"github.com/noypi/spatial/db"
	"github.com/noypi/spatial/osm"
)

func main() {
	//parsertest()
	search()
}

func search() {
	db, err := osm.New(spatial.OptKVName{leveldb.Name},
		spatial.OptKVDir{"/d/dev/datasets/osmtmp"})
	if nil != err {
		log.Fatal(err)
	}
	defer db.Close()

	log.Println("dbpath=", db.DbPath())

	var total int
	for i := 0; i < 1 && total < 2; i++ {
		results, _ := db.Search("residential", total, 3)
		fmt.Println("total=", results.Total)

		if 0 == len(results.Hits) {
			break
		}

		total += len(results.Hits)

		for _, d := range results.Hits {
			if d.ID[0] == 'w' {
				id, _ := strconv.ParseInt(d.ID[1:], 10, 64)
				v, err := db.GetWayInfo(id)
				pretty.Println("err=", err, ", id=", id, ", v=", v)
			}
		}
	}

}

func parsertest() {

	db, err := osm.New(spatial.OptKVName{leveldb.Name},
		spatial.OptKVDir{"/d/dev/datasets/osmtmp"})
	if nil != err {
		log.Fatal(err)
	}
	defer db.Close()
	parser := db.AsParser(2200)
	parser.PreserveTmp = true
	parser.LoadTempKV()

	f, err := os.Open("/d/dev/datasets/PH.tar.bz2")
	if nil != err {
		log.Fatal(err)
	}

	parser.SkipNodes = 23368917
	parser.SkipWays = 3767187
	//parser.SkipRelations = 9327

	if err = parser.ParseFromTarBz2(f); nil != err {
		log.Fatal(err)
	}

	/*fmt.Printf("count n=%d, w=%d, r=%d\n", parser.NodeCnt, parser.WayCnt, parser.RelationCnt)
	fmt.Printf("missed w=%d, r=%d\n", parser.MissedWays, parser.MissedRelations)
	fmt.Printf("%% missed w=%.2f%%, r=%.2f%%\n", parser.MissedWaysPercent(), parser.MissedRelationsPercent())

	fmt.Println("dbpath=", parser.DbPath())
	fmt.Println("nodes count=", parser.CountNodesInTmpKV())
	fmt.Println("ways count=", parser.CountWaysInTmpKV())
	*/
}
