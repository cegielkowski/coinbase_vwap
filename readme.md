## Real-Time VWAP (volume-weighted average price) calculation engine.
*What is VWAP?*

The volume-weighted average price (VWAP) is a measurement that shows the average price of a security, adjusted for its volume. It is calculated during a specific trading session by taking the total dollar value of trading in the security and dividing it by the volume of trades.

VWAP is calculated using the following formula:


<img src="https://wikimedia.org/api/rest_v1/media/math/render/svg/6c0a822a0a9e58a127105e818a07061a02851685">

### Prerequisite

Change .envexample to .env to be able to run the project.

### Code design and ideas
The code was made in a way that was not so poor, and not so over-engineered, the idea here was to achieve a good structure to make it easy to maintain and improve.

Channels and go routines are the core of the project, they make everything work as expected, it was used an external JSON library to improve the performance of the code because the goal of the project was performance, this library is way faster than the native one.

I used Make() with predefined sizes of the structures too, the goal here is to save memory instead of creating an automatic allocation.

Sometimes the WebSocket disconnect itself, so I created a logic to try to reconnect 3 times before canceling the script.

Errors were treated in channels too, not to lose essential problems!

The vwap side of the code was split into tiny functions to improve the readability and make it easier to create tests and maintain.

Instead of using the len(window) all the time, I create an INT variable to handle this information, to achieve better performance.

### Test
Run the below command to run the tests, and make sure that all tests are passing
```
make test
```

### Running
Run the commands below to run the app locally
```
go mod download
```
```
make run-server
```

Run the commands below to run app in docker
```
make create-docker-image
```
```
make run-docker
```

### Improvements

1. Improve test coverage
2. Spread subscriptions (especially full channel subscriptions) over more than one websocket client connection. For example, do not subscribe to BTC-USD and ETH-USD on the same channel if possible. Instead, open up two separate websocket connections to help load balance those inbound messages across separate connections.
3. Create .env validation
