package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

type leaderResp struct {
	Leader string `json:"leader"`
}

func main() {
	log.SetFlags(0)
	peers := strings.Split(os.Getenv("ETCDCTL_PEERS"), ",")
	leaderChan := make(chan string, 0)
	for _, peer := range peers {
		peer := peer
		go func() {
			resp, err := http.Get(peer + "/v2/stats/leader")
			if err != nil {
				log.Println(err)
				return
			}
			defer resp.Body.Close()
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Println(err)
				return
			}
			var leaderResp leaderResp
			if err = json.Unmarshal(body, &leaderResp); err != nil {
				log.Println(err)
			}
			if leaderResp.Leader != "" {
				leaderChan <- peer
			}
		}()
	}

	// Responses have 5 seconds to finish.
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(5 * time.Second)
		timeout <- true
	}()

	select {
	case leader := <-leaderChan:
		log.Println(leader)
	case <-timeout:
		log.Fatal("timed out waiting for a response")
	}
}
