package versioninfo

type Version struct {
	File    [4]uint
	Product [4]uint
}

func Get(fname string) (*Version, error) {
	return get(fname)
}
