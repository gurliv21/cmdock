package store

func (c Command) String() string{
	return c.Command + " " + c.Directory
}