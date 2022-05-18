package interfaces

type DiskI interface {
	CreateFolder(input string) error
	UploadToCloud(input string) error
	MakeDirectoryPath(key string) string
	MakeCdnPath(key string) string
}
