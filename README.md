# parser
parses a text file in a specific format into a dataset.  Not general enough to be useful to others at this time.

Build and run
```
docker build --tag parser .; docker run -v "$(pwd)/data":/data parser
```