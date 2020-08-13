# go-tdameritrade
go client for the tdameritrade api

[![Documentation](https://godoc.org/github.com/JonCooperWorks/go-tdameritrade?status.svg)](https://godoc.org/github.com/JonCooperWorks/go-tdameritrade)


```import "github.com/JonCooperWorks/go-tdameritrade"```

go-tdameritrade handles all interaction with the [TD Ameritrade REST API](https://developer.tdameritrade.com/apis).
See the TD Ameritrade [developer site](https://developer.tdameritrade.com/) to learn how their APIs work.
This is a very thin wrapper and does not perform any validation.
go-tdameritrade doesn't support streaming yet.


## Authentication with TD Ameritrade
There is an example of using OAuth2 to authenticate a user and use the services on the TD Ameritrade API in [examples/webauth/webauth.go](https://github.com/JonCooperWorks/go-tdameritrade/blob/master/examples/webauth/webauth.go).
Authentication is handled by the ```Authenticator``` struct and its methods ```StartOAuth2Flow``` and ```FinishOAuth2Flow``` (see [auth.go](https://github.com/JonCooperWorks/go-tdameritrade/blob/master/auth.go)).

Use at your own risk.
