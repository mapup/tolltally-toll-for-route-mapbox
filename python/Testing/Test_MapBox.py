# Importing modules
import json
import requests
import os

MAPBOX_API_KEY = os.environ.get("MAPBOX_API_KEY")
MAPBOX_API_URL = "https://api.mapbox.com/directions/v5/mapbox/driving"
MAPBOX_GEOCODE_API_URL = "https://api.mapbox.com/geocoding/v5/mapbox.places"

TOLLGURU_API_KEY = os.environ.get("TOLLGURU_API_KEY")
TOLLGURU_API_URL = "https://apis.tollguru.com/toll/v2"
POLYLINE_ENDPOINT = "complete-polyline-from-mapping-service"

# Explore https://tollguru.com/toll-api-docs to get best of all the parameter that TollGuru has to offer
request_parameters = {
    "vehicle": {
        "type": "2AxlesAuto"
    },
    # Visit https://en.wikipedia.org/wiki/Unix_time to know the time format
    "departure_time": "2021-01-05T09:46:08Z",
}

# Fetching geocode from Mapbox
def get_geocode_from_mapbox(address):
    address_actual = address
    address = address.replace(" ", "%20").replace(",", "%2C")
    url = (
        f"{MAPBOX_GEOCODE_API_URL}/{address}.json?limit=1&access_token={MAPBOX_API_KEY}"
    )
    # print(url)
    res = requests.get(url).json()
    # print(res)
    try:
        return res["features"][0]["geometry"]["coordinates"]
    except:
        print(f"error in name {address_actual}")
        return (False, False)


# Fetching Polyline from Mapbox
def get_polyline_from_mapbox(
    source_longitude, source_latitude, destination_longitude, destination_latitude
):
    # Query Mapbox with Key and Source-Destination coordinates
    url = (
        "{a}/{b},{c};{d},{e}?geometries=polyline&access_token={f}&overview=full".format(
            a=MAPBOX_API_URL,
            b=source_longitude,
            c=source_latitude,
            d=destination_longitude,
            e=destination_latitude,
            f=MAPBOX_API_KEY,
        )
    )
    # converting the response to json
    response_from_mapbox = requests.get(url, timeout=200).json()
    # checking for errors in response
    if str(response_from_mapbox).find("message") == -1:
        # Extracting polyline if no error
        polyline_from_mapbox = response_from_mapbox["routes"][0]["geometry"]
        return polyline_from_mapbox
    else:
        raise Exception("{}".format(response_from_mapbox["message"]))

# Calling Tollguru API
def get_rates_from_tollguru(polyline, count=0):
    # Tollguru querry url
    Tolls_URL = f"{TOLLGURU_API_URL}/{POLYLINE_ENDPOINT}"
    # Tollguru resquest parameters
    headers = {"Content-type": "application/json", "x-api-key": TOLLGURU_API_KEY}
    params = {
        # explore https://tollguru.com/developers/docs/ to get best off all the parameter that tollguru offers
        "source": "mapbox",
        "polyline": polyline,
        **request_parameters,
    }
    # Requesting Tollguru with parameters
    response_tollguru = requests.post(
        Tolls_URL, json=params, headers=headers, timeout=200
    ).json()
    # checking for errors or printing rates
    if str(response_tollguru).find("message") == -1:
        # return rates in dictionary is not error
        return response_tollguru["route"]["costs"]
    else:
        raise Exception("{} in row {}".format(response_tollguru["message"], count))

"""Testing"""
# Importing Functions
from csv import reader, writer
import time

temp_list = []
with open("testCases.csv", "r") as f:
    csv_reader = reader(f)
    for count, i in enumerate(csv_reader):
        # if count>2:
        #   break
        if count == 0:
            i.extend(
                (
                    "Input_polyline",
                    "Tollguru_Tag_Cost",
                    "Tollguru_Cash_Cost",
                    "Tollguru_QueryTime_In_Sec",
                )
            )
        else:
            try:
                source_longitude, source_latitude = get_geocode_from_mapbox(i[1])
                destination_longitude, destination_latitude = get_geocode_from_mapbox(
                    i[2]
                )
                polyline = get_polyline_from_mapbox(
                    source_longitude,
                    source_latitude,
                    destination_longitude,
                    destination_latitude,
                )
                i.append(polyline)
            except:
                i.append("Routing Error")

            start = time.time()
            try:
                rates = get_rates_from_tollguru(polyline)
            except:
                i.append(False)
                rates = {}
            time_taken = time.time() - start

            if not rates:
                i.append((None, None))
            else:
                try:
                    tag = rates.get("tag")
                except:
                    tag = None
                try:
                    cash = rates.get("cash")
                except:
                    cash = None
                i.extend((tag, cash))

            i.append(time_taken)

        # print(f"{len(i)}   {i}\n")
        temp_list.append(i)

with open("testCases_result.csv", "w") as f:
    writer(f).writerows(temp_list)

"""Testing Ends"""
