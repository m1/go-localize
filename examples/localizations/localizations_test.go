package localizations

import (
	"reflect"
	"testing"
)

func TestLocalizer_Get(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		key          string
		replacements []*Replacements
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "field",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{
				key:          "messages.hello",
				replacements: nil,
			},
			want: "hello",
		},
		{
			name: "no key",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{
				key:          "messages.hello2",
				replacements: nil,
			},
			want: "messages.hello2",
		},
		{
			name: "valid replacements",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{
				key: "messages.hello_my_name_is",
				replacements: []*Replacements{
					{"name": "test"},
				},
			},
			want: "Hello my name is test",
		},
		{
			name: "valid replacements multiple",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{
				key: "messages.hello_firstname_lastname",
				replacements: []*Replacements{
					{"firstname": "test"},
					{"lastname": "test"},
				},
			},
			want: "Hello test test",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.Get(tt.args.key, tt.args.replacements...); got != tt.want {
				t1.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizer_GetWithLocale(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		locale       string
		key          string
		replacements []*Replacements
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "valid",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{
				locale:       "es",
				key:          "messages.hello",
				replacements: nil,
			},
			want: "Hola",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.GetWithLocale(tt.args.locale, tt.args.key, tt.args.replacements...); got != tt.want {
				t1.Errorf("GetWithLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizer_SetFallbackLocale(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		fallback string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Localizer
	}{
		{
			name: "valid",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{fallback: "ru"},
			want: Localizer{
				Locale:         "en",
				FallbackLocale: "ru",
				Localizations:  localizations,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.SetFallbackLocale(tt.args.fallback); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SetFallbackLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizer_SetLocale(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		locale string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Localizer
	}{
		{
			name: "valid",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{locale: "ru"},
			want: Localizer{
				Locale:         "ru",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.SetLocale(tt.args.locale); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SetLocale() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizer_SetLocales(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		locale   string
		fallback string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Localizer
	}{
		{
			name: "valid",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
			args: args{locale: "ru", fallback: "ru"},
			want: Localizer{
				Locale:         "ru",
				FallbackLocale: "ru",
				Localizations:  localizations,
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.SetLocales(tt.args.locale, tt.args.fallback); !reflect.DeepEqual(got, tt.want) {
				t1.Errorf("SetLocales() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizer_getLocalizationKey(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		locale string
		key    string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "valid",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "en",
				Localizations:  nil,
			},
			args: args{
				locale: "en",
				key:    "test",
			},
			want: "en.test",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.getLocalizationKey(tt.args.locale, tt.args.key); got != tt.want {
				t1.Errorf("getLocalizationKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalizer_replace(t1 *testing.T) {
	type fields struct {
		Locale         string
		FallbackLocale string
		Localizations  map[string]string
	}
	type args struct {
		str          string
		replacements []*Replacements
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name: "valid",
			fields: fields{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  nil,
			},
			args: args{
				str:          "Hello {{.firstname}} {{.lastname}}",
				replacements: []*Replacements{{"firstname": "test", "lastname": "test"}},
			},
			want: "Hello test test",
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			t := Localizer{
				Locale:         tt.fields.Locale,
				FallbackLocale: tt.fields.FallbackLocale,
				Localizations:  tt.fields.Localizations,
			}
			if got := t.replace(tt.args.str, tt.args.replacements...); got != tt.want {
				t1.Errorf("replace() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		locale         string
		fallbackLocale string
	}
	tests := []struct {
		name string
		args args
		want *Localizer
	}{
		{
			name: "valid",
			args: args{
				locale:         "en",
				fallbackLocale: "es",
			},
			want: &Localizer{
				Locale:         "en",
				FallbackLocale: "es",
				Localizations:  localizations,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.locale, tt.args.fallbackLocale); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
