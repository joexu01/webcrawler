package config

const (
	// Parser Names
	ParseSchool         = "ParseSchool"
	ParseSchoolList     = "ParseSchoolList"
	ParseTeacherProfile = "ParseTeacherProfile"
	NilParser = "NilParser"

	ItemSaverPort = 1234
	ElasticIndex = "teacher"

	WorkerPort0 = 9000

	// RPC Endpoints
	ItemSaverRpc = "ItemSaveService.Save"
	CrawlServiceRpc = "CrawlService.Process"
)
