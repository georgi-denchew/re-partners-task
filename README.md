# RE Partners challenge

This repository contains the backend application for the RE Partners Software Engineering Challenge.

The application is written in Golang and exposes an HTTP API for interaction.
It is accessible through <http://3.70.72.6:3000> (example usage described bellow)

## Endpoints

### `POST /orders`

Creates a new order

#### Body Parameters

> | name        | type     | data type | description                                 |
> |-------------|----------|-----------|---------------------------------------------|
> | items_count | required | Number    | Number of items the API user wants to order |

#### Responses

> | http code | content-type       | response                                                   |
> |-----------|--------------------|------------------------------------------------------------|
> | `201`     | `application/json` | `{"item_packs": [{"items": 250,"packs": 1}]}`              |
> | `400`     | `application/json` | `{"message": "field 'items_count' cannot be less than 1"}` |

#### Example cURL

```sh
curl -X POST -H "Content-Type: application/json" --data '{"items_count": 4200}' http://3.70.72.6:3000/orders
```

#### Example Response

```json
{
  "item_packs": [
    {
      "items":250,
      "packs":1
    },
    {
      "items":2000,
      "packs":2
    }
  ]
}
```

### `GET /orders`

Retrieves orders

#### Query Parameters

> | name        | type     | data type | description                                     |
> |-------------|----------|-----------|-------------------------------------------------|
> | count       | optional | Number    | Number of orders the API user wants to retrieve |

#### Responses

> | http code | content-type       | response                                                   |
> |-----------|--------------------|------------------------------------------------------------|
> | `200`     | `application/json` | `{"item_packs": [{"items": 250,"packs": 1}]}`              |
> | `400`     | `application/json` | `{"message": "field 'items_count' cannot be less than 1"}` |

#### Example cURL

```sh
curl -X POST -H "Content-Type: application/json" --data '{"items_count": 4200}' http://3.70.72.6:3000/orders
```

#### Example Response

```json
{
  "item_packs": [
    {
      "items":250,
      "packs":1
    },
    {
      "items":2000,
      "packs":2
    }
  ]
}
```

## Local setup

### Prerequisites

- clone the repository
- set `PACK_SIZES` environment variable which specifies the sizes for available item packs, e.g.

  ```sh
  PACK_SIZES=5000,250,500,1000,2000
  ```

- you can also optionally provide `PORT` to specify which port will the application listen on

### Run application

## Implementation decisions

The following is a list of purposefully made decisions:

- the application is containerized using Docker and hosted on AWS ECS hosted
- the frontend is kept in the same repository for simplicity
  - it is also containerized and runs on ECS
- `PACK_SIZES` are configurable through an environment variable
- `PACK_SIZES` can also be replaced using `PUT /admin/packs`
- unit tests run as part of the image build process
- different unit testing practices have been utilized to showcase variety
  - some are written as whitebox tests, where the package name doesn't include the `_test` suffix, allowing the test to access private fields and functions
  - some are written as blackbox tests
  - some utilize testify's `suite.Suite`
- apart from the use of [gomock](https://github.com/uber-go/mock), no other code generation tools were used so the reader could evaluate the author's coding style
  - some useful tools would have been [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) and [cleanenv](https://github.com/ilyakaznacheev/cleanenv)

## List of potential improvements

- replace in-memory persistence with a Database and implement a repository for DB interactions
- CI/CD pipeline
- add an OpenAPI spec for API definition
  - could also be used with `oapi-codegen` to generate boilerplate types and http handlers
- there's a `createdAt` field set against each stored order which isn't displayed on the front-end
