package etcdstore

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/coreos/etcd/client"
	"gitlab.ricebook.net/platform/core/types"
	"gitlab.ricebook.net/platform/core/utils"
)

// get a pod from etcd
// storage path in etcd is `/eru-core/pod/:podname/info`
func (k *krypton) GetPod(name string) (*types.Pod, error) {
	key := fmt.Sprintf(podInfoKey, name)
	resp, err := k.etcd.Get(context.Background(), key, nil)
	if err != nil {
		return nil, err
	}
	if resp.Node.Dir {
		return nil, fmt.Errorf("Pod storage path %q in etcd is a directory", key)
	}

	pod := &types.Pod{}
	if err := json.Unmarshal([]byte(resp.Node.Value), pod); err != nil {
		return nil, err
	}

	return pod, nil
}

// add a pod
// save it to etcd
// storage path in etcd is `/eru-core/pod/:podname/info`
func (k *krypton) AddPod(name, favor, desc string) (*types.Pod, error) {
	key := fmt.Sprintf(podInfoKey, name)
	favor = strings.ToUpper(favor)
	if favor == "" {
		favor = types.MEMORY_PRIOR
	} else if favor != types.MEMORY_PRIOR && favor != types.CPU_PRIOR {
		return nil, fmt.Errorf("favor should be either CPU or MEM, got %s", favor)
	}
	pod := &types.Pod{Name: name, Desc: desc, Favor: favor}

	bytes, err := json.Marshal(pod)
	if err != nil {
		return nil, err
	}

	_, err = k.etcd.Create(context.Background(), key, string(bytes))
	if err != nil {
		return nil, err
	}

	return pod, nil
}

// get all pods in etcd
// any error will break and return error immediately
// storage path in etcd is `/eru-core/pod`
func (k *krypton) GetAllPods() ([]*types.Pod, error) {
	var (
		pods []*types.Pod
		err  error
	)

	resp, err := k.etcd.Get(context.Background(), allPodsKey, nil)
	if err != nil {
		return pods, err
	}
	if !resp.Node.Dir {
		return nil, fmt.Errorf("Pod storage path %q in etcd is not a directory", allPodsKey)
	}

	for _, node := range resp.Node.Nodes {
		name := utils.Tail(node.Key)
		p, err := k.GetPod(name)
		if err != nil {
			return pods, err
		}
		pods = append(pods, p)
	}
	return pods, err
}

// DeletePod if the pod has no nodes left, otherwise return an error
func (k *krypton) DeletePod(podname string, force bool) error {
	key := fmt.Sprintf("%s/%s", allPodsKey, podname)

	ns, err := k.GetNodesByPod(podname)
	if err != nil && !client.IsKeyNotFound(err) {
		return err
	}
	if len(ns) != 0 && force == false {
		return fmt.Errorf("[DeletePod] pod %s still has nodes, delete the nodes first", podname)
	}

	_, err = k.etcd.Delete(context.Background(), key, &client.DeleteOptions{Dir: true, Recursive: true})
	return err
}
