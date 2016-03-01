package scaffolding

const (
	ConfigPath     = "scaffolding.json"
	ScaffoldingDir = ".scaffolding/"
	HistoryDir     = ScaffoldingDir + "history/"
	TemplatesDir   = ScaffoldingDir + "templates/"

	GenerateValuesPath  = ScaffoldingDir + "values"
	GenerateSecretsPath = ScaffoldingDir + "secrets"

	CreateMode    = 0644
	CreateDirMode = 0777
)
