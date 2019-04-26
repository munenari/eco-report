package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/munenari/eco-report/api"
	"github.com/munenari/eco-report/config"
	"github.com/munenari/eco-report/model"
)

var (
	configFilePath = flag.String("c", "config.toml", "to load config filepath, .toml type")
)

func main() {
	flag.Parse()
	c, err := config.LoadByPath(*configFilePath)
	if err != nil {
		panic(err)
	}
	// leak?
	tick := time.NewTicker(20 * time.Second)
	defer tick.Stop()
	log.Println("agent started")
	for range tick.C {
		go func() {
			if err := execute(c); err != nil {
				log.Println("\nfailed to get ot post data:", err)
			} else {
				fmt.Print(".")
			}
		}()
	}
}

func execute(c *config.EcoReport) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("panic: %+v\n", err)
		}
	}()
	instantData, batteryData, err := getProperties(c)
	if err != nil {
		return fmt.Errorf("failed to get data from API: %s", err)
	}
	record := map[string]interface{}{
		"time":                time.Now().UTC(),
		"instantData":         instantData,
		"batteryData":         batteryData,
		"instantDataCircuits": instantData.GetCircuitsMap(),
	}
	b, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("ailed to marshal json from result map: %s", err)
	}
	record = nil
	instantData = nil
	batteryData = nil
	r := bytes.NewReader(b)
	err = postJSONData(c.ElasticOrigin, r)
	if err != nil {
		return fmt.Errorf("failed to post data to elasticsearch: %s", err)
	}
	// log.Printf("success to post data\n%s\n", b)
	b = nil
	r = nil
	return nil
}

func getProperties(c *config.EcoReport) (*model.InstantData, *map[string]interface{}, error) {
	n, err := api.GetNonce(c.Origin)
	if err != nil {
		return nil, nil, err
	}
	o, err := api.GetOneTimePassword(c.Origin, n, c.Password)
	if err != nil {
		return nil, nil, err
	}
	instantData, err := api.GetInstantData(c.Origin, o, c.FilterValue)
	if err != nil {
		return nil, nil, err
	}
	battery, err := api.GetBatteryData(c.Origin, o, c.BatteryDeviceID)
	if err != nil {
		return nil, nil, err
	}
	err = api.Logout(c.Origin, o)
	if err != nil {
		return nil, nil, err
	}
	return instantData, battery, nil
}

func postJSONData(elasticURL string, r io.Reader) error {
	resp, err := http.Post(elasticURL, "application/json", r)
	if err != nil {
		io.Copy(ioutil.Discard, r)
		r = nil
		return err
	}
	defer resp.Body.Close()
	io.Copy(ioutil.Discard, resp.Body)
	io.Copy(ioutil.Discard, r)
	r = nil
	return err
}
