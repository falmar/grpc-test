# grpc-test

Play around with GRPC

## Build
```bash
# compile protobuf
$ protoc -I pb todo.proto --go_out=plugins=grpc:pb

# build binaries
$ go build -o build/server ./server && ./build/server
$ go build -o build/client ./client && ./build/client
```
 
## Client CLI API

#### List TODOs

```text
Usage of list:
  -c    retrieve completed TODOs
  -n int
        max amount of items (default 15)
  -o int
        offset of items
  -q string
        search query string
```

```bash
$ ./build/client list
ID:"5af7b965-dbc3-4f00-a416-13614d0445b0" Name:"create a cli" 
ID:"bd6677dd-a5b5-4448-adb6-f52ca9863c16" Name:"update readme"

$ ./build/client list -q readme
ID:"bd6677dd-a5b5-4448-adb6-f52ca9863c16" Name:"update readme"

$ ./build/client list -n 1
ID:"5af7b965-dbc3-4f00-a416-13614d0445b0" Name:"create a cli"

$ ./build/client list -n 1 -o
ID:"bd6677dd-a5b5-4448-adb6-f52ca9863c16" Name:"update readme"
```

#### Create TODO

```text
Usage of create:
  -c    set TODO as completed
  -name string
        TODO name
```

```bash
$ ./build/client create -name "create a cli"
ID:"5af7b965-dbc3-4f00-a416-13614d0445b0" Name:"create a cli"

& ./build/client create -name "update readme"
ID:"bd6677dd-a5b5-4448-adb6-f52ca9863c16" Name:"update readme"
```

#### Mark TODO as Completed

```text
Usage of update:
  -c    set TODO as completed
  -id string
        TODO id
```

```bash
$ ./build/client mark-completed -c -id bd6677dd-a5b5-4448-adb6-f52ca9863c16
ID:"bd6677dd-a5b5-4448-adb6-f52ca9863c16" Completed:true

$ ./build/client mark-completed -id bd6677dd-a5b5-4448-adb6-f52ca9863c16
ID:"bd6677dd-a5b5-4448-adb6-f52ca9863c16"
```


#### Delete TODO

```text
Usage of delete:
  -id string
        TODO id
```

```bash
$ ./build/client delete -id bd6677dd-a5b5-4448-adb6-f52ca9863c16
```