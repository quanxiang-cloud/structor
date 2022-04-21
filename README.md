# Structor

## Overview

The structor project is component of QuanXiang, it's a abstract layer between bussiness layer and database, that will make users easy to use database without database knowledge.
We use json to define table's columns data type 

## Getting Started

Structor implements GRPC API for services call. gRPC is useful for low-latency, high performance scenarios and has language integration using the proto clients. If you want to use Structor for your project, refer [GRPC](!https://grpc.io/docs/) for more information about GRPC. Structor can be started up as a service or a sidecar with Kubernetes POD.

### Installation

```
git clone https://github.com/quanxiang-cloud/structor.git
cd structor
CGO_ENABLED=1 GOARCH=amd64 GOOS=Linux go build -o structor cmd/structor/main.go
```
> **Notice:**
>
> The Build example base on Linux-like OS, if you use other OS, please change the GOOS and GOARCH for your case.
>
> - GOOS: darwin, linux, windows, freebsd etc.
> - GOARCH: amd64, arm, i386, p360 etc.

### How to run

For MySQL connection, command example as below:

```
./structor --mysql-host=127.0.0.1:3306 --mysql-database=example --mysql-user=root --mysql-password=***** --mysql-log=false --mysql-log-level=0 --mysql-maxIdelConns=10 --mysql-maxOpenConns=20 --mysql-engine=innoDB --my-charset=utf8 --mysql-collate=utf8_unicode_ci
```

For MongoDB connection, command example as below:

```
./structor --mongo-host=127.0.0.1:27017 --mongo-direct=true --mongo-auth-mechanism=SCRAM-SHA-1 --mongo-auth-source=admin --mongo-password=***** --mongo-password-set=false --mongo-database=example
```

### How to use

Structor supports Data Definition Statements with DDL and Data Manipulation Statements with DSL. We refer ElasticSearch Query DSL to define our DSL, we just support a little of ElasticSearch queries.    

Below code  will be use in all code example,  just replace "// Add  code  here" in your case.

```
package main

import (
  "context"
	"encoding/json"
	"fmt"
	
	client "github.com/quanxiang-cloud/structor/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/structpb"
)

func main() {
  dsl, ddl, err := getConn("localhost:8080")
	if err != nil {
		panic(err)
	}
	// Add code here
	_, _ = dsl, ddl
}

func getConn(addr string) (client.DSLServiceClient, client.DDLServiceClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, nil, err
	}

	dslConn := client.NewDSLServiceClient(conn)
	ddlConn := client.NewDDLServiceClient(conn)
	return dslConn, ddlConn, nil
}

func anyToRaw(any *anypb.Any) (json.RawMessage, error) {
	out := structpb.NewNullValue()
	err := any.UnmarshalTo(out)
	if err != nil {
		return nil, err
	}

	body, err := out.MarshalJSON()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func rawToAny(raw []byte) (*anypb.Any, error) {
	in := structpb.NewNullValue()
	err := in.UnmarshalJSON(raw)
	if err != nil {
		return nil, err
	}
	any := &anypb.Any{}
	err = any.MarshalFrom(in)
	return any, err
}
```

##### Data Definition Statements

Create Table:

```
  _, err := ddl.Create(context.Background(), &client.CreateReq{
		TableName: "idtest",
		Fields: []*client.Field{
			{
				Title:   "bool",
				Type:    "bool",
				Comment: "bool",
			},
		},
	})
	if err != nil {
		panic(err)
	}
```

Add column definition:

```
  _, err := ddl.Add(context.Background(), &client.AddReq{
		TableName: "idtest",
		Fields: []*client.Field{
			{
				Title: "object",
				Type:  "object",
			},
			{
				Title: "date",
				Type:  "datetime",
			},
			{
				Title: "float",
				Max:   321321,
				Type:  "float",
			},
			{
				Title: "varchar",
				Max:   200,
				Type:  "string",
			},
			{
				Title: "text",
				Max:   500,
				Type:  "string",
			},
		},
	})
	if err != nil {
		panic(err)
	}
```

Update table:

```
	_, err := ddl.Modify(context.Background(), &client.ModifyReq{
		TableName: "idtest",
		Fields: []*client.Field{
			{
				Title: "float",
				Max:   321321,
				Type:  "int",
			},
			{
				Title: "varchar",
				Max:   200,
				Type:  "int",
			},
			{
				Title: "text",
				Max:   500,
				Type:  "int",
			},
		},
	})
	if err != nil {
		panic(err)
	}
```

##### Data Manipulation Statements

 Query match records :

```
  queryBody, err := rawToAny([]byte(`{
		"query": {
			"match": {
				"name": "st"
			}
		}
	}`))
	if err != nil {
		panic(err)
	}

	resp, err := dsl.FindOne(context.Background(), &client.FindOneReq{
		TableName: "user",
		Dsl:       queryBody,
	})
	if err != nil {
		panic(err)
	}

	body, err := anyToRaw(resp.GetData())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
```

Query records:

```
  any, err := rawToAny([]byte(`
	{
		"query": {
			"bool": {
				"should": [
					{
						"term": {
							"id": "129"
						}
					},
					{
						"bool": {
							"should": [
								{
									"term": {
										"name": "test2"
									}
								}
							]
						}
					}
				]
			}
		}
	}
	`))
	if err != nil {
		panic(err)
	}

	resp, err := dsl.Find(context.Background(), &client.FindReq{
		TableName: "user",
		Dsl:       any,
		Sort:      []string{"id", "-name"},
	})
	if err != nil {
		panic(err)
	}

	data, err := anyToRaw(resp.GetData())
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
```

Delete records:

```
	delBody, err := rawToAny([]byte(`{
		"query": {
			"range": {
				"id": {
					"gt": 131,
					"lt": 135
				}
			}	
		}
	}`))
	if err != nil {
		panic(err)
	}

	resp, err := dsl.Delete(context.Background(), &client.DeleteReq{
		TableName: "user",
		Dsl:       delBody,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Count)
```

Count records:

```
	cuBody, err := rawToAny([]byte(`{
		"query": {
			"match": {
				"name": "test"
			}
		}
	}`))
	if err != nil {
		panic(err)
	}

	resp, err := dsl.Count(context.Background(), &client.CountReq{
		TableName: "user",
		Dsl:       cuBody,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Data)
```

Insert record:

```
	entity, err := rawToAny([]byte(`{
		"id": "136"
	}`))
	if err != nil {
		panic(err)
	}

	entity2, err := rawToAny([]byte(`{
		"id": "137"
	}`))
	if err != nil {
		panic(err)
	}

	var entities []*anypb.Any
	entities = append(entities, entity, entity2)
	_, err = dsl.Insert(context.Background(), &client.InsertReq{
		TableName: "new_user",
		Entities:  entities,
	})
	if err != nil {
		panic(err)
	}
```

Update records

```
	updateBody, err := rawToAny([]byte(`{
		"query": {
			"term": {
				"name": "test"
			}
		}
	}`))
	if err != nil {
		panic(err)
	}

	entity, err := rawToAny([]byte(`{
		"name": "test909090"
	}`))
	if err != nil {
		panic(err)
	}

	resp, err := dsl.Update(context.Background(), &client.UpdateReq{
		TableName: "user",
		Dsl:       updateBody,
		Entity:    entity,
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Count)
```

## Roadmap

Folowing DataBases are supported.

| Database | CRUD | index | primary key| foreign key| version|
|---|----|----|---|---|---|
| MySQL| [x] | [x] | [x] | [x] | 5.7 or later |
| MongoDB | [x] | [x] | | | 4.0 or later |

Plan to support:

| Database | CRUD | index | primary key| foreign key| version|
|---|----|----|---|---|---|
| Oracle | [] | [] | [] | [] |  |
| PostgreSQL | [] | [] | [] | [] |  |
| TiDB | [] | [] | [] | [] |  |
| More | [] | [] | [] | [] |  |

## Contributing to Structor

Welcome!!!!
