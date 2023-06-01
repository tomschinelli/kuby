package cluster

func (c *Cluster) Logs() error {
	if c.err != nil {
		return c.err
	}

	//c.clientset.AppsV1().Deployments("")
	return nil
}
