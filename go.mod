module webserverRaspberry

go 1.13

require (
	github.com/gin-gonic/gin v1.6.3
	github.com/gorilla/websocket v1.4.2
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9
	golang.org/x/sync v0.0.0-20200317015054-43a5402ce75a // indirect
)

replace (
	cloud.google.com/go => github.com/GoogleCloudPlatform/google-cloud-go v0.58.0

	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20200604202706-70a84ac30bf9

	golang.org/x/exp => github.com/golang/exp v0.0.0-20200513190911-00229845015e

	golang.org/x/image => github.com/golang/image v0.0.0-20200609002522-3f4726a040e8

	golang.org/x/lint => github.com/golang/lint v0.0.0-20200302205851-738671d3881b

	golang.org/x/mobile => github.com/golang/mobile v0.0.0-20200329125638-4c31acba0007

	golang.org/x/mod => github.com/golang/mod v0.3.0

	golang.org/x/net => github.com/golang/net v0.0.0-20200602114024-627f9648deb9

	golang.org/x/oauth2 => github.com/golang/oauth2 v0.0.0-20200107190931-bf48bf16ab8d

	golang.org/x/sync => github.com/golang/sync v0.0.0-20200317015054-43a5402ce75a

	golang.org/x/sys => github.com/golang/sys v0.0.0-20200610111108-226ff32320da

	golang.org/x/text => github.com/golang/text v0.3.2

	golang.org/x/time => github.com/golang/time v0.0.0-20200416051211-89c76fbcd5d1

	golang.org/x/tools => github.com/golang/tools v0.0.0-20200612220849-54c614fe050c

	google.golang.org/api => github.com/googleapis/google-api-go-client v0.26.0

	google.golang.org/appengine => github.com/golang/appengine v1.6.6

	google.golang.org/genproto => github.com/ilisin/genproto v0.0.0-20181026194446-8b5d7a19e2d9

	google.golang.org/grpc => github.com/grpc/grpc v1.29.1
)
