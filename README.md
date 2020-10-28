## Amazgo:
REST APIs for Amazon scraping.\
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

