package main

func main() {
	c.Cyan("hi")
	client := readConfigTokensAndConnect()
	tweets := getTimeline(client)
	for i := 0; i < len(tweets); i++ {
		c.Cyan(tweets[i].User.ScreenName + ": " + tweets[i].Text)
		c.Green(tweets[i].Lang)
		if tweets[i].Lang == "en" {
			festivalEn(tweets[i].User.ScreenName + ": " + tweets[i].Text)
		} else {
			festivalCa(tweets[i].User.ScreenName + ": " + tweets[i].Text)
		}
	}
}
