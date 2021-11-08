package constants

const (
	baseURL      = "https://api.bigdatacloud.net"
	endpointData = baseURL + "/data"

	// endpointGeolocationAPIs https://www.bigdatacloud.com/ip-geolocation-apis
	endpointGeolocationAPIs = endpointData + "/ip-geolocation"
	// EndpointIPGeolocationFull https://www.bigdatacloud.com/ip-geolocation-apis/ip-address-geolocation-with-confidence-area-and-hazard-report-api
	EndpointIPGeolocationFull = endpointGeolocationAPIs + "-full"
	// EndpointIPGeolocationWithConfidence https://www.bigdatacloud.com/ip-geolocation-apis/ip-address-geolocation-with-confidence-area-api
	EndpointIPGeolocationWithConfidence = endpointGeolocationAPIs + "-with-confidence"
	// EndpointIPGeolocation https://www.bigdatacloud.com/ip-geolocation-apis/ip-address-geolocation-api
	EndpointIPGeolocation = endpointGeolocationAPIs

	// endpointReverseGeocode https://www.bigdatacloud.com/geocoding-apis
	endpointReverseGeocode = endpointData + "/reverse-geocode"
	// EndpointReverseGeocodeWithTimezone https://www.bigdatacloud.com/geocoding-apis/reverse-geocode-with-timezone
	EndpointReverseGeocodeWithTimezone = endpointReverseGeocode + "-with-timezone"
	// EndpointReverseGeocode https://www.bigdatacloud.com/geocoding-apis/reverse-geocode-to-city-api
	EndpointReverseGeocode = endpointReverseGeocode
	// EndpointReverseGeocodeClient https://www.bigdatacloud.com/geocoding-apis/free-reverse-geocode-to-city-api
	EndpointReverseGeocodeClient = endpointReverseGeocode + "-client" // Don't implement this in Golang, makes no sense

	// endpointCountry https://www.bigdatacloud.com/country-info-apis
	endpointCountry = endpointData + "/country"
	// EndpointCountryByIP https://www.bigdatacloud.com/country-info-apis/country-by-ip-address-api
	EndpointCountryByIP = endpointCountry + "-by-ip"
	// EndpointCountryInfo https://www.bigdatacloud.com/country-info-apis/country-info-api
	EndpointCountryInfo = endpointCountry + "-info"

	// https://www.bigdatacloud.com/network-apis
	endpointNetwork = endpointData + "/network"
	// EndpointPrefixesList https://www.bigdatacloud.com/network-apis/bgp-ipv4-active-prefixes-api
	EndpointPrefixesList = endpointData + "/prefixes-list"
	// EndpointNetworkByCIDR https://www.bigdatacloud.com/network-apis/networks-by-cidr-api
	EndpointNetworkByCIDR = endpointNetwork + "-by-cidr"
	// EndpointNetworkByIP https://www.bigdatacloud.com/network-apis/network-by-ip-address-api
	EndpointNetworkByIP = endpointNetwork + "-by-ip"

	// endpointASNInfo https://www.bigdatacloud.com/autonomous-system-info-apis
	endpointASNInfo = endpointData + "/asn-info"
	// EndpointASNInfoFull https://www.bigdatacloud.com/autonomous-system-info-apis/asn-info-extended-api
	EndpointASNInfoFull = endpointASNInfo + "-full"
	// EndpointASNInfo https://www.bigdatacloud.com/autonomous-system-info-apis/asn-short-info-api
	EndpointASNInfo = endpointASNInfo

	// endpointTimezone https://www.bigdatacloud.com/time-zone-apis
	endpointTimezone = endpointData + "/timezone"
	// EndpointTimezoneInfo https://www.bigdatacloud.com/time-zone-apis/timezone-info-api
	EndpointTimezoneInfo = endpointTimezone + "-info"
	// EndpointTimezoneByLocation https://www.bigdatacloud.com/time-zone-apis/timezone-by-location-api
	EndpointTimezoneByLocation = endpointTimezone + "-by-location"
	// EndpointTimezoneByIP https://www.bigdatacloud.com/time-zone-apis/timezone-by-ip-address-api
	EndpointTimezoneByIP = endpointTimezone + "-by-ip"

	// EndpointTORExitNodes https://www.bigdatacloud.com/insights-apis
	// EndpointTORExitNodes https://www.bigdatacloud.com/insights-apis/tor-exit-nodes-geolocated-api
	EndpointTORExitNodes = endpointData + "/tor-exit-nodes-list"
	// EndpointAddressSpaceIPv4 https://www.bigdatacloud.com/insights-apis/ipv4-address-space-monitoring-api
	EndpointAddressSpaceIPv4 = endpointData + "/address-space-stats-ipv4"

	// endpointClient https://www.bigdatacloud.com/client-info-apis
	endpointClient = endpointData + "/client"
	// EndpointClientIP https://www.bigdatacloud.com/client-info-apis/public-ip-address-api
	EndpointClientIP = endpointClient + "-ip"
	// EndpointUserAgentInfo https://www.bigdatacloud.com/client-info-apis/user-agent-parser-api
	EndpointUserAgentInfo = endpointData + "/user-agent-info"
	// EndpointClientInfo https://www.bigdatacloud.com/client-info-apis/client-info-api
	EndpointClientInfo = endpointClient + "-info"
	// EndpointAmIRoaming https://www.bigdatacloud.com/client-info-apis/free-am-i-roaming-api
	EndpointAmIRoaming = endpointData + "/am-i-roaming"

	// endpointPhoneNumber https://www.bigdatacloud.com/phone-number-apis
	endpointPhoneNumber = endpointData + "/phone-number"
	// EndpointPhoneNumberValidateByIP https://www.bigdatacloud.com/phone-number-apis/phone-number-validation-by-ip-address-api
	EndpointPhoneNumberValidateByIP = endpointPhoneNumber + "-validate-by-ip"
	// EndpointPhoneNumberValidate https://www.bigdatacloud.com/phone-number-apis/phone-number-validation-api
	EndpointPhoneNumberValidate = endpointPhoneNumber + "-validate"

	// EndpointEmailVerify https://www.bigdatacloud.com/email-address-validation-apis
	// EndpointEmailVerify https://www.bigdatacloud.com/email-address-validation-apis/email-verify-api
	EndpointEmailVerify = endpointData + "/email-verify"

	DefaultUserAgent      = "github.com/BRUHItsABunny/bigdatacloud-apis"
	HeaderKeyUserAgent    = "user-agent"
	ParameterKeyIP        = "ip"
	ParameterKeyLongitude = "longitude"
	ParameterKeyLatitude  = "latitude"
	ParameterKeyLanguage  = "localityLanguage"
	ParameterKeyAPIKey    = "key"
)
