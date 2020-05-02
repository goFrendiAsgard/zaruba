package example

import "testing"

type mockContext struct {
	param    string
	query    string
	postForm string
}

func (c *mockContext) Param(key string) string {
	return c.param
}

func (c *mockContext) Query(key string) string {
	return c.query
}

func (c *mockContext) PostForm(key string) string {
	return c.postForm
}

// TestGetNameEmptyContext test getName
func TestGetNameEmptyContext(t *testing.T) {
	c := &mockContext{}
	name := getName(c)
	if name != "" {
		t.Errorf("Unexpected result: %s", name)
	}
}

// TestGetNameParamContext test getName
func TestGetNameParamContext(t *testing.T) {
	c := &mockContext{param: "Kouga", query: "query", postForm: "postForm"}
	name := getName(c)
	if name != "Kouga" {
		t.Errorf("Unexpected result: %s", name)
	}
}

// TestGetNameQueryContext test getName
func TestGetNameQueryContext(t *testing.T) {
	c := &mockContext{query: "Kouga", postForm: "postForm"}
	name := getName(c)
	if name != "Kouga" {
		t.Errorf("Unexpected result: %s", name)
	}
}

// TestGetNamePostFormContext test getName
func TestGetNamePostFormContext(t *testing.T) {
	c := &mockContext{postForm: "Kouga"}
	name := getName(c)
	if name != "Kouga" {
		t.Errorf("Unexpected result: %s", name)
	}
}
