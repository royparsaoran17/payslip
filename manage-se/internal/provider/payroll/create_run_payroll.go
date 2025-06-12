package payroll

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"manage-se/internal/presentations"
	"manage-se/internal/provider/providererrors"
	"manage-se/pkg/httpx"
	"net/http"
)

func (c *client) RunPayroll(ctx context.Context, input presentations.RunPayroll) error {
	urlEndpoint := c.endpoint("/internal/v1/run-payroll")

	var request bytes.Buffer
	err := json.NewEncoder(&request).Encode(input)
	if err != nil {
		return errors.Wrap(err, "new encoder encode")
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, urlEndpoint, &request)
	if err != nil {
		return errors.Wrap(err, "new request with context")
	}

	req.Header.Set(httpx.ContentType, httpx.MediaTypeJSON)

	res, err := c.dep.HttpClient.Do(req)
	if err != nil {
		return errors.Wrap(err, "do request")
	}

	rawBody, _ := io.ReadAll(res.Body)
	res.Body.Close() // must close
	res.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch res.StatusCode {
	case http.StatusCreated, http.StatusOK:
		return nil

	default:
		bodyErr := providererrors.Error{}
		err := json.Unmarshal(rawBody, &bodyErr)
		if err != nil {
			return providererrors.NewErrRequestWithResponse(req, res)
		}

		bodyErr.Code = res.StatusCode
		return bodyErr

	}
}
