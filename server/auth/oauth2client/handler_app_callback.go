package oauth2client

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/oauth2"
)

// func CallbackHandler(c oauth2.Config) func(rw http.ResponseWriter, req *http.Request) {
// 	return func(rw http.ResponseWriter, req *http.Request) {
// 		codeVerifier := ResetPKCE(rw)
// 		rw.Write([]byte(`<h1>Callback site</h1><a href="/">Go back</a>`))
// 		rw.Header().Set("Content-Type", "text/html; charset=utf-8")
// 		if req.URL.Query().Get("error") != "" {
// 			rw.Write([]byte(fmt.Sprintf(`<h1>Error!</h1>
// 			Error: %s<br>
// 			Error Hint: %s<br>
// 			Description: %s<br>
// 			<br>`,
// 				req.URL.Query().Get("error"),
// 				req.URL.Query().Get("error_hint"),
// 				req.URL.Query().Get("error_description"),
// 			)))
// 			return
// 		}

// 		client := NewBasicClient(c.ClientID, c.ClientSecret)
// 		if req.URL.Query().Get("revoke") != "" {
// 			revokeURL := strings.Replace(c.Endpoint.TokenURL, "token", "revoke", 1)
// 			payload := url.Values{
// 				"token_type_hint": {"refresh_token"},
// 				"token":           {req.URL.Query().Get("revoke")},
// 			}
// 			resp, body, err := client.Post(revokeURL, payload)
// 			if err != nil {
// 				rw.Write([]byte(fmt.Sprintf(`<p>Could not revoke token %s</p>`, err)))
// 				return
// 			}

// 			rw.Write([]byte(fmt.Sprintf(`<p>Received status code from the revoke endpoint:<br><code>%d</code></p>`, resp.StatusCode)))
// 			if body != "" {
// 				rw.Write([]byte(fmt.Sprintf(`<p>Got a response from the revoke endpoint:<br><code>%s</code></p>`, body)))
// 			}

// 			rw.Write([]byte(fmt.Sprintf(`<p>These tokens have been revoked, try to use the refresh token by <br><a href="%s">by clicking here</a></p>`, "?refresh="+url.QueryEscape(req.URL.Query().Get("revoke")))))
// 			rw.Write([]byte(fmt.Sprintf(`<p>Try to use the access token by <br><a href="%s">by clicking here</a></p>`, "/protected?token="+url.QueryEscape(req.URL.Query().Get("access_token")))))

// 			return
// 		}

// 		if req.URL.Query().Get("refresh") != "" {
// 			payload := url.Values{
// 				"grant_type":    {"refresh_token"},
// 				"refresh_token": {req.URL.Query().Get("refresh")},
// 				"scope":         {"fosite"},
// 			}
// 			_, body, err := client.Post(c.Endpoint.TokenURL, payload)
// 			if err != nil {
// 				rw.Write([]byte(fmt.Sprintf(`<p>Could not refresh token %s</p>`, err)))
// 				return
// 			}
// 			rw.Write([]byte(fmt.Sprintf(`<p>Got a response from the refresh grant:<br><code>%s</code></p>`, body)))
// 			return
// 		}

// 		if req.URL.Query().Get("code") == "" {
// 			rw.Write([]byte(fmt.Sprintln(`<p>Could not find the authorize code. If you've used the implicit grant, check the
// 			browser location bar for the
// 			access token <small><a href="http://en.wikipedia.org/wiki/Fragment_identifier#Basics">(the server side does not have access to url fragments)</a></small>
// 			</p>`,
// 			)))
// 			return
// 		}

// 		rw.Write([]byte(fmt.Sprintf(`<p>Amazing! You just got an authorize code!:<br><code>%s</code></p>
// 		<p>Click <a href="/">here to return</a> to the front page</p>`,
// 			req.URL.Query().Get("code"),
// 		)))

// 		// We'll check whether we sent a code+PKCE request, and if so, send the code_verifier along when requesting the access token.
// 		var opts []oauth2.AuthCodeOption
// 		if IsPKCE(req) {
// 			opts = append(opts, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
// 		}

// 		token, err := c.Exchange(context.Background(), req.URL.Query().Get("code"), opts...)
// 		if err != nil {
// 			rw.Write([]byte(fmt.Sprintf(`<p>I tried to exchange the authorize code for an access token but it did not work but got error: %s</p>`, err.Error())))
// 			return
// 		}

// 		rw.Write([]byte(fmt.Sprintf(`<p>Cool! You are now a proud token owner.<br>
// 		<ul>
// 			<li>
// 				Access token (click to make <a href="%s">authorized call</a>):<br>
// 				<code>%s</code>
// 			</li>
// 			<li>
// 				Refresh token (click <a href="%s">here to use it</a>) (click <a href="%s">here to revoke it</a>):<br>
// 				<code>%s</code>
// 			</li>
// 			<li>
// 				Extra info: <br>
// 				<code>%s</code>
// 			</li>
// 		</ul>`,
// 			"/protected?token="+token.AccessToken,
// 			token.AccessToken,
// 			"?refresh="+url.QueryEscape(token.RefreshToken),
// 			"?revoke="+url.QueryEscape(token.RefreshToken)+"&access_token="+url.QueryEscape(token.AccessToken),
// 			token.RefreshToken,
// 			token,
// 		)))
// 	}
// }




func (s *accountHandler) CallbackHandler(gctx *gin.Context) {

	rw := gctx.Writer
	req := gctx.Request

	var c = auth.ClientConf
	codeVerifier := oauth2client.ResetPKCE(rw)
	// if req.URL.Query().Get("error") != "" {
	// 	logrus.Errorln("ResetPKCE error: %s, %s, %s", req.URL.Query().Get("error"),
	// 		req.URL.Query().Get("error_hint"),
	// 		req.URL.Query().Get("error_description"))
	// 	utils.ResponseMessage(gctx, http.StatusBadRequest, "ResetPKCE error")
	// 	return
	// }

	client := oauth2client.NewBasicClient(c.ClientID, c.ClientSecret)
	if req.URL.Query().Get("revoke") != "" {
		revokeURL := strings.Replace(c.Endpoint.TokenURL, "token", "revoke", 1)
		payload := url.Values{
			"token_type_hint": {"refresh_token"},
			"token":           {req.URL.Query().Get("revoke")},
		}
		resp, body, err := client.Post(revokeURL, payload)
		if err != nil {
			logrus.Errorln("Could not revoke token: %w", err)
			utils.ResponseMessage(gctx, http.StatusBadRequest, "Could not revoke token")
			return
		}
		logrus.Infoln("Received status code from the revoke endpoint: %d", resp.StatusCode)

		if body != "" {
			logrus.Infoln("Got a response from the revoke endpoint: %s", body)
		}

		logrus.Infoln("These tokens have been revoked, try to use the refresh token by %s", "?refresh="+url.QueryEscape(req.URL.Query().Get("revoke")))
		logrus.Infoln("Try to use the access token by  %s", "/protected?token="+url.QueryEscape(req.URL.Query().Get("access_token")))

		return
	}

	if req.URL.Query().Get("refresh") != "" {
		payload := url.Values{
			"grant_type":    {"refresh_token"},
			"refresh_token": {req.URL.Query().Get("refresh")},
			"scope":         {"fosite"},
		}
		_, body, err := client.Post(c.Endpoint.TokenURL, payload)
		if err != nil {
			logrus.Errorln("Could not refresh token: %w", err)
			utils.ResponseMessage(gctx, http.StatusBadRequest, "Could not refresh token")
			return
		}
		logrus.Infoln("Got a response from the refresh grant: %s", body)

		return
	}

	if req.URL.Query().Get("code") == "" {
		logrus.Errorln("Could not find the authorize code")
		return
	}
	logrus.Infoln("Amazing! You just got an authorize code: %s", req.URL.Query().Get("code"))

	var opts []oauth2.AuthCodeOption
	if oauth2client.IsPKCE(req) {
		opts = append(opts, oauth2.SetAuthURLParam("code_verifier", codeVerifier))
	}

	token, err := c.Exchange(context.Background(), req.URL.Query().Get("code"), opts...)
	if err != nil {
		logrus.Errorln("I tried to exchange the authorize code for an access token but it did not work but got error: %w", err)
		return
	}

	// rw.Write([]byte(fmt.Sprintf(`<p>Cool! You are now a proud token owner.<br>
	// <ul>
	// 	<li>
	// 		Access token (click to make <a href="%s">authorized call</a>):<br>
	// 		<code>%s</code>
	// 	</li>
	// 	<li>
	// 		Refresh token (click <a href="%s">here to use it</a>) (click <a href="%s">here to revoke it</a>):<br>
	// 		<code>%s</code>
	// 	</li>
	// 	<li>
	// 		Extra info: <br>
	// 		<code>%s</code>
	// 	</li>
	// </ul>`,
	// 	"/protected?token="+token.AccessToken,
	// 	token.AccessToken,
	// 	"?refresh="+url.QueryEscape(token.RefreshToken),
	// 	"?revoke="+url.QueryEscape(token.RefreshToken)+"&access_token="+url.QueryEscape(token.AccessToken),
	// 	token.RefreshToken,
	// 	token,
	// )))

	gctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
