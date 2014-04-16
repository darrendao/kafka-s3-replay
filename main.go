package main

import (
	"flag"
	"fmt"
	configfile "github.com/crowdmob/goconfig"
	"github.com/darrendao/kafka-go-have-fun/s3replay"
	"time"
)

var configFilename string
var keepBufferFiles bool
var debug bool
var shouldOutputVersion bool
var hostsStr string
var config *configfile.ConfigFile
var clusterId string
var topicsStr string

const (
	VERSION                            = "0.1"
	ONE_MINUTE_IN_NANOS                = 60000000000
	S3_REWIND_IN_DAYS_BEFORE_LONG_LOOP = 3
	DAY_IN_SECONDS                     = 24 * 60 * 60
)

func init() {
	flag.StringVar(&configFilename, "c", "consumer.properties", "path to config file")
	flag.BoolVar(&keepBufferFiles, "k", false, "keep buffer files around for inspection")
	flag.BoolVar(&shouldOutputVersion, "v", false, "output the current version and quit")
	flag.StringVar(&hostsStr, "h", "localhost:9092", "host:port comma separated list")
	flag.StringVar(&clusterId, "i", "", "ID of the Kafka cluster")
	flag.StringVar(&topicsStr, "t", "*", "comma separated list of topics. Defaults to all.")
}

func main() {
	config, _ = configfile.ReadConfigFile(configFilename)

	targets := []string{"localhost:9092"}

	startDate, err := time.Parse("2006-01-02", "2014-04-15")
	endDate, _ := time.Parse("2006-01-02", "2014-04-15")

	if err != nil {
		println(err.Error())
	}

	println("startdate is", startDate.String())
	fmt.Println(startDate)
	s3replay.Replay(config, targets, "", "test4", 0, startDate, endDate)
}
