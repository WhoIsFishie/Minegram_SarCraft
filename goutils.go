package main

func removePlayer(s []onlinePlayer, r string) []onlinePlayer {
	for i, v := range s {
		if v.inGameName == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func containsPlayer(s []onlinePlayer, e string) bool {
	for _, a := range s {
		if a.inGameName == e {
			return true
		}
	}
	return false
}

func getOnlinePlayer(ign string) onlinePlayer {
	for _, player := range online {
		if player.inGameName == ign {
			return player
		}
	}
	return onlinePlayer{}
}