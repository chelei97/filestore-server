package meta

//FileMeta: 文件的一些信息
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init(){
	fileMetas = make(map[string]FileMeta)
}

//UpdateFileMeta: 新增/更新文件信息的操作
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//GetFileMeta: 获取文件对象
func GetFileMeta(fileSha1 string) FileMeta{
	return fileMetas[fileSha1]
}
