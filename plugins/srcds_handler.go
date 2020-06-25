package main

import (
	"log"
	"regexp"
)

var PayloadHandlerDescription string = "Regex to Datacollector"
var PayloadHandlerVersion string = "0.0.0a"
var Debug bool = false

var LogEntryRegex = regexp.MustCompile(`(?P<prefix>.*): (?P<firstparta>(?P<initiator>[\w\s]+)|(?P<firstpartb>"(?P<name>[\w\d]+)<(?P<id>[\d]+)><(?P<steamid>[\[\]\w:]+)><(?P<team>[\w]?|[\w]+)>")) (?P<secondpart>.*)`)
var SecondPartRegex = regexp.MustCompile(`(?P<connected>connected, address "(?P<address>[\d\.]+)")|(?P<steamvalidated>STEAM USERID validated)|(?P<joinedteam>joined team "(?P<newteam>[\w]+)")|(?P<enteredgame>entered the game)|(?P<killed>killed (?P<killinfo>.*))|(?P<triggered>triggered (?P<triggeredinfo>.*))`)

func PayloadHandlerFunction(payload []byte) (bool, error) {
	log.Printf("Handling payload length=%d\n", len(payload))
	found := LogEntryRegex.FindSubmatch(payload)

	if Debug {
		log.Printf("Payload: %s\n", payload)
	}

	if len(found) == len(LogEntryRegex.SubexpNames()) {
		log.Print("Found %s matches", len(found))
		//
		// Fields:
		// Log entry: "prefix" "firstparta" "initiator" "firstpartb" "name" "id" "steamid" "team" "secondpart"
		// 2nd part: "connected" "address" "steamvalidated" "joinedteam" "newteam" "enteredgame" "killed" "killinfo" "triggered" "triggeredinfo"
		//

		if len(found[5]) > 0 {
			log.Printf("Player '%s' %s Team '%s'", found[5], found[7], found[8])

			secondpartFound := SecondPartRegex.FindSubmatch(found[9])
			log.Printf("Second part: %q", secondpartFound)
		}
	}

	return true, nil
}
