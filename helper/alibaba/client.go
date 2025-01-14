package alibaba

import (
	"errors"
	"regexp"

	"github.com/mitchellh/mapstructure"

	ac "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	au "github.com/alibabacloud-go/openapi-util/service"
	as "github.com/alibabacloud-go/tea-utils/v2/service"
	at "github.com/alibabacloud-go/tea/tea"
)

func Request(rq *Params) (any, error) {

	if ep, err := solveEndpoint(rq); ep != "" {
		rq.Endpoint = ep
	} else {
		return nil, err
	}

	resp, err := newClient(rq)

	if err != nil {
		return nil, getSDKError(err)
	}

	return resp["body"], nil

}

func newClient(rq *Params) (map[string]any, error) {

	config := &ac.Config{
		AccessKeyId:     &rq.SecretId,
		AccessKeySecret: &rq.SecretKey,
		RegionId:        &rq.RegionId,
		Endpoint:        &rq.Endpoint,
	}

	params := &ac.Params{
		Action:      at.String(rq.Action),
		Version:     at.String(rq.Version),
		Protocol:    at.String("HTTPS"),
		Pathname:    at.String("/"),
		Method:      at.String("POST"),
		AuthType:    at.String("AK"),
		Style:       at.String("RPC"),
		ReqBodyType: at.String("json"),
		BodyType:    at.String("json"),
	}

	request := &ac.OpenApiRequest{
		Query: au.Query(rq.Query),
		Body:  rq.Payload,
	}

	runtime := &as.RuntimeOptions{}

	if client, err := ac.NewClient(config); err == nil {
		return client.CallApi(params, request, runtime)
	} else {
		return nil, err
	}

}

func getSDKError(e error) error {

	se := at.SDKError{}

	if mapstructure.Decode(e, &se) != nil {
		return e
	}

	if se.Message == nil {
		return e
	}

	re, _ := regexp.Compile(`^code: \d+, (.+) request id.+$`)
	msg := re.ReplaceAllString(*se.Message, "$1")

	return errors.New(msg)

}
