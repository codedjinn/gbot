package main

import (
	"container/list"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"errors"
	"strings"
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
	errParse := json.Unmarshal(body, &timeZones)
	if errParse == nil {
		this.timeZones = timeZones
	}
}

func (this *Bot) Process() {
	this.executing = true

	fmt.Println("Start processing messages")

	for this.executing == true {

		inMsg := this.incoming.Front()
		if inMsg != nil {
			this.incoming.Remove(inMsg)

			result, err := parseMessage(inMsg.Value.(string));
			if err != nil {
				fmt.Println("ERROR!", err);
				return
			}

			fmt.Println(result);

			fmt.Println("incoming", inMsg.Value)
		}

		outMsg := this.outgoing.Front()
		if outMsg != nil {
			this.outgoing.Remove(outMsg)
			fmt.Println("outgoing", outMsg.Value)
		}

		this.executing = false
	}
	
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



// private
func parseMessage(msg string) (string, error) {
	if len(msg) > 512 || !strings.HasSuffix(msg, "\n") {
		return "", errors.New("Message must be smaller than 512 and contain a newline");
	}

	// extract username

	return msg, nil;
}

func getUsername(msg string) (string, byte) {
	if strings.
}
