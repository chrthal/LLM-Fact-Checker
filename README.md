# llm_fact_checker

## API Keys
- Get a google [custom search api key](https://developers.google.com/custom-search/v1/introduction)
- Create a [custom search engine](https://programmablesearchengine.google.com/controlpanel/all)
- Create a [bing search azure ressource](https://portal.azure.com/#create/Microsoft.BingSearch)
- Create a [OpenAI API Key](https://platform.openai.com/account/api-keys)

### Set the following environment variables:
- ```OPENAI_API_KEY = <API_KEY> ```
- ```GCS_KEY = <API_KEY> ```
- ```GCS_ID = <ID> ```
- ```BING_KEY = <API_KEY> ```
- ```OLLAMA_MODEL = <MODEL``` https://ollama.com/library
    
## Start with Makefile
- ```make app```

## Docker Build 
- ```docker-compose build```
- ```docker-compose up```

## App
- ```http://localhost:8080/```

