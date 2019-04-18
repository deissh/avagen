package avatar

import (
	"bytes"
	"image/color"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func Test_getColorByName(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want *color.RGBA
	}{
		{
			name: "should return default color",
			args: args{
				"test",
			},
			want: avatarBgColors["45BDF3"],
		},
		{
			name: "should return color",
			args: args{
				"",
			},
			want: avatarBgColors["7986CB"],
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getColorByName(tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getColorByName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getInitials(t *testing.T) {
	type args struct {
		name  string
		count int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "one symbol",
			args: args{
				name:  "deissh",
				count: 1,
			},
			want: "D",
		},
		{
			name: "one unicode symbol",
			args: args{
				name:  "что-то",
				count: 1,
			},
			want: "Ч",
		},
		{
			name: "two symbols from sentence",
			args: args{
				name:  "some sentence",
				count: 2,
			},
			want: "SS",
		},
		{
			name: "two symbols from unicode sentence",
			args: args{
				name:  "какое то",
				count: 2,
			},
			want: "КТ",
		},
		{
			name: "should return empty",
			args: args{
				name:  "",
				count: 2,
			},
			want: "",
		},
		{
			name: "should return empty",
			args: args{
				name:  "",
				count: 1,
			},
			want: "",
		},
		{
			name: "should return from email",
			args: args{
				name:  "joe@example.com",
				count: 1,
			},
			want: "J",
		},
		{
			name: "should return from email",
			args: args{
				name:  "joe@example.com",
				count: 2,
			},
			want: "J",
		},
		{
			name: "should return from name with ( & )",
			args: args{
				name:  "John Doe (dj)",
				count: 1,
			},
			want: "J",
		},
		{
			name: "should return from name with ( & )",
			args: args{
				name:  "John Doe (dj)",
				count: 2,
			},
			want: "JD",
		},
		{
			name: "should return from name with ( & )",
			args: args{
				name:  "Doe John (dj)",
				count: 3,
			},
			want: "dj",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getInitials(tt.args.name, tt.args.count); got != tt.want {
				t.Errorf("getInitials() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInitialsAvatar_DrawToBytes(t *testing.T) {
	av := New("../../Cousine-Bold.ttf")

	stuffs := []struct {
		name     string
		size     int
		count    int
		encoding string
	}{
		{"Sword", 128, 2, "png"},
		{"Condor Heroes", 30, 2, "png"},
		{"孔子", 128, 2, "png"},
		{"Art", 48, 2, "png"},
		{"*", 64, 2, "png"},
		{"хахаха", 64, 2, "png"},

		{"Sword", 128, 2, "jpeg"},
		{"Condor Heroes", 30, 2, "jpeg"},
		{"孔子", 128, 2, "jpeg"},
		{"Art", 48, 2, "jpeg"},
		{"*", 64, 2, "jpeg"},
		{"хахаха", 64, 2, "jpeg"},
	}

	for _, v := range stuffs {
		raw, err := av.DrawToBytes(v.name, v.size, v.count, v.encoding)
		if err != nil {
			if err == ErrUnsupportedChar {
				t.Skip("ErrUnsupportChar")
			}
			t.Error(err)
		}
		switch v.encoding {
		case "png":
			if _, perr := png.Decode(bytes.NewReader(raw)); perr != nil {
				t.Error(perr)
			}
		case "jpeg":
			if _, perr := jpeg.Decode(bytes.NewReader(raw)); perr != nil {
				t.Error(perr, v)
			}
		}
	}
}

func TestParseFont(t *testing.T) {
	fileNotExists := "xxxxxxx.ttf"
	_, err := parseFont(fileNotExists)
	if err == nil {
		t.Error("should return error")
	}

	_, err = newDrawer(fileNotExists, 75.0)
	if err == nil {
		t.Error("should return error")
	}

	fileExistsButNotTTF, _ := ioutil.TempFile(os.TempDir(), "prefix")
	defer os.Remove(fileExistsButNotTTF.Name())

	_, err = parseFont(fileExistsButNotTTF.Name())
	if err == nil {
		t.Error("should return error")
	}
	_, err = newDrawer(fileExistsButNotTTF.Name(), 75.0)
	if err == nil {
		t.Error("should return error")
	}
	_, err = newDrawer("", 75.0)
	if err == nil {
		t.Error("should return error")
	}

}
