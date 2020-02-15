package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func getAllCalls(logs []Log, span string) []Log {
	res := make([]Log, 0)
	for i := 0; i < len(logs); i++ {
		if logs[i].callerSpan == span {
			logs[i].Calls = getAllCalls(logs, logs[i].Span)
			res = append(res, logs[i])
		}
	}
	return res
}

// based on the root log span, get all logs whose callerSpan euqals span
// recursive for callers using the function above
func traceReconstruct(logs []Log) {
	id := logs[0].trace
	root := logs[len(logs)-1]
	root.Calls = getAllCalls(logs, root.Span)
	res := Result{id, root}
	json, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}

// split input, no validation
func constructLogFromInput(input string) (bool, Log) {
	inputs := strings.Fields(input)
	if (len(inputs)) < 5 {
		return false, Log{}
	}
	start := inputs[0]
	end := inputs[1]
	trace := inputs[2]
	service := inputs[3]
	spans := strings.Split(inputs[4], "->")
	if (len(spans)) < 2 {
		return false, Log{}
	}
	callerSpan := spans[0]
	span := spans[1]
	return true, Log{start, end, trace, service, make([]Log, 0), callerSpan, span}
}

func main() {
	m := make(map[string][]Log)
	reader := bufio.NewReader(os.Stdin)

	for {
		input, _ := reader.ReadString('\n')

		valid, log := constructLogFromInput(input)
		if valid == false {
			continue
		}
		m[log.trace] = append(m[log.trace], log)
		// span "null->xxx" always finishes last, ignore following entries for the same trace
		if log.callerSpan == "null" {
			traceReconstruct(m[log.trace])
		}
	}
}
