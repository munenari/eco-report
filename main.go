package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
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
	for range tick.C {
		err := execute(c)
		if err != nil {
			log.Println(err)
			continue
		}
	}
}

func execute(c *config.EcoReport) error {
	instantData, batteryData, err := getProperties(c)
	if err != nil {
		return fmt.Errorf("failed to get data from API: %s", err)
	}
	record := make(map[string]interface{})
	record["time"] = time.Now().UTC()
	record["instantData"] = instantData
	record["batteryData"] = batteryData
	record["instantDataCircuits"] = instantData.GetCircuitsMap()
	b, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("ailed to marshal json from result map: %s", err)
	}
	err = postJSONData(c.ElasticOrigin, b)
	if err != nil {
		return fmt.Errorf("failed to post data to elasticsearch: %s", err)
	}
	log.Println("success to post data", string(b))
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

func postJSONData(elasticURL string, b []byte) error {
	r := bytes.NewReader(b)
	resp, err := http.Post(elasticURL, "application/json", r)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
