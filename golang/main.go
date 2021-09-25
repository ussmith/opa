package main

import (
	"context"
	"encoding/json"
	"io/ioutil"

	"github.com/open-policy-agent/opa/rego"
	log "github.com/sirupsen/logrus"
)

func main() {
	r := rego.New(rego.Query("x = data.pigs.huff_puffable"),
		rego.Load([]string{"../rules.rego"}, nil))

	ctx := context.Background()
	query, err := r.PrepareForEval(ctx)

	if err != nil {
		log.WithError(err).Fatal("Failed to init rules")
	}

	byteSlice, err := ioutil.ReadFile("../piggies2.json")
	if err != nil {
		log.WithError(err).Fatal("Failed to load the input file")
	}

	var input interface{}

	err = json.Unmarshal(byteSlice, &input)

	if err != nil {
		log.WithError(err).Fatal("Failed to unmarshal")
	}

	results, err := query.Eval(ctx, rego.EvalInput(input))

	if err != nil {
		log.WithError(err).Fatal("Failed on the query evaluation")
	}

	//Just looking for one result
	log.Infof("Results: %v", results[0].Bindings["x"])
}
