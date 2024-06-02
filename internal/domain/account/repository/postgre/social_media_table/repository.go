package social_media_table

func NewSocialMediaTableRepository(config SocialMediaTableRepositoryConfig) SocialMediaTableRepositoryProvider {
	return &SocialmediaTableRepository{}
}
