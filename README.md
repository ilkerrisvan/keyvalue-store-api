![Go](https://img.shields.io/badge/Go-1.17.3-f21170?style=flat-square&logo=docker&logoColor=white)
![Docker](https://img.shields.io/badge/Docker-3.3.2-f21170?style=flat-square&logo=docker&logoColor=white)

# KeyValue Store Rest API

With this API, an inmemory key value store can be used using endpoints. Endpoints are described in API Doc. The key
value store is saved in a JSON file every minute under temp. When the API starts working, if there is already a JSON
file, the data in this file is saved in memory.

## Table of Contents:

- [Getting Started](#getting-started)
    - [Requirements](#requirements)
    - [Building with Docker](#with-docker)
- [API Endpoints and Documentations](#api-endpoints-and-documentations)
    - [GET `/api/get`](#get-allkeyvaluepairs)
    - [POST `/api/set`](#post-keyvaluepair)
    - [GET `/api/get-all`](#get-keyvaluepairs)
    - [GET `/api/delete`](#get-keyvaluepairs)
    - [GET `/api/flush`](#get-keyvaluepairs)
- [Contact Information](#contact-information)
- [License](#license)

<br/>

## Getting Started

<hr/>

### Requirements:

<hr/>

- Go v1.17.3 or higher -> [Go Installation Page](https://go.dev/dl/)
- Docker v3.3.2 or higher (optional) -> [Docker Get Started Page](https://www.docker.com/get-started)
  <br/>

  Before starting the application, fork/download/clone this repo.

<hr/>

### Building with Docker

- To run the application on [localhost:8000](http://localhost:8000):

```
docker-compose up --build
```


## API Endpoints and Documentations

<hr/>

### GET `/api/get?key=<foo>`

- Description: The key's value returns in response.


#### Request:

```
GET Request to '/api/get?key=foo' endpoint. //Foo means name of key, desired value can be entered.
```


#### Reponse:

```
{
    "key": "foo",
    "value": "bar",
    "status": "OK"
}
```

If the key is not used:

```
{
    "message": "Bad Request. The URL may be an incorrect or there may not be a value for the key value."
}
```

| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 200    | Success| The value of the key was searched |
| 400     | Bad Request       |   URL, JSON structure or request method is wrong |
<hr/>


### POST `/api/set`

- Description: Creates new key-value pair.

#### Request Body:

```
{
    "pair": [
        {
            "key": "foo",
            "value": "bar"
        },
         {
            "key": "yemek",
            "value": "sepeti"
        }
    ]
}
```

#### Reponse:

```
{
    "pair": [
        {
            "key": "foo",
            "value": "bar",
            "status": "Saved."
        },
        {
            "key": "yemek",
            "value": "sepeti",
            "status": "Saved."
        }
    ]
}
```
If the key-value store is already used.
```
{
    "pair": [
        {
            "key": "foo",
            "value": "bar",
            "status": "Key not saved. It was used or one of the tags is wrong."
        },
        {
            "key": "yemek",
            "value": "sepeti",
            "status": "Key not saved. It was used or one of the tags is wrong."
        }
    ]
}
```
| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 201    | Created| The key-value pair created |
| 400     | Bad Request       |    URL, JSON structure or request method is wrong |
<hr/>

### GET `/api/get-all`

- Description: Returns information about all key-values.
#### Request:

```
GET Request to '/api/get-all'
```

#### Reponse:

```
{
    "foo": "bar",
    "yemek": "sepeti"
}
```
If there is no key-value pair.
```
{
    "status": "There is no pair."
}
```
| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 200    | Success|Provided information about all key/value pairs |
| 400     | Bad Request       |  Request method is wrong |
<hr/>

### GET `/api/delete?key=<foo>`

- Description: Deletes  key-value pair.


#### Request:

```
GET Request to '/api/delete?key=foo'
```

#### Reponse:

```
{
    "status": "The key is deleted"
}
```

If there is no key-value pair:

```
{
    "message": "Bad Request. The URL may be an incorrect or there may not be a value for the key value."
}
```


| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 200    | Success| The desired key-value pair has been deleted|
| 400     | Bad Request       |  Request method is wrong or the key is not using|
<hr/>


### GET `/api/flush`

- Description: Deletes all key-value pairs.



#### Request:


```
GET Request to '/api/flush' endpoint.
```



#### Reponse:

```
{
    "status": "All datas are deleted"
}
```


| Status Code  | HTTP Meaning | API Meaning |
| :------------ |:---------------:| -----:|
| 200    | Success| All pairs deleted|
| 400     | Bad Request       |   Wrong request method|
|

## Contact Information

<hr/>

#### Author: İlker Rişvan

#### Github: ilkerrisvan

#### Email: ilkerrisvan@outlook.com

#### Date: December, 2021

## License

<hr/>

[MIT](https://choosealicense.com/licenses/mit/)
