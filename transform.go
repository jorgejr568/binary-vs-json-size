package main

import (
	"encoding/json"
	"google.golang.org/protobuf/proto"
	"os"
	protos_models "proto-diff/protos"
)

type CovidTimeSeriesEntry struct {
	Date      string `json:"date"`
	Confirmed int    `json:"confirmed"`
	Deaths    int    `json:"deaths"`
	Recovered int    `json:"recovered"`
}

type CovidTimeSeries map[string][]CovidTimeSeriesEntry

func ConvertEntry(entry CovidTimeSeriesEntry) (*protos_models.CovidTimeSeriesEntry, error) {
	return &protos_models.CovidTimeSeriesEntry{
		Date:      entry.Date,
		Confirmed: int32(entry.Confirmed),
		Deaths:    int32(entry.Deaths),
		Recovered: int32(entry.Recovered),
	}, nil
}

func ConvertTimeSeries(timeSeries CovidTimeSeries) (*protos_models.CovidTimeSeries, error) {
	resultEntries := make([]*protos_models.CovidTimeSeriesEntry, 0)
	for _, entries := range timeSeries {
		for _, entry := range entries {
			convertedEntry, err := ConvertEntry(entry)
			if err != nil {
				return nil, err
			}
			resultEntries = append(resultEntries, convertedEntry)
		}
	}

	return &protos_models.CovidTimeSeries{
		Entries: resultEntries,
	}, nil
}

func main() {
	file, err := os.OpenFile("covid-timeseries.json", os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}

	jsonTimeSeries := new(CovidTimeSeries)
	err = json.NewDecoder(file).Decode(jsonTimeSeries)
	if err != nil {
		panic(err)
	}

	protoTimeSeries, err := ConvertTimeSeries(*jsonTimeSeries)
	if err != nil {
		panic(err)
	}

	protoFile, err := os.OpenFile("covid-timeseries.pb", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	pbData, err := proto.Marshal(protoTimeSeries)
	if err != nil {
		panic(err)
	}

	_, err = protoFile.Write(pbData)
	if err != nil {
		panic(err)
	}
}
