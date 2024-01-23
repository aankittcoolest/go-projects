### Go-replay Fun

#### Motivation
- This project aims at demonstrating how PROD traffic can be captured and then played back into QA environment.
- For simulation, we have used three PROD http servers sitting over an nginx load balancer (nginx-prod).
- For simulation, we have used three QA http servers sitting over an nginx load balancer (nginx-qa).
- PROD environment is simulated on "http://localhost:8090"
- QA environment is simulated on "http://localhost:9000"

#### How to use

```shell
# Bring up the containers
docker-compose up --build

# Open a different terminal and run the following commands
wget https://github.com/buger/goreplay/releases/download/1.3.3/gor_1.3.3_mac.tar.gz -O gor.tar.gz
tar -xvf gor.tar.gz && rm gor.tar.gz

# Execute go-replay to listen to traffic
sudo ./gor --input-raw :8090 --output-file=requests.gor

# Simulate 1000 requests
for i in {1..1000};do curl http://localhost:8090/hello; done

# Look into PROD logs inside "./prod/logs/logs.txt". They should be populated.

# Replay the requests in QA environment
./gor --input-file requests_0.gor --output-http="http://localhost:9000"

# Look into QA logs inside "./qa/logs/logs.txt". They should be populated.
```

- If you start again and re-capture replays, you can use the clean-up script.

```shell
sh scripts/cleanup.sh
```


#### References
- https://goreplay.org/