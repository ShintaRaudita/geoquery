package gcf

import (
	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/ShintaRaudita/geoquery"
)

func init() {
	functions.HTTP("Init", geoquery.PostGeoIntersects)
}
