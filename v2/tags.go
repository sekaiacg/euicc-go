package sgp22

import "github.com/damonto/euicc-go/bertlv"

// region Request Tags

func (*PrepareDownloadRequest) Tag() bertlv.Tag           { return []byte{0xBF, 0x21} }
func (*ES9BoundProfilePackageRequest) Tag() bertlv.Tag    { return []byte{0xBF, 0x21} }
func (*ListNotificationRequest) Tag() bertlv.Tag          { return []byte{0xBF, 0x28} }
func (*ListNotificationResponse) Tag() bertlv.Tag         { return []byte{0xBF, 0x28} }
func (*SetNicknameRequest) Tag() bertlv.Tag               { return []byte{0xBF, 0x29} }
func (*SetNicknameResponse) Tag() bertlv.Tag              { return []byte{0xBF, 0x29} }
func (*ProfileInfoListRequest) Tag() bertlv.Tag           { return []byte{0xBF, 0x2D} }
func (*ProfileInfoListResponse) Tag() bertlv.Tag          { return []byte{0xBF, 0x2D} }
func (*GetEuiccChallengeRequest) Tag() bertlv.Tag         { return []byte{0xBF, 0x2E} }
func (*GetEuiccChallengeResponse) Tag() bertlv.Tag        { return []byte{0xBF, 0x2E} }
func (*NotificationMetadata) Tag() bertlv.Tag             { return []byte{0xBF, 0x2F} }
func (*NotificationSentRequest) Tag() bertlv.Tag          { return []byte{0xBF, 0x30} }
func (*NotificationSentResponse) Tag() bertlv.Tag         { return []byte{0xBF, 0x30} }
func (*EnableProfileRequest) Tag() bertlv.Tag             { return []byte{0xBF, 0x31} }
func (*EnableProfileResponse) Tag() bertlv.Tag            { return []byte{0xBF, 0x31} }
func (*DisableProfileRequest) Tag() bertlv.Tag            { return []byte{0xBF, 0x32} }
func (*DisableProfileResponse) Tag() bertlv.Tag           { return []byte{0xBF, 0x32} }
func (*DeleteProfileRequest) Tag() bertlv.Tag             { return []byte{0xBF, 0x33} }
func (*DeleteProfileResponse) Tag() bertlv.Tag            { return []byte{0xBF, 0x33} }
func (*EuiccMemoryResetRequest) Tag() bertlv.Tag          { return []byte{0xBF, 0x34} }
func (*EuiccMemoryResetResponse) Tag() bertlv.Tag         { return []byte{0xBF, 0x34} }
func (*AuthenticateServerRequest) Tag() bertlv.Tag        { return []byte{0xBF, 0x38} }
func (*EuiccConfiguredAddressesRequest) Tag() bertlv.Tag  { return []byte{0xBF, 0x3C} }
func (*EuiccConfiguredAddressesResponse) Tag() bertlv.Tag { return []byte{0xBF, 0x3C} }
func (*GetEuiccDataRequest) Tag() bertlv.Tag              { return []byte{0xBF, 0x3E} }
func (*GetEuiccDataResponse) Tag() bertlv.Tag             { return []byte{0xBF, 0x3E} }
func (*SetDefaultDPAddressRequest) Tag() bertlv.Tag       { return []byte{0xBF, 0x3F} }
func (*SetDefaultDPAddressResponse) Tag() bertlv.Tag      { return []byte{0xBF, 0x3F} }
func (*CancelSessionRequest) Tag() bertlv.Tag             { return []byte{0xBF, 0x41} }
func (*ES9CancelSessionRequest) Tag() bertlv.Tag          { return []byte{0xBF, 0x41} }
func (*ProfileInfo) Tag() bertlv.Tag                      { return []byte{0xE3} }

// endregion

// region GetProfilesInfo Tags

var (
	TagICCID                         = bertlv.Application.Primitive(26)
	TagISDPAID                       = bertlv.Application.Primitive(15)
	TagProfileState                  = bertlv.ContextSpecific.Primitive(112)
	TagNickname                      = bertlv.ContextSpecific.Primitive(16)
	TagServiceProviderName           = bertlv.ContextSpecific.Primitive(17)
	TagProfileName                   = bertlv.ContextSpecific.Primitive(18)
	TagProfileIconType               = bertlv.ContextSpecific.Primitive(19)
	TagProfileIcon                   = bertlv.ContextSpecific.Primitive(20)
	TagProfileClass                  = bertlv.ContextSpecific.Primitive(21)
	TagNotificationConfigurationInfo = bertlv.ContextSpecific.Constructed(22)
	TagProfileOwner                  = bertlv.ContextSpecific.Constructed(23)
	TagSMDPProprietaryData           = bertlv.ContextSpecific.Constructed(24)
	TagProfilePolicyRules            = bertlv.ContextSpecific.Primitive(25)
	TagServiceSpecificData           = bertlv.ContextSpecific.Constructed(34)
)

// endregion
