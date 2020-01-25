/**
 * Copyright 2020 Innodev LLC. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package common

import (
	"fmt"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

type FirebaseIDPRequest struct {
	RequestURI          string `json:"requestUri,omitempty"`
	PostBody            string `json:"postBody,omitempty"`
	ReturnSecureToken   bool   `json:"returnSecureToken,omitempty"`
	ReturnIdpCredential bool   `json:"returnIdpCredential,omitempty"`
}

func NewFirebaseIDPRequest(idToken, providerID, uri string, returnIdpCredential bool) FirebaseIDPRequest {
	return FirebaseIDPRequest{
		PostBody:            fmt.Sprintf("id_token=%s&providerId=%s", idToken, providerID),
		RequestURI:          uri,
		ReturnSecureToken:   true,
		ReturnIdpCredential: returnIdpCredential,
	}
}

type FirebaseIDPResponse struct {
	FederatedId      string `json:"federatedId,omitempty"`
	ProviderId       string `json:"providerId,omitempty"`
	LocalId          string `json:"localId,omitempty"`
	EmailVerified    bool   `json:"emailVerified,omitempty"`
	Email            string `json:"email,omitempty"`
	OauthIdToken     string `json:"oauthIdToken,omitempty"`
	OauthAccessToken string `json:"oauthAccessToken,omitempty"`
	OauthTokenSecret string `json:"oauthTokenSecret,omitempty"`
	RawUserInfo      string `json:"rawUserInfo,omitempty"`
	FirstName        string `json:"firstName,omitempty"`
	LastName         string `json:"lastName,omitempty"`
	FullName         string `json:"fullName,omitempty"`
	DisplayName      string `json:"displayName,omitempty"`
	PhotoUrl         string `json:"photoUrl,omitempty"`
	IdToken          string `json:"idToken,omitempty"`
	RefreshToken     string `json:"refreshToken,omitempty"`
	ExpiresIn        string `json:"expiresIn,omitempty"`
	NeedConfirmation bool   `json:"needConfirmation,omitempty"`
}

func (res FirebaseIDPResponse) Token() *oauth2.Token {
	secs, err := strconv.Atoi(res.ExpiresIn)
	if err != nil {
		secs = 0
	}
	return &oauth2.Token{
		AccessToken:  res.IdToken,
		RefreshToken: res.RefreshToken,
		Expiry:       time.Now().Add(time.Duration(secs) * time.Second),
	}
}
