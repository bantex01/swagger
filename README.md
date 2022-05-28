# Go API client for swagger

API to receive streamed updates. This is an ssl socket connection of CRLF delimited json messages (see RequestMessage & ResponseMessage)

## Overview
This API client was generated by the [swagger-codegen](https://github.com/swagger-api/swagger-codegen) project.  By using the [swagger-spec](https://github.com/swagger-api/swagger-spec) from a remote server, you can easily generate an API client.

- API version: 1.0.1423
- Package version: 1.0.0
- Build date: 2022-05-28T08:31:41.818+01:00
- Build package: class io.swagger.codegen.languages.GoClientCodegen
For more information, please visit [https://developer.betfair.com/support/](https://developer.betfair.com/support/)

## Installation
Put the package under your project folder and add the following in import:
```
    "./swagger"
```

## Documentation for API Endpoints

All URIs are relative to *http://stream-api.betfair.com:443/api*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*DefaultApi* | [**RequestPost**](docs/DefaultApi.md#requestpost) | **Post** /request | 


## Documentation For Models

 - [AllRequestTypesExample](docs/AllRequestTypesExample.md)
 - [AllResponseTypesExample](docs/AllResponseTypesExample.md)
 - [AuthenticationMessage](docs/AuthenticationMessage.md)
 - [ConnectionMessage](docs/ConnectionMessage.md)
 - [HeartbeatMessage](docs/HeartbeatMessage.md)
 - [KeyLineDefinition](docs/KeyLineDefinition.md)
 - [KeyLineSelection](docs/KeyLineSelection.md)
 - [MarketChange](docs/MarketChange.md)
 - [MarketChangeMessage](docs/MarketChangeMessage.md)
 - [MarketDataFilter](docs/MarketDataFilter.md)
 - [MarketDefinition](docs/MarketDefinition.md)
 - [MarketFilter](docs/MarketFilter.md)
 - [MarketSubscriptionMessage](docs/MarketSubscriptionMessage.md)
 - [Order](docs/Order.md)
 - [OrderChangeMessage](docs/OrderChangeMessage.md)
 - [OrderFilter](docs/OrderFilter.md)
 - [OrderMarketChange](docs/OrderMarketChange.md)
 - [OrderRunnerChange](docs/OrderRunnerChange.md)
 - [OrderSubscriptionMessage](docs/OrderSubscriptionMessage.md)
 - [PriceLadderDefinition](docs/PriceLadderDefinition.md)
 - [RequestMessage](docs/RequestMessage.md)
 - [ResponseMessage](docs/ResponseMessage.md)
 - [RunnerChange](docs/RunnerChange.md)
 - [RunnerDefinition](docs/RunnerDefinition.md)
 - [StatusMessage](docs/StatusMessage.md)
 - [StrategyMatchChange](docs/StrategyMatchChange.md)


## Documentation For Authorization

 All endpoints do not require authorization.


## Author

bdp@betfair.com

