package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Bot struct {
	executing  bool
	timeApiUrl string
	timeZones  []string
	messages   map[string]string
	incoming   *list.List
	outgoing   *list.List
}

func NewBot(timeApiUrl string) *Bot {
	result := Bot{
		timeApiUrl: timeApiUrl,
	}
	result.messages = make(map[string]string)
	result.incoming = list.New()
	result.outgoing = list.New()
	return &result
}

func (this *Bot) Initialize() {

	fmt.Println("Initializing bot...")

	response, err := http.Get(this.timeApiUrl + "/timezone")
	if err != nil {
		fmt.Println("Unable to get timezones")
		return
	}

	// make sure we close connection when method is finished
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading repsonse body")
		return
	}

	// get timezones
	var timeZones []string
	err1 := json.Unmarshal(body, &timeZones)
	if err1 == nil {
		this.timeZones = timeZones
	}
}

func (this *Bot) Process() {
	this.executing = true

	fmt.Println("Starting to process")
	// for this.executing {
	// 	fmt.Println("executing")
	// }
	fmt.Println("Finishing processing")
}

func (this *Bot) AddMessage(msg string) {
	this.incoming.PushFront(msg)
}

func (this *Bot) PrintAll() {
	for item := this.incoming.Front(); item != nil; item = item.Next() {
		fmt.Println(item.Value.(string))
	}
}
