require 'HTTParty'
require 'json'
require 'cgi'

MAPBOX_API_KEY = ENV["MAPBOX_API_KEY"]
MAPBOX_API_URL = "https://api.mapbox.com/directions/v5/mapbox/driving"
MAPBOX_GEOCODE_API_URL = "https://api.mapbox.com/geocoding/v5/mapbox.places"

TOLLGURU_API_KEY = ENV["TOLLGURU_API_KEY"]
TOLLGURU_API_URL = "https://apis.tollguru.com/toll/v2"
POLYLINE_ENDPOINT = "complete-polyline-from-mapping-service"

def get_coord_array(loc)
    geocoding_url = "#{MAPBOX_GEOCODE_API_URL}/#{CGI::escape(loc)}.json?limit=1&access_token=#{MAPBOX_API_KEY}"
    coord = JSON.parse(HTTParty.get(geocoding_url).body)
    return  coord['features'].pop['geometry']['coordinates']
end

source = get_coord_array("Philadelphia, PA")
destination = get_coord_array("New York, NY")

# Explore https://tollguru.com/toll-api-docs to get the best of all the parameters that tollguru has to offer
request_parameters = {
  "vehicle": {
    "type": "2AxlesAuto",
  },
  # Visit https://en.wikipedia.org/wiki/Unix_time to know the time format
  "departure_time": "2021-01-05T09:46:08Z",
}

# GET Request to Mapbox for Polyline
MAPBOX_URL = "#{MAPBOX_API_URL}/#{source[0]},#{source[1]};#{destination[0]},#{destination[1]}?geometries=polyline&access_token=#{MAPBOX_API_KEY}&overview=full"
RESPONSE = HTTParty.get(MAPBOX_URL).body
json_parsed = JSON.parse(RESPONSE)

# Extracting mapbox polyline from JSON
mapbox_polyline = json_parsed['routes'].map { |x| x['geometry'] }.pop

# Sending POST request to TollGuru
TOLLGURU_URL = "#{TOLLGURU_API_URL}/#{POLYLINE_ENDPOINT}"
headers = {'content-type' => 'application/json', 'x-api-key' => TOLLGURU_API_KEY}
body = {'source' => "mapbox", 'polyline' => mapbox_polyline, **request_parameters}
tollguru_response = HTTParty.post(TOLLGURU_URL,:body => body.to_json, :headers => headers)
puts tollguru_response