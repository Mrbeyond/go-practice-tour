package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"testing"
)

func BenchmarkJsonMarshal(b *testing.B) {
	data := map[string]interface{}{
		"Name": "John",
		"Age":  30,
		"Ade": map[string]interface{}{
			"a": 20,
			"b": 200,
			"c": 2000,
			"d": 20000,
			"e": 200000,
			"f": map[string]interface{}{
				"a": 20,
				"b": 200,
				"c": 2000,
				"d": 20000,
				"e": 200000,
				"f": map[string]interface{}{
					"a": 20,
					"b": 200,
					"c": 2000,
					"d": 20000,
					"e": 200000,
					"f": map[string]interface{}{
						"a": 20,
						"b": 200,
						"c": 2000,
						"d": 20000,
						"e": 200000,
						"f": "last",
					},
				},
			},
		},
	}

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkGobEncoder(b *testing.B) {
	type Dtype struct {
		Data map[string]interface{}
	}

	data := map[string]interface{}{
		"Name": "John",
		"Age":  30,
		"Ade": map[string]interface{}{
			"a": 20,
			"b": 200,
			"c": 2000,
			"d": 20000,
			"e": 200000,
			"f": map[string]interface{}{
				"a": 20,
				"b": 200,
				"c": 2000,
				"d": 20000,
				"e": 200000,
				"f": map[string]interface{}{
					"a": 20,
					"b": 200,
					"c": 2000,
					"d": 20000,
					"e": 200000,
					"f": map[string]interface{}{
						"a": 20,
						"b": 200,
						"c": 2000,
						"d": 20000,
						"e": 200000,
						"f": "last",
					},
				},
			},
		},
	}

	// gob.Register(Dtype{})

	gob.Register(map[string]interface{}{})

	ddata := Dtype{
		Data: data,
	}

	for i := 0; i < b.N; i++ {
		var buf bytes.Buffer
		encoder := gob.NewEncoder(&buf)
		err := encoder.Encode(ddata)
		if err != nil {
			b.Fatal(err)
		}
	}
}
