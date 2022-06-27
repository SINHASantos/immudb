/*
Copyright 2022 CodeNotary, Inc. All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package s3

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestOpen(t *testing.T) {
	s, err := Open(
		"http://localhost:9000",
		"minioadmin",
		"minioadmin",
		"immudb",
		"",
		"prefix",
	)
	require.NoError(t, err)
	require.NotNil(t, s)
	require.Equal(t, "s3:http://localhost:9000/immudb/prefix/", s.String())
}

func TestCornerCases(t *testing.T) {
	t.Run("bucket name can not be empty", func(t *testing.T) {
		s, err := Open(
			"http://localhost:9000",
			"minioadmin",
			"minioadmin",
			"",
			"",
			"",
		)
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsBucketEmpty)
		require.Nil(t, s)
	})

	t.Run("bucket name can not contain /", func(t *testing.T) {
		s, err := Open(
			"http://localhost:9000",
			"minioadmin",
			"minioadmin",
			"immudb/test",
			"",
			"",
		)
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsBucketSlash)
		require.Nil(t, s)
	})

	t.Run("prefix must be correctly normalized", func(t *testing.T) {
		s, err := Open(
			"http://localhost:9000",
			"minioadmin",
			"minioadmin",
			"immudb",
			"",
			"",
		)
		require.NoError(t, err)
		require.Equal(t, "", s.(*Storage).prefix)

		s, err = Open(
			"http://localhost:9000",
			"minioadmin",
			"minioadmin",
			"immudb",
			"",
			"/test/",
		)
		require.NoError(t, err)
		require.Equal(t, "test/", s.(*Storage).prefix)

		s, err = Open(
			"http://localhost:9000",
			"minioadmin",
			"minioadmin",
			"immudb",
			"",
			"/test",
		)
		require.NoError(t, err)
		require.Equal(t, "test/", s.(*Storage).prefix)
	})

	t.Run("invalid url", func(t *testing.T) {
		s, err := Open(
			"h**s://localhost:9000",
			"minioadmin",
			"minioadmin",
			"bucket",
			"",
			"",
		)
		require.NoError(t, err)
		require.Equal(t, "s3(misconfigured)", s.String())
	})

	t.Run("invalid get / put path", func(t *testing.T) {
		s, err := Open(
			"htts://localhost:9000",
			"minioadmin",
			"minioadmin",
			"bucket",
			"",
			"",
		)
		require.NoError(t, err)

		_, err = s.Get(context.Background(), "/file", 0, -1)
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsNameSlash)

		_, err = s.Get(context.Background(), "file/", 0, -1)
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsNameSlash)

		err = s.Put(context.Background(), "/file", "/tmp/test.txt")
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsNameSlash)

		err = s.Put(context.Background(), "file/", "/tmp/test.txt")
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsNameSlash)
	})

	t.Run("invalid get offset / size", func(t *testing.T) {
		s, err := Open(
			"htts://localhost:9000",
			"minioadmin",
			"minioadmin",
			"bucket",
			"",
			"",
		)
		require.NoError(t, err)

		_, err = s.Get(context.Background(), "file", 0, 0)
		require.ErrorIs(t, err, ErrInvalidArguments)
		require.ErrorIs(t, err, ErrInvalidArgumentsOffsSize)
	})
}

func TestSignatureV4(t *testing.T) {
	// Example request available at:
	//  https://docs.aws.amazon.com/AmazonS3/latest/API/sig-v4-header-based-auth.html
	s, err := Open(
		"https://examplebucket.s3.amazonaws.com",
		"AKIAIOSFODNN7EXAMPLE",
		"wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY",
		"examplebucket",
		"us-east-1",
		"",
	)
	require.NoError(t, err)

	st := s.(*Storage)

	url, err := st.originalRequestURL("test.txt")
	require.NoError(t, err)

	req, err := st.s3SignedRequestV4(
		context.Background(),
		url,
		"GET",
		nil,
		"",
		"e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		func(req *http.Request) error {
			req.Header.Add("Range", "bytes=0-9")
			return nil
		},
		time.Date(2013, 5, 24, 0, 0, 0, 0, time.UTC),
	)
	require.NoError(t, err)
	require.NotNil(t, req)

	require.Equal(t,
		"AWS4-HMAC-SHA256 "+
			"Credential=AKIAIOSFODNN7EXAMPLE/20130524/us-east-1/s3/aws4_request,"+
			"SignedHeaders=host;range;x-amz-content-sha256;x-amz-date,"+
			"Signature=f0e8bdb87c964420e857bd35b5d6ed310bd44f0170aba48dd91039c6036bdb41",
		req.Header.Get("Authorization"),
	)
}

func TestHandlingRedirects(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		switch r.URL.Path {
		case "/bucket/object1":
			require.Equal(t, "GET", r.Method)
			http.Redirect(w, r, "/bucket/object2", http.StatusSeeOther)

		case "/bucket/object2":
			require.Equal(t, "GET", r.Method)
			http.Redirect(w, r, "/bucket/object3", http.StatusPermanentRedirect)

		case "/bucket/object3":
			require.Equal(t, "GET", r.Method)

			_, err := w.Write([]byte("Hello world"))
			require.NoError(t, err)

		case "/bucket/object4":
			require.Equal(t, "GET", r.Method)
			http.Redirect(w, r, "/bucket/object4", http.StatusTemporaryRedirect)

		case "/bucket/object5":
			require.Equal(t, "GET", r.Method)
			http.Redirect(w, r, "h**p://invalid", http.StatusSeeOther)

		case "/bucket/object6":
			require.Equal(t, "GET", r.Method)
			http.Redirect(w, r, "h**p://invalid", http.StatusTemporaryRedirect)

		default:
			require.Fail(t, "Unknown request")
		}

	}))
	defer ts.Close()

	s, err := Open(ts.URL, "", "", "bucket", "", "")
	require.NoError(t, err)

	ctx := context.Background()

	t.Run("open bucket with redirects", func(t *testing.T) {

		io, err := s.Get(ctx, "object1", 0, -1)
		require.NoError(t, err)

		b, err := ioutil.ReadAll(io)
		require.NoError(t, err)

		err = io.Close()
		require.NoError(t, err)

		require.Equal(t, []byte("Hello world"), b)
	})

	t.Run("detect infinite redirect loop", func(t *testing.T) {
		_, err := s.Get(ctx, "object4", 0, -1)
		require.ErrorIs(t, err, ErrTooManyRedirects)
	})

	t.Run("error getting 303 redirect", func(t *testing.T) {
		_, err := s.Get(ctx, "object5", 0, -1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to parse Location header")
	})

	t.Run("error getting 307 redirect", func(t *testing.T) {
		_, err := s.Get(ctx, "object6", 0, -1)
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to parse Location header")
	})

}
