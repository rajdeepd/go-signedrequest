package main

type SignedRequestStruct struct {
	Algorithm string `json:"algorithm"`
	IssuedAt  int    `json:"issuedAt"`
	UserID    string `json:"userId"`
	Client    struct {
		RefreshToken interface{} `json:"refreshToken"`
		InstanceID   string      `json:"instanceId"`
		TargetOrigin string      `json:"targetOrigin"`
		InstanceURL  string      `json:"instanceUrl"`
		OauthToken   string      `json:"oauthToken"`
	} `json:"client"`
	Context struct {
		User struct {
			UserID                   string      `json:"userId"`
			UserName                 string      `json:"userName"`
			FirstName                string      `json:"firstName"`
			LastName                 string      `json:"lastName"`
			Email                    string      `json:"email"`
			FullName                 string      `json:"fullName"`
			Locale                   string      `json:"locale"`
			Language                 string      `json:"language"`
			TimeZone                 string      `json:"timeZone"`
			ProfileID                string      `json:"profileId"`
			RoleID                   interface{} `json:"roleId"`
			UserType                 string      `json:"userType"`
			CurrencyISOCode          string      `json:"currencyISOCode"`
			ProfilePhotoURL          string      `json:"profilePhotoUrl"`
			ProfileThumbnailURL      string      `json:"profileThumbnailUrl"`
			SiteURL                  interface{} `json:"siteUrl"`
			SiteURLPrefix            interface{} `json:"siteUrlPrefix"`
			NetworkID                interface{} `json:"networkId"`
			AccessibilityModeEnabled bool        `json:"accessibilityModeEnabled"`
			IsDefaultNetwork         bool        `json:"isDefaultNetwork"`
		} `json:"user"`
		Links struct {
			LoginURL            string `json:"loginUrl"`
			EnterpriseURL       string `json:"enterpriseUrl"`
			MetadataURL         string `json:"metadataUrl"`
			PartnerURL          string `json:"partnerUrl"`
			RestURL             string `json:"restUrl"`
			SobjectURL          string `json:"sobjectUrl"`
			SearchURL           string `json:"searchUrl"`
			QueryURL            string `json:"queryUrl"`
			RecentItemsURL      string `json:"recentItemsUrl"`
			ChatterFeedsURL     string `json:"chatterFeedsUrl"`
			ChatterGroupsURL    string `json:"chatterGroupsUrl"`
			ChatterUsersURL     string `json:"chatterUsersUrl"`
			ChatterFeedItemsURL string `json:"chatterFeedItemsUrl"`
			UserURL             string `json:"userUrl"`
		} `json:"links"`
		Application struct {
			Name                   string        `json:"name"`
			CanvasURL              string        `json:"canvasUrl"`
			ApplicationID          string        `json:"applicationId"`
			Version                string        `json:"version"`
			AuthType               string        `json:"authType"`
			ReferenceID            string        `json:"referenceId"`
			Options                []interface{} `json:"options"`
			SamlInitiationMethod   string        `json:"samlInitiationMethod"`
			Namespace              string        `json:"namespace"`
			IsInstalledPersonalApp bool          `json:"isInstalledPersonalApp"`
			DeveloperName          string        `json:"developerName"`
		} `json:"application"`
		Organization struct {
			OrganizationID       string      `json:"organizationId"`
			Name                 string      `json:"name"`
			MulticurrencyEnabled bool        `json:"multicurrencyEnabled"`
			NamespacePrefix      interface{} `json:"namespacePrefix"`
			CurrencyIsoCode      string      `json:"currencyIsoCode"`
		} `json:"organization"`
		Environment struct {
			Referer         interface{} `json:"referer"`
			LocationURL     string      `json:"locationUrl"`
			DisplayLocation string      `json:"displayLocation"`
			Sublocation     interface{} `json:"sublocation"`
			UITheme         string      `json:"uiTheme"`
			Dimensions      struct {
				Width        string `json:"width"`
				Height       string `json:"height"`
				MaxWidth     string `json:"maxWidth"`
				MaxHeight    string `json:"maxHeight"`
				ClientWidth  string `json:"clientWidth"`
				ClientHeight string `json:"clientHeight"`
			} `json:"dimensions"`
			Parameters struct {
			} `json:"parameters"`
			Record struct {
			} `json:"record"`
			Version struct {
				Season string `json:"season"`
				API    string `json:"api"`
			} `json:"version"`
		} `json:"environment"`
	} `json:"context"`
}
