/*Package xal defines defines structs that conform to the OASIS xAL spec

Specification

You can find and download the spec here: https://www.oasis-open.org/committees/ciq/download.html

The entry point for a XAL address is the top level XAL struct.

Fields annotated with Attr are what used to be attribute fields in the XML formatted specification.
*/
package xal

type (
	// XAL - Root element for a list of addresses
	XAL struct {
		AttrVersion    string            `json:"attr_version,omitempty"` // Specific to DTD to specify the version number of DTD
		AddressDetails []*AddressDetails `json:"address_details,omitempty"`
	}

	// AddressDetails - This container defines the details of the address.
	// Can define multiple addresses including tracking address history
	AddressDetails struct {
		AttrAddressType    string              `json:"attr_address_type,omitempty"`    // maxLength=23
		AttrCurrentStatus  string              `json:"attr_current_status,omitempty"`  // maxLength=10
		AttrUsage          string              `json:"attr_usage,omitempty"`           // maxLength=6
		AttrValidFromDate  string              `json:"attr_valid_from_date,omitempty"` // maxLength=11
		AttrValidToDate    string              `json:"attr_valid_to_date,omitempty"`   // maxLength=13
		AddressLines       *AddressLines       `json:"address_lines,omitempty"`
		AdministrativeArea *AdministrativeArea `json:"administrative_area,omitempty"`
		Country            *Country            `json:"country,omitempty"`
		Locality           *Locality           `json:"locality,omitempty"`
	}

	// AddressLine - Free format address representation.
	// An address can have more than one line.
	// The order of the AddressLine elements must be preserved.
	AddressLine struct {
		AttrType string `json:"attr_type,omitempty"` // Defines the type of address line. eg. Street, Address Line 1, etc.
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// AddressLines - Container for Address lines
	AddressLines []*AddressLine

	// AdministrativeArea - Examples of administrative areas are provinces counties,
	// special regions (such as "Rijnmond"), etc.
	AdministrativeArea struct {
		AttrType               string                    `json:"attr_type,omitempty"`       // maxLength=8; Province or State or County or Kanton, etc
		AttrUsageType          string                    `json:"attr_usage_type,omitempty"` // Postal or Political - Sometimes locations must be distinguished between postal system, and physical locations as defined by a political system
		AttrIndicator          string                    `json:"attr_indicator,omitempty"`  // Erode (Dist) where (Dist) is the Indicator
		AdministrativeAreaName []*AdministrativeAreaName `json:"administrative_area_name,omitempty"`
		Locality               *Locality                 `json:"locality,omitempty"`
	}

	// AdministrativeAreaName - Name of the administrative area. eg. MI in USA, NSW in Australia
	AdministrativeAreaName struct {
		AttrType string `json:"attr_type,omitempty"` // maxLength=12
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// BuildingName - Specification of the name of a building.
	BuildingName struct {
		AttrType           string `json:"attr_type,omitempty"`
		AttrTypeOccurrence string `json:"attr_type_occurrence,omitempty"` // Occurrence of the building name before/after the type. eg. EGIS BUILDING where name appears before type
		AttrCode           string `json:"attr_code,omitempty"`            // Used by postal services to encode the name of the element.
		Text               string `json:"text,omitempty"`
	}

	// Country - Specification of a country
	Country struct {
		AdministrativeArea *AdministrativeArea `json:"administrative_area,omitempty"`
		CountryName        *CountryName        `json:"country_name,omitempty"`
		CountryNameCode    *CountryNameCode    `json:"country_name_code,omitempty"`
		Locality           *Locality           `json:"locality,omitempty"`
		Thoroughfare       *Thoroughfare       `json:"thoroughfare,omitempty"`
	}

	// CountryName - Specification of the name of a country.
	CountryName struct {
		AttrType string `json:"attr_type,omitempty"` // Old name, new name, etc
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// CountryNameCode - A country code according to the specified scheme
	//
	// Country code scheme possible values, but not limited to:
	//  iso.3166-2,
	//  iso.3166-3 for two and three character country codes.
	CountryNameCode struct {
		AttrScheme string `json:"attr_scheme,omitempty"`
		AttrCode   string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text       string `json:"text,omitempty"`
	}

	// Locality - Locality is one level lower than administrative area.
	// Eg.: cities, reservations and any other built-up areas.
	//
	// AttrUsageType: Postal or Political - Sometimes locations must be distinguished between postal system, and physical locations as defined by a political system
	Locality struct {
		AttrType          string             `json:"attr_type,omitempty"` // maxLength=8; Possible values not limited to: City, IndustrialEstate, etc
		AttrUsageType     string             `json:"attr_usage_type,omitempty"`
		AttrIndicator     string             `json:"attr_indicator,omitempty"` // Erode (Dist) where (Dist) is the Indicator
		DependentLocality *DependentLocality `json:"dependent_locality,omitempty"`
		LargeMailUser     *LargeMailUser     `json:"large_mail_user,omitempty"`
		LocalityName      []*LocalityName    `json:"locality_name,omitempty"`
		PostBox           *PostBox           `json:"post_box,omitempty"`
		PostOffice        *PostOffice        `json:"post_office,omitempty"`
		PostalCode        *PostalCode        `json:"postal_code,omitempty"`
		Premise           *Premise           `json:"premise,omitempty"`
		Thoroughfare      *Thoroughfare      `json:"thoroughfare,omitempty"`
	}

	// LocalityName - Name of the locality
	LocalityName struct {
		AttrType string `json:"attr_type,omitempty"` // maxLength=12
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// Department - Subdivision in the firm: School of Physics at Victoria University (School of Physics is the department)
	Department struct {
		AttrType       string          `json:"attr_type,omitempty"` // School in Physics School, Division in Radiology division of school of physics
		DepartmentName *DepartmentName `json:"department_name,omitempty"`
	}

	// DepartmentName - Specification of the name of a department.
	DepartmentName struct {
		AttrType string `json:"attr_type,omitempty"`
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// DependentLocality - Dependent localities are Districts within cities/towns, locality divisions,
	// postal divisions of cities, suburbs, etc.
	//
	// DependentLocality is a recursive element,
	// but no nesting deeper than two exists (Locality-DependentLocality-DependentLocality).
	//
	// AttrUsageType: Postal or Political - Sometimes locations must be distinguished between postal system, and physical locations as defined by a political system
	//
	// AttrConnector: "VIA" as in Hill Top VIA Parish where Parish is a locality and Hill Top is a dependent locality
	//
	// AttrIndicator: Eg. Erode (Dist) where (Dist) is the Indicator
	DependentLocality struct {
		AttrConnector           string                     `json:"attr_connector,omitempty"` // maxLength=25
		AttrType                string                     `json:"attr_type,omitempty"`      // maxLength=12
		AttrUsageType           string                     `json:"attr_usage_type,omitempty"`
		DependentLocality       *DependentLocality         `json:"dependent_locality,omitempty"`
		DependentLocalityName   []*DependentLocalityName   `json:"dependent_locality_name,omitempty"`
		DependentLocalityNumber []*DependentLocalityNumber `json:"dependent_locality_number,omitempty"`
		LargeMailUser           *LargeMailUser             `json:"large_mail_user,omitempty"`
		PostOffice              *PostOffice                `json:"post_office,omitempty"`
		Premise                 *Premise                   `json:"premise,omitempty"`
		Thoroughfare            *Thoroughfare              `json:"thoroughfare,omitempty"`
	}

	// DependentLocalityName - Name of the dependent locality
	DependentLocalityName struct {
		AttrType string `json:"attr_type,omitempty"` // maxLength=12
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// DependentLocalityNumber - Number of the dependent locality. Some areas are numbered.
	// Eg. SECTOR 5 in a Suburb as in India or SOI SUKUMVIT 10 as in Thailand
	DependentLocalityNumber struct {
		AttrNameNumberOccurrence string `json:"attr_name_number_occurrence,omitempty"` // maxLength=6
		AttrCode                 string `json:"attr_code,omitempty"`                   // Used by postal services to encode the name of the element.
		Text                     string `json:"text,omitempty"`
	}

	// DependentThoroughfare is related to a street; occurs in GB, IE, ES, PT
	DependentThoroughfare struct {
		AttrType                 string                    `json:"attr_type,omitempty"`
		ThoroughfareName         *ThoroughfareName         `json:"thoroughfare_name,omitempty"`
		ThoroughfarePreDirection *ThoroughfarePreDirection `json:"thoroughfare_pre_direction,omitempty"`
		ThoroughfareTrailingType *ThoroughfareTrailingType `json:"thoroughfare_trailing_type,omitempty"`
	}

	// LargeMailUser - Specification of a large mail user address.
	//
	// Examples of large mail users are postal companies, companies in France with a cedex number,
	// hospitals and airports with their own post code.
	//
	// Large mail user addresses do not have a street name with premise name or premise number
	// in countries like Netherlands. But they have a POBox and street also in countries like France.
	LargeMailUser struct {
		AttrType                string                   `json:"attr_type,omitempty"` // maxLength=8
		BuildingName            *BuildingName            `json:"building_name,omitempty"`
		Department              *Department              `json:"department,omitempty"`
		LargeMailUserIdentifier *LargeMailUserIdentifier `json:"large_mail_user_identifier,omitempty"`
		LargeMailUserName       *LargeMailUserName       `json:"large_mail_user_name,omitempty"`
	}

	// LargeMailUserIdentifier - Specification of the identification number of a large mail user.
	//
	// An example are the Cedex codes in France.
	LargeMailUserIdentifier struct {
		AttrType      string `json:"attr_type,omitempty"`      // maxLength=14
		AttrIndicator string `json:"attr_indicator,omitempty"` // eg. Building 429 in which Building is the Indicator
		AttrCode      string `json:"attr_code,omitempty"`      // Used by postal services to encode the name of the element.
		Text          string `json:"text,omitempty"`
	}

	// LargeMailUserName - Name of the large mail user.
	//
	// eg. Smith Ford International airport
	LargeMailUserName struct {
		AttrType string `json:"attr_type,omitempty"` // Airport, Hospital, etc
		AttrCode string `json:"attr_code,omitempty"`
		Text     string `json:"text,omitempty"`
	}

	// PostBox - Specification of a postbox like mail delivery point.
	//
	// Only a single postbox number can be specified.
	//
	// Examples of postboxes are POBox, free mail numbers, etc.
	PostBox struct {
		AttrType      string         `json:"attr_type,omitempty"`      // maxLength=5
		AttrIndicator string         `json:"attr_indicator,omitempty"` // LOCKED BAG NO:1234 where the Indicator is NO: and Type is LOCKED BAG
		PostBoxNumber *PostBoxNumber `json:"post_box_number,omitempty"`
		PostalCode    *PostalCode    `json:"postal_code,omitempty"`
	}

	// PostBoxNumber - Specification of the number of a postbox
	PostBoxNumber struct {
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// PostOffice - Specification of a post office.
	//
	// Examples are a rural post office where post is delivered and a post office containing post office boxes.
	PostOffice struct {
		AttrType         string            `json:"attr_type,omitempty"`      // maxLength=14
		AttrIndicator    string            `json:"attr_indicator,omitempty"` // eg. Kottivakkam (P.O) here (P.O) is the Indicator
		PostOfficeName   *PostOfficeName   `json:"post_office_name,omitempty"`
		PostOfficeNumber *PostOfficeNumber `json:"post_office_number,omitempty"`
		PostalCode       *PostalCode       `json:"postal_code,omitempty"`
	}

	// PostOfficeName - Specification of the name of the post office.
	//
	// This can be a rural post office where post is delivered or a post office containing post office boxes.
	PostOfficeName struct {
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// PostOfficeNumber - Specification of the number of the post office.
	//
	// Common in rural post offices
	PostOfficeNumber struct {
		AttrIndicator string `json:"attr_indicator,omitempty"` // maxLength=3
		Text          string `json:"text,omitempty"`
	}

	// PostalCode - PostalCode is the container element for either simple or complex (extended) postal codes.
	//
	// Type: Area Code, Postcode, etc.
	PostalCode struct {
		AttrType                  string                     `json:"attr_type,omitempty"` // maxLength=9
		PostalCodeNumber          *PostalCodeNumber          `json:"postal_code_number,omitempty"`
		PostalCodeNumberExtension *PostalCodeNumberExtension `json:"postal_code_number_extension,omitempty"`
	}

	// PostalCodeNumber - Specification of a postcode.
	//
	// The postcode is formatted according to country-specific rules, example:
	//  SW3 0A8-1A, 600074, 2067
	PostalCodeNumber struct {
		AttrType string `json:"attr_type,omitempty"` // Old Postal Code, new code, etc
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// PostalCodeNumberExtension - Examples are:
	//  1234 (USA), 1G (UK), etc.
	PostalCodeNumberExtension struct {
		AttrType                     string `json:"attr_type,omitempty"`                       // maxLength=19
		AttrNumberExtensionSeparator string `json:"attr_number_extension_separator,omitempty"` // The separator between postal code number and the extension. Eg. "-"
		Text                         string `json:"text,omitempty"`
	}

	// Premise - Specification of a single premise, for example a house or a building.
	//
	// The premise as a whole has a unique premise (house) number or a premise name.
	// There could be more than one premise in a street referenced in an address.
	// For example a building address near a major shopping centre or raiwlay station
	//
	// AttrType: COMPLEXE in COMPLEX DES JARDINS, A building, station, etc
	//
	// AttrPremiseDependency: STREET, PREMISE, SUBPREMISE, PARK, FARM, etc
	//
	// AttrPremiseDependencyType: NEAR, ADJACENT TO, etc
	//
	// AttrPremiseThoroughfareConnector: DES, DE, LA, LA, DU in RUE DU BOIS. These terms connect a premise/thoroughfare type and premise/thoroughfare name. Terms may appear with names AVE DU BOIS
	Premise struct {
		AttrPremiseDependency            string               `json:"attr_premise_dependency,omitempty"`      // maxLength=7
		AttrPremiseDependencyType        string               `json:"attr_premise_dependency_type,omitempty"` // maxLength=19
		AttrType                         string               `json:"attr_type,omitempty"`                    // maxLength=18
		AttrPremiseThoroughfareConnector string               `json:"attr_premise_thoroughfare_connector,omitempty"`
		BuildingName                     *BuildingName        `json:"building_name,omitempty"`
		PostalCode                       *PostalCode          `json:"postal_code,omitempty"`
		Premise                          *Premise             `json:"premise,omitempty"`
		PremiseLocation                  *PremiseLocation     `json:"premise_location,omitempty"`
		PremiseName                      *PremiseName         `json:"premise_name,omitempty"`
		PremiseNumber                    *PremiseNumber       `json:"premise_number,omitempty"`
		PremiseNumberSuffix              *PremiseNumberSuffix `json:"premise_number_suffix,omitempty"`
		SubPremise                       []*SubPremise        `json:"sub_premise,omitempty"`
	}

	// PremiseLocation - LOBBY, BASEMENT, GROUND FLOOR, etc...
	PremiseLocation struct {
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// PremiseName - Specification of the name of the premise (house, building, park, farm, etc).
	//
	// A premise name is specified when the premise cannot be addressed using a street name plus premise (house) number.
	//
	// AttrTypeOccurrence: EGIS Building where EGIS occurs before Building, DES JARDINS occurs after COMPLEXE DES JARDINS
	PremiseName struct {
		AttrTypeOccurrence string `json:"attr_type_occurrence,omitempty"` // maxLength=5
		Text               string `json:"text,omitempty"`
	}

	// PremiseNumber - Specification of the identifier of the premise (house, building, etc).
	//
	// Premises in a street are often uniquely identified by means of consecutive identifiers.
	// The identifier can be a number, a letter or any combination of the two.
	PremiseNumber struct {
		AttrNumberType           string `json:"attr_number_type,omitempty"`            // Building 12-14 is "Range" and Building 12 is "Single"
		AttrType                 string `json:"attr_type,omitempty"`                   //
		AttrIndicator            string `json:"attr_indicator,omitempty"`              // No. in House No.12, # in #12, etc.
		AttrIndicatorOccurrence  string `json:"attr_indicator_occurrence,omitempty"`   // No. occurs before 12 No.12
		AttrNumberTypeOccurrence string `json:"attr_number_type_occurrence,omitempty"` // 12 in BUILDING 12 occurs "after" premise type BUILDING
		Code                     string `json:"code,omitempty"`                        // Used by postal services to encode the name of the element.
		Text                     string `json:"text,omitempty"`
	}

	// PremiseNumberSuffix - A in 12A
	PremiseNumberSuffix struct {
		AttrNumberPrefixSeparator string `json:"attr_number_prefix_separator,omitempty"` // A-12 where 12 is number and A is prefix and "-" is the separator
		AttrType                  string `json:"attr_type,omitempty"`                    //
		AttrCode                  string `json:"attr_code,omitempty"`                    // Used by postal services to encode the name of the element.
		Text                      string `json:"text,omitempty"`
	}

	// SubPremise - Specification of a single sub-premise.
	//
	// Examples of sub-premises are apartments and suites. Each sub-premise should be uniquely identifiable.
	SubPremise struct {
		AttrType               string                  `json:"attr_type,omitempty"` // maxLength=9
		SubPremise             []*SubPremise           `json:"sub_premise,omitempty"`
		SubPremiseName         []*SubPremiseName       `json:"sub_premise_name,omitempty"`
		SubPremiseNumber       []*SubPremiseNumber     `json:"sub_premise_number,omitempty"`
		SubPremiseNumberSuffix *SubPremiseNumberSuffix `json:"sub_premise_number_suffix,omitempty"`
	}

	// SubPremiseName -  Name of the SubPremise
	SubPremiseName struct {
		AttrType           string `json:"attr_type,omitempty"`
		AttrTypeOccurrence string `json:"attr_type_occurrence,omitempty"` //  EGIS Building where EGIS occurs before Building
		AttrCode           string `json:"attr_code,omitempty"`            //  Used by postal services to encode the name of the element.
		Text               string `json:"text,omitempty"`
	}

	// SubPremiseNumber -  Specification of the identifier of a sub-premise.
	//
	// Examples of sub-premises are apartments and suites. sub-premises in a building are often uniquely identified
	// by means of consecutive identifiers. The identifier can be a number, a letter or any combination of the two.
	// In the latter case, the identifier includes exactly one variable (range) part, which is either a number,
	// or a single letter that is surrounded by fixed parts at the left (prefix) or the right (postfix).
	SubPremiseNumber struct {
		AttrIndicator              string `json:"attr_indicator,omitempty"`                // "TH" in 12TH which is a floor number, "NO." in NO.1, "#" in APT #12, etc.
		AttrIndicatorOccurrence    string `json:"attr_indicator_occurrence,omitempty"`     // "No." occurs before 1 in No.1, or TH occurs after 12 in 12TH
		AttrNumberTypeOccurrence   string `json:"attr_number_type_occurrence,omitempty"`   // 12TH occurs "before" FLOOR (a type of subpremise) in 12TH FLOOR
		AttrPremiseNumberSeparator string `json:"attr_premise_number_separator,omitempty"` // "/" in 12/14 Archer Street where 12 is sub-premise number and 14 is premise number
		AttrType                   string `json:"attr_type,omitempty"`                     //
		AttrCode                   string `json:"attr_code,omitempty"`                     // Used by postal services to encode the name of the element.
		Text                       string `json:"text,omitempty"`
	}

	// SubPremiseNumberSuffix -  Prefix of the sub premise number. eg. A in A-12
	SubPremiseNumberSuffix struct {
		AttrNumberSuffixSeparator string `json:"attr_number_suffix_separator,omitempty"` //  12-A where 12 is number and A is suffix and "-" is the separator
		AttrType                  string `json:"attr_type,omitempty"`                    //
		AttrCode                  string `json:"attr_code,omitempty"`                    //  Used by postal services to encode the name of the element.
		Text                      string `json:"text,omitempty"`
	}

	// Thoroughfare - Specification of a thoroughfare.
	//
	// A thoroughfare could be a rd, street, canal, river, etc. Note dependentlocality in a street.
	// For example, in some countries, a large street will have many subdivisions with numbers.
	// Normally the subdivision name is the same as the road name, but with a number to identifiy it. Eg.
	//  SOI SUKUMVIT 3, SUKUMVIT RD, BANGKOK
	Thoroughfare struct {
		AttrDependentThoroughfares          string                     `json:"attr_dependent_thoroughfares,omitempty"`           // maxLength=3
		AttrDependentThoroughfaresConnector string                     `json:"attr_dependent_thoroughfares_connector,omitempty"` // maxLength=3
		AttrDependentThoroughfaresIndicator string                     `json:"attr_dependent_thoroughfares_indicator,omitempty"` // maxLength=9
		AttrDependentThoroughfaresType      string                     `json:"attr_dependent_thoroughfares_type,omitempty"`      // STS in GEORGE and ADELAIDE STS, RDS IN A and B RDS, etc. Use only when both the street types are the same
		AttrType                            string                     `json:"attr_type,omitempty"`                              // maxLength=6
		DependentLocality                   *DependentLocality         `json:"dependent_locality,omitempty"`
		DependentThoroughfare               *DependentThoroughfare     `json:"dependent_thoroughfare,omitempty"`
		PostalCode                          *PostalCode                `json:"postal_code,omitempty"`
		Premise                             *Premise                   `json:"premise,omitempty"`
		ThoroughfareLeadingType             *ThoroughfareLeadingType   `json:"thoroughfare_leading_type,omitempty"`
		ThoroughfareName                    *ThoroughfareName          `json:"thoroughfare_name,omitempty"`
		ThoroughfareNumber                  *ThoroughfareNumber        `json:"thoroughfare_number,omitempty"`
		ThoroughfareNumberRange             *ThoroughfareNumberRange   `json:"thoroughfare_number_range,omitempty"`
		ThoroughfareNumberSuffix            *ThoroughfareNumberSuffix  `json:"thoroughfare_number_suffix,omitempty"`
		ThoroughfarePostDirection           *ThoroughfarePostDirection `json:"thoroughfare_post_direction,omitempty"`
		ThoroughfarePreDirection            *ThoroughfarePreDirection  `json:"thoroughfare_pre_direction,omitempty"`
		ThoroughfareTrailingType            *ThoroughfareTrailingType  `json:"thoroughfare_trailing_type,omitempty"`
	}

	// ThoroughfareLeadingType - Appears before the thoroughfare name.
	//
	//  Spanish: Avenida Aurora, where Avenida is the leading type
	//  French: Rue Moliere, where Rue is the leading type.
	ThoroughfareLeadingType string

	// ThoroughfareName - Specification of the name of a Thoroughfare
	//
	// Also dependant street name: street name, canal name, etc.
	ThoroughfareName string

	// ThoroughfareNumber - Eg.: 23 Archer street or 25/15 Zero Avenue, etc
	ThoroughfareNumber struct {
		AttrType                string `json:"attr_type,omitempty"`
		AttrNumberType          string `json:"attr_number_type,omitempty"`          // 12 Archer Street is "Single" and 12-14 Archer Street is "Range"
		AttrIndicator           string `json:"attr_indicator,omitempty"`            // maxLength=3; No. in Street No.12 or "#" in Street # 12, etc.
		AttrIndicatorOccurrence string `json:"attr_indicator_occurrence,omitempty"` // maxLength=6; No.12 where "No." is before actual street number
		AttrNumberOccurrence    string `json:"attr_number_occurrence,omitempty"`    // 23 Archer St, Archer Street 23, St Archer 23
		Text                    string `json:"text,omitempty"`
	}

	// ThoroughfareNumberFrom - Starting number in the range
	ThoroughfareNumberFrom struct {
		AttrCode           string              `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		ThoroughfareNumber *ThoroughfareNumber `json:"thoroughfare_number,omitempty"`
	}

	// ThoroughfareNumberRange - A container to represent a range of numbers (from x thru y) for a thoroughfare.
	//
	//  eg. 1-2 Albert Av
	ThoroughfareNumberRange struct {
		AttrIndicator          string                  `json:"attr_indicator,omitempty"` // maxLength=2
		AttrType               string                  `json:"attr_type,omitempty"`      // maxLength=4
		ThoroughfareNumberFrom *ThoroughfareNumberFrom `json:"thoroughfare_number_from,omitempty"`
		ThoroughfareNumberTo   *ThoroughfareNumberTo   `json:"thoroughfare_number_to,omitempty"`
	}

	// ThoroughfareNumberSuffix - Suffix after the number. A in 12A Archer Street
	ThoroughfareNumberSuffix struct {
		AttrNumberSuffixSeparator string `json:"attr_number_suffix_separator,omitempty"` // 12-A where 12 is number and A is suffix and "-" is the separator
		AttrType                  string `json:"attr_type,omitempty"`                    // NEAR, ADJACENT TO, etc
		AttrCode                  string `json:"attr_code,omitempty"`                    // Used by postal services to encode the name of the element.
		Text                      string `json:"text,omitempty"`
	}

	// ThoroughfareNumberTo - Ending number in the range
	ThoroughfareNumberTo struct {
		AttrCode           string              `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		ThoroughfareNumber *ThoroughfareNumber `json:"thoroughfare_number,omitempty"`
	}

	// ThoroughfarePostDirection - 221-bis Baker Street North, where North is the post-direction.
	//
	// The post-direction appears after the name.
	ThoroughfarePostDirection struct {
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}

	// ThoroughfarePreDirection - North Baker Street, where North is the pre-direction.
	//
	// The direction appears before the name.
	ThoroughfarePreDirection struct {
		AttrCode string `json:"attr_code,omitempty"`
		Text     string `json:"text,omitempty"`
	}

	// ThoroughfareTrailingType - Appears after the thoroughfare name. Ed. British: Baker Lane, where Lane is the trailing type.
	ThoroughfareTrailingType struct {
		AttrCode string `json:"attr_code,omitempty"` // Used by postal services to encode the name of the element.
		Text     string `json:"text,omitempty"`
	}
)
