package main

import (
	"fmt"
	"testing"

	"github.com/robfig/cron/v3"
	"github.com/stretchr/testify/require"
)

func TestData(t *testing.T) {
	c := cron.New()
	id, err := c.AddFunc("* * * * *", func() {
		fmt.Println("hi")
	})
	require.NoError(t, err)
	fmt.Println(id)
	c.Start()
}
