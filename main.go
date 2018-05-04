package main

import (
	"net/http"

	"github.com/davecgh/go-spew/spew"
)

func main() {

	app := App{}
	app.initApp()

	spew.Dump(app)

	// http.HandleFunc("/upload", TorrentUploadHandler)
	// http.HandleFunc("/stats", connectionStatsHandler)
	http.ListenAndServe(":8890", nil)
	// fmt.Println(cl.PeerID())
}

// func handleTorrentFile(cl *torrent.Client, path string) {
// 	var err error
// 	t, err = cl.AddTorrentFromFile(path)
// 	if err != nil {
// 		log.Println("Error adding torrent")
// 	}
// 	<-t.GotInfo()
// 	t.DownloadAll()
// }

// func connectionStatsHandler(w http.ResponseWriter, req *http.Request) {
// 	if req.Method == "GET" {
// 		stats := t.Stats()
// 		dat := spew.Sdump(stats)
// 		spew.Dump(dat)
// 		w.Write([]byte(dat))
// 	} else {
// 		http.Error(w, "Method not supported", http.StatusNotImplemented)
// 	}
// }

//TorrentUploadHandler handles all requests to /torrent
// func TorrentUploadHandler(w http.ResponseWriter, req *http.Request) {
// 	defer req.Body.Close()
// 	if req.Method == "POST" {
// 		file, header, err := req.FormFile("torrentFile")
// 		defer file.Close()
// 		if err != nil {
// 			log.Println(err.Error())
// 			http.Error(w, "Bad File", http.StatusBadRequest)
// 			return
// 		}
// 		data, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			log.Println(err.Error())
// 			http.Error(w, "Error with file upload", http.StatusBadRequest)
// 			return
// 		}
// 		// writePath := filepath.Join(cfg.torrentsDir, header.Filename)
// 		// ok := ioutil.WriteFile(writePath, data, os.ModePerm)
// 		// if ok != nil {
// 		// 	log.Println(err.Error())
// 		// 	http.Error(w, "Error writing file to disk", http.StatusBadRequest)
// 		// 	return
// 		// }
// 		// handleTorrentFile(cl, writePath)
// 		// w.WriteHeader(http.StatusOK)
// 		return
// 	}
// }
