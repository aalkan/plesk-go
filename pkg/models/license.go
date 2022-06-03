package models

import "time"

type LicenseGet struct {
	Nickname string `json:"nickname"`
}

type LicenseResponse struct {
	OwnerId     string `json:"owner_id"`
	KeyId       string `json:"key_id"`
	ParentKeyId string `json:"parent_key_id"`
}

type RetrieveLicenseGet struct {
	KeyId string `json:"key_id"` //for endpoint path
}

type KeyIdentify struct {
	KeyId          int    `json:"keyId"`
	KeyNumber      string `json:"keyNumber"`
	ActivationCode string `json:"activationCode"`
}

type RetrieveLicenseResponse struct {
	OwnerId                string         `json:"ownerId"`
	KeyIdentifiers         *KeyIdentify   `json:"keyIdentifiers"`
	ParentKeyIdentifiers   interface{}    `json:"parentKeyIdentifiers"`
	ActivationInfo         interface{}    `json:"activationInfo"`
	IpaddressBinding       interface{}    `json:"ipAddressBinding"`
	Nickname               string         `json:"nickname"`
	ProductConfigurationId int            `json:"productConfigurationId"`
	Items                  []*Items       `json:"items"`
	CreationDate           time.Time      `json:"creationDate"`
	LastModificationDate   time.Time      `json:"lastModificationDate"`
	UpdateDate             time.Time      `json:"updateDate"`
	ExpirationDate         time.Time      `json:"expirationDate"`
	SusexpirationDate      time.Time      `json:"susExpirationDate"`
	SusStatus              string         `json:"susStatus"`
	SupportExpirationDate  interface{}    `json:"supportExpirationDate"`
	SupportStatus          interface{}    `json:"supportStatus"`
	Status                 string         `json:"status"`
	Terminated             bool           `json:"terminated"`
	Suspended              bool           `json:"suspended"`
	OwnerSuspended         bool           `json:"ownerSuspended"`
	AutoRenew              bool           `json:"autoRenew"`
	RestrictIpBinding      bool           `json:"restrictIPBinding"`
	ChildKeyIdentifiers    []*KeyIdentify `json:"childKeyIdentifiers"`
}

type Items struct {
	ExternalId interface{} `json:"externalId"`
	Item       string      `json:"item"`
}

type RenewingLicenseGet struct {
	KeyId string `json:"key_id"`
}

type RenewingLicenseResponse struct {
	KeyId string `json:"key_id"` //for endpoint path
}
