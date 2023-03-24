# Connection interface for Ascendex

The pre-interview testing task

### Dependencies

The tool uses [Gorilla WebSocket](https://github.com/gorilla/websocket) package.

### Installation and usage

```
$ git clone https://github.com/aageorg/ascendex
$ cd ascendex
$ go build -o apiclient
$ ./apiclient BTC_USDT
bid price 27744.980000 amout 0.013570, offer price 27760.390000 amount 0.001500
bid price 27745.000000 amout 0.013570, offer price 27760.390000 amount 0.001500
bid price 27745.020000 amout 0.013570, offer price 27760.390000 amount 0.001500
bid price 27745.040000 amout 0.013570, offer price 27760.380000 amount 0.004900
bid price 27745.060000 amout 0.013570, offer price 27760.380000 amount 0.004900
bid price 27745.080000 amout 0.013570, offer price 27760.380000 amount 0.004900
bid price 27745.090000 amout 0.007200, offer price 27760.380000 amount 0.001500
bid price 27745.120000 amout 0.013570, offer price 27760.380000 amount 0.001500
bid price 27745.140000 amout 0.013570, offer price 27760.380000 amount 0.001500
bid price 27745.160000 amout 0.013570, offer price 27760.380000 amount 0.001500
bid price 27745.170000 amout 0.007200, offer price 27760.370000 amount 0.004900
bid price 27744.670000 amout 0.013570, offer price 27760.370000 amount 0.004900
bid price 27744.690000 amout 0.013570, offer price 27760.370000 amount 0.004900
bid price 27744.710000 amout 0.013570, offer price 27760.370000 amount 0.004900
bid price 27744.730000 amout 0.013550, offer price 27760.370000 amount 0.001500
...
```
