package forecastquery_test

// tests for the forecastquery service interface for this ../../../service/forecastquery/iface.go

import (
	"context"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/forecastquery"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/forecastquery/mocks"
	"github.com/nmccready/aws-sdk-go-v2-ifaces/service/forecastquery/forecastquery_iface"
	"github.com/stretchr/testify/assert"
)

func TestForecastqueryServiceCanBeMocked(t *testing.T) {
	var iface forecastquery_iface.IClient
	iface = &forecastquery.Client{}
	assert.NotNil(t, iface, "not nil")
}

func TestIClient(t *testing.T) {
    mockClient := &mocks.IClient{}
    ctx := context.TODO()

    t.Run("TestOptions", func(t *testing.T) {
        output := forecastquery.Options{}
        mockClient.On("Options").Return(output)

        result := mockClient.Options()
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQueryForecast", func(t *testing.T) {
        input := &forecastquery.QueryForecastInput{}
        output := &forecastquery.QueryForecastOutput{}

        mockClient.On("QueryForecast", ctx, input).Return(output, nil)

        result, err := mockClient.QueryForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })

    t.Run("TestQueryWhatIfForecast", func(t *testing.T) {
        input := &forecastquery.QueryWhatIfForecastInput{}
        output := &forecastquery.QueryWhatIfForecastOutput{}

        mockClient.On("QueryWhatIfForecast", ctx, input).Return(output, nil)

        result, err := mockClient.QueryWhatIfForecast(ctx, input)
        assert.NoError(t, err)
        assert.Equal(t, output, result)

        mockClient.AssertExpectations(t)
    })
}
