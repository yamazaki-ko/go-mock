# go-architecture
## golang reccomend architecture

<pre>
.
├── main.go
├── go.mod
├── go.sum
|
├── cmd
└── lib
└── app
    └── app
    └── config
    └── domain
    |    └── entity
    |    └── repository
    └── ui  
    |    └── cli
    |    └── http     
    |        └── middleware
    |        └── rest 
    └── infra
        └── db
</pre>

## golang with lamda reccomend architecture 

<pre>
.
├── go.mod
├── go.sum
├── serverless.yml
|
├── config
└── lib
└── app
    └── domain
    |    └── model
    |    └── repository
    └── handler
    |    └── cmd
    |    └── parametor
    └── usecase  
    └── infra
        └── mysql
        └── s3
        └── ex...
</pre>
