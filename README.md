I decided to run the mini app in a docker container. The image is built locally and the container traffic simply served on localhost.

local image build:
docker build -t itinerary-app .

run the container on localhost:8080
docker run -p 8080:8080 --name my-itinerary-app itinerary-app

Additionally, one could mount the local folder as volume to help real time development
docker run -p 8080:8080 --name my-itinerary-app -v .:/app itinerary-app

