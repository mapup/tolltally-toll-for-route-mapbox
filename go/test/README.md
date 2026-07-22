# [Mapbox](https://www.mapbox.com/)

### Get token to access Mapbox APIs (if you have an API token skip this)
#### Step 1: Login/Signup
* Create an accont to access [Mapbox Account Dashboard](https://account.mapbox.com/)
* go to signup/login link https://account.mapbox.com/auth/signin/

#### Step 2: Creating a token
* You will be presented with a default token.
* If you want you can create an application specific token.


To get the route polyline make a GET request on https://api.mapbox.com/directions/v5/mapbox/driving/${source.longitude},${source.latitude};${destination.longitude},${destination.latitude}?geometries=polyline&access_token=${token}&overview=full

### Note:
* we will be sending `geometries` as `polyline` and `overview` as `full`.
* Setting overview as full sends us complete route. Default value for `overview` is `simplified`, which is an approximate (smoothed) path of the resulting directions.
* Mapbox accepts source and destination, as semicolon seperated
  `${longitude,latitude}`.

```Go
The user has to enter the from and to in the CSV file and the program will give the output.
```

Whole working code can be found in test.go file.

### Running the tests

The API keys are read from the environment, so export them before running:

```sh
export MAPBOX_API_KEY="your-mapbox-token"
export TOLLGURU_API_KEY="your-tollguru-key"

cd go/test
go run test.go
```

Never hardcode the keys in `test.go` — the repository is public and committed secrets stay in git history even after they are deleted.
