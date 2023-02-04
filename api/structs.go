package api

import "fmt"

type IPGeolocation struct {
	IP                        string           `json:"ip"`
	LocalityLanguageRequested string           `json:"localityLanguageRequested"`
	IsReachableGlobally       bool             `json:"isReachableGlobally"`
	Country                   *Country         `json:"country"`
	Location                  *Location        `json:"location"`
	LastUpdated               string           `json:"lastUpdated"`
	Network                   *Network         `json:"network"`
	Confidence                *string          `json:"confidence,omitempty"`
	ConfidenceArea            []ConfidenceArea `json:"confidenceArea,omitempty"`
	SecurityThreat            *string          `json:"securityThreat,omitempty"`
	HazardReport              *HazardReport    `json:"hazardReport,omitempty"`
}

type ConfidenceArea struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Country struct {
	ISOAlpha2         string              `json:"isoAlpha2"`
	ISOAlpha3         string              `json:"isoAlpha3"`
	M49Code           int64               `json:"m49Code"`
	Name              string              `json:"name"`
	ISOName           string              `json:"isoName"`
	ISONameFull       string              `json:"isoNameFull"`
	ISOAdminLanguages []*ISOAdminLanguage `json:"isoAdminLanguages"`
	UnRegion          string              `json:"unRegion"`
	Currency          *Currency           `json:"currency"`
	WbRegion          *WBNode             `json:"wbRegion"`
	WbIncomeLevel     *WBNode             `json:"wbIncomeLevel"`
	CallingCode       string              `json:"callingCode"`
	CountryFlagEmoji  string              `json:"countryFlagEmoji"`
	WikidataID        string              `json:"wikidataId"`
	GeonameID         int64               `json:"geonameId"`
	Continents        []*Continent        `json:"continents"`
}

type Continent struct {
	Continent     string `json:"continent"`
	ContinentCode string `json:"continentCode"`
}

type Currency struct {
	NumericCode int64  `json:"numericCode"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	MinorUnits  int64  `json:"minorUnits"`
}

type ISOAdminLanguage struct {
	ISOAlpha3  string `json:"isoAlpha3"`
	ISOAlpha2  string `json:"isoAlpha2"`
	ISOName    string `json:"isoName"`
	NativeName string `json:"nativeName"`
}

type WBNode struct {
	ID       string `json:"id"`
	Iso2Code string `json:"iso2Code"`
	Value    string `json:"value"`
}

type HazardReport struct {
	IsKnownAsTorServer       bool  `json:"isKnownAsTorServer"`
	IsKnownAsVPN             bool  `json:"isKnownAsVpn"`
	IsKnownAsProxy           bool  `json:"isKnownAsProxy"`
	IsSpamhausDrop           bool  `json:"isSpamhausDrop"`
	IsSpamhausEdrop          bool  `json:"isSpamhausEdrop"`
	IsSpamhausAsnDrop        bool  `json:"isSpamhausAsnDrop"`
	IsBlacklistedUceprotect  bool  `json:"isBlacklistedUceprotect"`
	IsBlacklistedBlocklistDe bool  `json:"isBlacklistedBlocklistDe"`
	IsKnownAsMailServer      bool  `json:"isKnownAsMailServer"`
	IsKnownAsPublicRouter    bool  `json:"isKnownAsPublicRouter"`
	IsBogon                  bool  `json:"isBogon"`
	IsUnreachable            bool  `json:"isUnreachable"`
	HostingLikelihood        int64 `json:"hostingLikelihood"`
	IsHostingAsn             bool  `json:"isHostingAsn"`
	IsCellular               bool  `json:"isCellular"`
}

type Location struct {
	Continent                   string        `json:"continent"`
	ContinentCode               string        `json:"continentCode"`
	ISOPrincipalSubdivision     string        `json:"isoPrincipalSubdivision"`
	ISOPrincipalSubdivisionCode string        `json:"isoPrincipalSubdivisionCode"`
	City                        string        `json:"city"`
	LocalityName                string        `json:"localityName"`
	Postcode                    string        `json:"postcode"`
	Latitude                    float64       `json:"latitude"`
	Longitude                   float64       `json:"longitude"`
	PlusCode                    string        `json:"plusCode"`
	TimeZone                    *TimeZone     `json:"timeZone"`
	LocalityInfo                *LocalityInfo `json:"localityInfo"`
}

type LocalityInfo struct {
	Administrative []*Locality `json:"administrative"`
	Informative    []*Locality `json:"informative"`
}

type Locality struct {
	Order       int64   `json:"order"`
	AdminLevel  *int64  `json:"adminLevel,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	ISOName     *string `json:"isoName,omitempty"`
	ISOCode     *string `json:"isoCode,omitempty"`
	WikidataID  *string `json:"wikidataId,omitempty"`
	GeonameID   *int64  `json:"geonameId,omitempty"`
}

type TimeZone struct {
	IANATimeID             string `json:"ianaTimeId"`
	DisplayName            string `json:"displayName"`
	EffectiveTimeZoneFull  string `json:"effectiveTimeZoneFull"`
	EffectiveTimeZoneShort string `json:"effectiveTimeZoneShort"`
	UTCOffsetSeconds       int64  `json:"utcOffsetSeconds"`
	UTCOffset              string `json:"utcOffset"`
	IsDaylightSavingTime   bool   `json:"isDaylightSavingTime"`
	LocalTime              string `json:"localTime"`
}

type Network struct {
	Registry                string     `json:"registry"`
	RegistryStatus          string     `json:"registryStatus"`
	RegisteredCountry       string     `json:"registeredCountry"`
	RegisteredCountryName   string     `json:"registeredCountryName"`
	Organisation            string     `json:"organisation"`
	IsReachableGlobally     bool       `json:"isReachableGlobally"`
	IsBogon                 bool       `json:"isBogon"`
	BGPPrefix               string     `json:"bgpPrefix"`
	BGPPrefixNetworkAddress string     `json:"bgpPrefixNetworkAddress"`
	BGPPrefixLastAddress    string     `json:"bgpPrefixLastAddress"`
	TotalAddresses          int64      `json:"totalAddresses"`
	Carriers                []*Carrier `json:"carriers"`
	ViaCarriers             []*Carrier `json:"viaCarriers"`
}

type Carrier struct {
	Asn                    string           `json:"asn"`
	AsnNumeric             int64            `json:"asnNumeric"`
	Organisation           string           `json:"organisation"`
	Name                   string           `json:"name"`
	Registry               string           `json:"registry"`
	RegisteredCountry      string           `json:"registeredCountry"`
	RegisteredCountryName  string           `json:"registeredCountryName"`
	RegistrationDate       *string          `json:"registrationDate,omitempty"`
	RegistrationLastChange *string          `json:"registrationLastChange,omitempty"`
	TotalIpv4Addresses     int64            `json:"totalIpv4Addresses"`
	TotalIpv4Prefixes      *int64           `json:"totalIpv4Prefixes,omitempty"`
	TotalIpv4BogonPrefixes *int64           `json:"totalIpv4BogonPrefixes,omitempty"`
	Rank                   int64            `json:"rank"`
	RankText               string           `json:"rankText"`
	TotalReceivingFrom     int64            `json:"totalReceivingFrom,omitempty"`
	ReceivingFrom          []*Carrier       `json:"receivingFrom,omitempty"`
	TotalTransitTo         int64            `json:"totalTransitTo,omitempty"`
	TransitTo              []*Carrier       `json:"transitTo,omitempty"`
	ConfidenceArea         []ConfidenceArea `json:"confidenceArea,omitempty"`
}

type ReverseGeolocation struct {
	Latitude                  float64       `json:"latitude"`
	Longitude                 float64       `json:"longitude"`
	PlusCode                  string        `json:"plusCode"`
	LocalityLanguageRequested string        `json:"localityLanguageRequested"`
	Continent                 string        `json:"continent"`
	ContinentCode             string        `json:"continentCode"`
	CountryName               string        `json:"countryName"`
	CountryCode               string        `json:"countryCode"`
	PrincipalSubdivision      string        `json:"principalSubdivision"`
	PrincipalSubdivisionCode  string        `json:"principalSubdivisionCode"`
	City                      string        `json:"city"`
	Locality                  string        `json:"locality"`
	Postcode                  string        `json:"postcode"`
	LocalityInfo              *LocalityInfo `json:"localityInfo"`
	TimeZone                  *TimeZone     `json:"timeZone,omitempty"`
}

type PhoneValidation struct {
	IsValid             bool     `json:"isValid"`
	E164Format          string   `json:"e164Format,omitempty"`
	InternationalFormat string   `json:"internationalFormat,omitempty"`
	NationalFormat      string   `json:"nationalFormat,omitempty"`
	Location            string   `json:"location,omitempty"`
	LineType            string   `json:"lineType,omitempty"`
	Country             *Country `json:"country,omitempty"`
}

type EmailVerification struct {
	InputData            string `json:"inputData"`
	IsValid              bool   `json:"isValid"`
	IsSyntaxValid        bool   `json:"isSyntaxValid"`
	IsMailServerDefined  bool   `json:"isMailServerDefined"`
	IsKnownSpammerDomain bool   `json:"isKnownSpammerDomain"`
	IsDisposable         bool   `json:"isDisposable"`
}

type UserAgentInfo struct {
	Device           string `json:"device"`
	OS               string `json:"os"`
	UserAgent        string `json:"userAgent"`
	Family           string `json:"family"`
	VersionMajor     string `json:"versionMajor"`
	VersionMinor     string `json:"versionMinor"`
	IsSpider         bool   `json:"isSpider"`
	IsMobile         bool   `json:"isMobile"`
	UserAgentDisplay string `json:"userAgentDisplay"`
}

type UserRisk struct {
	Risk        string `json:"risk"`
	Description string `json:"description"`
}

type ErrorResponse struct {
	ID          string `json:"$id"`
	Status      int64  `json:"status"`
	Description string `json:"description"`
}

func (err *ErrorResponse) Error() string {
	return fmt.Sprintf("[%d] %s", err.Status, err.Description)
}

type IPStatsPerCountry struct {
	FullUpdate                string          `json:"fullUpdate"`
	PartialUpdate             string          `json:"partialUpdate"`
	TotalReachableAddresses   int64           `json:"totalReachableAddresses"`
	TotalHighQuality          int64           `json:"totalHighQuality"`
	Countries                 []*CountryStats `json:"countries"`
	TotalModerateQuality      int64           `json:"totalModerateQuality"`
	TotalLowQuality           int64           `json:"totalLowQuality"`
	Unresolved                int64           `json:"unresolved"`
	LocalityLanguageRequested string          `json:"localityLanguageRequested"`
}

type CountryStats struct {
	Rank                 int64  `json:"rank"`
	CountryName          string `json:"countryName"`
	CountryCode          string `json:"countryCode"`
	TotalAddresses       int64  `json:"totalAddresses"`
	TotalHighQuality     int64  `json:"totalHighQuality"`
	TotalModerateQuality int64  `json:"totalModerateQuality"`
	TotalLowQuality      int64  `json:"totalLowQuality"`
}
