// AWS Cognito client
package client

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/amplify"
)

type Client struct {
	*amplify.Amplify
}

// New inits the client.
func New(region string) *Client {
	return &Client{
		amplify.New(
			session.Must(session.NewSession()),
			aws.NewConfig().WithRegion(region),
		),
	}
}

func putZip(d io.Reader, url string) error {
	c := http.Client{Timeout: 5 * time.Second}
	req, err := http.NewRequest("PUT", url, d)
	if err != nil {
		return err
	}
	res, err := c.Do(req)
	if err != nil {
		return err
	}
	if fmt.Sprintf("%d", res.StatusCode)[:1] != "2" {
		return fmt.Errorf("zip submission error: %d", res.StatusCode)
	}
	return nil
}

func (c *Client) DeployZip(ctx context.Context, appID, branch, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	s, err := c.CreateDeploymentWithContext(ctx,
		&amplify.CreateDeploymentInput{
			AppId:      aws.String(appID),
			BranchName: aws.String(branch),
		},
	)
	if err != nil {
		return err
	}
	if err := putZip(f, *s.ZipUploadUrl); err != nil {
		return err
	}
	_, err = c.StartDeployment(&amplify.StartDeploymentInput{
		AppId:      aws.String(appID),
		BranchName: aws.String(branch),
		JobId:      s.JobId,
	})
	if err != nil {
		return err
	}
	return nil
}
