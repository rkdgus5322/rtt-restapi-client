package main

import (
	"encoding/json"
	"fmt"
	"openscrap/domain"
	"testing"

	_ "embed"

	"github.com/stretchr/testify/require"
)

//go:embed sample.json
var data []byte

func TestData(t *testing.T) {
	var model domain.Model
	err := json.Unmarshal(data, &model)
	require.NoError(t, err)

	fmt.Println(model.Response.Body.Items)

}
