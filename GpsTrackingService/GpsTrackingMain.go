package main

import (
	"crypto/tls"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/iot"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	svc := iot.New(sess, aws.NewConfig().WithRegion("us-west-2"))
	fmt.Println(svc)

	cer, err := tls.LoadX509KeyPair("path/f8527f7924-certificate.pem", "path/f8527f7924-private.pem.key")
	if err!=nil{

	}

}

