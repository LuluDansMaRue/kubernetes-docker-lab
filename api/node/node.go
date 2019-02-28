package node

import "os"

// GetNodeName return the node name
func GetNodeName() string {
	node := os.Getenv("NODE_NAME")

	if len(node) > 0 {
		return node
	}

	return "minikube / local"
}
