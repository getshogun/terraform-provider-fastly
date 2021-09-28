package fastly

import (
	"context"

	gofastly "github.com/fastly/go-fastly/v5/fastly"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var computeAttributes = ServiceMetadata{
	ServiceTypeCompute,
}

// Ordering is important - stored is processing order
// Some objects may need to be updated first, as they can be referenced by other
// configuration objects (Backends, Request Headers, etc).
var computeService = &BaseServiceDefinition{
	Type: computeAttributes.serviceType,
	Attributes: []ServiceAttributeDefinition{
		NewServiceSettings(),
		NewServiceDomain(computeAttributes),
		newFakeLoggingHandler(computeAttributes, "healthcheck"),
		NewServiceBackend(computeAttributes),
		newFakeLoggingHandler(computeAttributes, "director"),
		NewServiceDictionary(computeAttributes),
		NewServicePackage(computeAttributes),
		NewServiceLoggingDatadog(computeAttributes),
		newFakeLoggingHandler(computeAttributes, "s3logging"),
		newFakeLoggingHandler(computeAttributes, "papertrail"),
		newFakeLoggingHandler(computeAttributes, "sumologic"),
		newFakeLoggingHandler(computeAttributes, "gcslogging"),
		newFakeLoggingHandler(computeAttributes, "bigquerylogging"),
		newFakeLoggingHandler(computeAttributes, "syslog"),
		newFakeLoggingHandler(computeAttributes, "logentries"),
		newFakeLoggingHandler(computeAttributes, "splunk"),
		newFakeLoggingHandler(computeAttributes, "blobstoragelogging"),
		newFakeLoggingHandler(computeAttributes, "httpslogging"),
		newFakeLoggingHandler(computeAttributes, "logging_elasticsearch"),
		newFakeLoggingHandler(computeAttributes, "logging_ftp"),
		newFakeLoggingHandler(computeAttributes, "logging_sftp"),
		newFakeLoggingHandler(computeAttributes, "logging_loggly"),
		newFakeLoggingHandler(computeAttributes, "logging_googlepubsub"),
		newFakeLoggingHandler(computeAttributes, "logging_scalyr"),
		newFakeLoggingHandler(computeAttributes, "logging_newrelic"),
		newFakeLoggingHandler(computeAttributes, "logging_kafka"),
		newFakeLoggingHandler(computeAttributes, "logging_heroku"),
		newFakeLoggingHandler(computeAttributes, "logging_honeycomb"),
		newFakeLoggingHandler(computeAttributes, "logging_logshuttle"),
		newFakeLoggingHandler(computeAttributes, "logging_openstack"),
		newFakeLoggingHandler(computeAttributes, "logging_digitalocean"),
		newFakeLoggingHandler(computeAttributes, "logging_cloudfiles"),
		newFakeLoggingHandler(computeAttributes, "logging_kinesis"),
	},
}

func resourceServiceComputeV1() *schema.Resource {
	return resourceService(computeService)
}

type fakeLoggingHandler struct {
	*DefaultServiceAttributeHandler
}

func newFakeLoggingHandler(sa ServiceMetadata, key string) ServiceAttributeDefinition {
	return ToServiceAttributeDefinition(&fakeLoggingHandler{
		&DefaultServiceAttributeHandler{
			key:             key,
			serviceMetadata: sa,
		},
	})
}

func (h *fakeLoggingHandler) Key() string { return h.key }

func (h *fakeLoggingHandler) GetSchema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeSet,
		Optional: true,
		Elem: &schema.Resource{
			Schema: make(map[string]*schema.Schema),
		},
	}
}

func (h *fakeLoggingHandler) Create(_ context.Context, d *schema.ResourceData, resource map[string]interface {
}, serviceVersion int, conn *gofastly.Client) error {
	return nil
}

func (h *fakeLoggingHandler) Read(_ context.Context, d *schema.ResourceData, _ map[string]interface{}, serviceVersion int, conn *gofastly.Client) error {

	return nil
}

func (h *fakeLoggingHandler) Update(_ context.Context, d *schema.ResourceData, resource, modified map[string]interface {
}, serviceVersion int, conn *gofastly.Client) error {
	return nil
}

func (h *fakeLoggingHandler) Delete(_ context.Context, d *schema.ResourceData, resource map[string]interface {
}, serviceVersion int, conn *gofastly.Client) error {
	return nil
}
