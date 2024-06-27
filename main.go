package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bznein/deduplicarr/sonarr/types"
	"github.com/bznein/deduplicarr/sqlite"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	// v, err := mediainfo.Open("/home/Plex/Plex/MoviesITA/Inside Out (2015)/Inside Out (2015) 2160p H265 10 bit DV HDR10+ ita eng AC3 5.1 sub ita eng Licdom.mkv")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//audioTrackes := v.Parameter(mediainfo.StreamAudio, 1, "Language")
	//	fmt.Printf("%+v\n", audioTrackes)
	//
	http.HandleFunc("/import", onImport)

	if err := sqlite.EnsureDBConnection(); err != nil {
		log.Fatal(err)
	}
	defer sqlite.CloseDB()

	// TODO move somewhere lese

	tableName := "TVSeries"
	idColumn := sqlite.NewColumn("ID", "int", false, true)
	if err := sqlite.EnsureTableExists(tableName, idColumn); err != nil {
		log.Fatal(err)
	}

	if err := http.ListenAndServe(":4455", nil); err != nil {
		log.Fatal(err)
	}
}

func onImport(w http.ResponseWriter, req *http.Request) {
	parseSonarrReq(w, req)
}

func parseSonarrReq(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading request body",
			http.StatusInternalServerError)
	}

	var res types.Series

	if err := json.Unmarshal(body, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", res)
}

// TODO:
// Parse req header to determine sonarr vs radarr vs other
// Define a common interface for the return types
