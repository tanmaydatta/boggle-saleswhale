package store

import (
	"bufio"
	"github.com/deckarep/golang-set"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"os"
)

type DictionaryStore struct {
	mapset.Set
}

func SetupDictionaryStore() DictionaryStore {
	d := DictionaryStore{mapset.NewSet()}
	file, err := os.Open(viper.GetString("dictionary_path"))
	if err != nil {
		logrus.Fatal("Error opening dictionary file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return d
}
