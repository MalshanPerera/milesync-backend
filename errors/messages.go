package errpkg

var (
	Unauthorized               = "Unauthorized"
	UserNotFound               = "UserNotFound"
	InvalidPassword            = "InvalidPassword"
	OrganizationNameExists     = "OrganizationNameExists"
	UserAlreadyHasOrganization = "UserAlreadyHasOrganization"
	OrganizationNameRequired   = "OrganizationNameRequired"
	OrganizationExists         = "OrganizationExists"
	ProjectExists              = "ProjectExists"
	ProjectKeyPrefixUsed       = "ProjectKeyPrefixUsed"
	InvalidProjectType         = "InvalidProjectType"
	SessionExpired             = "SessionExpired"
	KeyPrefixTooLong           = "KeyPrefixTooLong"
)
