package port

type Adapter interface {
	CreateBranchLink(linkData map[string]string) (string, error)
}