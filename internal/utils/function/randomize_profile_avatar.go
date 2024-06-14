package function

import (
	"math/rand"
	"time"
)

func RandomizeProfileAvatar() string {
	avatars := []string{
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296292736.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296296048.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296299084.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296302274.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296304603.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296308421.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296311867.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296314440.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296318676.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296321058.svg",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296323807.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296325789.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296329189.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296332069.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296334812.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296337381.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296340323.svg",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296342587.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296345057.svg?alt=media",
		"https://firebasestorage.googleapis.com/v0/b/account-service-bucket.appspot.com/o/default-avatar%2FadventurerNeutral-1718296347289.svg?alt=media",
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(avatars))
	return avatars[randomIndex]
}
