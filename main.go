package main

func main() {
	c.Cyan("connecting to twitter api and getting the last tweets")
	client := readConfigTokensAndConnect()
	tweets := getTimeline(client)
	for i := 0; i < len(tweets); i++ {
		c.Cyan(tweets[i].User.ScreenName + ": " + tweets[i].Text)
		c.Green(tweets[i].Lang)
		content := removeLinks(tweets[i].Text)
		if tweets[i].Lang == "en" {
			festivalEn(tweets[i].User.ScreenName + ": " + content)
		} else {
			festivalCa(tweets[i].User.ScreenName + ": " + content)
		}
	}
}
