package mbim

import "fmt"

// MBIMStatusError represents status errors from the MBIM device.
type MBIMStatusError uint32

const (
	MBIMStatusErrorNone MBIMStatusError = iota // 0x00000000
	MBIMStatusErrorBusy
	MBIMStatusErrorFailure
	MBIMStatusErrorSimNotInserted
	MBIMStatusErrorBadSim
	MBIMStatusErrorPinRequired
	MBIMStatusErrorPinDisabled
	MBIMStatusErrorNotRegistered
	MBIMStatusErrorProvidersNotFound
	MBIMStatusErrorNoDeviceSupport
	MBIMStatusErrorProviderNotVisible
	MBIMStatusErrorDataClassNotAvailable
	MBIMStatusErrorPacketServiceDetached
	MBIMStatusErrorMaxActivatedContexts
	MBIMStatusErrorNotInitialized
	MBIMStatusErrorVoiceCallInProgress
	MBIMStatusErrorContextNotActivated
	MBIMStatusErrorServiceNotActivated
	MBIMStatusErrorInvalidAccessString
	MBIMStatusErrorInvalidUserNamePwd
	MBIMStatusErrorRadioPowerOff
	MBIMStatusErrorInvalidParameters
	MBIMStatusErrorReadFailure
	MBIMStatusErrorWriteFailure
	MBIMStatusErrorReserved
	MBIMStatusErrorNoPhonebook
	MBIMStatusErrorParameterTooLong
	MBIMStatusErrorStkBusy
	MBIMStatusErrorOperationNotAllowed
	MBIMStatusErrorMemoryFailure
	MBIMStatusErrorInvalidMemoryIndex
	MBIMStatusErrorMemoryFull
	MBIMStatusErrorFilterNotSupported
	MBIMStatusErrorDssInstanceLimit
	MBIMStatusErrorInvalidDeviceServiceOperation
	MBIMStatusErrorAuthIncorrectAutn
	MBIMStatusErrorAuthSyncFailure
	MBIMStatusErrorAuthAmfNotSet
	MBIMStatusErrorContextNotSupported
	MBIMStatusErrorSmsUnknownSmscAddress          MBIMStatusError = 0x00000064 // 100
	MBIMStatusErrorSmsNetworkTimeout              MBIMStatusError = 0x00000065 // 101
	MBIMStatusErrorSmsLangNotSupported            MBIMStatusError = 0x00000066 // 102
	MBIMStatusErrorSmsEncodingNotSupported        MBIMStatusError = 0x00000067 // 103
	MBIMStatusErrorSmsFormatNotSupported          MBIMStatusError = 0x00000068 // 104
	MBIMStatusErrorMsNoLogicalChannels            MBIMStatusError = 0x87430001
	MBIMStatusErrorMsSelectFailed                 MBIMStatusError = 0x87430002
	MBIMStatusErrorMsInvalidLogicalChannel        MBIMStatusError = 0x87430003
	MBIMStatusErrorInvalidSignature               MBIMStatusError = 0x91000001
	MBIMStatusErrorInvalidImei                    MBIMStatusError = 0x91000002
	MBIMStatusErrorInvalidTimestamp               MBIMStatusError = 0x91000003
	MBIMStatusErrorNetworkListTooLarge            MBIMStatusError = 0x91000004
	MBIMStatusErrorSignatureAlgorithmNotSupported MBIMStatusError = 0x91000005
	MBIMStatusErrorFeatureNotSupported            MBIMStatusError = 0x91000006
	MBIMStatusErrorDecodeOrParsingError           MBIMStatusError = 0x91000007
)

func (e MBIMStatusError) Error() string {
	switch e {
	case MBIMStatusErrorNone:
		return "Success"
	case MBIMStatusErrorBusy:
		return "Busy"
	case MBIMStatusErrorFailure:
		return "Failure"
	case MBIMStatusErrorSimNotInserted:
		return "Sim Not Inserted"
	case MBIMStatusErrorBadSim:
		return "Bad Sim"
	case MBIMStatusErrorPinRequired:
		return "Pin Required"
	case MBIMStatusErrorPinDisabled:
		return "Pin Disabled"
	case MBIMStatusErrorNotRegistered:
		return "Not Registered"
	case MBIMStatusErrorProvidersNotFound:
		return "Providers Not Found"
	case MBIMStatusErrorNoDeviceSupport:
		return "No Device Support"
	case MBIMStatusErrorProviderNotVisible:
		return "Provider Not Visible"
	case MBIMStatusErrorDataClassNotAvailable:
		return "Data Class Not Available"
	case MBIMStatusErrorPacketServiceDetached:
		return "Packet Service Detached"
	case MBIMStatusErrorMaxActivatedContexts:
		return "Max Activated Contexts"
	case MBIMStatusErrorNotInitialized:
		return "Not Initialized"
	case MBIMStatusErrorVoiceCallInProgress:
		return "Voice Call In Progress"
	case MBIMStatusErrorContextNotActivated:
		return "Context Not Activated"
	case MBIMStatusErrorServiceNotActivated:
		return "Service Not Activated"
	case MBIMStatusErrorInvalidAccessString:
		return "Invalid Access String"
	case MBIMStatusErrorInvalidUserNamePwd:
		return "Invalid User Name Pwd"
	case MBIMStatusErrorRadioPowerOff:
		return "Radio Power Off"
	case MBIMStatusErrorInvalidParameters:
		return "Invalid Parameters"
	case MBIMStatusErrorReadFailure:
		return "Read Failure"
	case MBIMStatusErrorWriteFailure:
		return "Write Failure"
	case MBIMStatusErrorNoPhonebook:
		return "No Phonebook"
	case MBIMStatusErrorParameterTooLong:
		return "Parameter Too Long"
	case MBIMStatusErrorStkBusy:
		return "Stk Busy"
	case MBIMStatusErrorOperationNotAllowed:
		return "Operation Not Allowed"
	case MBIMStatusErrorMemoryFailure:
		return "Memory Failure"
	case MBIMStatusErrorInvalidMemoryIndex:
		return "Invalid Memory Index"
	case MBIMStatusErrorMemoryFull:
		return "Memory Full"
	case MBIMStatusErrorFilterNotSupported:
		return "Filter Not Supported"
	case MBIMStatusErrorDssInstanceLimit:
		return "Dss Instance Limit"
	case MBIMStatusErrorInvalidDeviceServiceOperation:
		return "Invalid Device Service Operation"
	case MBIMStatusErrorAuthIncorrectAutn:
		return "Auth Incorrect Autn"
	case MBIMStatusErrorAuthSyncFailure:
		return "Auth Sync Failure"
	case MBIMStatusErrorAuthAmfNotSet:
		return "Auth Amf Not Set"
	case MBIMStatusErrorContextNotSupported:
		return "Context Not Supported"
	case MBIMStatusErrorSmsUnknownSmscAddress:
		return "Sms Unknown Smsc Address"
	case MBIMStatusErrorSmsNetworkTimeout:
		return "Sms Network Timeout"
	case MBIMStatusErrorSmsLangNotSupported:
		return "Sms Lang Not Supported"
	case MBIMStatusErrorSmsEncodingNotSupported:
		return "Sms Encoding Not Supported"
	case MBIMStatusErrorSmsFormatNotSupported:
		return "Sms Format Not Supported"
	case MBIMStatusErrorMsNoLogicalChannels:
		return "Ms No Logical Channels"
	case MBIMStatusErrorMsSelectFailed:
		return "Ms Select Failed"
	case MBIMStatusErrorMsInvalidLogicalChannel:
		return "Ms Invalid Logical Channel"
	case MBIMStatusErrorInvalidSignature:
		return "Invalid Signature"
	case MBIMStatusErrorInvalidImei:
		return "Invalid Imei"
	case MBIMStatusErrorInvalidTimestamp:
		return "Invalid Timestamp"
	case MBIMStatusErrorNetworkListTooLarge:
		return "Network List Too Large"
	case MBIMStatusErrorSignatureAlgorithmNotSupported:
		return "Signature Algorithm Not Supported"
	case MBIMStatusErrorFeatureNotSupported:
		return "Feature Not Supported"
	case MBIMStatusErrorDecodeOrParsingError:
		return "Decode Or Parsing Error"
	default:
		return fmt.Sprintf("Unknown MBIM Status Error: %d", e)
	}
}

type MBIMError struct {
	Code MBIMStatusError
}

func (e MBIMError) Error() string {
	return fmt.Sprintf("MBIM Error: %s (%d) ", e.Code.Error(), e.Code)
}
