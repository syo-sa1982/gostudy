package main
import "testing"


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
	var gravatarAvatar GravataAvatar
	client := new(client)
	client.userData =
	map[string]interface{}{"email": "fujii1308@gmail.com"}
	url, err := gravatarAvatar.GetAvatarURL(client)
	if err != ErrNoAvatarURL {
		t.Error("GravataAvatar.GetAvatarURLはエラーを返すべきでない")
	}
	if url != "//www.gravatar.com/avatar/abc" {
		t.Errorf("GravatarAvitar.GetAvatarURL wrongly returned %s", url)
	}
}

