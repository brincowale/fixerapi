# Fixer API
Cache the rates from [fixer.io](https://fixer.io/) and cache it locally

## Docker
The environment variables required can be found in the `EXAMPLE.env` file

Run: `docker run -d -p 8000:8000 --env-file fixerapi.env --restart always --name fixerapi brincowale/fixerapi:1.0.0`


## Usage
`http://127.0.0.1:80/fixer?apikey=YOUR_API_KEY`  
`http://127.0.0.1:80/fixer/update?apikey=YOUR_API_KEY`

This API doesn't refresh automatically, so any other task should call the /update endpoint every 24 hours

## Build
`go build`
