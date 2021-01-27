/*
Copyright 2020 The Magma Authors.
This source code is licensed under the BSD-style license found in the
LICENSE file in the root directory of this source tree.
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package servicers

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/gogo/protobuf/proto"
	"github.com/golang/glog"
	"github.com/golang/snappy"

	"magma/orc8r/cloud/go/services/metricsd/protos"
	protos2 "magma/orc8r/cloud/go/services/orchestrator/protos"
	"magma/orc8r/lib/go/registry"

	"github.com/pkg/errors"
	dto "github.com/prometheus/client_model/go"
)

// RemoteWriteExporterServicer pushes metrics to a datasink via the Prometheus
// Remote Write interface. Here we convert metrics from the old protobuf types
// to the new one, which consists of only a set of metadata (metric names, types,
// and help text), and a list of TimeSeries which is just labels + samples.
type RemoteWriteExporterServicer struct {
	// registry is the servicer's local service registry.
	// Local registry since the gRPC servicer is not a proper Orchestrator
	// service.
	registry *registry.ServiceRegistry
	pushAddr string
	client   RemoteWriteClient
}

func NewRemoteWriteExporterServicer(pushAddr string, client RemoteWriteClient) protos.MetricsExporterServer {
	glog.Errorf("New Remote Write Exporter: %s", pushAddr)
	srv := &RemoteWriteExporterServicer{registry: registry.NewWithMode(registry.YamlRegistryMode), pushAddr: pushAddr, client: client}
	srv.registry.AddService(registry.ServiceLocation{Name: serviceName, Host: pushAddr})
	return srv
}

func (s *RemoteWriteExporterServicer) Submit(ctx context.Context, req *protos.SubmitMetricsRequest) (*protos.SubmitMetricsResponse, error) {
	metricsToSend := preprocessRemoteWriteMetrics(req.GetMetrics())
	if len(metricsToSend.Timeseries) == 0 {
		return &protos.SubmitMetricsResponse{}, nil
	}
	var shardKey string
	switch metCtx := req.GetContext().(type) {
	case *protos.SubmitMetricsRequest_GatewayContext:
		glog.Errorf("Submitting %d Gateway Metrics", len(metricsToSend.Timeseries))
		shardKey = metCtx.GatewayContext.GatewayId
	case *protos.SubmitMetricsRequest_CloudContext:
		glog.Errorf("Submitting %d Cloud Metrics", len(metricsToSend.Timeseries))
		shardKey = metCtx.CloudContext.CloudHost
	case *protos.SubmitMetricsRequest_PushedContext:
		glog.Errorf("Submitting %d Pushed Metrics", len(metricsToSend.Timeseries))
		shardKey = metCtx.PushedContext.NetworkId
	default:
		glog.Errorf("Unknown metric context type: %T", metCtx)
	}

	err := s.pushFamilies(metricsToSend, s.client, shardKey)
	return &protos.SubmitMetricsResponse{}, err
}

func (s *RemoteWriteExporterServicer) pushFamilies(writeReq *protos2.WriteRequest, client RemoteWriteClient, shardKey string) error {
	glog.Error("Pushing families")
	writeData, err := packageWriteData(writeReq, nil)
	if err != nil {
		glog.Errorf("%v\n", err)
		return err
	}
	err = makeRequest(writeData, client, s.pushAddr)
	if err != nil {
		glog.Errorf("Error writing: %v\n", err)
		return err
	}
	return nil
}

func packageWriteData(writeReq *protos2.WriteRequest, buf []byte) ([]byte, error) {
	data, err := proto.Marshal(writeReq)
	if err != nil {
		return nil, fmt.Errorf("error marshaling write request: %v", err)
	}

	// snappy uses len() to see if it needs to allocate a new slice. Make the
	// buffer as long as possible.
	if buf != nil {
		buf = buf[0:cap(buf)]
	}
	compressed := snappy.Encode(buf, data)
	return compressed, nil
}

func preprocessRemoteWriteMetrics(families []*dto.MetricFamily) *protos2.WriteRequest {
	if families == nil {
		return &protos2.WriteRequest{}
	}
	fmt.Printf("Preprocessing %d families", len(families))
	mdata := buildMetadata(families)
	series := familiesToTimeSeries(families)
	return &protos2.WriteRequest{Metadata: mdata, Timeseries: series}
}

func buildMetadata(families []*dto.MetricFamily) []*protos2.MetricMetadata {
	mdata := make([]*protos2.MetricMetadata, 0, len(families))
	for _, fam := range families {
		mdata = append(mdata, &protos2.MetricMetadata{
			Type:             metricTypeToMetricTypeProto(fam.GetType()),
			MetricFamilyName: fam.GetName(),
			Help:             fam.GetHelp(),
		})
	}
	return mdata
}

func familiesToTimeSeries(families []*dto.MetricFamily) []*protos2.TimeSeries {
	timeSeries := []*protos2.TimeSeries{}
	for _, fam := range families {
		timeSeries = append(timeSeries, buildSeries(fam)...)
	}
	return timeSeries
}

func buildSeriesLabels(met *dto.Metric, name string) []*protos2.Label {
	labels := make([]*protos2.Label, 0, len(met.Label)+1)
	labels = append(labels, &protos2.Label{
		Name:  "__name__",
		Value: name,
	})
	for _, label := range met.Label {
		labels = append(labels, &protos2.Label{
			Name:  *label.Name,
			Value: *label.Value,
		})
	}
	return labels
}

func makeRequest(req []byte, client RemoteWriteClient, url string) error {
	httpReq, err := http.NewRequest("POST", url, bytes.NewReader(req))
	if err != nil {
		return err
	}
	httpReq.Header.Add("Content-Encoding", "snappy")
	httpReq.Header.Set("Content-Type", "application/x-protobuf")
	httpReq.Header.Set("X-Prometheus-Remote-Write-Version", "0.1.0")
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()
	httpReq = httpReq.WithContext(ctx)

	httpResp, err := client.Do(httpReq)
	if err != nil {
		return err
	}
	resp, err := ioutil.ReadAll(httpResp.Body)
	fmt.Printf("RESPONSE: %s\n", string(resp))
	defer func() {
		io.Copy(ioutil.Discard, httpResp.Body)
		httpResp.Body.Close()
	}()
	if httpResp.StatusCode/100 != 2 {
		scanner := bufio.NewScanner(io.LimitReader(httpResp.Body, 512))
		line := ""
		if scanner.Scan() {
			line = scanner.Text()
		}
		err = errors.Errorf("server returned HTTP status %s: %s", httpResp.Status, line)
	}
	return nil
}

func buildSeries(fam *dto.MetricFamily) []*protos2.TimeSeries {
	switch fam.GetType() {
	case dto.MetricType_GAUGE:
		return buildSingleValueSamples(fam)
	case dto.MetricType_COUNTER:
		return buildSingleValueSamples(fam)
	case dto.MetricType_HISTOGRAM:
		return buildHistogramSamples(fam)
	case dto.MetricType_SUMMARY:
		return buildSummarySamples(fam)
	case dto.MetricType_UNTYPED:
		return buildSingleValueSamples(fam)
	default:
		return []*protos2.TimeSeries{}
	}
}

func buildSingleValueSamples(fam *dto.MetricFamily) []*protos2.TimeSeries {
	series := make([]*protos2.TimeSeries, 0, len(fam.Metric))
	metType := fam.GetType()
	for _, met := range fam.Metric {
		var val float64
		switch metType {
		case dto.MetricType_GAUGE:
			if met.Gauge == nil {
				continue
			}
			val = met.Gauge.GetValue()
		case dto.MetricType_COUNTER:
			if met.Counter == nil {
				continue
			}
			val = met.Counter.GetValue()
		case dto.MetricType_UNTYPED:
			if met.Untyped == nil {
				continue
			}
			val = met.Untyped.GetValue()
		}
		series = append(series, &protos2.TimeSeries{
			Labels: buildSeriesLabels(met, fam.GetName()),
			Samples: []*protos2.Sample{{
				Value:     val,
				Timestamp: met.GetTimestampMs(),
			}},
		})
	}
	return series
}

func buildHistogramSamples(fam *dto.MetricFamily) []*protos2.TimeSeries {
	series := []*protos2.TimeSeries{}
	baseName := fam.GetName()
	for _, met := range fam.Metric {
		if met.Histogram == nil {
			continue
		}
		// SampleCount
		series = append(series, &protos2.TimeSeries{
			Labels: buildSeriesLabels(met, fmt.Sprintf("%s_count", baseName)),
			Samples: []*protos2.Sample{{
				Value:     float64(*met.Histogram.SampleCount),
				Timestamp: met.GetTimestampMs(),
			}},
		})
		// SampleSum
		series = append(series, &protos2.TimeSeries{
			Labels: buildSeriesLabels(met, fmt.Sprintf("%s_sum", baseName)),
			Samples: []*protos2.Sample{{
				Value:     float64(*met.Histogram.SampleSum),
				Timestamp: met.GetTimestampMs(),
			}},
		})
		// Buckets
		for _, bucket := range met.Histogram.Bucket {
			labels := buildSeriesLabels(met, baseName)
			labels = append(labels, &protos2.Label{
				Name:  "le",
				Value: strconv.FormatFloat(bucket.GetUpperBound(), 'f', -1, 64),
			})
			series = append(series, &protos2.TimeSeries{
				// TODO: Probably need to add label for the bucket upper bound
				Labels: labels,
				Samples: []*protos2.Sample{{
					Value:     float64(bucket.GetCumulativeCount()),
					Timestamp: met.GetTimestampMs(),
				}},
			})
		}

	}
	return series
}

func buildSummarySamples(fam *dto.MetricFamily) []*protos2.TimeSeries {
	series := make([]*protos2.TimeSeries, 0, len(fam.Metric))
	baseName := fam.GetName()
	for _, met := range fam.Metric {
		if met.Summary == nil {
			continue
		}
		// SampleCount
		series = append(series, &protos2.TimeSeries{
			Labels: buildSeriesLabels(met, fmt.Sprintf("%s_count", baseName)),
			Samples: []*protos2.Sample{{
				Value:     float64(*met.Summary.SampleCount),
				Timestamp: met.GetTimestampMs(),
			}},
		})
		// SampleSum
		series = append(series, &protos2.TimeSeries{
			Labels: buildSeriesLabels(met, fmt.Sprintf("%s_sum", baseName)),
			Samples: []*protos2.Sample{{
				Value:     float64(*met.Summary.SampleSum),
				Timestamp: met.GetTimestampMs(),
			}},
		})
		// Quantile
		for _, quant := range met.Summary.Quantile {
			labels := buildSeriesLabels(met, baseName)
			labels = append(labels, &protos2.Label{
				Name:  "quantile",
				Value: strconv.FormatFloat(quant.GetQuantile(), 'f', -1, 64),
			})
			series = append(series, &protos2.TimeSeries{
				// TODO: Probably need to add label for the quantile
				Labels: buildSeriesLabels(met, baseName),
				Samples: []*protos2.Sample{{
					Value:     float64(quant.GetValue()),
					Timestamp: met.GetTimestampMs(),
				}},
			})
		}

	}
	return series

}

func metricTypeToMetricTypeProto(t dto.MetricType) protos2.MetricMetadata_MetricType {
	v, ok := protos2.MetricMetadata_MetricType_value[t.String()]
	if !ok {
		return protos2.MetricMetadata_UNKNOWN
	}

	return protos2.MetricMetadata_MetricType(v)
}

type RemoteWriteClient interface {
	Do(req *http.Request) (*http.Response, error)
}
