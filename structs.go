package main

// Log is the struct for each line of input, also used in the output
type Log struct {
	Start      string `json:"start"`
	End        string `json:"end"`
	trace      string
	Service    string `json:"service"`
	Calls      []Log  `json:"calls"`
	callerSpan string
	Span       string `json:"span"`
}

// Result is the output struct
type Result struct {
	ID   string `json:"id"`
	Root Log    `json:"root"`
}
