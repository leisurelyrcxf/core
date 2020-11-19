// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import enginetypes "github.com/projecteru2/core/engine/types"
import mock "github.com/stretchr/testify/mock"
import types "github.com/projecteru2/core/types"

// Cluster is an autogenerated mock type for the Cluster type
type Cluster struct {
	mock.Mock
}

// AddNode provides a mock function with given fields: _a0, _a1
func (_m *Cluster) AddNode(_a0 context.Context, _a1 *types.AddNodeOptions) (*types.Node, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, *types.AddNodeOptions) *types.Node); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.AddNodeOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddPod provides a mock function with given fields: ctx, podname, desc
func (_m *Cluster) AddPod(ctx context.Context, podname string, desc string) (*types.Pod, error) {
	ret := _m.Called(ctx, podname, desc)

	var r0 *types.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *types.Pod); ok {
		r0 = rf(ctx, podname, desc)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, podname, desc)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BuildImage provides a mock function with given fields: ctx, opts
func (_m *Cluster) BuildImage(ctx context.Context, opts *types.BuildOptions) (chan *types.BuildImageMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.BuildImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.BuildOptions) chan *types.BuildImageMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.BuildImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.BuildOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CacheImage provides a mock function with given fields: ctx, podname, nodenames, images, step
func (_m *Cluster) CacheImage(ctx context.Context, podname string, nodenames []string, images []string, step int) (chan *types.CacheImageMessage, error) {
	ret := _m.Called(ctx, podname, nodenames, images, step)

	var r0 chan *types.CacheImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, []string, int) chan *types.CacheImageMessage); ok {
		r0 = rf(ctx, podname, nodenames, images, step)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.CacheImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, []string, int) error); ok {
		r1 = rf(ctx, podname, nodenames, images, step)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CalculateCapacity provides a mock function with given fields: _a0, _a1
func (_m *Cluster) CalculateCapacity(_a0 context.Context, _a1 *types.DeployOptions) (*types.CapacityMessage, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *types.CapacityMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions) *types.CapacityMessage); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.CapacityMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ConnectNetwork provides a mock function with given fields: ctx, network, target, ipv4, ipv6
func (_m *Cluster) ConnectNetwork(ctx context.Context, network string, target string, ipv4 string, ipv6 string) ([]string, error) {
	ret := _m.Called(ctx, network, target, ipv4, ipv6)

	var r0 []string
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, string) []string); ok {
		r0 = rf(ctx, network, target, ipv4, ipv6)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, string) error); ok {
		r1 = rf(ctx, network, target, ipv4, ipv6)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ContainerStatusStream provides a mock function with given fields: ctx, appname, entrypoint, nodename, labels
func (_m *Cluster) ContainerStatusStream(ctx context.Context, appname string, entrypoint string, nodename string, labels map[string]string) chan *types.ContainerStatus {
	ret := _m.Called(ctx, appname, entrypoint, nodename, labels)

	var r0 chan *types.ContainerStatus
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, map[string]string) chan *types.ContainerStatus); ok {
		r0 = rf(ctx, appname, entrypoint, nodename, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.ContainerStatus)
		}
	}

	return r0
}

// ControlContainer provides a mock function with given fields: ctx, IDs, t, force
func (_m *Cluster) ControlContainer(ctx context.Context, IDs []string, t string, force bool) (chan *types.ControlContainerMessage, error) {
	ret := _m.Called(ctx, IDs, t, force)

	var r0 chan *types.ControlContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, []string, string, bool) chan *types.ControlContainerMessage); ok {
		r0 = rf(ctx, IDs, t, force)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.ControlContainerMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, string, bool) error); ok {
		r1 = rf(ctx, IDs, t, force)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Copy provides a mock function with given fields: ctx, opts
func (_m *Cluster) Copy(ctx context.Context, opts *types.CopyOptions) (chan *types.CopyMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.CopyMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.CopyOptions) chan *types.CopyMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.CopyMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.CopyOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateContainer provides a mock function with given fields: ctx, opts
func (_m *Cluster) CreateContainer(ctx context.Context, opts *types.DeployOptions) (chan *types.CreateContainerMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.CreateContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions) chan *types.CreateContainerMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.CreateContainerMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DisconnectNetwork provides a mock function with given fields: ctx, network, target, force
func (_m *Cluster) DisconnectNetwork(ctx context.Context, network string, target string, force bool) error {
	ret := _m.Called(ctx, network, target, force)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) error); ok {
		r0 = rf(ctx, network, target, force)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DissociateContainer provides a mock function with given fields: ctx, IDs
func (_m *Cluster) DissociateContainer(ctx context.Context, IDs []string) (chan *types.DissociateContainerMessage, error) {
	ret := _m.Called(ctx, IDs)

	var r0 chan *types.DissociateContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, []string) chan *types.DissociateContainerMessage); ok {
		r0 = rf(ctx, IDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.DissociateContainerMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExecuteContainer provides a mock function with given fields: ctx, opts, inCh
func (_m *Cluster) ExecuteContainer(ctx context.Context, opts *types.ExecuteContainerOptions, inCh <-chan []byte) chan *types.AttachContainerMessage {
	ret := _m.Called(ctx, opts, inCh)

	var r0 chan *types.AttachContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ExecuteContainerOptions, <-chan []byte) chan *types.AttachContainerMessage); ok {
		r0 = rf(ctx, opts, inCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.AttachContainerMessage)
		}
	}

	return r0
}

// Finalizer provides a mock function with given fields:
func (_m *Cluster) Finalizer() {
	_m.Called()
}

// GetContainer provides a mock function with given fields: ctx, ID
func (_m *Cluster) GetContainer(ctx context.Context, ID string) (*types.Container, error) {
	ret := _m.Called(ctx, ID)

	var r0 *types.Container
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Container); ok {
		r0 = rf(ctx, ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainers provides a mock function with given fields: ctx, IDs
func (_m *Cluster) GetContainers(ctx context.Context, IDs []string) ([]*types.Container, error) {
	ret := _m.Called(ctx, IDs)

	var r0 []*types.Container
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*types.Container); ok {
		r0 = rf(ctx, IDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetContainersStatus provides a mock function with given fields: ctx, IDs
func (_m *Cluster) GetContainersStatus(ctx context.Context, IDs []string) ([]*types.StatusMeta, error) {
	ret := _m.Called(ctx, IDs)

	var r0 []*types.StatusMeta
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*types.StatusMeta); ok {
		r0 = rf(ctx, IDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.StatusMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, IDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNode provides a mock function with given fields: ctx, nodename
func (_m *Cluster) GetNode(ctx context.Context, nodename string) (*types.Node, error) {
	ret := _m.Called(ctx, nodename)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Node); ok {
		r0 = rf(ctx, nodename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, nodename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPod provides a mock function with given fields: ctx, podname
func (_m *Cluster) GetPod(ctx context.Context, podname string) (*types.Pod, error) {
	ret := _m.Called(ctx, podname)

	var r0 *types.Pod
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.Pod); ok {
		r0 = rf(ctx, podname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, podname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListContainers provides a mock function with given fields: ctx, opts
func (_m *Cluster) ListContainers(ctx context.Context, opts *types.ListContainersOptions) ([]*types.Container, error) {
	ret := _m.Called(ctx, opts)

	var r0 []*types.Container
	if rf, ok := ret.Get(0).(func(context.Context, *types.ListContainersOptions) []*types.Container); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ListContainersOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNetworks provides a mock function with given fields: ctx, podname, driver
func (_m *Cluster) ListNetworks(ctx context.Context, podname string, driver string) ([]*enginetypes.Network, error) {
	ret := _m.Called(ctx, podname, driver)

	var r0 []*enginetypes.Network
	if rf, ok := ret.Get(0).(func(context.Context, string, string) []*enginetypes.Network); ok {
		r0 = rf(ctx, podname, driver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*enginetypes.Network)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, podname, driver)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListNodeContainers provides a mock function with given fields: ctx, nodename, labels
func (_m *Cluster) ListNodeContainers(ctx context.Context, nodename string, labels map[string]string) ([]*types.Container, error) {
	ret := _m.Called(ctx, nodename, labels)

	var r0 []*types.Container
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string) []*types.Container); ok {
		r0 = rf(ctx, nodename, labels)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Container)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string) error); ok {
		r1 = rf(ctx, nodename, labels)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPodNodes provides a mock function with given fields: ctx, podname, labels, all
func (_m *Cluster) ListPodNodes(ctx context.Context, podname string, labels map[string]string, all bool) ([]*types.Node, error) {
	ret := _m.Called(ctx, podname, labels, all)

	var r0 []*types.Node
	if rf, ok := ret.Get(0).(func(context.Context, string, map[string]string, bool) []*types.Node); ok {
		r0 = rf(ctx, podname, labels, all)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, map[string]string, bool) error); ok {
		r1 = rf(ctx, podname, labels, all)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListPods provides a mock function with given fields: ctx
func (_m *Cluster) ListPods(ctx context.Context) ([]*types.Pod, error) {
	ret := _m.Called(ctx)

	var r0 []*types.Pod
	if rf, ok := ret.Get(0).(func(context.Context) []*types.Pod); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.Pod)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// LogStream provides a mock function with given fields: ctx, opts
func (_m *Cluster) LogStream(ctx context.Context, opts *types.LogStreamOptions) (chan *types.LogStreamMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.LogStreamMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.LogStreamOptions) chan *types.LogStreamMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.LogStreamMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.LogStreamOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeResource provides a mock function with given fields: ctx, nodename, fix
func (_m *Cluster) NodeResource(ctx context.Context, nodename string, fix bool) (*types.NodeResource, error) {
	ret := _m.Called(ctx, nodename, fix)

	var r0 *types.NodeResource
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *types.NodeResource); ok {
		r0 = rf(ctx, nodename, fix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.NodeResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, bool) error); ok {
		r1 = rf(ctx, nodename, fix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PodResource provides a mock function with given fields: ctx, podname
func (_m *Cluster) PodResource(ctx context.Context, podname string) (*types.PodResource, error) {
	ret := _m.Called(ctx, podname)

	var r0 *types.PodResource
	if rf, ok := ret.Get(0).(func(context.Context, string) *types.PodResource); ok {
		r0 = rf(ctx, podname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.PodResource)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, podname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReallocResource provides a mock function with given fields: ctx, opts
func (_m *Cluster) ReallocResource(ctx context.Context, opts *types.ReallocOptions) error {
	ret := _m.Called(ctx, opts)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *types.ReallocOptions) error); ok {
		r0 = rf(ctx, opts)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveContainer provides a mock function with given fields: ctx, IDs, force, step
func (_m *Cluster) RemoveContainer(ctx context.Context, IDs []string, force bool, step int) (chan *types.RemoveContainerMessage, error) {
	ret := _m.Called(ctx, IDs, force, step)

	var r0 chan *types.RemoveContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, []string, bool, int) chan *types.RemoveContainerMessage); ok {
		r0 = rf(ctx, IDs, force, step)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.RemoveContainerMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []string, bool, int) error); ok {
		r1 = rf(ctx, IDs, force, step)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveImage provides a mock function with given fields: ctx, podname, nodenames, images, step, prune
func (_m *Cluster) RemoveImage(ctx context.Context, podname string, nodenames []string, images []string, step int, prune bool) (chan *types.RemoveImageMessage, error) {
	ret := _m.Called(ctx, podname, nodenames, images, step, prune)

	var r0 chan *types.RemoveImageMessage
	if rf, ok := ret.Get(0).(func(context.Context, string, []string, []string, int, bool) chan *types.RemoveImageMessage); ok {
		r0 = rf(ctx, podname, nodenames, images, step, prune)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.RemoveImageMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, []string, []string, int, bool) error); ok {
		r1 = rf(ctx, podname, nodenames, images, step, prune)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveNode provides a mock function with given fields: ctx, nodename
func (_m *Cluster) RemoveNode(ctx context.Context, nodename string) error {
	ret := _m.Called(ctx, nodename)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, nodename)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemovePod provides a mock function with given fields: ctx, podname
func (_m *Cluster) RemovePod(ctx context.Context, podname string) error {
	ret := _m.Called(ctx, podname)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, podname)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ReplaceContainer provides a mock function with given fields: ctx, opts
func (_m *Cluster) ReplaceContainer(ctx context.Context, opts *types.ReplaceOptions) (chan *types.ReplaceContainerMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.ReplaceContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.ReplaceOptions) chan *types.ReplaceContainerMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.ReplaceContainerMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.ReplaceOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RunAndWait provides a mock function with given fields: ctx, opts, inCh
func (_m *Cluster) RunAndWait(ctx context.Context, opts *types.DeployOptions, inCh <-chan []byte) (<-chan *types.AttachContainerMessage, error) {
	ret := _m.Called(ctx, opts, inCh)

	var r0 <-chan *types.AttachContainerMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.DeployOptions, <-chan []byte) <-chan *types.AttachContainerMessage); ok {
		r0 = rf(ctx, opts, inCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *types.AttachContainerMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.DeployOptions, <-chan []byte) error); ok {
		r1 = rf(ctx, opts, inCh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Send provides a mock function with given fields: ctx, opts
func (_m *Cluster) Send(ctx context.Context, opts *types.SendOptions) (chan *types.SendMessage, error) {
	ret := _m.Called(ctx, opts)

	var r0 chan *types.SendMessage
	if rf, ok := ret.Get(0).(func(context.Context, *types.SendOptions) chan *types.SendMessage); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(chan *types.SendMessage)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.SendOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetContainersStatus provides a mock function with given fields: ctx, status, ttls
func (_m *Cluster) SetContainersStatus(ctx context.Context, status []*types.StatusMeta, ttls map[string]int64) ([]*types.StatusMeta, error) {
	ret := _m.Called(ctx, status, ttls)

	var r0 []*types.StatusMeta
	if rf, ok := ret.Get(0).(func(context.Context, []*types.StatusMeta, map[string]int64) []*types.StatusMeta); ok {
		r0 = rf(ctx, status, ttls)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.StatusMeta)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, []*types.StatusMeta, map[string]int64) error); ok {
		r1 = rf(ctx, status, ttls)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetNode provides a mock function with given fields: ctx, opts
func (_m *Cluster) SetNode(ctx context.Context, opts *types.SetNodeOptions) (*types.Node, error) {
	ret := _m.Called(ctx, opts)

	var r0 *types.Node
	if rf, ok := ret.Get(0).(func(context.Context, *types.SetNodeOptions) *types.Node); ok {
		r0 = rf(ctx, opts)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Node)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *types.SetNodeOptions) error); ok {
		r1 = rf(ctx, opts)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchServiceStatus provides a mock function with given fields: _a0
func (_m *Cluster) WatchServiceStatus(_a0 context.Context) (<-chan types.ServiceStatus, error) {
	ret := _m.Called(_a0)

	var r0 <-chan types.ServiceStatus
	if rf, ok := ret.Get(0).(func(context.Context) <-chan types.ServiceStatus); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan types.ServiceStatus)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
