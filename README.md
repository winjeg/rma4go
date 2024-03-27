# rma4go
[![CircleCI](https://circleci.com/gh/winjeg/rma4go/tree/master.svg?style=svg)](https://circleci.com/gh/winjeg/rma4go/?branch=master)
[![GithubWorkflow](https://github.com/winjeg/rma4go/actions/workflows/go.yml/badge.svg)](https://github.com/winjeg/rma4go/actions)
[![GolangCI](https://golangci.com/badges/github.com/winjeg/rma4go.svg)](https://golangci.com/r/github.com/winjeg/rma4go)
[![codecov](https://codecov.io/gh/winjeg/rma4go/branch/master/graph/badge.svg)](https://codecov.io/gh/winjeg/rma4go)
[![Mit License](https://img.shields.io/badge/license-MIT-blue)]('./LICENSE')
![MinVer](https://img.shields.io/badge/Golang-1.13-blue)

`rma4go` (redis memory analyzer for golang) is a online redis memory analyzer.
It's light and easy to use, and prints out a pretty markdown table

## functionality
rma4go provide statistics for redis in many dimensions, the mainly supported dimensions are listed below
1. key size, data size statistics
2. key expire statistics
3. big key statistics
4. different types of key statistics


## usage

### how to use in your code
```
go get github.com/winjeg/rma4go
```

```go

func testFunc() {
	h := "localhost"
	a := ""
	p := 6379
	cli := client.BuildRedisClient(client.ConnInfo{
		Host: h,
		Auth: a,
		Port: p,
	}, cmder.GetDb())

	stat := analyzer.ScanAllKeys(cli)
    // print in command line
	stat.Print()
	// the object is ready to use
}

```
### how to build
before build
1. make sure your go version >= 1.11.0
2. make sure you have internet access and can access google.com
```bash
git clone git@github.com:winjeg/rma4go.git
cd rma4go
go build .
```

### command line usage
usage is quite simple, it prints out a table in command line in markdown format
```
rma4go usage:
standalone:  rma4go -r some_host -p 6379 -a password -d 0
cluster:    rma4go -c "localhost:123,localhost:124,localhost:1234" -a "password"
======================================================
  -h help content
  -H string
    	address of a redis (default "localhost")
  -a string
    	password/auth of the redis
  -d int
    	db of the redis to analyze
  -p int
    	port of the redis (default 6379)
  -r string
    	address of a redis (default "localhost")
  -m match a pattern to scan 
```

## sample output
rendered by markdown 
total count 4004

all keys statistics

|                    PATTERN                    | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-----------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| TOP_TEN_NEW_XXXXXXXX*                         |       1 |       20 |      1529 |              0 |             0 |              0 |               0 |            1 |
| XXXXXXXXXXXXXX_STATISTICS_MIGRATION_LIST*     |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
| time-root:*                                   |      23 |      272 |       299 |              0 |             0 |              0 |               0 |           23 |
| DS_AXXXXXXXX_CORRECT*                         |       2 |       45 |        46 |              0 |             0 |              0 |               0 |            2 |
| time-2*                                       |     761 |     7528 |      9893 |              0 |             0 |              0 |               0 |          761 |
| time-level:*                                  |     537 |     8461 |      6981 |              0 |             0 |              0 |               0 |          537 |
| time-9*                                       |     102 |      901 |      1326 |              0 |             0 |              0 |               0 |          102 |
| time-7*                                       |     153 |     1372 |      1989 |              0 |             0 |              0 |               0 |          153 |
| DS_MAGIC_SUCC_2017-06-22*                     |       1 |       24 |       415 |              0 |             0 |              0 |               0 |            1 |
| tersssss*                                     |       5 |      124 |         0 |              0 |             0 |              0 |               0 |            5 |
| appoint_abcdefg_msgid*                        |       1 |       21 |         0 |              0 |             0 |              0 |               0 |            1 |
| BUSSINESSXXXXXXX_STATISTICS_NEED_CALC_RECENT* |       1 |       44 |         1 |              0 |             0 |              0 |               0 |            1 |
| switch_abcd_abcde*                            |       3 |       69 |         3 |              0 |             0 |              0 |               0 |            3 |
| abcdeferCounter_201*                          |       3 |       78 |         0 |              0 |             0 |              0 |               0 |            3 |
| diy1234567flag*                               |       1 |       14 |         1 |              0 |             0 |              0 |               0 |            1 |
| DS_PRXXBCD_LIST*                              |       1 |       15 |     17208 |              0 |             0 |              0 |               0 |            1 |
| time-4*                                       |     133 |     1194 |      1729 |              0 |             0 |              0 |               0 |          133 |
| datastatistics_switch_version0*               |       1 |       30 |         1 |              0 |             0 |              0 |               0 |            1 |
| register_count_2_201*                         |     592 |    15984 |       640 |              0 |             0 |              0 |               0 |          592 |
| canVisitNewabcdef1234PageLevels*              |       1 |       31 |         0 |              0 |             0 |              0 |               0 |            1 |
| YOUR_WEEK_VITALITY_INFO*                      |       1 |       23 |     75782 |              0 |             0 |              0 |               0 |            1 |
| time-8*                                       |     101 |      894 |      1313 |              0 |             0 |              0 |               0 |          101 |
| EXPERTS_APPOINT_INFO_MAP*                     |       1 |       24 |         0 |              0 |             0 |              0 |               0 |            1 |
| time-3*                                       |     130 |     1215 |      1690 |              0 |             0 |              0 |               0 |          130 |
| time-1*                                       |     943 |     9456 |     12259 |              0 |             0 |              0 |               0 |          943 |
| time-64*                                      |      87 |      781 |      1131 |              0 |             0 |              0 |               0 |           87 |
| time-5*                                       |     168 |     1516 |      2184 |              0 |             0 |              0 |               0 |          168 |
| total                                         |    4004 |    53422 |   7832490 |              0 |             0 |              0 |               0 |         4004 |


string keys statistics

|                    PATTERN                    | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-----------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| BUSSINESSXXXXXXX_STATISTICS_NEED_CALC_RECENT* |       1 |       44 |         1 |              0 |             0 |              0 |               0 |            1 |
| time-5*                                       |     130 |     1174 |      1690 |              0 |             0 |              0 |               0 |          130 |
| datastatistics_switch_version0*               |       1 |       30 |         1 |              0 |             0 |              0 |               0 |            1 |
| time-7*                                       |      39 |      348 |       507 |              0 |             0 |              0 |               0 |           39 |
| time-level:*                                  |     567 |     8939 |      7371 |              0 |             0 |              0 |               0 |          567 |
| diy1234567flag*                               |       1 |       14 |         1 |              0 |             0 |              0 |               0 |            1 |
| switch_abcd_abcde*                            |       3 |       69 |         3 |              0 |             0 |              0 |               0 |            3 |
| time-2*                                       |     598 |     5918 |      7774 |              0 |             0 |              0 |               0 |          598 |
| time-6*                                       |     125 |     1118 |      1625 |              0 |             0 |              0 |               0 |          125 |
| time-4*                                       |     136 |     1225 |      1768 |              0 |             0 |              0 |               0 |          136 |
| time-8*                                       |      72 |      636 |       936 |              0 |             0 |              0 |               0 |           72 |
| time-1*                                       |    1176 |    11814 |     15288 |              0 |             0 |              0 |               0 |         1176 |
| time-9*                                       |     100 |      880 |      1300 |              0 |             0 |              0 |               0 |          100 |
| time-root:*                                   |      23 |      272 |       299 |              0 |             0 |              0 |               0 |           23 |
| register_count_2_201*                         |     592 |    15984 |       640 |              0 |             0 |              0 |               0 |          592 |
| DS_AXXXXXXXX_CORRECT*                         |       1 |       20 |        20 |              0 |             0 |              0 |               0 |            1 |
| TOP_TEN_NEW_tersssss*                         |       1 |       20 |      1529 |              0 |             0 |              0 |               0 |            1 |
| time-3*                                       |     202 |     1925 |      2626 |              0 |             0 |              0 |               0 |          202 |
| total                                         |    3989 |    53042 |     46253 |              0 |             0 |              0 |               0 |         3989 |


list keys statistics

|                  PATTERN                  | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| XXXXXXXXXXXXXX_STATISTICS_MIGRATION_LIST* |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
| DS_MAGIC_SUCC_2017-06-22*                 |       1 |       24 |       415 |              0 |             0 |              0 |               0 |            1 |
| DS_PRXXBCD_LIST*                          |       1 |       15 |     17208 |              0 |             0 |              0 |               0 |            1 |
| total                                     |       3 |       79 |   7710455 |              0 |             0 |              0 |               0 |            3 |


hash keys statistics

|           PATTERN            | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| tersssss_action_prepage_new* |       1 |       27 |         0 |              0 |             0 |              0 |               0 |            1 |
| YOUR_WEEK_VITALITY_INFO*     |       1 |       23 |     75782 |              0 |             0 |              0 |               0 |            1 |
| EXPERTS_APPOINT_INFO_MAP*    |       1 |       24 |         0 |              0 |             0 |              0 |               0 |            1 |
| abcdeferCounter_2017-06-11*  |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| tersssssHardTaskCounter*     |       1 |       23 |         0 |              0 |             0 |              0 |               0 |            1 |
| abcdeferCounter_2018-04-27*  |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| abcdeferCounter_2017-09-01*  |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| tersssssEasyTaskCounter*     |       1 |       23 |         0 |              0 |             0 |              0 |               0 |            1 |
| total                        |       8 |      198 |     75782 |              0 |             0 |              0 |               0 |            8 |


set keys statistics

|             PATTERN              | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|----------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| tersssss_bind_phone_phone*       |       1 |       25 |         0 |              0 |             0 |              0 |               0 |            1 |
| appoint_abcdefg_msgid*           |       1 |       21 |         0 |              0 |             0 |              0 |               0 |            1 |
| canVisitNewabcdef1234PageLevels* |       1 |       31 |         0 |              0 |             0 |              0 |               0 |            1 |
| tersssss_bind_phone_userid*      |       1 |       26 |         0 |              0 |             0 |              0 |               0 |            1 |
| total                            |       4 |      103 |         0 |              0 |             0 |              0 |               0 |            4 |


zset keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
other keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |


big keys statistics

|                  PATTERN                  | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|-------------------------------------------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| XXXXXXXXXXXXXX_STATISTICS_MIGRATION_LIST* |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
| total                                     |       1 |       40 |   7692832 |              0 |             0 |              0 |               0 |            1 |
## maintain
1. welcome to contribute to this repo
2. welcome to submit issues for this repo
3. if you have any questions,  feel free to ask
