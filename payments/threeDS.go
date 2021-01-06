package payments

import "github.com/checkout/checkout-sdk-go/common"

// ThreeDS ...
type ThreeDS struct {
	Enabled            *bool                     `json:"enabled,omitempty"`
	AttemptN3d         *bool                     `json:"attempt_n3d,omitempty"`
	ChallengeIndicator common.ChallengeIndicator `json:"challenge_indicator,omitempty"`
	Exemption          common.Exemption          `json:"exemption,omitempty"`
	ECI                string                    `json:"eci,omitempty"`
	Cryptogram         string                    `json:"cryptogram,omitempty"`
	XID                string                    `json:"xid,omitempty"`
	Version            string                    `json:"version,omitempty"`
}

// ThreeDSEnrollment : 3D-Secure Enrollment Data
type ThreeDSEnrollment struct {
	Downgraded             *bool  `json:"downgraded,omitempty"`
	Enrolled               string `json:"enrolled,omitempty"`
	SignatureValid         string `json:"signature_valid,omitempty"`
	AuthenticationResponse string `json:"authentication_response,omitempty"`
	Cryptogram             string `json:"cryptogram,omitempty"`
	XID                    string `json:"xid,omitempty"`
}
