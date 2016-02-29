package braintree

import (
	"encoding/base64"
	"encoding/xml"
	"strings"
)

type WebhookNotificationGateway struct {
	*Braintree
}

func parseSignature(publicKey, signature string) (string, error) {

	signatureKeyPairs := strings.Split(signature, "&")

	for _, signatureKeyPair := range signatureKeyPairs {
		if !strings.Contains(signatureKeyPair, "|") {
			return "", SignatureError{"Signature-key pair does not contain |"}
		}
		split := strings.Split(signatureKeyPair, "|")
		if len(split) != 2 {
			return "", SignatureError{"Signature-key pair contains more than one |"}
		}
		if split[0] == publicKey {
			return split[1], nil
		}

	}

	return "", SignatureError{"No Signature-key pair found with matching public key"}
}

func (w *WebhookNotificationGateway) Parse(signatures, payload string) (*WebhookNotification, error) {

	matchedSignature, err := parseSignature(w.Braintree.PublicKey, signatures)
	if err != nil {
		return nil, err
	}

	hmacer := newHmacer(w.Braintree)
	if verified, err := hmacer.verifySignature(matchedSignature, payload); err != nil {
		return nil, err
	} else if !verified {
		return nil, SignatureError{}
	}

	xmlNotification, err := base64.StdEncoding.DecodeString(payload)
	if err != nil {
		return nil, err
	}

	var n WebhookNotification
	err = xml.Unmarshal(xmlNotification, &n)
	if err != nil {
		return nil, err
	}
	return &n, nil

}

func (w *WebhookNotificationGateway) Encode(notification *WebhookNotification) (signature, payload string, err error) {

	xmlNotification, err := xml.Marshal(notification)
	if err != nil {
		return "", "", err
	}

	payload = base64.StdEncoding.EncodeToString(xmlNotification)

	hmacer := newHmacer(w.Braintree)
	hmacedPayload, err := hmacer.hmac(payload)
	if err != nil {
		return "", "", err
	}
	signature = w.PublicKey + "|" + hmacedPayload

	return signature, payload, nil
}

func (w *WebhookNotificationGateway) Verify(challenge string) (string, error) {
	hmacer := newHmacer(w.Braintree)
	digest, err := hmacer.hmac(challenge)
	if err != nil {
		return ``, err
	}
	return hmacer.PublicKey + `|` + digest, nil
}
