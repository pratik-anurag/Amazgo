## Amazgo:
REST APIs for Amazon scraping using colly.
\
Contains POST (REST) API in Golang
Which scrapes an Amazon web page given its URL
The code should scrape following data from the page
- Product Name/Title
- Product image url
- Product description
- Product price
- Product total number of reviews.
\
API handler takes the JSON structure from the request payload,\
Handler writes the obtained payload to a new/updated Document,\
Also write timestamp of the create/update within the document.\

To run:-
```
git clone https://github.com/pratik-anurag/Amazgo.git
cd Amazego
docker-compose up
```
\
Sample cURL to insert a single product in mongoDB:-
```
curl -X POST \
  http://localhost:8000/products \
  -H 'cache-control: no-cache' \
  -H 'content-type: application/json' \
  -H 'postman-token: c6a68fb0-2b20-db8a-d3bc-1bf189dbd567' \
  -d '{
        "Name": "PlayStation 2 Slim Console PS2 Renewed",
        "ImageUrl": "https://m.media-amazon.com/images/I/4116HYcMH0L._AC_UY218_.jpg",
        "TotalReview": "273",
        "Description": "test",
        "Price": "209.99"
    }'
```
/

Sample cURL to fetch document from given url and insert it to DB:-
```
curl -X POST \
  'http://localhost:8000/scrapes?url=https%3A%2F%2Fwww.amazon.com%2Fs%3Fk%3Dps2%26ref%3Dnb_sb_noss_2' \
  -H 'cache-control: no-cache' \
  -H 'postman-token: 7a628957-c693-03c0-58cc-0578b677e87e'
```

