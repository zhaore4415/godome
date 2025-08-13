package utils

import (
	"context"
	"encoding/json"
	"errors"

	"google.golang.org/grpc/metadata"
)

const (
	UserInfoKey = "X-User-Info"
	LanguageKey = "X-Language"
)

type UserInfo struct {
	AppId       string  `json:"app_id"`
	Id          string  `json:"id"`
	Username    string  `json:"username"`
	CompanyId   string  `json:"company_id"`
	CompanyName string  `json:"company_name"`
	Roles       []int64 `json:"roles"`
	OrgId       string  `json:"org_id"`
}

func SetLanguage(ctx context.Context, language string) (context.Context, error) {
	md := metadata.New(map[string]string{
		LanguageKey: language,
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, nil
}

func GetLanguage(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	return md.Get(LanguageKey)[0]
}

func SetUserMetadata(ctx context.Context, userMetadata *UserInfo) (context.Context, error) {
	if userMetadata == nil {
		return ctx, errors.New("user metadata is nil")
	}
	b, err := json.Marshal(userMetadata)
	if err != nil {
		return ctx, err
	}
	md := metadata.New(map[string]string{
		UserInfoKey: string(b),
	})

	ctx = metadata.NewOutgoingContext(ctx, md)

	return ctx, nil
}

func GetUserMetadata(ctx context.Context) *UserInfo {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil
	}

	userInfo := md.Get(UserInfoKey)
	if len(userInfo) == 0 {
		return nil
	}

	var userMetadata UserInfo
	err := json.Unmarshal([]byte(userInfo[0]), &userMetadata)
	if err != nil {
		return nil
	}
	return &userMetadata
}
