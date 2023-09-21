# go-wasm-http-client-gzip-issue-reproduce
The behavior of go client is somewhat inconsistent between wasm run and native run.

```
go version go1.21.0 darwin/arm64
chrome version 116.0.5845.187 (Official Build) (arm64)
```
### step 1
#### run server
> The server provides an API (/gzip) to return "hello world" gzip string 
```sh
go run server/main.go
```

### step 2
> in main.go. requests /gzip api, and print content-encoding.
#### 1. native go run
```sh
go run main.go
# output
# Content-Encoding: 
# Content-Length: 
# ContentLength: -1
# body: [104 101 108 108 111 44 32 119 111 114 108 100]

# response body has been unzipped and header['Content Encoding'] removed
# go source: https://github.com/golang/go/blob/ace1494d9235be94f1325ab6e45105a446b3224c/src/net/http/transport.go#L2245
```

#### 2. compile into wasm and run it in a browser
```sh
# 1. compile
GOOS=js GOARCH=wasm go build -o main.wasm ./main.go
# 2. open browser. go to http://localhost:8080
# 3 click "run" button on the page
# 3. open browser devtools to check console outputs


# output
# Content-Encoding: gzip
# Content-Length: 36
# ContentLength: 36
# body: [104 101 108 108 111 44 32 119 111 114 108 100]

# response body has been unzipped and header['Content Encoding'] has not been deleted
```


## attempt to analyze the root cause
if run in wasm_js go uses fetch as the implementation of http.client

the behavior of the fetch Api is different with the behavior of the go http.client, fetch will automatically unzip but not delete the 'content-encoding' header.
