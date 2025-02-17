package schema

import (
	"context"
	"sync"

	"github.com/kform-dev/kform-plugin/kfprotov1/kfplugin1"
	"github.com/kform-dev/kform-sdk-go/pkg/diag"
	"github.com/henderiw/logger/log"
)

func NewGRPCProviderServer(p *Provider) *GRPCProviderServer {
	return &GRPCProviderServer{
		provider: p,
		stopCh:   make(chan struct{}),
	}
}

// GRPCProviderServer handles the server, or plugin side of the rpc connection.
type GRPCProviderServer struct {
	provider *Provider
	stopCh   chan struct{}
	m        sync.Mutex
}

func (r *GRPCProviderServer) Capabilities(ctx context.Context, req *kfplugin1.Capabilities_Request) (*kfplugin1.Capabilities_Response, error) {
	// todo add ctx + tracing
	rpc := "capabilities"
	log := log.FromContext(ctx)
	log.Info(rpc)

	return &kfplugin1.Capabilities_Response{
		Diagnostics:        []*kfplugin1.Diagnostic{},
		ReadDataSources:    r.provider.getDataSources(),
		ListDataSources:    r.provider.getListDataSources(),
		Resources:          r.provider.getResources(),
		ServerCapabilities: &kfplugin1.ServerCapabilities{},
	}, nil
}

func (r *GRPCProviderServer) Configure(ctx context.Context, req *kfplugin1.Configure_Request) (*kfplugin1.Configure_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("configure...")

	if req == nil || req.Config == nil {
		return &kfplugin1.Configure_Response{
			Diagnostics: diag.Errorf("cannot configure a provider with empty config"),
		}, nil
	}

	diags := r.provider.Configure(ctx, req.Config)
	log.Info("configure done")
	return &kfplugin1.Configure_Response{
		Diagnostics: diags,
	}, nil
}
func (r *GRPCProviderServer) StopProvider(ctx context.Context, req *kfplugin1.StopProvider_Request) (*kfplugin1.StopProvider_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("stopProvider...")

	r.m.Lock()
	defer r.m.Unlock()

	// stop provider
	close(r.stopCh)
	// reset the stop signal
	r.stopCh = make(chan struct{})

	log.Info("stopProvider done")
	return &kfplugin1.StopProvider_Response{}, nil
}

func (r *GRPCProviderServer) ReadDataSource(ctx context.Context, req *kfplugin1.ReadDataSource_Request) (*kfplugin1.ReadDataSource_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("readDataSource...")

	res, ok := r.provider.DataSourcesMap[req.GetName()]
	if !ok {
		return &kfplugin1.ReadDataSource_Response{
			Diagnostics: diag.Errorf("cannot read data source, resourceType not found, got: %s", req.GetName()),
		}, nil
	}

	if res.ReadContext == nil {
		return &kfplugin1.ReadDataSource_Response{
			Diagnostics: diag.Errorf("cannot read data source, readContext not initialized, for: %s", req.GetName()),
		}, nil
	}

	d, diags := res.ReadContext(ctx, &ResourceObject{Scope: req.Scope, Obj: req.Obj}, r.provider.providerMetaConfig)

	log.Info("readDataSource done")

	return &kfplugin1.ReadDataSource_Response{
		Diagnostics: diags,
		Obj:         d,
	}, nil
}

func (r *GRPCProviderServer) ListDataSource(ctx context.Context, req *kfplugin1.ListDataSource_Request) (*kfplugin1.ListDataSource_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("listDataSource...")

	res, ok := r.provider.ListDataSourcesMap[req.GetName()]
	if !ok {
		return &kfplugin1.ListDataSource_Response{
			Diagnostics: diag.Errorf("cannot list data source, resourceType not found, got: %s", req.GetName()),
		}, nil
	}

	if res.ListContext == nil {
		return &kfplugin1.ListDataSource_Response{
			Diagnostics: diag.Errorf("cannot list data source, listContext not initialized, for: %s", req.GetName()),
		}, nil
	}

	obj, diags := res.ListContext(ctx, &ResourceObject{Scope: req.Scope, Obj: req.Obj}, r.provider.providerMetaConfig)

	log.Info("listDataSource done")

	return &kfplugin1.ListDataSource_Response{
		Diagnostics: diags,
		Obj:         obj,
	}, nil
}

func (r *GRPCProviderServer) ReadResource(ctx context.Context, req *kfplugin1.ReadResource_Request) (*kfplugin1.ReadResource_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("readResource...")

	res, ok := r.provider.ResourceMap[req.GetName()]
	if !ok {
		return &kfplugin1.ReadResource_Response{
			Diagnostics: diag.Errorf("cannot read resource, resourceType not found, got: %s", req.GetName()),
		}, nil
	}

	if res.ReadContext == nil {
		return &kfplugin1.ReadResource_Response{
			Diagnostics: diag.Errorf("cannot read resource, readContext not initialized, for: %s", req.GetName()),
		}, nil
	}

	obj, diags := res.ReadContext(ctx, &ResourceObject{Scope: req.Scope, Obj: req.Obj}, r.provider.providerMetaConfig)

	log.Info("readResource done")

	return &kfplugin1.ReadResource_Response{
		Diagnostics: diags,
		Obj:         obj,
	}, nil
}

func (r *GRPCProviderServer) CreateResource(ctx context.Context, req *kfplugin1.CreateResource_Request) (*kfplugin1.CreateResource_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("createResource...")

	res, ok := r.provider.ResourceMap[req.GetName()]
	if !ok {
		return &kfplugin1.CreateResource_Response{
			Diagnostics: diag.Errorf("cannot create resource, resourceType not found, got: %s", req.GetName()),
		}, nil
	}

	if res.CreateContext == nil {
		return &kfplugin1.CreateResource_Response{
			Diagnostics: diag.Errorf("cannot create resource, createContext not initialized, for: %s", req.GetName()),
		}, nil
	}

	obj, diags := res.CreateContext(ctx, &ResourceObject{Scope: req.Scope, DryRun: req.DryRun, Obj: req.Obj}, r.provider.providerMetaConfig)

	log.Info("createResource done")

	return &kfplugin1.CreateResource_Response{
		Diagnostics: diags,
		Obj:         obj,
	}, nil
}

func (r *GRPCProviderServer) UpdateResource(ctx context.Context, req *kfplugin1.UpdateResource_Request) (*kfplugin1.UpdateResource_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("updateResource...")

	res, ok := r.provider.ResourceMap[req.GetName()]
	if !ok {
		return &kfplugin1.UpdateResource_Response{
			Diagnostics: diag.Errorf("cannot update resource, resourceType not found, got: %s", req.GetName()),
		}, nil
	}

	if res.UpdateContext == nil {
		return &kfplugin1.UpdateResource_Response{
			Diagnostics: diag.Errorf("cannot update resource, updateContext not initialized, for: %s", req.GetName()),
		}, nil
	}

	obj, diags := res.UpdateContext(ctx, &ResourceObject{Scope: req.Scope, DryRun: req.DryRun, Obj: req.NewObj, OldObj: req.OldObj}, r.provider.providerMetaConfig)

	log.Info("updateResource done")

	return &kfplugin1.UpdateResource_Response{
		Diagnostics: diags,
		Obj:         obj,
	}, nil
}

func (r *GRPCProviderServer) DeleteResource(ctx context.Context, req *kfplugin1.DeleteResource_Request) (*kfplugin1.DeleteResource_Response, error) {
	// todo add ctx + tracing
	log := log.FromContext(ctx)
	log.Info("deleteResource...")

	res, ok := r.provider.ResourceMap[req.GetName()]
	if !ok {
		return &kfplugin1.DeleteResource_Response{
			Diagnostics: diag.Errorf("cannot delete resource, resourceType not found, got: %s", req.GetName()),
		}, nil
	}

	if res.DeleteContext == nil {
		return &kfplugin1.DeleteResource_Response{
			Diagnostics: diag.Errorf("cannot delete resource, createContext not initialized, for: %s", req.GetName()),
		}, nil
	}

	diags := res.DeleteContext(ctx, &ResourceObject{Scope: req.Scope, DryRun: req.DryRun, Obj: req.Obj}, r.provider.providerMetaConfig)

	log.Info("deleteResource done")

	return &kfplugin1.DeleteResource_Response{
		Diagnostics: diags,
	}, nil
}
