# the registry api

## developing setup

```
$ brew install elasticsearch
```

```
curl -XPOST localhost:9200/registry -d '{
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
```

```
$ go get github.com/tools/godep
$ godep restore
```

## run locally

```
$ go run main.go
```

## api endpoints

```
POST /types/:type/packages
GET /types/:type/packages
GET /types/:type/packages/:name
DELETE /types/:type/packages/:name
GET /types/:type/packages/search
```
