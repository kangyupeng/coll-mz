package transferOldFiles

import "../core"

//该模块用于将collmz1.0数据文件向新版本迁移
//新版本文件格式将有所改变:
// * 每个文件、合集将对应有JSON数据，避免数据库遗失、迁移数据库后出现错误
// * JSON数据将包含更多文件、合集信息，这样可脱离数据库实现读写
// * 新版本中，数据迁移无需转移数据库，而是直接拷贝走文件数据即可

func Run(){
	core.SendLog("启动旧版本向新版本数据迁移模块。")
	//
	core.SendLog("迁移结束。")
}