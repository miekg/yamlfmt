package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatStream(t *testing.T) {
	var out bytes.Buffer
	assert.NoError(t, formatStream(strings.NewReader(input), &out, 2))
	assert.Equal(t, output, out.String())
}

const (
	input = `
apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
spec:
  selector:
         app:    sqlflow-mysql
---
apiVersion: v1   
metadata:
    name: sqlflow-mysql
`
	output = `apiVersion: apps/v1 # for versions before 1.9.0 use apps/v1beta2
spec:
  selector:
    app: sqlflow-mysql
---
apiVersion: v1
metadata:
  name: sqlflow-mysql
`
)
