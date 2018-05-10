package main

import (
	"fmt"
	"log"
	"net/url"

	"github.com/intello-io/oktasdk-go/okta"
)

func logListQuery() {
	defer printEnd(printStart("logListQuery"))
	orgURL, err := url.Parse(orgName)
	if err != nil {
		log.Panicln(err)
	}
	client := okta.NewClientWithBaseURL(nil, orgURL, apiToken)
	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	logFilter := &okta.LogListFilterOptions{}
	logFilter.Limit = 1000
	logFilter.Query = "shlomo@intello.io"

	var aggLogs []okta.Log

	i := 0
	for {
		i++
		logs, resp, err := client.Logs.ListWithFilters(logFilter)
		if err != nil {
			log.Panicf("client.Log.ListWithFilters %v\n", err)
		}
		fmt.Printf("request for: %s resp status code: %d\n", resp.Request.URL, resp.StatusCode)
		aggLogs = append(aggLogs, logs...)
		if resp.NextURL == nil || len(logs) == 0 {
			break
		}
		logFilter.NextURL = resp.NextURL
	}
	fmt.Printf("len(all_logs) = %v\n", len(aggLogs))
	printLogArray(aggLogs)
}

func logListEventType() {
	defer printEnd(printStart("logListEventType"))
	orgURL, err := url.Parse(orgName)
	if err != nil {
		log.Panicln(err)
	}
	client := okta.NewClientWithBaseURL(nil, orgURL, apiToken)
	fmt.Printf("Client Base URL: %v\n\n", client.BaseURL)

	logFilter := &okta.LogListFilterOptions{}
	logFilter.Limit = 1000
	logFilter.EventType = "user.authentication.sso"

	var aggLogs []okta.Log

	i := 0
	for {
		i++
		logs, resp, err := client.Logs.ListWithFilters(logFilter)
		if err != nil {
			log.Panicf("client.Log.ListWithFilters %v\n", err)
		}
		fmt.Printf("request for: %s resp status code: %d\n", resp.Request.URL, resp.StatusCode)
		aggLogs = append(aggLogs, logs...)
		if resp.NextURL == nil || len(logs) == 0 {
			break
		}
		logFilter.NextURL = resp.NextURL
	}
	fmt.Printf("len(all_logs) = %v\n", len(aggLogs))
	printLogArray(aggLogs)
}
