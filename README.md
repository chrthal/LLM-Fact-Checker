# llm_fact_checker

## API Keys
Get a google [custom search api key](https://developers.google.com/custom-search/v1/introduction)
Create a [custom search engine](https://programmablesearchengine.google.com/controlpanel/all)
Create a [bing search azure ressource](https://portal.azure.com/#create/Microsoft.BingSearch)
Set the following environment variables:
- ```GCS_KEY = <API_KEY> ```
- ```GCS_ID = <ID> ```

## Start with Makefile
- ```make app```

## Docker Build 
- ```docker-compose build```
- ```docker-compose up```

## App
- ```http://localhost:8080/```

