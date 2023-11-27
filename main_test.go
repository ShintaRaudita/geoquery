package geoquery

import (
	"fmt"
	"testing"

	"github.com/ShintaRaudita/geoquery/gq"
	"github.com/ShintaRaudita/geoquery/helper"
)

func TestUpdateGetData(t *testing.T) {
	mconn := helper.SetConnection("mongodb+srv://exuvia3021:10mongo11db@cluster10.invzrgf.mongodb.net/", "gis2")
	datagedung := gq.GeoIntersects(mconn, 107.53182998271274, -6.867870150307709)
	fmt.Println(datagedung)
}
