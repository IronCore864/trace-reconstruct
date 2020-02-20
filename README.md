# Trace Reconstruct from Logs

## Dependency

go version go1.13.4 darwin/amd64

## Build

```
go build
```
## Unit Test

```
go test ./...`
```

Just an example, only tested the config module.

## Run

```
# change default interval (20s) by env var if needed:
export Interval=10
# then execute:
./trace-reconstruct
```

## Explanation

- there is an error in the documentation, which is the "span" in the example output, which does not correspond to the example input, caused a little confusion, please update :)
- total time taken: 2 hour (thinking + coding), plus 10 min (documentation)
- another around 1 hour to work on the 20s limit, clear cache, mutex part
- don't want to spend more time, so lots of things not implemented, but the core is there
- didn't pass traces-evaluator, it might be related to input format, but tested locally

First test case used is the example input in the doc;

Second test case is what I recovered from the output of the traces-evaluator:

```
2020-02-15T08:39:33.044Z 2020-02-15T08:39:33.123Z 5ry2k4vf service5 vlkh4jn2->v55vu7ab
2020-02-15T08:39:32.965Z 2020-02-15T08:39:33.285Z 5ry2k4vf service3 mym6zdfr->vlkh4jn2
2020-02-15T08:39:33.600Z 2020-02-15T08:39:33.602Z 5ry2k4vf service2 jl2dyo5v->wyqewehm
2020-02-15T08:39:33.600Z 2020-02-15T08:39:33.603Z 5ry2k4vf service1 oltcykrh->jl2dyo5v
2020-02-15T08:39:33.598Z 2020-02-15T08:39:33.603Z 5ry2k4vf service6 yb2bnwbx->oltcykrh
2020-02-15T08:39:33.543Z 2020-02-15T08:39:33.797Z 5ry2k4vf service9 mym6zdfr->yb2bnwbx
2020-02-15T08:39:34.656Z 2020-02-15T08:39:34.938Z 5ry2k4vf service7 mym6zdfr->znjk6pze
2020-02-15T08:39:35.073Z 2020-02-15T08:39:35.088Z 5ry2k4vf service1 53egpuym->bmmyuw5a
2020-02-15T08:39:35.029Z 2020-02-15T08:39:35.093Z 5ry2k4vf service7 mym6zdfr->53egpuym
2020-02-15T08:39:32.811Z 2020-02-15T08:39:35.561Z 5ry2k4vf service9 null->mym6zdfr
```

The timestamps showed when I finished my code and started testing.

What is missing:
- very basic input validation
- no optional stderr for statistics

If I work more on it, I will do:
- add more UT
- add more input validation to ignore malformed input lines
- add statistics about orphan lines and other required stuff
