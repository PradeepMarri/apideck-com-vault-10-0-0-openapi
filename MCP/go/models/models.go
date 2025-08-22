package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// ResourceExample represents the ResourceExample schema from the OpenAPI specification
type ResourceExample struct {
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Unified_api string `json:"unified_api,omitempty"` // Name of Apideck Unified API
	Example_response map[string]interface{} `json:"example_response,omitempty"` // Example response from the downstream API
	Resource LinkedConnectorResource `json:"resource,omitempty"`
}

// UpdateCustomMappingRequest represents the UpdateCustomMappingRequest schema from the OpenAPI specification
type UpdateCustomMappingRequest struct {
	Value string `json:"value"` // Target Field Mapping value
}

// CreateCustomMappingRequest represents the CreateCustomMappingRequest schema from the OpenAPI specification
type CreateCustomMappingRequest struct {
	Value string `json:"value"` // Target Field Mapping value
}

// GetConsumersResponse represents the GetConsumersResponse schema from the OpenAPI specification
type GetConsumersResponse struct {
	Data []map[string]interface{} `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetConnectionsResponse represents the GetConnectionsResponse schema from the OpenAPI specification
type GetConnectionsResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Connection `json:"data"`
	Status string `json:"status"` // HTTP Response Status
}

// ProxyRequest represents the ProxyRequest schema from the OpenAPI specification
type ProxyRequest struct {
}

// UnprocessableResponse represents the UnprocessableResponse schema from the OpenAPI specification
type UnprocessableResponse struct {
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// DeleteConsumerResponse represents the DeleteConsumerResponse schema from the OpenAPI specification
type DeleteConsumerResponse struct {
	Data interface{} `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// ConnectionEvent represents the ConnectionEvent schema from the OpenAPI specification
type ConnectionEvent struct {
	Execution_attempt float64 `json:"execution_attempt,omitempty"` // The current count this request event has been attempted
	Occurred_at string `json:"occurred_at,omitempty"` // ISO Datetime for when the original event occurred
	Service_id string `json:"service_id,omitempty"` // Service provider identifier
	Entity ConsumerConnection `json:"entity,omitempty"`
	Entity_id string `json:"entity_id,omitempty"` // The service provider's ID of the entity that triggered this event
	Entity_type string `json:"entity_type,omitempty"` // The type entity that triggered this event
	Event_id string `json:"event_id,omitempty"` // Unique reference to this request event
	Event_type string `json:"event_type,omitempty"`
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Theme map[string]interface{} `json:"theme,omitempty"` // Theming options to change the look and feel of Vault.
	Consumer_metadata ConsumerMetadata `json:"consumer_metadata,omitempty"` // The metadata of the consumer. This is used to display the consumer in the sidebar. This is optional, but recommended.
	Custom_consumer_settings map[string]interface{} `json:"custom_consumer_settings,omitempty"` // Custom consumer settings that are passed as part of the session.
	Redirect_uri string `json:"redirect_uri,omitempty"` // The URL to redirect the user to after the session has been configured.
	Settings map[string]interface{} `json:"settings,omitempty"` // Settings to change the way the Vault is displayed.
}

// WebhookSubscription represents the WebhookSubscription schema from the OpenAPI specification
type WebhookSubscription struct {
	Created_at string `json:"created_at,omitempty"` // The date and time the webhook subscription was created downstream
	Downstream_event_types []string `json:"downstream_event_types,omitempty"` // The list of downstream Events this connection is subscribed to
	Downstream_id string `json:"downstream_id,omitempty"` // The ID of the downstream service
	Execute_url string `json:"execute_url,omitempty"` // The URL the downstream is sending to when the event is triggered
	Unify_event_types []string `json:"unify_event_types,omitempty"` // The list of Unify Events this connection is subscribed to
}

// Links represents the Links schema from the OpenAPI specification
type Links struct {
	Current string `json:"current,omitempty"` // Link to navigate to the current page through the API
	Next string `json:"next,omitempty"` // Link to navigate to the previous page through the API
	Previous string `json:"previous,omitempty"` // Link to navigate to the previous page through the API
}

// Connection represents the Connection schema from the OpenAPI specification
type Connection struct {
	Logo string `json:"logo,omitempty"` // The logo of the connection, that will be shown in the Vault
	Settings_required_for_authorization []string `json:"settings_required_for_authorization,omitempty"` // List of settings that are required to be configured on integration before authorization can occur
	Auth_type string `json:"auth_type,omitempty"` // Type of authorization used by the connector
	Authorize_url string `json:"authorize_url,omitempty"` // The OAuth redirect URI. Redirect your users to this URI to let them authorize your app in the connector's UI. Before you can use this URI, you must add `redirect_uri` as a query parameter to the `authorize_url`. Be sure to URL encode the `redirect_uri` part. Your users will be redirected to this `redirect_uri` after they granted access to your app in the connector's UI.
	Id string `json:"id,omitempty"` // The unique identifier of the connection.
	Resource_schema_support []string `json:"resource_schema_support,omitempty"`
	Website string `json:"website,omitempty"` // The website URL of the connection
	Icon string `json:"icon,omitempty"` // A visual icon of the connection, that will be shown in the Vault
	Revoke_url string `json:"revoke_url,omitempty"` // The OAuth revoke URI. Redirect your users to this URI to revoke this connection. Before you can use this URI, you must add `redirect_uri` as a query parameter. Your users will be redirected to this `redirect_uri` after they granted access to your app in the connector's UI.
	State string `json:"state,omitempty"` // [Connection state flow](#section/Connection-state)
	Form_fields []FormField `json:"form_fields,omitempty"` // The settings that are wanted to create a connection.
	Name string `json:"name,omitempty"` // The name of the connection
	Enabled bool `json:"enabled,omitempty"` // Whether the connection is enabled or not. You can enable or disable a connection using the Update Connection API.
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Attach your own consumer specific metadata
	Oauth_grant_type string `json:"oauth_grant_type,omitempty"` // OAuth grant type used by the connector. More info: https://oauth.net/2/grant-types
	Settings map[string]interface{} `json:"settings,omitempty"` // Connection settings. Values will persist to `form_fields` with corresponding id
	Tag_line string `json:"tag_line,omitempty"`
	Subscriptions []WebhookSubscription `json:"subscriptions,omitempty"`
	Status string `json:"status,omitempty"` // Status of the connection.
	Schema_support bool `json:"schema_support,omitempty"`
	Unified_api string `json:"unified_api,omitempty"` // The unified API category where the connection belongs to.
	Has_guide bool `json:"has_guide,omitempty"` // Whether the connector has a guide available in the developer docs or not (https://docs.apideck.com/connectors/{service_id}/docs/consumer+connection).
	Integration_state string `json:"integration_state,omitempty"` // The current state of the Integration.
	Service_id string `json:"service_id,omitempty"` // The ID of the service this connection belongs to.
	Validation_support bool `json:"validation_support,omitempty"`
	Configuration []map[string]interface{} `json:"configuration,omitempty"`
	Configurable_resources []string `json:"configurable_resources,omitempty"`
	Resource_settings_support []string `json:"resource_settings_support,omitempty"`
	Created_at float64 `json:"created_at,omitempty"`
	Updated_at float64 `json:"updated_at,omitempty"`
	Custom_mappings []CustomMapping `json:"custom_mappings,omitempty"` // List of custom mappings configured for this connection
}

// UpdateCustomMappingResponse represents the UpdateCustomMappingResponse schema from the OpenAPI specification
type UpdateCustomMappingResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data CustomMapping `json:"data"`
}

// CustomFieldFinder represents the CustomFieldFinder schema from the OpenAPI specification
type CustomFieldFinder struct {
	Finder string `json:"finder,omitempty"` // JSONPath finder for retrieving this value when mapping a response payload from downstream
	Id string `json:"id,omitempty"` // Custom Field ID
	Name string `json:"name,omitempty"` // Custom Field name to use as a label if provided
	Value interface{} `json:"value,omitempty"` // Custom Field value
	Description string `json:"description,omitempty"` // More information about the custom field
}

// GetCustomFieldsResponse represents the GetCustomFieldsResponse schema from the OpenAPI specification
type GetCustomFieldsResponse struct {
	Data []CustomFieldFinder `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// ConsumerConnection represents the ConsumerConnection schema from the OpenAPI specification
type ConsumerConnection struct {
	Auth_type string `json:"auth_type,omitempty"` // Type of authorization used by the connector
	Website string `json:"website,omitempty"`
	Id string `json:"id,omitempty"`
	State string `json:"state,omitempty"`
	Created_at string `json:"created_at,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Attach your own consumer specific metadata
	Updated_at string `json:"updated_at,omitempty"`
	Enabled bool `json:"enabled,omitempty"`
	Icon string `json:"icon,omitempty"`
	Service_id string `json:"service_id,omitempty"`
	Tag_line string `json:"tag_line,omitempty"`
	Logo string `json:"logo,omitempty"`
	Unified_api string `json:"unified_api,omitempty"`
	Consumer_id string `json:"consumer_id,omitempty"`
	Name string `json:"name,omitempty"`
	Settings map[string]interface{} `json:"settings,omitempty"` // Connection settings. Values will persist to `form_fields` with corresponding id
}

// GetResourceSchemaResponse represents the GetResourceSchemaResponse schema from the OpenAPI specification
type GetResourceSchemaResponse struct {
	Data ResourceSchema `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// CustomMapping represents the CustomMapping schema from the OpenAPI specification
type CustomMapping struct {
	Value string `json:"value,omitempty"` // Target Field Mapping value
	Consumer_id string `json:"consumer_id,omitempty"` // Consumer ID
	Custom_field bool `json:"custom_field,omitempty"` // This mapping represents a finder for a custom field
	Description string `json:"description,omitempty"` // Target Field description
	Id string `json:"id,omitempty"` // Target Field ID
	Key string `json:"key,omitempty"` // Target Field Key
	Label string `json:"label,omitempty"` // Target Field name to use as a label
	Required bool `json:"required,omitempty"` // Target Field Mapping is required
}

// Consumer represents the Consumer schema from the OpenAPI specification
type Consumer struct {
	Connections []ConsumerConnection `json:"connections,omitempty"`
	Request_count_updated string `json:"request_count_updated,omitempty"`
	Request_counts RequestCountAllocation `json:"request_counts,omitempty"`
	Services []string `json:"services,omitempty"`
	Aggregated_request_count float64 `json:"aggregated_request_count,omitempty"`
	Application_id string `json:"application_id,omitempty"` // ID of your Apideck Application
	Created string `json:"created,omitempty"`
	Modified string `json:"modified,omitempty"`
	Consumer_id string `json:"consumer_id"` // Unique consumer identifier. You can freely choose a consumer ID yourself. Most of the time, this is an ID of your internal data model that represents a user or account in your system (for example account:12345). If the consumer doesn't exist yet, Vault will upsert a consumer based on your ID.
	Metadata ConsumerMetadata `json:"metadata,omitempty"` // The metadata of the consumer. This is used to display the consumer in the sidebar. This is optional, but recommended.
}

// ConsumerRequestCountsInDateRangeResponse represents the ConsumerRequestCountsInDateRangeResponse schema from the OpenAPI specification
type ConsumerRequestCountsInDateRangeResponse struct {
	Data map[string]interface{} `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// LinkedConnectorResource represents the LinkedConnectorResource schema from the OpenAPI specification
type LinkedConnectorResource struct {
	Downstream_id string `json:"downstream_id,omitempty"` // ID of the resource in the Connector's API (downstream)
	Downstream_name string `json:"downstream_name,omitempty"` // Name of the resource in the Connector's API (downstream)
	Id string `json:"id,omitempty"` // ID of the resource, typically a lowercased version of name.
	Name string `json:"name,omitempty"` // Name of the resource (plural)
	Status string `json:"status,omitempty"` // Status of the resource. Resources with status live or beta are callable.
}

// SimpleFormFieldOption represents the SimpleFormFieldOption schema from the OpenAPI specification
type SimpleFormFieldOption struct {
	Label string `json:"label,omitempty"`
	Value string `json:"value,omitempty"`
}

// BadRequestResponse represents the BadRequestResponse schema from the OpenAPI specification
type BadRequestResponse struct {
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
}

// CreateConsumerResponse represents the CreateConsumerResponse schema from the OpenAPI specification
type CreateConsumerResponse struct {
	Data Consumer `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// NotFoundResponse represents the NotFoundResponse schema from the OpenAPI specification
type NotFoundResponse struct {
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
}

// GetCustomMappingResponse represents the GetCustomMappingResponse schema from the OpenAPI specification
type GetCustomMappingResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data CustomMapping `json:"data"`
}

// UpdateConsumerResponse represents the UpdateConsumerResponse schema from the OpenAPI specification
type UpdateConsumerResponse struct {
	Data Consumer `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// ConnectionImportData represents the ConnectionImportData schema from the OpenAPI specification
type ConnectionImportData struct {
	Credentials map[string]interface{} `json:"credentials,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"` // Attach your own consumer specific metadata
	Settings map[string]interface{} `json:"settings,omitempty"` // Connection settings. Values will persist to `form_fields` with corresponding id
}

// UpdateConsumerRequest represents the UpdateConsumerRequest schema from the OpenAPI specification
type UpdateConsumerRequest struct {
	Metadata ConsumerMetadata `json:"metadata,omitempty"` // The metadata of the consumer. This is used to display the consumer in the sidebar. This is optional, but recommended.
}

// GetLogsResponse represents the GetLogsResponse schema from the OpenAPI specification
type GetLogsResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data []Log `json:"data"`
	Links Links `json:"links,omitempty"` // Links to navigate to previous or next pages through the API
	Meta Meta `json:"meta,omitempty"` // Response metadata
}

// GetResourceExampleResponse represents the GetResourceExampleResponse schema from the OpenAPI specification
type GetResourceExampleResponse struct {
	Data ResourceExample `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// ConnectionMetadata represents the ConnectionMetadata schema from the OpenAPI specification
type ConnectionMetadata struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Log represents the Log schema from the OpenAPI specification
type Log struct {
	Http_method string `json:"http_method"` // HTTP Method of request.
	Id string `json:"id"` // UUID acting as Request Identifier.
	Api_style string `json:"api_style"` // Indicates if the request was made via REST or Graphql endpoint.
	Has_children bool `json:"has_children"` // When request is a parent request, this indicates if there are child requests associated.
	Latency float64 `json:"latency"` // Latency added by making this request via Unified Api.
	Parent_id string `json:"parent_id"` // When request is a child request, this UUID indicates it's parent request.
	Source_ip string `json:"source_ip,omitempty"` // The IP address of the source of the request.
	Timestamp string `json:"timestamp"` // ISO Date and time when the request was made.
	Path string `json:"path"` // The path component of the URI the request was made to.
	Duration float64 `json:"duration"` // The entire execution time in milliseconds it took to call the Apideck service provider.
	Success bool `json:"success"` // Whether or not the request was successful.
	Status_code int `json:"status_code"` // HTTP Status code that was returned.
	Operation map[string]interface{} `json:"operation"` // The request as defined in OpenApi Spec.
	Error_message string `json:"error_message,omitempty"` // If error occurred, this is brief explanation
	Execution int `json:"execution"` // The entire execution time in milliseconds it took to make the request.
	Consumer_id string `json:"consumer_id"` // The consumer Id associated with the request.
	Unified_api string `json:"unified_api"` // Which Unified Api request was made to.
	Sandbox bool `json:"sandbox"` // Indicates whether the request was made using Apidecks sandbox credentials or not.
	Service map[string]interface{} `json:"service"` // Apideck service provider associated with request.
	Child_request bool `json:"child_request"` // Indicates whether or not this is a child or parent request.
	Base_url string `json:"base_url"` // The Apideck base URL the request was made to.
}

// UnexpectedErrorResponse represents the UnexpectedErrorResponse schema from the OpenAPI specification
type UnexpectedErrorResponse struct {
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
}

// ResourceSchema represents the ResourceSchema schema from the OpenAPI specification
type ResourceSchema struct {
}

// ConsumerMetadata represents the ConsumerMetadata schema from the OpenAPI specification
type ConsumerMetadata struct {
	Image string `json:"image,omitempty"` // The avatar of the user in the sidebar. Must be a valid URL
	User_name string `json:"user_name,omitempty"` // The name of the user as shown in the sidebar.
	Account_name string `json:"account_name,omitempty"` // The name of the account as shown in the sidebar.
	Email string `json:"email,omitempty"` // The email of the user as shown in the sidebar.
}

// CreateCustomMappingResponse represents the CreateCustomMappingResponse schema from the OpenAPI specification
type CreateCustomMappingResponse struct {
	Data CustomMapping `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetConnectionResponse represents the GetConnectionResponse schema from the OpenAPI specification
type GetConnectionResponse struct {
	Data Connection `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// ConnectionWebhook represents the ConnectionWebhook schema from the OpenAPI specification
type ConnectionWebhook struct {
	Unified_api string `json:"unified_api"` // Name of Apideck Unified API
	Disabled_reason string `json:"disabled_reason,omitempty"` // Indicates if the webhook has has been disabled as it reached its retry limit or if account is over the usage allocated by it's plan.
	Events []string `json:"events"` // The list of subscribed events for this webhook. [`*`] indicates that all events are enabled.
	Status string `json:"status"` // The status of the webhook.
	Created_at string `json:"created_at,omitempty"` // The date and time when the object was created.
	Delivery_url string `json:"delivery_url"` // The delivery url of the webhook endpoint.
	Description string `json:"description,omitempty"` // A description of the object.
	Updated_at string `json:"updated_at,omitempty"` // The date and time when the object was last updated.
	Execute_base_url string `json:"execute_base_url"` // The Unify Base URL events from connectors will be sent to after service id is appended.
	Id string `json:"id,omitempty"`
}

// FormField represents the FormField schema from the OpenAPI specification
type FormField struct {
	Allow_custom_values bool `json:"allow_custom_values,omitempty"` // Only applicable to select fields. Allow the user to add a custom value though the option select if the desired value is not in the option select list.
	Hidden bool `json:"hidden,omitempty"` // Indicates if the form field is not displayed but the value that is being stored on the connection.
	Options []interface{} `json:"options,omitempty"`
	Placeholder string `json:"placeholder,omitempty"` // The placeholder for the form field
	Disabled bool `json:"disabled,omitempty"` // Indicates if the form field is displayed in a “read-only” mode.
	Required bool `json:"required,omitempty"` // Indicates if the form field is required, which means it must be filled in before the form can be submitted
	Custom_field bool `json:"custom_field,omitempty"`
	Prefix string `json:"prefix,omitempty"` // Prefix to display in front of the form field.
	Sensitive bool `json:"sensitive,omitempty"` // Indicates if the form field contains sensitive data, which will display the value as a masked input.
	TypeField interface{} `json:"type,omitempty"`
	Description string `json:"description,omitempty"` // The description of the form field
	Id string `json:"id,omitempty"` // The unique identifier of the form field.
	Label string `json:"label,omitempty"` // The label of the field
	Suffix string `json:"suffix,omitempty"` // Suffix to display next to the form field.
}

// UnauthorizedResponse represents the UnauthorizedResponse schema from the OpenAPI specification
type UnauthorizedResponse struct {
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
}

// CreateSessionResponse represents the CreateSessionResponse schema from the OpenAPI specification
type CreateSessionResponse struct {
	Data map[string]interface{} `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// NotImplementedResponse represents the NotImplementedResponse schema from the OpenAPI specification
type NotImplementedResponse struct {
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail interface{} `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
}

// Meta represents the Meta schema from the OpenAPI specification
type Meta struct {
	Items_on_page int `json:"items_on_page,omitempty"` // Number of items returned in the data property of the response
	Cursors map[string]interface{} `json:"cursors,omitempty"` // Cursors to navigate to previous or next pages through the API
}

// LogsFilter represents the LogsFilter schema from the OpenAPI specification
type LogsFilter struct {
	Connector_id string `json:"connector_id,omitempty"`
	Exclude_unified_apis string `json:"exclude_unified_apis,omitempty"`
	Status_code float64 `json:"status_code,omitempty"`
}

// RequestCountAllocation represents the RequestCountAllocation schema from the OpenAPI specification
type RequestCountAllocation struct {
	Vault float64 `json:"vault,omitempty"`
	Proxy float64 `json:"proxy,omitempty"`
	Unify float64 `json:"unify,omitempty"`
}

// UpdateConnectionResponse represents the UpdateConnectionResponse schema from the OpenAPI specification
type UpdateConnectionResponse struct {
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Connection `json:"data"`
	Status string `json:"status"` // HTTP Response Status
}

// FormFieldOptionGroup represents the FormFieldOptionGroup schema from the OpenAPI specification
type FormFieldOptionGroup struct {
	Id string `json:"id,omitempty"`
	Label string `json:"label,omitempty"`
	Options []SimpleFormFieldOption `json:"options,omitempty"`
}

// CreateConnectionResponse represents the CreateConnectionResponse schema from the OpenAPI specification
type CreateConnectionResponse struct {
	Data Connection `json:"data"`
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
}

// GetConsumerResponse represents the GetConsumerResponse schema from the OpenAPI specification
type GetConsumerResponse struct {
	Status string `json:"status"` // HTTP Response Status
	Status_code int `json:"status_code"` // HTTP Response Status Code
	Data Consumer `json:"data"`
}

// PaymentRequiredResponse represents the PaymentRequiredResponse schema from the OpenAPI specification
type PaymentRequiredResponse struct {
	Ref string `json:"ref,omitempty"` // Link to documentation of error type
	Status_code float64 `json:"status_code,omitempty"` // HTTP status code
	Type_name string `json:"type_name,omitempty"` // The type of error returned
	Detail string `json:"detail,omitempty"` // Contains parameter or domain specific information related to the error and why it occurred.
	ErrorField string `json:"error,omitempty"` // Contains an explanation of the status_code as defined in HTTP/1.1 standard (RFC 7231)
	Message string `json:"message,omitempty"` // A human-readable message providing more details about the error.
}
