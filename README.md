I decided to run the mini app in a very simple docker container. The image is built locally and the container traffic simply served on localhost. 

The code can be tested via a simple ``go test -v`` command from the project root folder. 

local image build:
``docker build -t itinerary-app .``

run the container on localhost:8080
``docker run -p 8080:8080 --name my-itinerary-app itinerary-app``

Additionally, one could mount the local folder as volume to help real time development
``docker run -p 8080:8080 --name my-itinerary-app -v .:/app itinerary-app``

simlpy run this curl request with the following payloads to hit the itinerary endpoint

``curl -X POST http://localhost:8080/itinerary \
     -H "Content-Type: application/json" \
     -d '{"tickets": [["JFK", "SFO"], ["SFO", "ATL"], ["ATL", "JFK"]]}'``