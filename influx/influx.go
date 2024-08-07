package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/influxdata/influxdb-client-go/v2"
)

func main() {
	url := os.Getenv("INFLUX_URL")
	token := os.Getenv("INFLUX_TOKEN")
	org := "11"
	bucket := "12"

	// Create a new client using an InfluxDB server base URL and an authentication token
	client := influxdb2.NewClient(url, token)

	defer client.Close()

	// Use blocking write client for writes to desired bucket
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// Create point using full params constructor
	p := influxdb2.NewPoint("go_test_influx",
		map[string]string{"unit": "temperature"},
		map[string]interface{}{"avg": 5.5, "max": 30.0},
		time.Now())

	// write point immediately
	writeAPI.WritePoint(context.Background(), p)

	// Create point using fluent style
	p = influxdb2.NewPointWithMeasurement("go_test_influx").
		AddTag("unit", "temperature").
		AddField("avg", 17.2).
		AddField("max", 25.0).
		SetTime(time.Now())

	err := writeAPI.WritePoint(context.Background(), p)
	if err != nil {
		panic(err)
	}

	// Or write directly line protocol
	line := fmt.Sprintf("stat,unit=temperature avg=%f,max=%f", 5.5, 20.0)
	err = writeAPI.WriteRecord(context.Background(), line)
	if err != nil {
		panic(err)
	}

	// Get query client
	//queryAPI := client.QueryAPI(org)
	//// Get parser flux query result
	//result, err := queryAPI.Query(context.Background(), `from(bucket:"11")|> range(start: -1h) |> filter(fn: (r) => r._measurement == "stat")`)
	//if err == nil {
	//	// Use Next() to iterate over query result lines
	//	for result.Next() {
	//		// Observe when there is new grouping key producing new table
	//		if result.TableChanged() {
	//			fmt.Printf("table: %s\n", result.TableMetadata().String())
	//		}
	//		// read result
	//		fmt.Printf("row: %s\n", result.Record().String())
	//	}
	//	if result.Err() != nil {
	//		fmt.Printf("Query error: %s\n", result.Err().Error())
	//	}
	//} else {
	//	panic(err)
	//}
	//// Ensures background processes finishes
}
