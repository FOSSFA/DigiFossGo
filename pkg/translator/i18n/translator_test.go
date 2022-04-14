package i18n

import (
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator"
	"github.com/amirhossein-ka/DigiFossGo/pkg/translator/messages"
	"testing"
)

const dataPath string = "./testdata/"

func TestTranslator(t *testing.T) {

	tr, err := New(dataPath)
	if err != nil {
		t.Errorf("New() error: %v, arguments: %s", err, dataPath)
	}

	type args struct {
		key  string
		lang translator.Lang
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "translate fa",
			args: args{
				key:  messages.NoUserFound,
				lang: translator.GetLanguage("fa"),
			},
			want: "کاربر مورد نظر یافت نشد.",
		},
		{
			name: "translate en",
			args: args{
				key:  messages.NoUserFound,
				lang: translator.GetLanguage("en"),
			},
			want: "no user found",
		},
		{
			name: "key does not exists",
			args: args{
				key:  "DoesNotExists",
				lang: translator.GetLanguage("en"),
			},
			want: "DoesNotExists",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(tt *testing.T) {
			if out := tr.Translate(test.args.key, test.args.lang); out != test.want {
				tt.Errorf("Translate() want: %s, got: %s", test.want, out)
			}
		})
	}

}
