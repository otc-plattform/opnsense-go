package haproxy

import (
	"context"
	"fmt"

	"github.com/browningluke/opnsense-go/pkg/api"
)

type HAProxyObject map[string]interface{}

type HAProxyObjectSearchResult struct {
	Rows     []HAProxyObject `json:"rows"`
	RowCount int             `json:"rowCount"`
	Total    int             `json:"total"`
	Current  int             `json:"current"`
}

type Settings struct {
	General GeneralGet `json:"general"`
}

type SettingsGetResponse struct {
	HAProxy Settings `json:"haproxy"`
}

type SettingsSetRequest struct {
	General GeneralSet `json:"general"`
}

type GeneralGet struct {
	Enabled         string      `json:"enabled"`
	GracefulStop    string      `json:"gracefulStop"`
	HardStopAfter   string      `json:"hardStopAfter"`
	CloseSpreadTime string      `json:"closeSpreadTime"`
	SeamlessReload  string      `json:"seamlessReload"`
	StoreOCSP       string      `json:"storeOcsp"`
	ShowIntro       string      `json:"showIntro"`
	Peers           Peers       `json:"peers"`
	Tuning          TuningGet   `json:"tuning"`
	Defaults        DefaultsGet `json:"defaults"`
	Logging         LoggingGet  `json:"logging"`
	Stats           StatsGet    `json:"stats"`
	Cache           Cache       `json:"cache"`
}

type GeneralSet struct {
	Enabled         string       `json:"enabled"`
	GracefulStop    string       `json:"gracefulStop"`
	HardStopAfter   string       `json:"hardStopAfter"`
	CloseSpreadTime string       `json:"closeSpreadTime"`
	SeamlessReload  string       `json:"seamlessReload"`
	StoreOCSP       string       `json:"storeOcsp"`
	ShowIntro       string       `json:"showIntro"`
	Peers           *Peers       `json:"peers,omitempty"`
	Tuning          *TuningSet   `json:"tuning,omitempty"`
	Defaults        *DefaultsSet `json:"defaults,omitempty"`
	Logging         *LoggingSet  `json:"logging,omitempty"`
	Stats           *StatsSet    `json:"stats,omitempty"`
	Cache           *Cache       `json:"cache,omitempty"`
}

type Peers struct {
	Enabled string `json:"enabled"`
	Name1   string `json:"name1"`
	Listen1 string `json:"listen1"`
	Port1   string `json:"port1"`
	Name2   string `json:"name2"`
	Listen2 string `json:"listen2"`
	Port2   string `json:"port2"`
}

type TuningGet struct {
	Root                           string           `json:"root"`
	MaxConnections                 string           `json:"maxConnections"`
	Nbthread                       string           `json:"nbthread"`
	ResolversPrefer                api.FieldOptions `json:"resolversPrefer"`
	SSLServerVerify                api.FieldOptions `json:"sslServerVerify"`
	MaxDHSize                      string           `json:"maxDHSize"`
	BufferSize                     string           `json:"bufferSize"`
	SpreadChecks                   string           `json:"spreadChecks"`
	BogusProxyEnabled              string           `json:"bogusProxyEnabled"`
	LuaMaxMem                      string           `json:"luaMaxMem"`
	CustomOptions                  string           `json:"customOptions"`
	OCSPUpdateEnabled              string           `json:"ocspUpdateEnabled"`
	OCSPUpdateMinDelay             string           `json:"ocspUpdateMinDelay"`
	OCSPUpdateMaxDelay             string           `json:"ocspUpdateMaxDelay"`
	SSLDefaultsEnabled             string           `json:"ssl_defaultsEnabled"`
	SSLBindOptions                 api.FieldOptions `json:"ssl_bindOptions"`
	SSLMinVersion                  api.FieldOptions `json:"ssl_minVersion"`
	SSLMaxVersion                  api.FieldOptions `json:"ssl_maxVersion"`
	SSLCipherList                  string           `json:"ssl_cipherList"`
	SSLCipherSuites                string           `json:"ssl_cipherSuites"`
	H2InitialWindowSize            string           `json:"h2_initialWindowSize"`
	H2InitialWindowSizeOutgoing    string           `json:"h2_initialWindowSizeOutgoing"`
	H2InitialWindowSizeIncoming    string           `json:"h2_initialWindowSizeIncoming"`
	H2MaxConcurrentStreams         string           `json:"h2_maxConcurrentStreams"`
	H2MaxConcurrentStreamsOutgoing string           `json:"h2_maxConcurrentStreamsOutgoing"`
	H2MaxConcurrentStreamsIncoming string           `json:"h2_maxConcurrentStreamsIncoming"`
}

type TuningSet struct {
	Root                           string `json:"root"`
	MaxConnections                 string `json:"maxConnections"`
	Nbthread                       string `json:"nbthread"`
	ResolversPrefer                string `json:"resolversPrefer"`
	SSLServerVerify                string `json:"sslServerVerify"`
	MaxDHSize                      string `json:"maxDHSize"`
	BufferSize                     string `json:"bufferSize"`
	SpreadChecks                   string `json:"spreadChecks"`
	BogusProxyEnabled              string `json:"bogusProxyEnabled"`
	LuaMaxMem                      string `json:"luaMaxMem"`
	CustomOptions                  string `json:"customOptions"`
	OCSPUpdateEnabled              string `json:"ocspUpdateEnabled"`
	OCSPUpdateMinDelay             string `json:"ocspUpdateMinDelay"`
	OCSPUpdateMaxDelay             string `json:"ocspUpdateMaxDelay"`
	SSLDefaultsEnabled             string `json:"ssl_defaultsEnabled"`
	SSLBindOptions                 string `json:"ssl_bindOptions"`
	SSLMinVersion                  string `json:"ssl_minVersion"`
	SSLMaxVersion                  string `json:"ssl_maxVersion"`
	SSLCipherList                  string `json:"ssl_cipherList"`
	SSLCipherSuites                string `json:"ssl_cipherSuites"`
	H2InitialWindowSize            string `json:"h2_initialWindowSize"`
	H2InitialWindowSizeOutgoing    string `json:"h2_initialWindowSizeOutgoing"`
	H2InitialWindowSizeIncoming    string `json:"h2_initialWindowSizeIncoming"`
	H2MaxConcurrentStreams         string `json:"h2_maxConcurrentStreams"`
	H2MaxConcurrentStreamsOutgoing string `json:"h2_maxConcurrentStreamsOutgoing"`
	H2MaxConcurrentStreamsIncoming string `json:"h2_maxConcurrentStreamsIncoming"`
}

type DefaultsGet struct {
	MaxConnections        string           `json:"maxConnections"`
	MaxConnectionsServers string           `json:"maxConnectionsServers"`
	TimeoutClient         string           `json:"timeoutClient"`
	TimeoutConnect        string           `json:"timeoutConnect"`
	TimeoutCheck          string           `json:"timeoutCheck"`
	TimeoutServer         string           `json:"timeoutServer"`
	Retries               string           `json:"retries"`
	Redispatch            api.FieldOptions `json:"redispatch"`
	InitAddr              api.FieldOptions `json:"init_addr"`
	CustomOptions         string           `json:"customOptions"`
}

type DefaultsSet struct {
	MaxConnections        string `json:"maxConnections"`
	MaxConnectionsServers string `json:"maxConnectionsServers"`
	TimeoutClient         string `json:"timeoutClient"`
	TimeoutConnect        string `json:"timeoutConnect"`
	TimeoutCheck          string `json:"timeoutCheck"`
	TimeoutServer         string `json:"timeoutServer"`
	Retries               string `json:"retries"`
	Redispatch            string `json:"redispatch"`
	InitAddr              string `json:"init_addr"`
	CustomOptions         string `json:"customOptions"`
}

type LoggingGet struct {
	Host     string           `json:"host"`
	Facility api.FieldOptions `json:"facility"`
	Level    api.FieldOptions `json:"level"`
	Length   string           `json:"length"`
}

type LoggingSet struct {
	Host     string `json:"host"`
	Facility string `json:"facility"`
	Level    string `json:"level"`
	Length   string `json:"length"`
}

type StatsGet struct {
	Enabled           string           `json:"enabled"`
	Port              string           `json:"port"`
	RemoteEnabled     string           `json:"remoteEnabled"`
	RemoteBind        api.FieldOptions `json:"remoteBind"`
	AuthEnabled       string           `json:"authEnabled"`
	Users             api.FieldOptions `json:"users"`
	AllowedUsers      []string         `json:"allowedUsers"`
	AllowedGroups     []string         `json:"allowedGroups"`
	CustomOptions     string           `json:"customOptions"`
	PrometheusEnabled string           `json:"prometheus_enabled"`
	PrometheusBind    api.FieldOptions `json:"prometheus_bind"`
	PrometheusPath    string           `json:"prometheus_path"`
}

type StatsSet struct {
	Enabled           string `json:"enabled"`
	Port              string `json:"port"`
	RemoteEnabled     string `json:"remoteEnabled"`
	RemoteBind        string `json:"remoteBind"`
	AuthEnabled       string `json:"authEnabled"`
	AllowedUsers      string `json:"allowedUsers"`
	AllowedGroups     string `json:"allowedGroups"`
	CustomOptions     string `json:"customOptions"`
	PrometheusEnabled string `json:"prometheus_enabled"`
	PrometheusBind    string `json:"prometheus_bind"`
	PrometheusPath    string `json:"prometheus_path"`
}

type Cache struct {
	Enabled             string `json:"enabled"`
	TotalMaxSize        string `json:"totalMaxSize"`
	MaxAge              string `json:"maxAge"`
	MaxObjectSize       string `json:"maxObjectSize"`
	ProcessVary         string `json:"processVary"`
	MaxSecondaryEntries string `json:"maxSecondaryEntries"`
}

type Server struct {
	Enabled              string `json:"enabled"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Type                 string `json:"type"`
	Address              string `json:"address"`
	ServiceName          string `json:"serviceName"`
	Number               string `json:"number"`
	LinkedResolver       string `json:"linkedResolver"`
	ResolverOpts         string `json:"resolverOpts"`
	UnixSocket           string `json:"unix_socket"`
	Port                 string `json:"port"`
	Mode                 string `json:"mode"`
	MultiplexerProtocol  string `json:"multiplexer_protocol"`
	ResolvePrefer        string `json:"resolvePrefer"`
	SSL                  string `json:"ssl"`
	SSLSNI               string `json:"sslSNI"`
	SSLSNIExpr           string `json:"sslSNIExpr"`
	SSLVerify            string `json:"sslVerify"`
	SSLCA                string `json:"sslCA"`
	SSLCRL               string `json:"sslCRL"`
	SSLClientCertificate string `json:"sslClientCertificate"`
	MaxConnections       string `json:"maxConnections"`
	Weight               string `json:"weight"`
	CheckInterval        string `json:"checkInterval"`
	CheckDownInterval    string `json:"checkDownInterval"`
	Checkport            string `json:"checkport"`
	Source               string `json:"source"`
	Advanced             string `json:"advanced"`
}

type ServerGetResponse struct {
	Server ServerGet `json:"server"`
}

type ServerGet struct {
	ID                   string           `json:"id"`
	Enabled              string           `json:"enabled"`
	Name                 string           `json:"name"`
	Description          string           `json:"description"`
	Address              string           `json:"address"`
	Port                 string           `json:"port"`
	Checkport            string           `json:"checkport"`
	Mode                 api.FieldOptions `json:"mode"`
	MultiplexerProtocol  api.FieldOptions `json:"multiplexer_protocol"`
	Type                 api.FieldOptions `json:"type"`
	ServiceName          string           `json:"serviceName"`
	Number               string           `json:"number"`
	LinkedResolver       api.FieldOptions `json:"linkedResolver"`
	ResolverOpts         api.FieldOptions `json:"resolverOpts"`
	ResolvePrefer        api.FieldOptions `json:"resolvePrefer"`
	SSL                  string           `json:"ssl"`
	SSLSNI               string           `json:"sslSNI"`
	SSLSNIExpr           string           `json:"sslSNIExpr"`
	SSLVerify            string           `json:"sslVerify"`
	SSLCA                api.FieldOptions `json:"sslCA"`
	SSLCRL               api.FieldOptions `json:"sslCRL"`
	SSLClientCertificate api.FieldOptions `json:"sslClientCertificate"`
	MaxConnections       string           `json:"maxConnections"`
	Weight               string           `json:"weight"`
	CheckInterval        string           `json:"checkInterval"`
	CheckDownInterval    string           `json:"checkDownInterval"`
	Source               string           `json:"source"`
	Advanced             string           `json:"advanced"`
	UnixSocket           api.FieldOptions `json:"unix_socket"`
}

type ServerSearchItem struct {
	UUID                 string `json:"uuid"`
	ID                   string `json:"id"`
	Enabled              string `json:"enabled"`
	Name                 string `json:"name"`
	Description          string `json:"description"`
	Address              string `json:"address"`
	Port                 string `json:"port"`
	Checkport            string `json:"checkport"`
	Mode                 string `json:"mode"`
	MultiplexerProtocol  string `json:"multiplexer_protocol"`
	Type                 string `json:"type"`
	ServiceName          string `json:"serviceName"`
	Number               string `json:"number"`
	LinkedResolver       string `json:"linkedResolver"`
	ResolverOpts         string `json:"resolverOpts"`
	ResolvePrefer        string `json:"resolvePrefer"`
	SSL                  string `json:"ssl"`
	SSLSNI               string `json:"sslSNI"`
	SSLSNIExpr           string `json:"sslSNIExpr"`
	SSLVerify            string `json:"sslVerify"`
	SSLCA                string `json:"sslCA"`
	SSLCRL               string `json:"sslCRL"`
	SSLClientCertificate string `json:"sslClientCertificate"`
	MaxConnections       string `json:"maxConnections"`
	Weight               string `json:"weight"`
	CheckInterval        string `json:"checkInterval"`
	CheckDownInterval    string `json:"checkDownInterval"`
	Source               string `json:"source"`
	Advanced             string `json:"advanced"`
	UnixSocket           string `json:"unix_socket"`
}

type ServerSearchResult struct {
	Rows     []ServerSearchItem `json:"rows"`
	RowCount int                `json:"rowCount"`
	Total    int                `json:"total"`
	Current  int                `json:"current"`
}

func (c *Controller) HAProxyGetSettings(ctx context.Context) (*SettingsGetResponse, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint: "/haproxy/settings/get",
		Method:       "GET",
	}

	resultData := &SettingsGetResponse{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("GetSettings call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxyAddServer(ctx context.Context, server Server) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint: "/haproxy/settings/add_server",
		Method:       "POST",
		BodyParameters: map[string]interface{}{
			"server": server,
		},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("AddServer call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxySearchServers(ctx context.Context) (*ServerSearchResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint: "/haproxy/settings/search_servers",
		Method:       "POST",
		BodyParameters: map[string]interface{}{
			"current":  1,
			"sort":     map[string]interface{}{},
			"rowCount": 9999,
		},
	}

	resultData := &ServerSearchResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("SearchServers call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxyGetServer(ctx context.Context, uuid string) (*ServerGetResponse, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:   "/haproxy/settings/get_server",
		Method:         "GET",
		PathParameters: []string{uuid},
	}

	resultData := &ServerGetResponse{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("GetServer call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxyEditServer(ctx context.Context, uuid string, server Server) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:   "/haproxy/settings/set_server",
		Method:         "POST",
		PathParameters: []string{uuid},
		BodyParameters: map[string]interface{}{
			"server": server,
		},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("EditServer call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxyDeleteServer(ctx context.Context, uuid string) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:   "/haproxy/settings/del_server",
		Method:         "POST",
		PathParameters: []string{uuid},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("DeleteServer call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxySetSettings(ctx context.Context, haproxy SettingsSetRequest) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint: "/haproxy/settings/set",
		Method:       "POST",
		BodyParameters: map[string]interface{}{
			"haproxy": haproxy,
		},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("SetSettings call failed: %w", err)
	}
	return result, nil
}

func (c *Controller) HAProxySearchBackends(ctx context.Context) (*HAProxyObjectSearchResult, error) {
	return c.searchObjects(ctx, "/haproxy/settings/search_backends", "SearchBackends")
}

func (c *Controller) HAProxyGetBackend(ctx context.Context, uuid string) (HAProxyObject, error) {
	return c.getObject(ctx, "/haproxy/settings/get_backend", "backend", "GetBackend", uuid)
}

func (c *Controller) HAProxyAddBackend(ctx context.Context, backend HAProxyObject) (*api.ActionResult, error) {
	return c.addObject(ctx, "/haproxy/settings/add_backend", "backend", "AddBackend", backend)
}

func (c *Controller) HAProxyEditBackend(ctx context.Context, uuid string, backend HAProxyObject) (*api.ActionResult, error) {
	return c.editObject(ctx, "/haproxy/settings/set_backend", "backend", "EditBackend", uuid, backend)
}

func (c *Controller) HAProxyDeleteBackend(ctx context.Context, uuid string) (*api.ActionResult, error) {
	return c.deleteObject(ctx, "/haproxy/settings/del_backend", "DeleteBackend", uuid)
}

func (c *Controller) HAProxySearchFrontends(ctx context.Context) (*HAProxyObjectSearchResult, error) {
	return c.searchObjects(ctx, "/haproxy/settings/search_frontends", "SearchFrontends")
}

func (c *Controller) HAProxyGetFrontend(ctx context.Context, uuid string) (HAProxyObject, error) {
	return c.getObject(ctx, "/haproxy/settings/get_frontend", "frontend", "GetFrontend", uuid)
}

func (c *Controller) HAProxyAddFrontend(ctx context.Context, frontend HAProxyObject) (*api.ActionResult, error) {
	return c.addObject(ctx, "/haproxy/settings/add_frontend", "frontend", "AddFrontend", frontend)
}

func (c *Controller) HAProxyEditFrontend(ctx context.Context, uuid string, frontend HAProxyObject) (*api.ActionResult, error) {
	return c.editObject(ctx, "/haproxy/settings/set_frontend", "frontend", "EditFrontend", uuid, frontend)
}

func (c *Controller) HAProxyDeleteFrontend(ctx context.Context, uuid string) (*api.ActionResult, error) {
	return c.deleteObject(ctx, "/haproxy/settings/del_frontend", "DeleteFrontend", uuid)
}

func (c *Controller) HAProxySearchACLs(ctx context.Context) (*HAProxyObjectSearchResult, error) {
	return c.searchObjects(ctx, "/haproxy/settings/search_acls", "SearchACLs")
}

func (c *Controller) HAProxyGetACL(ctx context.Context, uuid string) (HAProxyObject, error) {
	return c.getObject(ctx, "/haproxy/settings/get_acl", "acl", "GetACL", uuid)
}

func (c *Controller) HAProxyAddACL(ctx context.Context, acl HAProxyObject) (*api.ActionResult, error) {
	return c.addObject(ctx, "/haproxy/settings/add_acl", "acl", "AddACL", acl)
}

func (c *Controller) HAProxyEditACL(ctx context.Context, uuid string, acl HAProxyObject) (*api.ActionResult, error) {
	return c.editObject(ctx, "/haproxy/settings/set_acl", "acl", "EditACL", uuid, acl)
}

func (c *Controller) HAProxyDeleteACL(ctx context.Context, uuid string) (*api.ActionResult, error) {
	return c.deleteObject(ctx, "/haproxy/settings/del_acl", "DeleteACL", uuid)
}

func (c *Controller) HAProxySearchActions(ctx context.Context) (*HAProxyObjectSearchResult, error) {
	return c.searchObjects(ctx, "/haproxy/settings/search_actions", "SearchActions")
}

func (c *Controller) HAProxyGetAction(ctx context.Context, uuid string) (HAProxyObject, error) {
	return c.getObject(ctx, "/haproxy/settings/get_action", "action", "GetAction", uuid)
}

func (c *Controller) HAProxyAddAction(ctx context.Context, action HAProxyObject) (*api.ActionResult, error) {
	return c.addObject(ctx, "/haproxy/settings/add_action", "action", "AddAction", action)
}

func (c *Controller) HAProxyEditAction(ctx context.Context, uuid string, action HAProxyObject) (*api.ActionResult, error) {
	return c.editObject(ctx, "/haproxy/settings/set_action", "action", "EditAction", uuid, action)
}

func (c *Controller) HAProxyDeleteAction(ctx context.Context, uuid string) (*api.ActionResult, error) {
	return c.deleteObject(ctx, "/haproxy/settings/del_action", "DeleteAction", uuid)
}

func (c *Controller) searchObjects(ctx context.Context, endpoint string, operation string) (*HAProxyObjectSearchResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint: endpoint,
		Method:       "POST",
		BodyParameters: map[string]interface{}{
			"current":  1,
			"sort":     map[string]interface{}{},
			"rowCount": 9999,
		},
	}

	resultData := &HAProxyObjectSearchResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("%s call failed: %w", operation, err)
	}
	return result, nil
}

func (c *Controller) getObject(ctx context.Context, endpoint string, objectKey string, operation string, uuid string) (HAProxyObject, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:   endpoint,
		Method:         "GET",
		PathParameters: []string{uuid},
	}

	resultData := map[string]HAProxyObject{}
	result, err := api.Call(c.Client(), ctx, callOpts, &resultData)
	if err != nil {
		return nil, fmt.Errorf("%s call failed: %w", operation, err)
	}
	return (*result)[objectKey], nil
}

func (c *Controller) addObject(ctx context.Context, endpoint string, objectKey string, operation string, object HAProxyObject) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint: endpoint,
		Method:       "POST",
		BodyParameters: map[string]interface{}{
			objectKey: object,
		},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("%s call failed: %w", operation, err)
	}
	return result, nil
}

func (c *Controller) editObject(ctx context.Context, endpoint string, objectKey string, operation string, uuid string, object HAProxyObject) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:   endpoint,
		Method:         "POST",
		PathParameters: []string{uuid},
		BodyParameters: map[string]interface{}{
			objectKey: object,
		},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("%s call failed: %w", operation, err)
	}
	return result, nil
}

func (c *Controller) deleteObject(ctx context.Context, endpoint string, operation string, uuid string) (*api.ActionResult, error) {
	callOpts := api.RPCOpts{
		BaseEndpoint:   endpoint,
		Method:         "POST",
		PathParameters: []string{uuid},
	}

	resultData := &api.ActionResult{}
	result, err := api.Call(c.Client(), ctx, callOpts, resultData)
	if err != nil {
		return nil, fmt.Errorf("%s call failed: %w", operation, err)
	}
	return result, nil
}
