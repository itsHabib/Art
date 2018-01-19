# Art - API Resource Tool

## Overview
Art is a command line tool that can be used to test an API to ensure the proper
responses are sent back. Given a base URL the API struct maps resources to specific
RestResources. Each RestResource also contains a Router which maps status codes to
callback response functions. 

## Usage
Testing an API with a base URL of https://httpbin.org and an enpoint of /get.
The callback response function in this example just returns a nil error when a 200 status code
is given and also prints the contents from the request. 
