package tiktoken

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
)

func TestEncoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := EncodingForModel("gpt-3.5-turbo-16k")
	ass.Nil(err, "Encoding  init should not be nil")
	tokens := enc.Encode("hello world!你好，世界！", []string{"all"}, []string{"all"})
	// these tokens are converted from the original python code
	sourceTokens := []int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, nil)
	sourceTokens = []int{15339, 220, 100257}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, []string{"all"})
	sourceTokens = []int{15339, 220, 100257}
	ass.ElementsMatch(sourceTokens, tokens, "Encoding should be equal")

	ass.Panics(func() {
		tokens = enc.Encode("hello <|endoftext|><|endofprompt|>", []string{"<|endoftext|>"}, []string{"all"})
	})
	ass.Panics(func() {
		tokens = enc.Encode("hello <|endoftext|>", []string{"<|endoftext|>"}, []string{"<|endoftext|>"})
	})
}

type urlRewriteLoader struct {
	realBase string
	fakeBase string
	inner    BpeLoader
}

func (u *urlRewriteLoader) LoadTiktokenBpe(url string) (map[string]int, error) {
	url = strings.Replace(url, u.realBase, u.fakeBase, 1)
	return u.inner.LoadTiktokenBpe(url)
}

func TestGetEncoding_ErrorResponseNotCached(t *testing.T) {
	cacheDir := t.TempDir()
	t.Setenv("TIKTOKEN_CACHE_DIR", cacheDir)

	ass := assert.New(t)
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal server error", http.StatusNotFound)
	}))

	loader := &urlRewriteLoader{
		realBase: "https://openaipublic.blob.core.windows.net",
		fakeBase: ts.URL,
		inner:    NewDefaultBpeLoader(),
	}
	SetBpeLoader(loader)

	t.Cleanup(func() {
		SetBpeLoader(NewDefaultBpeLoader())
	})

	got, err := GetEncoding(MODEL_O200K_BASE)
	ass.Nil(got)
	ass.Error(err, "expected error when fetching encoding with bad response")

	// This fails right now: bad response body is cached despite the error
	entries, _ := os.ReadDir(cacheDir)
	ass.Empty(entries, "expected empty cache dir after error")
}

func TestDecoding(t *testing.T) {
	ass := assert.New(t)
	enc, err := GetEncoding(MODEL_CL100K_BASE)
	ass.Nil(err, "Encoding  init should not be nil")
	sourceTokens := []int{15339, 1917, 0, 57668, 53901, 3922, 3574, 244, 98220, 6447}
	txt := enc.Decode(sourceTokens)
	ass.Equal("hello world!你好，世界！", txt, "Decoding should be equal")
}

func TestEncodingForModel_Names(t *testing.T) {
	for model := range MODEL_TO_ENCODING {
		// we don't support gpt2 model so far
		if model == "gpt2" {
			continue
		}
		t.Run("Check model "+model, func(t *testing.T) {
			t.Parallel()
			testEncodingForModel(t, model)
		})
	}
}

func TestEncodingForModel_Prefixes(t *testing.T) {
	for prefix := range MODEL_PREFIX_TO_ENCODING {
		t.Run("Check prefix "+prefix, func(t *testing.T) {
			t.Parallel()
			testEncodingForModel(t, prefix)
		})
	}
}

func testEncodingForModel(t *testing.T, model string) {
	t.Helper()

	text := "hello world"
	ass := assert.New(t)
	req := require.New(t)

	tkm, err := EncodingForModel(model)
	req.NoErrorf(err, "error getting encoding for model %q: %v", model, err)
	ass.NotNil(tkm, "Encoding for model %s should not be nil", model)

	encText := tkm.Encode(text, nil, nil)
	ass.Len(encText, 2, "Encoding len should be equal")

	decText := tkm.Decode(encText)
	ass.Equal(text, decText, "decoding mismatch - want: %s, got: %s", text, decText)
}
