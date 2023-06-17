package main

func main() {
	// Initialize env vars
	Init()

	// Initialize and pull the git repo
	var wg WikiGit
	wg.Init()
	wg.LoopFetchAndPull()
}
