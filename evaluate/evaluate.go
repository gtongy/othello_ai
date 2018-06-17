package evaluate

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type Evaluation struct {
	Rows [][]int `json:"rows"`
}

func (evaluation *Evaluation) Set() {
	// JSONファイル読み込み
	bytes, err := ioutil.ReadFile("./evaluate/evaluation.json")
	if err != nil {
		log.Fatal(err)
	}
	if err := json.Unmarshal(bytes, &evaluation); err != nil {
		log.Fatal(err)
	}
}
