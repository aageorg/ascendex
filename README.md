# Connection interface for Ascendex

The pre-interview testing task

### Dependencies

The tool uses [Gorilla WebSocket](https://github.com/gorilla/websocket) package.

### Installation and usage

```
$git clone https://github.com/aageorg/ascendex
$cd ascendex
$go build -o apiclient
$./apiclient BTC/USDC
BTC/USDT: bid 27723.4 offer 27775.3
BTC/USDT: bid 27723.41 offer 27771.8
BTC/USDT: bid 27730.8 offer 27771.79
BTC/USDT: bid 27730.8 offer 27771.79
BTC/USDT: bid 27730.8 offer 27771.79
BTC/USDT: bid 27730.82 offer 27774.59
BTC/USDT: bid 27732.72 offer 27774.59
BTC/USDT: bid 27732.77 offer 27771.78
BTC/USDT: bid 27732.73 offer 27771.78
BTC/USDT: bid 27732.78 offer 27771.78
...
```
