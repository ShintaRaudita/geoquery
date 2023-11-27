package geoquery

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ShintaRaudita/geoquery/config"
	"github.com/ShintaRaudita/geoquery/gq"
	"github.com/ShintaRaudita/geoquery/helper"
	"github.com/aiteung/atmessage"
	"github.com/aiteung/module/model"
)

func PostGeoIntersects(w http.ResponseWriter, r *http.Request) {
	var msg model.IteungMessage
	var resp atmessage.Response
	json.NewDecoder(r.Body).Decode(&msg)
	if r.Header.Get("Secret") == config.EndpointSecret {
		resp.Response = gq.GeoIntersects(config.Mongocon, msg.Longitude, msg.Latitude)
	} else {
		resp.Response = "Secret Salah"
	}
	fmt.Fprintf(w, helper.Json(resp))
}
