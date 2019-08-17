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
	ticker := time.NewTicker(100 * time.Millisecond)
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
	// marshal particle data package
	j, _ := json.Marshal(particle.Data.(map[string]interface{}))

	// send particle endpoint request
	req, _ := http.NewRequest("POST", particle.Endpoint, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		// handle particle return [failure]
		reqFail, _ := http.NewRequest("POST", particle.Return.Fail, bytes.NewBuffer(j))
		reqFail.Header.Set("Content-Type", "application/json")
		clientFail := &http.Client{}
		respFail, err := clientFail.Do(reqFail)
		if err != nil {
			fmt.Println(err)
		}
		defer respFail.Body.Close()
		return
	}
	defer resp.Body.Close()

	// handle particle return [success]
	body, _ := ioutil.ReadAll(resp.Body)
	reqSuc, _ := http.NewRequest("POST", particle.Return.Success, bytes.NewBuffer(body))
	reqSuc.Header.Set("Content-Type", "application/json")
	clientSuc := &http.Client{}
	respSuc, err := clientSuc.Do(reqSuc)
	if err != nil {
		fmt.Println(err)
	}
	defer respSuc.Body.Close()
}
