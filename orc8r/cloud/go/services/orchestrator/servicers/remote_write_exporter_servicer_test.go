package servicers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"magma/orc8r/cloud/go/services/metricsd"
	"magma/orc8r/cloud/go/services/metricsd/exporters"
	tests "magma/orc8r/cloud/go/services/metricsd/test_common"
	"magma/orc8r/cloud/go/services/metricsd/test_init"
	"magma/orc8r/cloud/go/services/orchestrator/protos"
	"magma/orc8r/cloud/go/services/orchestrator/servicers/mocks"
)

var (
	sampleGatewayContext = &exporters.GatewayMetricContext{
		NetworkID: "net1",
		GatewayID: "gw1",
	}
)

func TestMetricTypeToMetricTypeProto(t *testing.T) {
	assert.Equal(t, metricTypeToMetricTypeProto(dto.MetricType_GAUGE), protos.MetricMetadata_GAUGE)
	assert.Equal(t, metricTypeToMetricTypeProto(dto.MetricType_COUNTER), protos.MetricMetadata_COUNTER)
	assert.Equal(t, metricTypeToMetricTypeProto(dto.MetricType_SUMMARY), protos.MetricMetadata_SUMMARY)
	assert.Equal(t, metricTypeToMetricTypeProto(dto.MetricType_HISTOGRAM), protos.MetricMetadata_HISTOGRAM)
	assert.Equal(t, metricTypeToMetricTypeProto(dto.MetricType_UNTYPED), protos.MetricMetadata_UNKNOWN)
}

type exporterTestCase struct {
	name               string
	metrics            []*dto.MetricFamily
	metricContext      exporters.MetricContext
	assertExpectations func(t *testing.T, client *mocks.RemoteWriteClient)
}

func (tc exporterTestCase) RunTest(t *testing.T) {
	// Set client return
	client := &mocks.RemoteWriteClient{}
	r := ioutil.NopCloser(bytes.NewReader([]byte("")))
	client.On("Do", mock.Anything, mock.Anything).Return(&http.Response{
		StatusCode: 200,
		Body:       r,
	}, nil)

	exporter := makeExporter(t, client)

	err := exporter.Submit(tc.metrics, tc.metricContext)
	assert.NoError(t, err)
	tc.assertExpectations(t, client)
}

func TestRemoteWriteExporter(t *testing.T) {
	tcs := []exporterTestCase{
		{
			name:    "submit no metrics",
			metrics: nil,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertNotCalled(t, "Do")
			},
		},
		{
			name:          "submit gauge",
			metrics:       []*dto.MetricFamily{tests.MakeTestMetricFamily(dto.MetricType_GAUGE, 1, []*dto.LabelPair{})},
			metricContext: sampleGatewayContext,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertCalled(t, "Do", mock.Anything, mock.Anything)
				srv.AssertNumberOfCalls(t, "Do", 1)
			},
		},
		{
			name:          "submit counter",
			metrics:       []*dto.MetricFamily{tests.MakeTestMetricFamily(dto.MetricType_COUNTER, 1, []*dto.LabelPair{})},
			metricContext: sampleGatewayContext,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertCalled(t, "Do", mock.Anything, mock.Anything)
				srv.AssertNumberOfCalls(t, "Do", 1)
			},
		},
		{
			name:          "submit untyped",
			metrics:       []*dto.MetricFamily{tests.MakeTestMetricFamily(dto.MetricType_UNTYPED, 1, []*dto.LabelPair{})},
			metricContext: sampleGatewayContext,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertCalled(t, "Do", mock.Anything, mock.Anything)
				srv.AssertNumberOfCalls(t, "Do", 1)
			},
		},
		{
			name:          "submit histogram",
			metrics:       []*dto.MetricFamily{tests.MakeTestMetricFamily(dto.MetricType_HISTOGRAM, 1, []*dto.LabelPair{})},
			metricContext: sampleGatewayContext,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertCalled(t, "Do", mock.Anything, mock.Anything)
				srv.AssertNumberOfCalls(t, "Do", 1)
			},
		},
		{
			name:          "submit summary",
			metrics:       []*dto.MetricFamily{tests.MakeTestMetricFamily(dto.MetricType_SUMMARY, 1, []*dto.LabelPair{})},
			metricContext: sampleGatewayContext,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertCalled(t, "Do", mock.Anything, mock.Anything)
				srv.AssertNumberOfCalls(t, "Do", 1)
			},
		},
		{
			name:          "submit many",
			metrics:       []*dto.MetricFamily{tests.MakeTestMetricFamily(dto.MetricType_GAUGE, 10, []*dto.LabelPair{})},
			metricContext: sampleGatewayContext,
			assertExpectations: func(t *testing.T, srv *mocks.RemoteWriteClient) {
				srv.AssertCalled(t, "Do", mock.Anything, mock.Anything)
				srv.AssertNumberOfCalls(t, "Do", 1)
			},
		},
	}

	for _, tc := range tcs {
		tc.RunTest(t)
	}
}

// makeExporter creates the following
//	- remote write servicer (standalone service)
//	- grpc metrics exporter servicer (standalone service)
//
// The returned exporter forwards to the metrics exporter, which in turn
// forwards to the remote write server.
func makeExporter(t *testing.T, client RemoteWriteClient) exporters.Exporter {
	// writeSrv, lis := test_utils.NewTestService(t, orc8r.ModuleName, remoteWriteControllerServiceName)
	// protos.RegisterRemoteWriterServer(writeSrv.GrpcServer, mockWrite)
	// go writeSrv.RunTest(lis)

	srv := NewRemoteWriteExporterServicer("http://test-remote-write-url:19291", client)
	test_init.StartTestServiceInternal(t, srv)

	exporter := exporters.NewRemoteExporter(metricsd.ServiceName)
	return exporter
}
