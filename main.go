package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"trace-reconstruct/config"
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
func traceReconstruct(m map[string][]Log, lock *sync.Mutex) {
	for trace, logs := range m {
		if time.Now().Sub(logs[0].insertTime) < time.Duration(config.Conf.Interval)*time.Second {
			continue
		}

		root := logs[len(logs)-1]
		root.Calls = getAllCalls(logs, root.Span)
		res := Result{trace, root}
		json, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(json))
		lock.Lock()
		delete(m, trace)
		lock.Unlock()
	}
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
	return true, Log{time.Now(), start, end, trace, service, make([]Log, 0), callerSpan, span}
}

func startPolling(m map[string][]Log, lock *sync.Mutex) {
	for range time.Tick(1 * time.Second) {
		go traceReconstruct(m, lock)
	}
}

func main() {
	m := make(map[string][]Log)
	lock := &sync.Mutex{}
	go startPolling(m, lock)

	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		valid, log := constructLogFromInput(input)
		if valid == false {
			continue
		}
		lock.Lock()
		m[log.trace] = append(m[log.trace], log)
		lock.Unlock()
	}
}
