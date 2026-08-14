# [Mapbox](https://www.mapbox.com/)

### Get token to access Mapbox APIs (if you have an API token skip this)
#### Step 1: Login/Signup
* Create an accont to access [Mapbox Account Dashboard](https://account.mapbox.com/)
* Go to signup/login link https://account.mapbox.com/auth/signin/

#### Step 2: Creating a token
* You will be presented with a default token.
* If you want you can create an application specific token.


To get the route polyline make a GET request on 'https://api.mapbox.com/directions/v5/mapbox/driving/'source_longitude+','+source_latitude+';'+destination_longitude+','+destination_latitude+'?geometries=polyline&access_token='+token+'&overview=full'

Example of GET request : https://api.mapbox.com/directions/v5/mapbox/driving/-96.7970,32.7767;-74.0060,40.7128?geometries=polyline&access_token=jk.evgggiejdjks2ZWxjbWFwdXAiLCJhIjoiY2tQ&overview=full

### Note:
* We will be sending `geometries` as `polyline` and `overview` as `full`.
* Setting overview as full sends us complete route. Default value for `overview` is `simplified`, which is an approximate (smoothed) path of the resulting directions.
* Mapbox accepts source and destination, as semicolon seperated
  `{longitude},{latitud}`.

```php

//using mapbox API

//Source and Destination Coordinates..
$source_longitude='-96.79448';
$source_latitude='32.78165';
$destination_longitude='-96.818';
$destination_latitude='32.95399';
$key = 'mapbox.token';

$url='https://api.mapbox.com/directions/v5/mapbox/driving/'.$source_longitude.','.$source_latitude.';'.$destination_longitude.','.$destination_latitude.'?geometries=polyline&access_token='.$key.'&overview=full';

//connection..
$mapbox = curl_init();

curl_setopt($mapbox, CURLOPT_SSL_VERIFYHOST, 2);
curl_setopt($mapbox, CURLOPT_SSL_VERIFYPEER, true);

curl_setopt($mapbox, CURLOPT_URL, $url);
curl_setopt($mapbox, CURLOPT_RETURNTRANSFER, true);

//getting response from mapboxapis..
$response = curl_exec($mapbox);
$err = curl_error($mapbox);

curl_close($mapbox);

if ($err) {
	  echo "cURL Error #:" . $err;
} else {
	  echo "200 : OK\n";
}

//extracting polyline from the JSON response..
$data_mapbox = json_decode($response, true);

//polyline..
$polyline_mapbox = $data_mapbox['routes']['0']['geometry'];


```

Note:

We extracted the polyline for a route from Mapbox API

We need to send this route polyline to TollGuru API to receive toll information

## [TollGuru API](https://tollguru.com/developers/docs/)

### Get key to access TollGuru polyline API
* Create a dev account to receive a free key from TollGuru https://tollguru.com/developers/get-api-key
* Suggest adding `vehicleType` parameter. Tolls for cars are different than trucks and therefore if `vehicleType` is not specified, may not receive accurate tolls. For example, tolls are generally higher for trucks than cars. If `vehicleType` is not specified, by default tolls are returned for 2-axle cars. 
* Similarly, `departure_time` is important for locations where tolls change based on time-of-the-day and can be passed through `$postdata`.

This snippet can be added at end of the above code to get rates and other details.
```php

//using tollguru API..
$curl = curl_init();

curl_setopt($curl, CURLOPT_SSL_VERIFYHOST, 2);
curl_setopt($curl, CURLOPT_SSL_VERIFYPEER, true);


$postdata = array(
	"source" => "google",
	"polyline" => $polyline_mapbox
);

//json encoding source and polyline to send as postfields..
$encode_postData = json_encode($postdata);

curl_setopt_array($curl, array(
CURLOPT_URL => "https://dev.tollguru.com/v1/calc/route",
CURLOPT_RETURNTRANSFER => true,
CURLOPT_ENCODING => "",
CURLOPT_MAXREDIRS => 10,
CURLOPT_TIMEOUT => 30,
CURLOPT_HTTP_VERSION => CURL_HTTP_VERSION_1_1,
CURLOPT_CUSTOMREQUEST => "POST",


//sending mapbox polyline to tollguru
CURLOPT_POSTFIELDS => $encode_postData,
CURLOPT_HTTPHEADER => array(
				      "content-type: application/json",
				      "x-api-key: tollguru.api"),
));

$response = curl_exec($curl);
$err = curl_error($curl);

curl_close($curl);

if ($err) {
	  echo "cURL Error #:" . $err;
} else {
	  echo "200 : OK\n";
}

//response from tollguru..
$data = var_dump(json_decode($response, true));
print_r($data);
    
```

Whole working code can be found in `php_curl_mapbox.php` file.
