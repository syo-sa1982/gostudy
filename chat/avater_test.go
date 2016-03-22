package main
import (
	"testing"
	"path/filepath"
	"io/ioutil"
	"os"
)


func TestAuthAvatar(t *testing.T) {
	var authAvatar AuthAvatar
	client := new(client)
	url, err := authAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("値がないなら、AuthAvatar.GetAvatarURLはErrNoAvatarURLを返すべき")
	}
	// 値をSET
	testUrl := "http://url-to-avatar/"
	client.userData = map[string]interface{}{"avatar_url": testUrl}
	url, err = authAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("値があるなら、AuthAvatar.GetAvatarURLはエラーを返すべきでない")
	} else {
		if url != testUrl {
			t.Error("AuthAvatar.GetAvatarURLは正しいURLを返すべき")
		}
	}
}

func TestGravatarAvatar(t *testing.T) {
	var gravatarAvatar GravatarAvatar
	client := new(client)
	client.userData = map[string]interface{}{
		"userid": "6e2c7867e96ea21953e6be4c6e1ee507",
	}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error(err)
		t.Error(url)
		t.Error("GravataAvatar.GetAvatarURLはエラーを返すべきでない")
	}
	if url != "//www.gravatar.com/avatar/6e2c7867e96ea21953e6be4c6e1ee507" {
		t.Errorf("GravatarAvitar.GetAvatarURL wrongly returned %s", url)
	}
}

func TestFileSystemAvatar(t *testing.T) {
	// テスト用アバタファイル
	filename := filepath.Join("avatars", "abc.jpg")
	ioutil.WriteFile(filename, []byte{}, 0777)
	defer func() { os.Remove(filename) }()

	var fileSystemAvatar FileSystemAvatar
	client := new(client)
	client.userData = map[string]interface{}{"userid": "abc"}
	url, err := fileSystemAvatar.GetAvatarURL(client)
	if err != nil {
		t.Error("FileSystemAvatar.GetAvaterURLはエラーを返すべきでない")
	}
	if url != "/avaters/abc.jpg" {
		t.Errorf("FileSystemAvatar.GetAvaterURLが%sという誤った値", url)
	}
}

