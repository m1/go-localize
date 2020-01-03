package main

import (
	"reflect"
	"testing"
)

func Test_run(t *testing.T) {
	type args struct {
		in  *string
		out *string
	}

	dirBlank := ""
	dirValid := "examples/localizations_src"
	dirTestFiles := "test_files"
	dirWithBad := "mock"
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				in:  &dirValid,
				out: &dirTestFiles,
			},
		},
		{
			name: "not valid",
			args: args{
				in:  &dirBlank,
				out: &dirBlank,
			},
			wantErr: true,
		},
		{
			name: "not valid",
			args: args{
				in:  &dirWithBad,
				out: &dirTestFiles,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := run(tt.args.in, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_generateLocalizations(t *testing.T) {
	type args struct {
		files []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				files: []string{
					"mock/dir/sub/valid_json.json",
					"mock/dir/valid_json.json",
					"mock/dir/valid_yaml.yaml",
					"mock/dir/valid_csv.csv",
					"mock/dir/valid_toml.toml",
					"mock/dir/dont_parse.txt",
				},
			},
			want: map[string]string{
				"mock.dir.sub.valid_json.test": "test",
				"mock.dir.valid_json.test":     "test",
				"mock.dir.valid_yaml.test":     "test",
				"mock.dir.valid_csv.test":      "test",
				"mock.dir.valid_toml.test":     "test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := generateLocalizations(tt.args.files)
			if (err != nil) != tt.wantErr {
				t.Errorf("generateLocalizations() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateLocalizations() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_generateFile(t *testing.T) {
	type args struct {
		output       string
		translations map[string]string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				output:       "test_files",
				translations: map[string]string{"hello": "one"},
			},
		},
		{
			name: "invalid dir",
			args: args{
				output:       "",
				translations: map[string]string{"hello": "one"},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := generateFile(tt.args.output, tt.args.translations); (err != nil) != tt.wantErr {
				t.Errorf("generateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getLocalizationsFromFile(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]string
		wantErr bool
	}{
		{
			name: "valid json",
			args: args{"mock/valid.json"},
			want: map[string]string{"mock.valid.test1": "test2"},
		},
		{
			name: "valid yaml",
			args: args{"mock/valid.yaml"},
			want: map[string]string{"mock.valid.test1": "test2"},
		},
		{
			name:    "file not exist",
			args:    args{"mock/non_exist.json"},
			wantErr: true,
		},
		{
			name:    "invalid json",
			args:    args{"mock/invalid.json"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocalizationsFromFile(tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalizationsFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocalizationsFromFile() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSlicePath(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "valid",
			args: args{"mock/valid.json"},
			want: []string{"mock", "valid"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSlicePath(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getSlicePath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseFlags(t *testing.T) {
	type args struct {
		input  *string
		output *string
	}

	dirBlank := ""
	dirOk := "input"

	tests := []struct {
		name      string
		args      args
		inputDir  string
		outputDir string
		wantErr   error
	}{
		{
			name: "valid",
			args: args{
				input:  &dirOk,
				output: &dirOk,
			},
			inputDir:  dirOk,
			outputDir: dirOk,
		},
		{
			name: "default output dir",
			args: args{
				input:  &dirOk,
				output: &dirBlank,
			},
			inputDir:  dirOk,
			outputDir: defaultOutputDir,
		},
		{
			name: "invalid input",
			args: args{
				input:  &dirBlank,
				output: &dirBlank,
			},
			wantErr: errFlagInputNotSet,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inputDir, outputDir, err := parseFlags(tt.args.input, tt.args.output)
			if (err != nil) != (tt.wantErr != nil) || err != tt.wantErr {
				t.Errorf("parseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if inputDir != tt.inputDir {
				t.Errorf("parseFlags() got = %v, want %v", inputDir, tt.inputDir)
			}
			if outputDir != tt.outputDir {
				t.Errorf("parseFlags() got1 = %v, want %v", outputDir, tt.outputDir)
			}
		})
	}
}

func Test_getLocalizationFiles(t *testing.T) {
	type args struct {
		dir string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "valid",
			args: args{"mock/dir"},
			want: []string{
				"mock/dir/sub/valid_json.json",
				"mock/dir/valid_json.json",
				"mock/dir/valid_yaml.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getLocalizationFiles(tt.args.dir)
			if (err != nil) != tt.wantErr {
				t.Errorf("getLocalizationFiles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getLocalizationFiles() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseCSV(t *testing.T) {
	type args struct {
		value []byte
		l     *localizationFile
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		want    *localizationFile
	}{
		{
			name: "valid",
			args: args{
				value: []byte("test,test"),
				l:     &localizationFile{},
			},
			want: &localizationFile{
				"test": "test",
			},
		},
		{
			name: "not valid",
			args: args{
				value: []byte("test,test\ntest,test,test"),
				l:     &localizationFile{},
			},
			wantErr: true,
		},
		{
			name: "record length above 2",
			args: args{
				value: []byte("test,test,test"),
				l:     &localizationFile{},
			},
			want: &localizationFile{"test": "test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := parseCSV(tt.args.value, tt.args.l)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseCSV() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.args.l, tt.want) && !tt.wantErr {
				t.Errorf("parseCSV() got = %v, want %v", tt.args.l, tt.want)
			}
		})
	}
}
