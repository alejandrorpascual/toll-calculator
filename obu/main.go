package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	types "github.com/yosonoronosoy/tolling/types"
)

const sendInterval = time.Second
const wsEndpoint = "ws://localhost:30000/ws"

func genCoord() float64 {
	return float64(rand.Intn(100)+1) + rand.Float64()
}

func genLocation() (float64, float64) {
	return genCoord(), genCoord()
}

func main() {
	obuIDs := generateOBUIDS(20)
	conn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for _, id := range obuIDs {
			lat, long := genLocation()
			data := types.OBUData{
				OBUID: id,
				Lat:   lat,
				Long:  long,
			}
			if err := conn.WriteJSON(data); err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(sendInterval)
	}
}

func generateOBUIDS(n int) []int {
	ids := make([]int, n)
	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}
	return ids
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
