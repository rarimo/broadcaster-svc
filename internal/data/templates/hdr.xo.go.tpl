{{- $tags := tags -}}
{{- $inject := inject -}}
{{- if $tags -}}
//go:build{{ range $tags }} {{ . }}{{ end }}

{{ end -}}
{{- if first -}}
// Package {{ pkg }} contains generated code for schema '{{ schema }}'.
{{ end -}}
package {{ pkg }}

// Code generated by xo. DO NOT EDIT.

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"strings"
	"time"
	"gitlab.com/rarimo/near-saver-svc/internal/data"
{{- if driver "postgres" }}
	"github.com/lib/pq/hstore"
	"github.com/xo/xo/types/xo"
{{ end }}{{ range imports }}
	{{ . }}
{{ end }}
)

{{- if $inject }}
{{ $inject }}
{{- end }}

