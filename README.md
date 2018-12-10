# rma4go
rma4go (redis memory analyzer for golang) is a online redis memory analyzer.
It's light and easy to use, and prints out an very pretty markdown table

## usage
usage is quite simple, it prints out a table in command line in markdown format
```
rma4go usage:
rma4go -r some_host -p 6379 -a password -d 0
======================================================
  -H string
        address of a redis (default "localhost")
  -a string
        password/auth of the redis
  -d int
        db of the redis to analyze
  -h    help content
  -p int
        port of the redis (default 6379)
  -r string
        address of a redis (default "localhost")
```

## sample output

```
all keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
string keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
list keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
hash keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
set keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
zset keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
other keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
big keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |

```
rendered by markdown 

all keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
string keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
list keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
hash keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
set keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
zset keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
other keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |
big keys statistics

| PATTERN | KEY NUM | KEY SIZE | DATA SIZE | EXPIRE IN HOUR | EXPIRE IN DAY | EXPIRE IN WEEK | EXPIRE OUT WEEK | NEVER EXPIRE |
|---------|---------|----------|-----------|----------------|---------------|----------------|-----------------|--------------|
| total   |       0 |        0 |         0 |              0 |             0 |              0 |               0 |            0 |