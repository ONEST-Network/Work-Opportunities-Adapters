package workerProfile

// WorkerProfile represents a worker profile in the database
type WorkerProfile struct {
	ID                 string            `bson:"id"`
	Name               string            `bson:"name"`
	Phone              string            `bson:"phone"`
	Email              string            `bson:"email"`
	Age                int               `bson:"age"`
	Gender             Gender            `bson:"gender"`
	PreferredLanguages []Language        `bson:"preferred_languages"`
	PreferredJobRoles  []JobRole         `bson:"preferred_job_roles"`
	Experience         int               `bson:"experience"`
	Skills             []string          `bson:"skills"`
	Certifications     []Certification   `bson:"certifications"`
	Credentials        []Credential      `bson:"credentials"`
	Location           Location          `bson:"location"`
	TransactionID      string            `bson:"transaction_id"`
	MessageID          string            `bson:"message_id"`
	ApplicantionID     map[string]string `bson:"application_id"`
}

type Credential struct {
	Name      string `bson:"name"`
	URL       string `bson:"url"`
	Type      string `bson:"type"`
	ShortDesc string `bson:"short_desc"`
	LongDesc  string `bson:"long_desc"`
}

// Location represents the location of a job
type Location struct {
	Coordinates Coordinates `bson:"coordinates" json:"coordinates"`
	PostalCode  string      `bson:"postal_code" json:"postalCode"`
	City        string      `bson:"city" json:"city"`
	State       string      `bson:"state" json:"state"`
}

// Coordinates represents the latitude and longitude of a location
type Coordinates struct {
	Latitude  float64 `bson:"latitude" json:"latitude"`
	Longitute float64 `bson:"longitude" json:"longitude"`
}

// Gender represents the gender of a worker
type Gender string

const (
	GenderMale   Gender = "Male"
	GenderFemale Gender = "Female"
	GenderOther  Gender = "Other"
)

// Language represents the preferred language for communication
type Language string

const (
	LanguageEnglish Language = "English"
	LanguageHindi   Language = "Hindi"
	LanguageKannada Language = "Kannada"
	// Add more languages as needed
)

// JobRole represents different types of job roles
type JobRole string

const (
	JobRoleHelper JobRole = "Helper"
	JobRoleWelder JobRole = "Welder"
	JobRoleDriver JobRole = "Driver"
)

// Experience represents past work experience
type Experience struct {
	JobTitle    string `bson:"job_title" json:"jobTitle"`
	Company     string `bson:"company" json:"company"`
	StartDate   string `bson:"start_date" json:"startDate"`
	EndDate     string `bson:"end_date" json:"endDate"`
	DocumentURL string `bson:"document_url" json:"documentUrl"`
}

// Certification represents qualifications and certifications
type Certification struct {
	Name        string `bson:"name" json:"name"`
	IssuedBy    string `bson:"issued_by" json:"issuedBy"`
	IssuedDate  string `bson:"issued_date" json:"issuedDate"`
	ExpiryDate  string `bson:"expiry_date" json:"expiryDate"`
	DocumentURL string `bson:"document_url" json:"documentUrl"`
}
