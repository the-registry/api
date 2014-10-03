# the-registry/api

The Registry web service.

## API

### GET

- __GET__ /types/:type/packages - all packages of a type
- __GET__ /types/:type/packages/:name - info of the package
- __GET__ /types/:type/packages/search - search for packages

### POST

- __POST__ /types/:type/packages - create a package

### DELETE

- __DELETE__ /types/:type/packages/:name - delete the package

## Developing The Registry

```
$ mkdir -p ~/dev/the-registry
$ cd ~/dev/the-registry
$ git clone https://github.com/the-registry/api.git
$ cd api
$ brew install elasticsearch
$ curl -XPOST localhost:9200/registry -d '{
   "settings": {
      "analysis": {
         "analyzer": {
            "package_name": {
               "tokenizer": "whitespace",
               "filter": [
                  "lowercase",
                  "word_delimiter"
               ]
            }
         }
      }
   },
   "mappings": {
      "packages": {
         "properties": {
            "name": {
               "type": "multi_field",
               "fields": {
                  "name": {
                     "type": "string",
                     "index": "analyzed",
                     "analyzer": "package_name"
                  },
                  "untouched": {
                     "type": "string",
                     "index": "not_analyzed"
                  }
               }
            }
         }
      }
   }
}'
$ go get github.com/tools/godep
$ godep restore
$ go run main.go
```
