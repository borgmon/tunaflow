package main

type Upstream struct {
	Enabled bool   `json:"enabled"`
	Msg     string `json:"msg"`
	Nested  struct {
		Value int64 `json:"value"`
	} `json:"nested"`
}
