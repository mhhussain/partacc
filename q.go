package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var q = make(chan Particle, 1000)

func runq() {
	ticker := time.NewTicker(10 * time.Second)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				qend := false
				for !qend {
					var tparticle Particle
					select {
					case tparticle = <-q:
						makeR(tparticle)
					default:
						qend = true
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()
}

func makeR(particle Particle) {
	//fmt.Println(json.Marshal(particle.Data.(map[string]interface{})))
	j, err := json.Marshal(particle.Data.(map[string]interface{}))
	req, err := http.NewRequest("POST", particle.Endpoint, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("wrong")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body: ", string(body))
}
