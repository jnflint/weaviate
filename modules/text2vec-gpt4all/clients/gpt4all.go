//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2025 Weaviate B.V. All rights reserved.
//
//  CONTACT: hello@weaviate.io
//

package clients

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/weaviate/weaviate/modules/text2vec-gpt4all/ent"
)

type client struct {
	origin     string
	httpClient *http.Client
	logger     logrus.FieldLogger
}

func New(origin string, timeout time.Duration, logger logrus.FieldLogger) *client {
	return &client{
		origin: origin,
		httpClient: &http.Client{
			Timeout: timeout,
		},
		logger: logger,
	}
}

func (c *client) Vectorize(ctx context.Context, text string) (*ent.VectorizationResult, error) {
	body, err := json.Marshal(vecRequest{text})
	if err != nil {
		return nil, errors.Wrapf(err, "marshal body")
	}

	req, err := http.NewRequestWithContext(ctx, "POST", c.url("/vectorize"),
		bytes.NewReader(body))
	if err != nil {
		return nil, errors.Wrap(err, "create POST request")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "send POST request")
	}
	defer res.Body.Close()

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read response body")
	}

	var resBody vecResponse
	if err := json.Unmarshal(bodyBytes, &resBody); err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("unmarshal response body. Got: %v", string(bodyBytes)))
	}

	if res.StatusCode != 200 {
		if resBody.Error != "" {
			return nil, errors.Errorf("failed with status: %d error: %v", res.StatusCode, resBody.Error)
		}
		return nil, errors.Errorf("failed with status: %d", res.StatusCode)
	}

	return &ent.VectorizationResult{
		Vector:     resBody.Vector,
		Dimensions: resBody.Dim,
		Text:       resBody.Text,
	}, nil
}

func (c *client) url(path string) string {
	return fmt.Sprintf("%s%s", c.origin, path)
}

type vecRequest struct {
	Text string `json:"text"`
}

type vecResponse struct {
	Text   string    `json:"text"`
	Vector []float32 `json:"vector"`
	Dim    int       `json:"dim"`
	Error  string    `json:"error"`
}
