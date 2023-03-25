# Connection interface for Ascendex

The pre-interview testing task

### Dependencies

The tool uses [Gorilla WebSocket](https://github.com/gorilla/websocket) package.

### Installation and usage

```
$ git clone https://github.com/aageorg/ascendex
$ cd ascendex
$ go build -o ascendex
$ ./ascendex BTC_USDT
bid price 27399.45 amout 0.014150, offer price 27416.02 amount 0.014150
bid price 27399.45 amout 0.014150, offer price 27416.00 amount 0.014150
bid price 27399.45 amout 0.014150, offer price 27415.99 amount 0.007300
bid price 27400.37 amout 0.008650, offer price 27415.99 amount 0.007300
bid price 27403.38 amout 0.008650, offer price 27415.98 amount 0.014370
bid price 27403.38 amout 0.008650, offer price 27415.97 amount 0.007300
bid price 27403.39 amout 0.014370, offer price 27415.94 amount 0.014370
bid price 27403.39 amout 0.014370, offer price 27415.92 amount 0.014370
bid price 27403.39 amout 0.023020, offer price 27415.90 amount 0.014370
bid price 27403.40 amout 0.014370, offer price 27415.88 amount 0.014370
...
```

### Tests
```
$ go test -v
=== RUN   TestConnection_Success
--- PASS: TestConnection_Success (1.05s)
=== RUN   TestConnection_Failure
--- PASS: TestConnection_Failure (0.00s)
=== RUN   TestSubscribeToChannel_Success
--- PASS: TestSubscribeToChannel_Success (0.00s)
=== RUN   TestSubscribeToChannel_Failure
--- PASS: TestSubscribeToChannel_Failure (0.00s)
=== RUN   TestWriteMessagesToChannel_Success
--- PASS: TestWriteMessagesToChannel_Success (0.26s)
=== RUN   TestReadMessagesFromChannel_Success
--- PASS: TestReadMessagesFromChannel_Success (0.25s)
PASS
ok      github.com/aageorg/ascendex    2.667s

```