package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func main() {

	if len(os.Args) > 1 {

		// profile := "default"
		region := "us-east-1"

		conf := &aws.Config{
			Region: &region,
		}

		sess, err := session.NewSession(conf)
		fatalErr(err)
		_sts := sts.New(sess)

		arg := os.Args[1]

		sn := "arn:aws:iam::500531898377:mfa/bartlomiej.specjalny@reeftechnology.com"

		res, err := _sts.GetSessionToken(&sts.GetSessionTokenInput{
			TokenCode:    &arg,
			SerialNumber: &sn,
		})

		fatalErr(err)

		fmt.Println("SET AWS_ACCESS_KEY_ID=" + *res.Credentials.AccessKeyId)
		fmt.Println("SET AWS_SESSION_TOKEN=" + *res.Credentials.SessionToken)
		fmt.Println("SET AWS_SECRET_ACCESS_KEY=" + *res.Credentials.SecretAccessKey)

	} else {
		log.Fatal("no auth code")
	}

}

func fatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
