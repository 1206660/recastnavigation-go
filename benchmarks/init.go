package benchmarks

import (
	"io/ioutil"

	"github.com/fananchong/recastnavigation-go/Detour"
	"github.com/fananchong/recastnavigation-go/DetourTileCache"
	"github.com/fananchong/recastnavigation-go/tests"
)

const RAND_MAX_COUNT int = 20000000
const PATH_MAX_NODE int = 2048

var tempdata []byte
var mesh1 *detour.DtNavMesh
var mesh2 *detour.DtNavMesh
var tilecache2 *dtcache.DtTileCache

func init() {
	tempdata, _ = ioutil.ReadFile("../tests/randpos.bin")
	mesh1 = tests.LoadStaticMesh("../tests/nav_test.obj.tile.bin")
	mesh2, tilecache2 = tests.LoadDynamicMesh("../tests/nav_test.obj.tilecache.bin")
	detour.DtIgnoreUnused(tilecache2)
}