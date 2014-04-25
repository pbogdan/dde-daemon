// This file is automatically generated, please don't edit manully.
package main

import (
	"fmt"
)

// Get key type
func getSettingVpnOpenconnectKeyType(key string) (t ktype) {
	switch key {
	default:
		t = ktypeUnknown
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		t = ktypeBoolean
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		t = ktypeString
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		t = ktypeString
	}
	return
}

// Check is key in current setting field
func isKeyInSettingVpnOpenconnect(key string) bool {
	switch key {
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		return true
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		return true
	}
	return false
}

// Get key's default value
func getSettingVpnOpenconnectKeyDefaultValueJSON(key string) (valueJSON string) {
	switch key {
	default:
		Logger.Error("invalid key:", key)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		valueJSON = `false`
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		valueJSON = `false`
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		valueJSON = `""`
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		valueJSON = `""`
	}
	return
}

// Get JSON value generally
func generalGetSettingVpnOpenconnectKeyJSON(data _ConnectionData, key string) (value string) {
	switch key {
	default:
		Logger.Error("generalGetSettingVpnOpenconnectKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		value = getSettingVpnOpenconnectKeyGatewayJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		value = getSettingVpnOpenconnectKeyCacertJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		value = getSettingVpnOpenconnectKeyProxyJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		value = getSettingVpnOpenconnectKeyCsdEnableJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		value = getSettingVpnOpenconnectKeyCsdWrapperJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		value = getSettingVpnOpenconnectKeyUsercertJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		value = getSettingVpnOpenconnectKeyPrivkeyJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		value = getSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		value = getSettingVpnOpenconnectKeyCookieJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		value = getSettingVpnOpenconnectKeyGwcertJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		value = getSettingVpnOpenconnectKeyAuthtypeJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		value = getSettingVpnOpenconnectKeyMtuJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		value = getSettingVpnOpenconnectKeyStokenSourceJSON(data)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		value = getSettingVpnOpenconnectKeyStokenStringJSON(data)
	}
	return
}

// Set JSON value generally
func generalSetSettingVpnOpenconnectKeyJSON(data _ConnectionData, key, valueJSON string) {
	switch key {
	default:
		Logger.Error("generalSetSettingVpnOpenconnectKeyJSON: invalide key", key)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY:
		setSettingVpnOpenconnectKeyGatewayJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CACERT:
		setSettingVpnOpenconnectKeyCacertJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PROXY:
		setSettingVpnOpenconnectKeyProxyJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE:
		setSettingVpnOpenconnectKeyCsdEnableJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER:
		setSettingVpnOpenconnectKeyCsdWrapperJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT:
		setSettingVpnOpenconnectKeyUsercertJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY:
		setSettingVpnOpenconnectKeyPrivkeyJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID:
		setSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE:
		setSettingVpnOpenconnectKeyCookieJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT:
		setSettingVpnOpenconnectKeyGwcertJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE:
		setSettingVpnOpenconnectKeyAuthtypeJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_MTU:
		setSettingVpnOpenconnectKeyMtuJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE:
		setSettingVpnOpenconnectKeyStokenSourceJSON(data, valueJSON)
	case NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING:
		setSettingVpnOpenconnectKeyStokenStringJSON(data, valueJSON)
	}
	return
}

// Check if key exists
func isSettingVpnOpenconnectKeyGatewayExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY)
}
func isSettingVpnOpenconnectKeyCacertExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT)
}
func isSettingVpnOpenconnectKeyProxyExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY)
}
func isSettingVpnOpenconnectKeyCsdEnableExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE)
}
func isSettingVpnOpenconnectKeyCsdWrapperExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER)
}
func isSettingVpnOpenconnectKeyUsercertExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT)
}
func isSettingVpnOpenconnectKeyPrivkeyExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY)
}
func isSettingVpnOpenconnectKeyPemPassphraseFsidExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID)
}
func isSettingVpnOpenconnectKeyCookieExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE)
}
func isSettingVpnOpenconnectKeyGwcertExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT)
}
func isSettingVpnOpenconnectKeyAuthtypeExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE)
}
func isSettingVpnOpenconnectKeyMtuExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU)
}
func isSettingVpnOpenconnectKeyStokenSourceExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE)
}
func isSettingVpnOpenconnectKeyStokenStringExists(data _ConnectionData) bool {
	return isSettingKeyExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING)
}

// Ensure field and key exists and not empty
func ensureFieldSettingVpnOpenconnectExists(data _ConnectionData, errs FieldKeyErrors, relatedKey string) {
	if !isSettingFieldExists(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME) {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_MISSING_SECTION, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME))
	}
	fieldData, _ := data[NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME]
	if len(fieldData) == 0 {
		rememberError(errs, relatedKey, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, fmt.Sprintf(NM_KEY_ERROR_EMPTY_SECTION, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME))
	}
}
func ensureSettingVpnOpenconnectKeyGatewayNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyGatewayExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyGateway(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCacertNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyCacertExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyCacert(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyProxyNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyProxyExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyProxy(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCsdEnableNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyCsdEnableExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCsdWrapperNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyCsdWrapperExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyCsdWrapper(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyUsercertNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyUsercertExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyUsercert(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyPrivkeyNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyPrivkeyExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyPrivkey(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyPemPassphraseFsidNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyPemPassphraseFsidExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, NM_KEY_ERROR_MISSING_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyCookieNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyCookieExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyCookie(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyGwcertNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyGwcertExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyGwcert(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyAuthtypeNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyAuthtypeExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyAuthtype(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyMtuNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyMtuExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyMtu(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyStokenSourceNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyStokenSourceExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyStokenSource(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, NM_KEY_ERROR_EMPTY_VALUE)
	}
}
func ensureSettingVpnOpenconnectKeyStokenStringNoEmpty(data _ConnectionData, errs FieldKeyErrors) {
	if !isSettingVpnOpenconnectKeyStokenStringExists(data) {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, NM_KEY_ERROR_MISSING_VALUE)
	}
	value := getSettingVpnOpenconnectKeyStokenString(data)
	if len(value) == 0 {
		rememberError(errs, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, NM_KEY_ERROR_EMPTY_VALUE)
	}
}

// Getter
func getSettingVpnOpenconnectKeyGateway(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY).(string)
	return
}
func getSettingVpnOpenconnectKeyCacert(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT).(string)
	return
}
func getSettingVpnOpenconnectKeyProxy(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY).(string)
	return
}
func getSettingVpnOpenconnectKeyCsdEnable(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE).(bool)
	return
}
func getSettingVpnOpenconnectKeyCsdWrapper(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER).(string)
	return
}
func getSettingVpnOpenconnectKeyUsercert(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT).(string)
	return
}
func getSettingVpnOpenconnectKeyPrivkey(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY).(string)
	return
}
func getSettingVpnOpenconnectKeyPemPassphraseFsid(data _ConnectionData) (value bool) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID).(bool)
	return
}
func getSettingVpnOpenconnectKeyCookie(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE).(string)
	return
}
func getSettingVpnOpenconnectKeyGwcert(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT).(string)
	return
}
func getSettingVpnOpenconnectKeyAuthtype(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE).(string)
	return
}
func getSettingVpnOpenconnectKeyMtu(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU).(string)
	return
}
func getSettingVpnOpenconnectKeyStokenSource(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE).(string)
	return
}
func getSettingVpnOpenconnectKeyStokenString(data _ConnectionData) (value string) {
	value, _ = getSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING).(string)
	return
}

// Setter
func setSettingVpnOpenconnectKeyGateway(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, value)
}
func setSettingVpnOpenconnectKeyCacert(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, value)
}
func setSettingVpnOpenconnectKeyProxy(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, value)
}
func setSettingVpnOpenconnectKeyCsdEnable(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, value)
}
func setSettingVpnOpenconnectKeyCsdWrapper(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, value)
}
func setSettingVpnOpenconnectKeyUsercert(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, value)
}
func setSettingVpnOpenconnectKeyPrivkey(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, value)
}
func setSettingVpnOpenconnectKeyPemPassphraseFsid(data _ConnectionData, value bool) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, value)
}
func setSettingVpnOpenconnectKeyCookie(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, value)
}
func setSettingVpnOpenconnectKeyGwcert(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, value)
}
func setSettingVpnOpenconnectKeyAuthtype(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, value)
}
func setSettingVpnOpenconnectKeyMtu(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, value)
}
func setSettingVpnOpenconnectKeyStokenSource(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, value)
}
func setSettingVpnOpenconnectKeyStokenString(data _ConnectionData, value string) {
	setSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, value)
}

// JSON Getter
func getSettingVpnOpenconnectKeyGatewayJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY))
	return
}
func getSettingVpnOpenconnectKeyCacertJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CACERT))
	return
}
func getSettingVpnOpenconnectKeyProxyJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PROXY))
	return
}
func getSettingVpnOpenconnectKeyCsdEnableJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE))
	return
}
func getSettingVpnOpenconnectKeyCsdWrapperJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER))
	return
}
func getSettingVpnOpenconnectKeyUsercertJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT))
	return
}
func getSettingVpnOpenconnectKeyPrivkeyJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY))
	return
}
func getSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID))
	return
}
func getSettingVpnOpenconnectKeyCookieJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE))
	return
}
func getSettingVpnOpenconnectKeyGwcertJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT))
	return
}
func getSettingVpnOpenconnectKeyAuthtypeJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE))
	return
}
func getSettingVpnOpenconnectKeyMtuJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_MTU))
	return
}
func getSettingVpnOpenconnectKeyStokenSourceJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE))
	return
}
func getSettingVpnOpenconnectKeyStokenStringJSON(data _ConnectionData) (valueJSON string) {
	valueJSON = getSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING))
	return
}

// JSON Setter
func setSettingVpnOpenconnectKeyGatewayJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY))
}
func setSettingVpnOpenconnectKeyCacertJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CACERT))
}
func setSettingVpnOpenconnectKeyProxyJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PROXY))
}
func setSettingVpnOpenconnectKeyCsdEnableJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE))
}
func setSettingVpnOpenconnectKeyCsdWrapperJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER))
}
func setSettingVpnOpenconnectKeyUsercertJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT))
}
func setSettingVpnOpenconnectKeyPrivkeyJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY))
}
func setSettingVpnOpenconnectKeyPemPassphraseFsidJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID))
}
func setSettingVpnOpenconnectKeyCookieJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE))
}
func setSettingVpnOpenconnectKeyGwcertJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT))
}
func setSettingVpnOpenconnectKeyAuthtypeJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE))
}
func setSettingVpnOpenconnectKeyMtuJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_MTU))
}
func setSettingVpnOpenconnectKeyStokenSourceJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE))
}
func setSettingVpnOpenconnectKeyStokenStringJSON(data _ConnectionData, valueJSON string) {
	setSettingKeyJSON(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING, valueJSON, getSettingVpnOpenconnectKeyType(NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING))
}

// Remover
func removeSettingVpnOpenconnectKeyGateway(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GATEWAY)
}
func removeSettingVpnOpenconnectKeyCacert(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CACERT)
}
func removeSettingVpnOpenconnectKeyProxy(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PROXY)
}
func removeSettingVpnOpenconnectKeyCsdEnable(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_ENABLE)
}
func removeSettingVpnOpenconnectKeyCsdWrapper(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_CSD_WRAPPER)
}
func removeSettingVpnOpenconnectKeyUsercert(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_USERCERT)
}
func removeSettingVpnOpenconnectKeyPrivkey(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PRIVKEY)
}
func removeSettingVpnOpenconnectKeyPemPassphraseFsid(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_PEM_PASSPHRASE_FSID)
}
func removeSettingVpnOpenconnectKeyCookie(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_COOKIE)
}
func removeSettingVpnOpenconnectKeyGwcert(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_GWCERT)
}
func removeSettingVpnOpenconnectKeyAuthtype(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_AUTHTYPE)
}
func removeSettingVpnOpenconnectKeyMtu(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_MTU)
}
func removeSettingVpnOpenconnectKeyStokenSource(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_SOURCE)
}
func removeSettingVpnOpenconnectKeyStokenString(data _ConnectionData) {
	removeSettingKey(data, NM_SETTING_VF_VPN_OPENCONNECT_SETTING_NAME, NM_SETTING_VPN_OPENCONNECT_KEY_STOKEN_STRING)
}