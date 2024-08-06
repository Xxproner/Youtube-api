package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	com := "curl"
	a1 := "https://www.youtube.com/youtubei/v1/backstage/create_post?prettyPrint=false"
	a2 := "-H"
	a3 := "accept: */*"
	a4 := "accept-language: ru,en;q=0.9"
	a5 := "authorization: SAPISIDHASH 1722869818_bc839736fa009e80529812912cce29f6e9bec074"
	a6 := "content-type: application/json"
	a7 := "cookie: VISITOR_INFO1_LIVE=DCcZUa687EE; VISITOR_PRIVACY_METADATA=CgJSVRIEGgAgVg^%^3D^%^3D; PREF=tz=Europe.Moscow; GPS=1; YSC=clSjNp1q654; HSID=ArPAeDTEX3a4PcxTP; SSID=AcVcySaEcJOxUeZC2; APISID=ZQBsox9eVDJsKE7M/AI2E71tDejLBWwqh1; SAPISID=iU81snPmRiZ4FAX7/AaIUeZc0X5Etk6Dlw; __Secure-1PAPISID=iU81snPmRiZ4FAX7/AaIUeZc0X5Etk6Dlw; __Secure-3PAPISID=iU81snPmRiZ4FAX7/AaIUeZc0X5Etk6Dlw; SID=g.a000mgiQALksn4LxzHcixbJhJybjh_a5K4t2cpnymNdLGdawvtSXBzKMCoiqBSGEkndZmfjUOAACgYKAa4SARASFQHGX2MiyhT3HuM-UeBzp0GM9lUt6RoVAUF8yKoMu4graJHtf0ztFb0JLKVF0076; __Secure-1PSID=g.a000mgiQALksn4LxzHcixbJhJybjh_a5K4t2cpnymNdLGdawvtSXvHYVVMQL6PYQSI7YMY0Y-wACgYKAfgSARASFQHGX2Mip8nDzTpo424dKwOt8GVD6BoVAUF8yKqXvfMEstP4L62j3OwopQFX0076; __Secure-3PSID=g.a000mgiQALksn4LxzHcixbJhJybjh_a5K4t2cpnymNdLGdawvtSX88oSmQ4KOLkZuMgleuS74wACgYKASMSARASFQHGX2MiyX5o8s9TncfDD7PcQPzt_hoVAUF8yKpiupdNIizHYp2wyBLwf3Hl0076; LOGIN_INFO=AFmmF2swRQIgdXY_2rXmXb_bmV-eXmFmu36ufD2OydwOilWm9KnezFgCIQCpXMLlPe9j5rhs-MyaZU8c0-N7lVAD6Y63Y5VPPckDgA:QUQ3MjNmendCcVFONm1JYmhvQ19iLVNNTUJhajV2UTZ6ZlBTRXlFMUFYWm52OHI2UzNjRnBxdzVpTGJ6dUp2MmJfTW96OXFfZmpJSlpPcWQ4WkZkVTV6RUZ3dUpwbkVxNkd1MTN4OFhXZW9rZThqOTBaVTc0X0NDZGFOanoxTVZCTnlzcS0zWjBiVjRIc3VFUHlmeEJaamx1VXFzTHloZkNB; __Secure-1PSIDTS=sidts-CjIB4E2dkXAASmdewvSRpvqI5th4E7qfB6MQ5zf781MW8CamwlTIHhk9R9lfKyT9rMnB8BAA; __Secure-3PSIDTS=sidts-CjIB4E2dkXAASmdewvSRpvqI5th4E7qfB6MQ5zf781MW8CamwlTIHhk9R9lfKyT9rMnB8BAA; SIDCC=AKEyXzUj5b7HefSx4ZYJmB42jjPTV8jDtCq8-lVs_WRDTvEZOZ_vDISFAmK7eIVUPj3cZVhx9g; __Secure-1PSIDCC=AKEyXzWXn3XjrJFJHuUkQkpIiNP7lb3VRtXGLSSt-l5rXQu49gmQKhZEAVXoTJ6B6PjKxR9k; __Secure-3PSIDCC=AKEyXzW0-m4DboKAQunHPzClKkB4Ep-M3ndmEt71GizN7fhFiaJb1RVGIwy9Rn3ygYpBOUGS; ST-19pu767=session_logininfo=AFmmF2swRQIgdXY_2rXmXb_bmV-eXmFmu36ufD2OydwOilWm9KnezFgCIQCpXMLlPe9j5rhs-MyaZU8c0-N7lVAD6Y63Y5VPPckDgA^%^3AQUQ3MjNmendCcVFONm1JYmhvQ19iLVNNTUJhajV2UTZ6ZlBTRXlFMUFYWm52OHI2UzNjRnBxdzVpTGJ6dUp2MmJfTW96OXFfZmpJSlpPcWQ4WkZkVTV6RUZ3dUpwbkVxNkd1MTN4OFhXZW9rZThqOTBaVTc0X0NDZGFOanoxTVZCTnlzcS0zWjBiVjRIc3VFUHlmeEJaamx1VXFzTHloZkNB^"
	a8 := "origin: https://www.youtube.com"
	a9 := "priority: u=1, i"
	a10 := "referer: https://www.youtube.com/watch?v=uBGsOezQ3vw"
	a11 := "sec-ch-ua: ^\\^\"Not-A.Brand^\\^\";v=^\\^\"99^\\^\", ^\\^\"Chromium^\\^\";v=^\\^\"124^\\^\"^"
	a12 := "sec-ch-ua-arch: ^\\^\"x86^\\^\"^"
	a13 := "sec-ch-ua-bitness: ^\\^\"64^\\^\"^"
	a14 := "sec-ch-ua-form-factors: ^\\^\"Desktop^\\^\"^"
	a15 := "sec-ch-ua-full-version: ^\\^\"124.0.6367.243^\\^\"^"
	a16 := "sec-ch-ua-full-version-list: ^\\^\"Not-A.Brand^\\^\";v=^\\^\"99.0.0.0^\\^\", ^\\^\"Chromium^\\^\";v=^\\^\"124.0.6367.243^\\^\"^"
	a17 := "sec-ch-ua-mobile: ?0"
	a18 := "sec-ch-ua-model: ^\\^\"^\\^\"^"
	a19 := "sec-ch-ua-platform: ^\\^\"Windows^\\^\"^"
	a20 := "sec-ch-ua-platform-version: ^\\^\"0.1.0^\\^\"^"
	a21 := "sec-ch-ua-wow64: ?0"
	a22 := "sec-fetch-dest: empty"
	a23 := "sec-fetch-mode: same-origin"
	a24 := "sec-fetch-site: same-origin"
	a25 := "user-agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/124.0.0.0 Safari/537.36"
	a26 := "x-goog-authuser: 0"
	a27 := "x-goog-visitor-id: CgtEQ2NaVWE2ODdFRSjFx8O1BjIKCgJSVRIEGgAgVg^%^3D^%^3D^"
	a28 := "x-origin: https://www.youtube.com"
	a29 := "x-youtube-bootstrap-logged-in: true"
	a30 := "x-youtube-client-name: 1"
	a31 := "x-youtube-client-version: 2.20240731.04.00"
	a32 := "--data-raw"
	a33 := "^{^^\"context^\\^\":^{^\n^\"client^\\^\":^{^\\^\"hl^\\^\":^\\^\"ru^\\^\",^\\^\"gl^\\^\":^\\^\"RU^\\^\",^\\^\"remoteHost^\\^\":^\\^\"177.126.48.14^\\^\",^\\^\"deviceMake^\\^\":^\\^\"^\\^\",^\\^\"deviceModel^\\^\":^\\^\"^\\^\",^\\^\"visitorData^\\^\":^\\^\"CgtEQ2NaVWE2ODdFRSjFx8O1BjIKCgJSVRIEGgAgVg^%^3D^%^3D^\\^\",^\\^\"userAgent^\\^\":^\\^\"Mozilla5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome4.0.0.0 Safari/537.36,gzip(gfe)^\\^\",^\\^\"clientName^\\^\":^\\^\"WEB^\\^\",^\\^\"clientVersion^\\^\":^\\^\"2.20240731.04.00^\\^\",^\\^\"osName^\\^\":^\\^\"Windows^\\^\",^\\^\"osVersion^\\^\":^\\^\"10.0^\\^\",^\\^\"originalUrl^\\^\":^\\^\"https://www.youtube.com/watch?v=X3aR5xk9fIM^\\^\",^\\^\"platform^\\^\":^\\^\"DESKTOP^\\^\",^\\^\"clientFormFactor^\\^\":^\\^\"UNKNOWN_FORM_FACTOR^\\^\",^\\^\"configInfo^\\^\":^{^\\^\"appInstallData^\\^\":^\\^\"CMXHw7UGEOrDrwUQkJKxBRCa8K8FENPhrwUQqJqwBRCNlLEFELfq_hIQyvmwBRCmkrEFEPGcsAUQ2_63IhDN17AFEIvPsAUQsO6wBRC1sf8SEJT-sAUQmZixBRCx3LAFEOX0sAUQ1-mvBRDQjbAFEL2KsAUQqLewBRCI468FEIiHsAUQ3ej-EhD0q7AFENWIsAUQieiuBRCok7EFEKensQUQ2qCxBRCmmrAFEMnmsAUQ782wBRCdprAFEI3MsAUQ4tSuBRCPxLAFEL22rgUQop2xBRDJ968FEP-IsQUQpcL-EhD2hrEFEOPRsAUQ3_WwBRCUibEFEJaVsAUQqJKxBRCTwP8SEJKdsQUQ65mxBRDZya8FEO6irwUQ65OuBRCPlLEFEOHssAUQooGwBRC9mbAFEJiNsQUQx5-xBRDuiLEFEJajsQUQntCwBRDJ17AFEPyFsAUQ4bz_EhDEkrEFEParsAUQ26-vBRDapbEFENShrwUQ-vCwBRDW3bAFEOvo_hIQqtiwBRDQ-rAFEK-hsQUQt--vBRCO2rAFELn4sAUQwaWxBRDM364FEIO5_xIQppOxBRDPnbEFEO2__xIQqdawBSooQ0FNU0dSVVNvTDJ3RE5Ia0J2YnFBTms3cE1mcEM0SDNBZm81SFFjPQ^%^3D^%^3D^\\^\"^},^\\^\"userInterfaceTheme^\\^\":^\\^\"USER_INTERFACE_THEME_LIGHT^\\^\",^\\^\"timeZone^\\^\":^\\^\"Europe/Moscow^\\^\",^\\^\"browserName^\\^\":^\\^\"Chrome^\\^\",^\\^\"browserVersion^\\^\":^\\^\"124.0.0.0^\\^\",^\\^\"acceptHeader^\\^\":^\\^\"text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7^\\^\",^\\^\"deviceExperimentId^\\^\":^\\^\"ChxOek01T1RZMk5EWXlNek0wTURVMU5UUTVNZz09EMXHw7UGGMXHw7UG^\\^\",^\\^\"screenWidthPoints^\\^\":811,^\\^\"screenHeightPoints^\\^\":652,^\\^\"screenPixelDensity^\\^\":1,^\\^\"screenDensityFloat^\\^\":1,^\\^\"utcOffsetMinutes^\\^\":180,^\\^\"connectionType^\\^\":^\\^\"CONN_CELLULAR_4G^\\^\",^\\^\"memoryTotalKbytes^\\^\":^\\^\"4000000^\\^\",^\\^\"mainAppWebInfo^\\^\":^{^\\^\"graftUrl^\\^\":^\\^\"https://www.youtube.com/watch?v=uBGsOezQ3vw^\\^\",^\\^\"pwaInstallabilityStatus^\\^\":^\\^\"PWA_INSTALLABILITY_STATUS_CAN_BE_INSTALLED^\\^\",^\\^\"webDisplayMode^\\^\":^\\^\"WEB_DISPLAY_MODE_BROWSER^\\^\",^\\^\"isWebNativeShareAvailable^\\^\":false^}^},^\\^\"user^\\^\":^{^\\^\"lockedSafetyMode^\\^\":false^},^\\^\"request^\\^\":^{^\\^\"useSsl^\\^\":true,^\\^\"internalExperimentFlags^\\^\":^[^],^\\^\"consistencyTokenJars^\\^\":^[^]^},^\\^\"clickTracking^\\^\":^{^\\^\"clickTrackingParams^\\^\":^\\^\"CF4Q2NgFIhMIw9eH-4vehwMVx8BJBx1AXSaD^\\^\"^},^\\^\"adSignalsInfo^\\^\":^{^\\^\"params^\\^\":^[^{^\\^\"key^\\^\":^\\^\"dt^\\^\",^\\^\"value^\\^\":^\\^\"1722868681194^\\^\"^},^{^\\^\"key^\\^\":^\\^\"flash^\\^\",^\\^\"value^\\^\":^\\^\"0^\\^\"^},^{^\\^\"key^\\^\":^\\^\"frm^\\^\",^\\^\"value^\\^\":^\\^\"0^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_tz^\\^\",^\\^\"value^\\^\":^\\^\"180^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_his^\\^\",^\\^\"value^\\^\":^\\^\"9^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_h^\\^\",^\\^\"value^\\^\":^\\^\"768^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_w^\\^\",^\\^\"value^\\^\":^\\^\"1366^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_ah^\\^\",^\\^\"value^\\^\":^\\^\"728^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_aw^\\^\",^\\^\"value^\\^\":^\\^\"1366^\\^\"^},^{^\\^\"key^\\^\":^\\^\"u_cd^\\^\",^\\^\"value^\\^\":^\\^\"24^\\^\"^},^{^\\^\"key^\\^\":^\\^\"bc^\\^\",^\\^\"value^\\^\":^\\^\"31^\\^\"^},^{^\\^\"key^\\^\":^\\^\"bih^\\^\",^\\^\"value^\\^\":^\\^\"652^\\^\"^},^{^\\^\"key^\\^\":^\\^\"biw^\\^\",^\\^\"value^\\^\":^\\^\"794^\\^\"^},^{^\\^\"key^\\^\":^\\^\"brdim^\\^\",^\\^\"value^\\^\":^\\^\"0,0,0,0,1366,0,1366,728,811,652^\\^\"^},^{^\\^\"key^\\^\":^\\^\"vis^\\^\",^\\^\"value^\\^\":^\\^\"1^\\^\"^},^{^\\^\"key^\\^\":^\\^\"wgl^\\^\",^\\^\"value^\\^\":^\\^\"true^\\^\"^},^{^\\^\"key^\\^\":^\\^\"ca_type^\\^\",^\\^\"value^\\^\":^\\^\"image^\\^\"^}^]^}^},^\\^\"createBackstagePostParams^\\^\":^\\^\"ChhVQ2RFMS05YzhGRzdvZ3ZkTFRIWHBkcGcQAQ^%^3D^%^3D^\\^\",^\\^\"commentText^\\^\":^\\^\"Hello^\\^\",^\\^\"clipAttachment^\\^\":^{^\\^\"externalVideoId^\\^\":^\\^\"-36t7tYGNT0^\\^\",^\\^\"offsetMs^\\^\":^\\^\"10000^\\^\",^\\^\"durationMs^\\^\":^\\^\"9000^\\^\"^}^}^"

	cmd := exec.Command(com, a1, a2, a3, a2, a4, a2, a5, a2, a6, a2, a7, a2, a8, a2, a9, a2, a10, a2, a11, a2, a12, a2, a13, a2, a14, a2, a15, a2, a16, a2, a17, a2, a18, a2, a19, a2, a20, a2, a21, a2, a22, a2, a23, a2, a24, a2, a25, a2, a26, a2, a27, a2, a28, a2, a29, a2, a30, a2, a31, a32, a33)

	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("%v", err)
	}

	fmt.Printf("%s\n", b)
}
