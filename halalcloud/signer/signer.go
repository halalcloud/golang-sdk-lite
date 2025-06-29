package signer

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"
)

// SigningConfig 包含API签名所需的配置
type SigningConfig struct {
	SecretID    string
	SecretKey   string
	AccessToken string
	RequestBody []byte // 使用[]byte以支持二进制数据
	Method      string
	// ContentType string
	// Nonce       string
	// Timestamp     string
	UtcTime       time.Time
	ApiPath       string
	Params        map[string]string
	Headers       map[string]string
	HeadersToSign []string
}

// WithTimeout 设置超时时间的选项

// DefaultConfig 返回一个默认的配置，用于测试或示例
func NewConfig(apiHost string, secretID, secretKey, accessToken string, requestBody []byte, method string, apiPath string, params map[string]string, headers map[string]string, headersToSign []string) *SigningConfig {
	utcTime := time.Now().UTC()
	//if err != nil {
	//	return "", fmt.Errorf("解析时间戳错误: %w", err)
	// }
	tx := &SigningConfig{
		SecretID:    secretID,
		SecretKey:   secretKey,
		AccessToken: accessToken,
		RequestBody: requestBody,
		Method:      method,
		// ContentType: "application/json; charset=utf-8",
		// Nonce:       ,
		UtcTime: utcTime,
		ApiPath: apiPath,
		Params:  params,
		Headers: map[string]string{
			// "content-type":   "application/json; charset=utf-8",
			"host":           apiHost,
			"x-hl-nonce":     strconv.FormatInt(utcTime.UnixNano(), 36),
			"x-hl-timestamp": utcTime.Format(time.RFC3339),
			"other-header":   "other-value",
		},
		HeadersToSign: []string{
			// "content-type",
			"host",
			"x-hl-nonce",
			"x-hl-timestamp",
			"other-header",
		},
	}
	for k, v := range headers {
		lowerKey := strings.ToLower(k)
		if lowerKey == "authorization" {
			continue
		}
		tx.Headers[lowerKey] = v
		if lowerKey == "content-type" || strings.HasPrefix(lowerKey, "x-hl-") {
			tx.HeadersToSign = append(tx.HeadersToSign, lowerKey)
		}
	}
	for _, v := range headersToSign {
		lowerKey := strings.ToLower(v)
		if lowerKey == "authorization" {
			continue
		}
		tx.HeadersToSign = append(tx.HeadersToSign, lowerKey)
	}
	// 确保HeadersToSign唯一
	tx.HeadersToSign = uniqueStrings(tx.HeadersToSign)
	return tx
}

func uniqueStrings(s []string) []string {
	uniqueMap := make(map[string]struct{})
	for _, str := range s {
		uniqueMap[str] = struct{}{}
	}
	uniqueList := make([]string, 0, len(uniqueMap))
	for str := range uniqueMap {
		uniqueList = append(uniqueList, str)
	}
	return uniqueList
}

// Signer 处理API请求签名
type Signer struct {
	Config            *SigningConfig
	sortedQueryString string // 用于存储排序后的查询字符串
}

// NewSigner 创建一个新的签名器
func NewSigner(config *SigningConfig) *Signer {
	sg := &Signer{Config: config}
	sg.sign()
	return sg
}

// GenerateAuthorization 生成完整的授权头信息
func (s *Signer) sign() {
	// 解析时间

	// 确保UTC时间并获取日期字符串
	// utcTime = utcTime.UTC()
	dateString := s.Config.UtcTime.Format("2006-01-02")

	// 构建签名所需的各个组件
	credentialScope := s.createCredentialScope(dateString)
	signedHeaders, canonicalHeaders := s.createCanonicalHeaders()
	canonicalRequest := s.createCanonicalRequest(signedHeaders, canonicalHeaders)
	hashedCanonicalRequest := sha256hex(canonicalRequest)
	stringToSign := s.createStringToSign(credentialScope, hashedCanonicalRequest)

	// 派生签名密钥
	signature := s.calculateSignature(dateString, stringToSign)

	// 构建授权头
	authorization := fmt.Sprintf("%s Credential=%s/%s, SignedHeaders=%s, Signature=%s",
		DefaultSignAlgorithm,
		s.Config.SecretID,
		credentialScope,
		signedHeaders,
		signature,
	)

	// 准备返回的头部信息
	s.Config.Headers["authorization"] = authorization

}

func (s *Signer) GetHeaders() map[string]string {
	// 确保签名已生成

	return s.Config.Headers
}

func (s *Signer) GetQueryString() string {
	// 确保签名已生成

	return s.sortedQueryString
}

func (s *Signer) GetRequestURL(SSL bool) string {
	queryString := s.GetQueryString()
	sb := strings.Builder{}
	scheme := "http"
	if SSL {
		scheme = "https"
	}
	sb.WriteString(fmt.Sprintf("%s://%s%s", scheme, s.Config.Headers["host"], s.Config.ApiPath))
	if queryString != "" {
		sb.WriteString("?")
		sb.WriteString(queryString)
	}
	return sb.String()
}

// 以下是辅助方法

func (s *Signer) createCanonicalHeaders() (string, string) {
	// 复制并排序待签名的头部
	headersToSign := make([]string, len(s.Config.HeadersToSign))
	copy(headersToSign, s.Config.HeadersToSign)
	sort.Strings(headersToSign)

	// 构建规范头部
	var sb strings.Builder
	for _, header := range headersToSign {
		val, exists := s.Config.Headers[header]
		if exists {
			sb.WriteString(fmt.Sprintf("%s:%s\n", header, val))
		}
	}

	return strings.Join(headersToSign, ";"), sb.String()
}

func (s *Signer) createCanonicalRequest(signedHeaders, canonicalHeaders string) string {
	// 构建并返回规范请求
	return strings.Join([]string{
		s.Config.Method,
		s.Config.ApiPath,
		s.createSortedQueryString(),
		canonicalHeaders,
		signedHeaders,
		sha256hexBytes(s.Config.RequestBody),
	}, "\n")
}

func (s *Signer) createSortedQueryString() string {
	if len(s.Config.Params) == 0 {
		return ""
	}

	var sb strings.Builder
	encodedParams := map[string]string{}
	unsortedKeys := make([]string, 0, len(s.Config.Params))
	for k, v := range s.Config.Params {
		key := rfc3986Encode(k)
		if v == "" {
			encodedParams[key] = ""
		} else {
			encodedParams[key] = rfc3986Encode(v)
		}
		unsortedKeys = append(unsortedKeys, key)
	}
	// Sort the keys
	// Note: In a real implementation, you would sort the keys alphabetically.
	// Here we assume the keys are already sorted for simplicity.
	sort.Strings(unsortedKeys)
	for i, key := range unsortedKeys {
		if i > 0 {
			sb.WriteString("&")
		}
		sb.WriteString(fmt.Sprintf("%s=%s", key, encodedParams[key]))
	}

	sx := sb.String()
	s.sortedQueryString = sx
	return sx
}

func (s *Signer) createStringToSign(credentialScope, hashedCanonicalRequest string) string {
	// 构建待签名字符串
	return strings.Join([]string{
		DefaultSignAlgorithm,
		s.Config.UtcTime.Format(time.RFC3339),
		credentialScope,
		hashedCanonicalRequest,
	}, "\n")
}

func (s *Signer) createCredentialScope(dateString string) string {
	return fmt.Sprintf("%s/%s/%s", dateString, s.Config.AccessToken, RequestSuffix)
}

func (s *Signer) calculateSignature(dateString, stringToSign string) string {
	// 派生密钥
	secretKey := []byte(SignPrefix + s.Config.SecretKey)
	dateKey := hmacSha256(secretKey, []byte(dateString))
	accessTokenKey := hmacSha256(dateKey, []byte(s.Config.AccessToken))
	signingKey := hmacSha256(accessTokenKey, []byte(RequestSuffix))

	// 计算签名
	signature := hex.EncodeToString(hmacSha256(signingKey, []byte(stringToSign)))

	return signature
}

// 工具函数

func sha256hex(s string) string {
	b := sha256.Sum256([]byte(s))
	return hex.EncodeToString(b[:])
}

func sha256hexBytes(s []byte) string {
	b := sha256.Sum256(s)
	return hex.EncodeToString(b[:])
}

func hmacSha256(key, data []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}

// rfc3986Encode 实现RFC3986标准的URL编码
func rfc3986Encode(s string) string {
	return strings.ReplaceAll(url.QueryEscape(s), "+", "%20")
}
