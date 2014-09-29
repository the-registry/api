# setup

$ brew install elasticsearch

## setup index

PUTS /registry

POST /registry/_close

PUT /register/_settings
{
   "settings": {
       "analysis" : {
            "analyzer" : {
                "default" : {
                    "tokenizer" : "whitespace",
                    "filter" : ["lowercase"]
                }
            }
        }
   }
}

POST /registry/_open
