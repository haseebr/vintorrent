package main

import (
	"log"
	"path/filepath"

	"github.com/anacrolix/torrent"
	"github.com/anacrolix/torrent/storage"
	homedir "github.com/mitchellh/go-homedir"
)

// App Global App
type App struct {
	config *torrent.Config
	client *torrent.Client
}

// NewApp NewApp
func (app *App) initApp() {

	config, err := defaultConfig()
	if err != nil {
		log.Fatalln("Failed to initialize defaultConfig")
	}

	app.config = config

	client, err := torrent.NewClient(config)
	if err != nil {
		log.Fatalln("Failed to initialize Client")
	}

	app.client = client
}

func defaultConfig() (cfg *torrent.Config, err error) {
	hd, err := homedir.Dir()

	if err != nil {
		log.Fatal(err.Error())
	}

	baseDir := filepath.Join(hd, ".vin")
	defaultStorage := storage.NewFileByInfoHash(baseDir)

	cfg = &torrent.Config{
		ListenAddr:     "localhost:9000",
		PeerID:         "VinTorrent1234567890",
		DefaultStorage: defaultStorage,
	}

	return cfg, nil

}
