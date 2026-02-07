package clay

const (
	DefaultAPIBase         = "https://api.clay.com/v3"
	DefaultFrontendVersion = "v20260205_151537Z_19e7945c5e"

	MixrankActionPackageID = "e251a70e-46d7-4f3a-b3ef-a211ad3d8bd2"
	MixrankActionKey       = "find-lists-of-companies-with-mixrank-source"

	GoogleMapsActionPackageID = "3282a1c7-6bb0-497e-a34b-32268e104e55"
	GoogleMapsActionKey       = "google-review-source-v3"
)

type CreateWorkbookParams struct {
	Name     string
	FolderID string
}

type CreateWorkbookResult struct {
	ID   string
	Name string
}

type SearchCompaniesParams struct {
	WorkbookID         string
	Keywords           []string
	Countries          []string
	CompanySizes       []string
	AnnualRevenues     []string
	MinLinkedInMembers *int
	MaxLinkedInMembers *int
	Limit              *int
}

type SearchCompaniesResult struct {
	TableID     string
	TableName   string
	RecordCount float64
}

type SearchBusinessesParams struct {
	WorkbookID    string
	Latitude      float64
	Longitude     float64
	ProximityKm   float64
	BusinessTypes []string
	Query         string
	NumResults    int
	TableName     string
	TableEmoji    string
}

type SearchBusinessesResult struct {
	TableID   string
	TableName string
}
