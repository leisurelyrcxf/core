package virt

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	log "github.com/sirupsen/logrus"

	virtapi "github.com/projecteru2/yavirt/api/client"
	virtypes "github.com/projecteru2/yavirt/api/types"

	"github.com/projecteru2/core/cluster"
	enginetypes "github.com/projecteru2/core/engine/types"
	coresource "github.com/projecteru2/core/source"
	coretypes "github.com/projecteru2/core/types"
)

// Virt implements the core engine.API interface.
type Virt struct {
	client *virtapi.Client
	config coretypes.Config
}

// MakeClient makes a virt. client which wraps yavirt API client.
func MakeClient(config coretypes.Config, endpoint, apiversion string) (*Virt, error) {
	cli, err := virtapi.New(endpoint, apiversion)
	if err != nil {
		return nil, err
	}
	return &Virt{cli, config}, nil
}

// Info shows a connected node's information.
func (v *Virt) Info(ctx context.Context) (*enginetypes.Info, error) {
	resp, err := v.client.Info(ctx)
	if err != nil {
		return nil, err
	}

	return &enginetypes.Info{
		ID:       resp.ID,
		NCPU:     resp.Cpu,
		MemTotal: resp.Mem,
	}, nil
}

// ExecCreate creates an execution.
func (v *Virt) ExecCreate(ctx context.Context, target string, config *enginetypes.ExecConfig) (id string, err error) {
	log.Warnf("does not implement")
	return
}

// ExecAttach executes an attachment.
func (v *Virt) ExecAttach(ctx context.Context, execID string, detach, tty bool) (io.ReadCloser, error) {
	return nil, fmt.Errorf("does not implement")
}

// ExecExitCode gets return code of a specific execution.
func (v *Virt) ExecExitCode(ctx context.Context, execID string) (code int, err error) {
	log.Warnf("does not implement")
	return
}

// NetworkConnect connects to a network.
func (v *Virt) NetworkConnect(ctx context.Context, network, target, ipv4, ipv6 string) (err error) {
	log.Warnf("does not implement")
	return
}

// NetworkDisconnect disconnects from one network.
func (v *Virt) NetworkDisconnect(ctx context.Context, network, target string, force bool) (err error) {
	log.Warnf("does not implement")
	return
}

// NetworkList lists all of networks.
func (v *Virt) NetworkList(ctx context.Context, drivers []string) (nets []*enginetypes.Network, err error) {
	log.Warnf("does not implement")
	return
}

// BuildRefs builds references, it's not necessary for virt. presently.
func (v *Virt) BuildRefs(ctx context.Context, name string, tags []string) (refs []string) {
	log.Warnf("does not implement")
	return
}

// BuildContent builds content, the use of it is similar to BuildRefs.
func (v *Virt) BuildContent(ctx context.Context, scm coresource.Source, opts *enginetypes.BuildOptions) (io.ReadCloser, error) {
	return nil, fmt.Errorf("does not implement")
}

// VirtualizationCreate creates a guest.
func (v *Virt) VirtualizationCreate(ctx context.Context, opts *enginetypes.VirtualizationCreateOptions) (guest *enginetypes.VirtualizationCreated, err error) {
	req := virtypes.CreateGuestReq{
		Cpu:       int(opts.Quota),
		Mem:       opts.Memory,
		ImageName: opts.Image,
	}

	var resp virtypes.Guest
	if resp, err = v.client.CreateGuest(ctx, req); err != nil {
		return nil, err
	}

	return &enginetypes.VirtualizationCreated{ID: resp.ID, Name: opts.Name}, nil
}

// VirtualizationCopyTo copies one.
func (v *Virt) VirtualizationCopyTo(ctx context.Context, ID, path string, content io.Reader, AllowOverwriteDirWithFile, CopyUIDGID bool) (err error) {
	log.Warnf("does not implement")
	return
}

// VirtualizationStart boots a guest.
func (v *Virt) VirtualizationStart(ctx context.Context, ID string) (err error) {
	_, err = v.client.StartGuest(ctx, ID)
	return
}

// VirtualizationStop stops it.
func (v *Virt) VirtualizationStop(ctx context.Context, ID string) (err error) {
	_, err = v.client.StopGuest(ctx, ID)
	return
}

// VirtualizationRemove removes a guest.
func (v *Virt) VirtualizationRemove(ctx context.Context, ID string, volumes, force bool) (err error) {
	_, err = v.client.DestroyGuest(ctx, ID)
	return
}

// VirtualizationInspect gets a guest.
func (v *Virt) VirtualizationInspect(ctx context.Context, ID string) (*enginetypes.VirtualizationInfo, error) {
	guest, err := v.client.GetGuest(ctx, ID)
	if err != nil {
		return nil, err
	}

	bytes, err := json.Marshal(coretypes.EruMeta{Publish: []string{"PORT"}})
	if err != nil {
		return nil, err
	}

	return &enginetypes.VirtualizationInfo{
		ID:       guest.ID,
		Image:    guest.ImageName,
		Running:  guest.Status == "running",
		Networks: guest.Networks,
		Labels:   map[string]string{cluster.ERUMeta: string(bytes)},
	}, nil
}

// VirtualizationLogs streams a specific guest's log.
func (v *Virt) VirtualizationLogs(ctx context.Context, ID string, follow, stdout, stderr bool) (io.Reader, error) {
	return nil, fmt.Errorf("does not implement")
}

// VirtualizationAttach attaches something to a guest.
func (v *Virt) VirtualizationAttach(ctx context.Context, ID string, stream, stdin bool) (io.ReadCloser, io.WriteCloser, error) {
	return nil, nil, fmt.Errorf("does not implement")
}

// VirtualizationWait is waiting for a shut-off
func (v *Virt) VirtualizationWait(ctx context.Context, ID, state string) (*enginetypes.VirtualizationWaitResult, error) {
	return nil, fmt.Errorf("does not implement")
}

// VirtualizationUpdateResource updates resource.
func (v *Virt) VirtualizationUpdateResource(ctx context.Context, ID string, opts *enginetypes.VirtualizationResource) (err error) {
	log.Warnf("does not implement")
	return
}

// VirtualizationCopyFrom copies from another.
func (v *Virt) VirtualizationCopyFrom(ctx context.Context, ID, path string) (io.ReadCloser, string, error) {
	return nil, "", fmt.Errorf("does not implement")
}
