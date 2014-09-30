# the registry api

```
$ brew install elasticsearch
```

## setup the registry index

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
