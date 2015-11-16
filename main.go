// (c) lsls

package main

import (
	"fmt"
	"log"
)

func main() {

	cfg := LoadConfig()

	server := cfg.UString("server", "localhost")
	port := cfg.UInt("port", 443)
	collectors := cfg.UList("collectors")
	ncollectors := len(collectors)

	fmt.Printf("Server: %s\n", server)
	fmt.Printf("Port: %d\n", port)

	for i := 0; i < ncollectors; i++ {
		collectorCfg, err := cfg.Get(fmt.Sprintf("collectors.%d", i))
		if err != nil {
			log.Fatal(err)
		}

		t, err := collectorCfg.String("type")
		if err != nil {
			log.Printf("\"type\" not set for collector[%d]! Ignoring\n", i)
			continue
		}

		switch t {
		case "sql":
			cs := SqlCollector{}
			if err := cs.SetConfig(collectorCfg); err != nil {
				log.Printf("Cannot parse config for collector[%d]: %s\nIgnoring\n", i, err)
				continue
			}
			fmt.Println(cs)
		}
	}

}
