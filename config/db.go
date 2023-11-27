package config

import "github.com/ShintaRaudita/geoquery/helper"

var Mongocon = helper.SetConnection(Mongostring, "geojson")
