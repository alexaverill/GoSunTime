# GoSunTime
A Sunrise and Sunset API written in GO for fun. Implenting the algorith described here [Sunrise/Sunset Algorith](https://edwilliams.org/sunrise_sunset_algorithm.htm)

This is simply a test run of writing some GO code and then deploying it to Heroku. It may be available on Heroku, but it may not. [Potential Sunrise Sunset Endpoint](https://sunset-sunrise-api.herokuapp.com/sunrise) 

# Usage
This API can be used by sending a GET request with URL encoded values to either /sunrise or /sunset to get the UTC time of an event at a specific latitude and longitude on a specified day, 

Example Query: ``` /sunrise?lat=-40&lng=74.3&day=6&month=4&year=2020```

Parameters:
```
lat = Latitude of the location
lng = Longitude of the location
day = day you want to know about
month  = month you want to know about
year = year you want to know about
```

Response: 
```
{"Hour":1,"Minute":24,"TimeZone":"UTC"}

```