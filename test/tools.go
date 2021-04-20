package test

import (
	"database/sql"
	"io/ioutil"
	"strings"

	"github.com/google/uuid"
	_ "github.com/lib/pq" // Driver for sql
)

func MigrateTestData(mode string, dbPath string, filesPath string) { //TODO: Add abs paths to this func
	if mode != "up" && mode != "down" {
		panic("")
	}

	db, err := sql.Open("postgres", dbPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	files, err := ioutil.ReadDir(filesPath)
	if err != nil {
		panic(err)
	}

	modeFiles := []string{}
	for _, f := range files {
		n := f.Name()
		names := strings.Split(n, ".")
		if names[len(names)-2] == mode {
			modeFiles = append(modeFiles, n)
		}
	}

	for _, f := range modeFiles {
		file, err := ioutil.ReadFile(filesPath + "/" + f)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(string(file))
		if err != nil {
			panic(err)
		}
	}
}

func RenewTestData(dbPath string, filesPath string) {
	MigrateTestData("down", dbPath, filesPath)
	MigrateTestData("up", dbPath, filesPath)
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
