package ds

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type model struct {
	Name string
	Ctx  map[string]string `json:"ctx"`
}

func TestMap(t *testing.T) {
	t.Run("serialise struct with nil map (default)", func(t *testing.T) {
		m := model{Name: "abc"}
		t.Logf("## access ctx map %v", m.Ctx)
		r, _ := json.Marshal(m)
		assert.Nil(t, m.Ctx)
		assert.Equal(t, 0, len(m.Ctx))
		t.Log(string(r))
		assert.Contains(t, string(r), `"ctx":null`)
	})
	t.Run("serialise with empty map", func(t *testing.T) {
		m := model{Name: "abc", Ctx: map[string]string{}}
		// m := model{Name: "abc", Ctx: map[string]string{"ref": "ref", "topic": "t"}}
		r, _ := json.Marshal(m)
		t.Log(string(r))
	})

	t.Run("deserialise struct with empty map", func(t *testing.T) {
		jd := `{"Name":"yu"}`
		m := model{}
		json.Unmarshal([]byte(jd), &m)

		assert.Equal(t, 0, len(m.Ctx))
		t.Log("print empty map:", m)
	})

	t.Run("serialise nil", func(t *testing.T) {
		var ref []int64
		ref = nil
		r, _ := json.Marshal(ref)
		t.Log(r)
	})
}

// func TestTypeAssert(t *testing.T) {
// 	input := "daisy-hill-qld-4127"

// 	t.Run("is ok", func(t *testing.T) {
// 		parts := strings.Split(input, "-")
// 		l := len(parts)
// 		state, postcode := parts[l-2], parts[l-1]
// 		sub := strings.Join(parts[:l-2], "-") + "-" + postcode
// 	})
// }

func TestStructJSON(t *testing.T) {
	type Ent struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
		Link string `json:"link"`
	}
	t.Run("should decode with missing field in json", func(t *testing.T) {
		jsonstr := `{"id":123, "name":"ent1"}`
		e := Ent{}
		json.Unmarshal([]byte(jsonstr), &e)
		t.Log(e)
	})

	t.Run("should decode with unknown struct field in json", func(t *testing.T) {
		jsonstr := `{"id":123, "name":"ent1", "link":"googole.com", "extra":123}`
		e := Ent{}
		json.Unmarshal([]byte(jsonstr), &e)
		t.Log(e)
	})
}
