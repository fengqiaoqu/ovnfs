//go:build !darwin

package base

// locator is unused here: only macOS maps an inode to its path (fsgetpath(2)).
// Elsewhere locate uses the paths of snapshot entries the caches have already
// served (see [Caches.remember]).
type locator struct{}

// locate returns the tree-relative slash path of a served snapshot entry
// with inode ino.
func (c *Caches) locate(ino uint64) (string, bool) {
	c.servedMu.Lock()
	defer c.servedMu.Unlock()
	p, ok := c.served[ino]
	return p, ok
}
