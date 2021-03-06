// Package adapter provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package adapter

// 権限問い合わせリクエスト
type RequestRights struct {
	// リソース
	Resource string `json:"resource"`

	// ユーザーID
	UserId string `json:"user_id"`
}

// PostV1RightsJSONBody defines parameters for PostV1Rights.
type PostV1RightsJSONBody RequestRights

// PostV1RightsJSONRequestBody defines body for PostV1Rights for application/json ContentType.
type PostV1RightsJSONRequestBody PostV1RightsJSONBody

