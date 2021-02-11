package omex

/*
 constants
 @author Tony Tian
 @date 2018-03-17
 @version 1.0.0
*/

const (
	/*
	  http headers
	*/
	OM_ACCESS_KEY        = "OM-ACCESS-KEY"
	OM_ACCESS_SIGN       = "OM-ACCESS-SIGN"
	OM_ACCESS_TIMESTAMP  = "OM-ACCESS-TIMESTAMP"
	OM_ACCESS_PASSPHRASE = "OM-ACCESS-PASSPHRASE"

	/**
	  paging params
	*/
	OM_FROM  = "OM-FROM"
	OM_TO    = "OM-TO"
	OM_LIMIT = "OM-LIMIT"

	CONTENT_TYPE = "Content-Type"
	ACCEPT       = "Accept"
	COOKIE       = "Cookie"
	LOCALE       = "locale="

	APPLICATION_JSON      = "application/json"
	APPLICATION_JSON_UTF8 = "application/json; charset=UTF-8"

	/*
	  i18n: internationalization
	*/
	ENGLISH            = "en_US"
	SIMPLIFIED_CHINESE = "zh_CN"
	//zh_TW || zh_HK
	TRADITIONAL_CHINESE = "zh_HK"

	/*
	  http methods
	*/
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

	/*
	 others
	*/
	ResultDataJsonString = "resultDataJsonString"
	ResultPageJsonString = "resultPageJsonString"
)
