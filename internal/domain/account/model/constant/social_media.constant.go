package constant

import "encoding/json"

type SocialMedia string

const (
	Linkedin SocialMedia = "linkedin"
	Github   SocialMedia = "github"
)

func (socialMedia SocialMedia) MarshalJSON() ([]byte, error) {
	return json.Marshal(socialMedia)
}
