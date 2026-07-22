# Platform

`platform` is an unofficial Go SDK collection for Facebook, Instagram, Threads, TikTok, X, and YouTube. Each platform is an independent package with typed request and response contracts, OAuth helpers, endpoint groups, scopes, named errors, and an injectable HTTP client.

This project is not affiliated with, endorsed by, or supported by any of the platforms it integrates with. You are responsible for complying with each provider's terms, developer policies, and API requirements.

## Install

```sh
go get github.com/social-ally/platform
```

The module currently targets Go 1.26.

## Packages

| Platform | Go package | Included groups |
| --- | --- | --- |
| Facebook | `facebook` | OAuth, users, pages, publishing, analytics |
| Instagram | `instagram` | OAuth, users, media, publishing, analytics |
| Threads | `threads` | OAuth, users, media, publishing, analytics |
| TikTok | `tiktok` | OAuth, users, videos, publishing |
| X | `x` | OAuth, users, posts, media, analytics |
| YouTube | `youtube` | OAuth, channels, videos, playlists, analytics |

Each package exports `DisplayName`, a primary `BaseURL`, scoped base URL constants where a provider has multiple API origins, and a platform-specific client type.

## Quick start: X

Create a client with the scopes your application needs. Provide `WithAccessToken` for calls to authenticated endpoint groups.

```go
package main

import (
    "context"
    "fmt"

    "github.com/social-ally/platform/x"
)

func main() {
    client, err := x.NewXClient(
        "client-id",
        "",
        "https://example.com/oauth/callback",
        x.WithScopes(x.ScopeUsersRead, x.ScopeTweetRead),
        x.WithAccessToken("user-access-token"),
    )
    if err != nil {
        panic(err)
    }

    response, err := x.NewUsers(client).GetAuthenticatedUser(
        context.Background(),
        &x.RequestGetAuthenticatedUser{},
    )
    if err != nil {
        panic(err)
    }
    fmt.Println(response.Success.Data.Username)
}
```

For an X confidential client, supply the client secret and `x.WithConfidentialClient()`. Token requests will use HTTP Basic authentication.

## OAuth authorization-code flow

OAuth groups create the provider authorization URL and exchange the returned code. Each package’s request and response types are named `Request<Method>` and `Response<Method>`.

```go
client, err := x.NewXClient(
    "client-id",
    "",
    "https://example.com/oauth/callback",
    x.WithScopes(x.ScopeUsersRead),
)
if err != nil {
    return err
}

authorize, err := x.NewOAuth(client).Authorize(ctx, &x.RequestAuthorize{
    Query: x.RequestAuthorizeQuery{
        State:         "csrf-state",
        CodeChallenge: "pkce-code-challenge",
    },
})
if err != nil {
    return err
}
// Redirect the user to authorize.URL.

tokens, err := x.NewOAuth(client).ExchangeCode(ctx, &x.RequestExchangeCode{
    Body: x.RequestExchangeCodeBody{
        Code:         returnedCode,
        CodeVerifier: "pkce-code-verifier",
    },
})
if err != nil {
    return err
}
_ = tokens.Success.AccessToken
```

TikTok supports PKCE with `tiktok.WithPKCE()`. X authorization-code flows require a PKCE challenge and verifier. Scope constants live in each package’s `scope.go`.

## Endpoint groups

Construct a group from its platform client and call its typed method:

```go
text := "Hello from Go"
posts := x.NewPosts(client)
response, err := posts.CreatePost(ctx, &x.RequestCreatePost{
    Body: x.RequestCreatePostBody{Text: &text},
})
```

All authenticated calls attach the configured bearer token. Clients can use a custom transport for proxies, retries, tracing, or tests:

```go
client, err := youtube.NewYouTubeClient(
    "client-id", "client-secret", "https://example.com/callback",
    youtube.WithScopes(youtube.ScopeYoutubeUpload),
    youtube.WithAccessToken("access-token"),
    youtube.WithHTTPClient(customHTTPClient),
)
```

YouTube video uploads accept media through `RequestUploadVideo.Media` (`io.Reader`) and use a multipart upload when media is provided.

## Errors

Packages expose named sentinel errors for configuration, missing tokens, missing identifiers, nil endpoint clients, and OAuth validation. API responses outside the 2xx range return that package’s `*APIError`, which matches `ErrUnexpectedStatus` through `errors.Is`.

```go
if errors.Is(err, x.ErrUnexpectedStatus) {
    var apiErr *x.APIError
    if errors.As(err, &apiErr) {
        fmt.Println(apiErr.StatusCode, string(apiErr.Body))
    }
}
```

## Development

```sh
go test ./...
go vet ./...
```
