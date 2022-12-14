package helper

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var App *app

type app struct {
	DB *DbHandle
}

func InitApp(db *DbHandle) error {
	App = &app{
		DB: db,
	}

	return nil
}

func BodyToJsonReq(r *http.Request, data interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return errors.New(fmt.Sprintf("Body unmarshall error %s", string(body)))
	}

	defer r.Body.Close()

	return nil
}

func StrToInt64(aval string) int64 {
	aval = strings.Trim(strings.TrimSpace(aval), "\n")
	i, err := strconv.ParseInt(aval, 10, 64)
	if err != nil {
		return 0
	}
	return i
}

func NewTicker(delay, repeat time.Duration) *time.Ticker {
	ticker := time.NewTicker(repeat)
	oc := ticker.C
	nc := make(chan time.Time, 1)
	go func() {
		nc <- time.Now()
		for tm := range oc {
			nc <- tm
		}
	}()
	ticker.C = nc
	return ticker
}

func InitDb() error {
	InitConf()
	db, err := NewPgSqlxDbHandle(*ConInfo, 10)
	if err != nil {
		errors.New("create db handle error.")
		return err
	}
	err = db.Ping()
	if err != nil {
		errors.New("ping db error.")
		return err
	}

	// Create Appplication Service
	err = InitApp(db)
	if err != nil {
		errors.New("init app error.")
		return err
	}

	return nil
}
