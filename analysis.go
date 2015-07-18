package main

import "fmt"

type analysis struct {
	nrepo int

	mostStarred string
	mostForked  string
	mostWatched string

	highestStars   int
	highestForks   int
	highestWatches int

	avStarsPerRepo   float64
	avForksPerRepo   float64
	avWatchesPerRepo float64

	totalStars   int
	totalForks   int
	totalWatches int

	totalCommits     int // TODO
	avCommitsPerRepo int
	forkedRepos      int
}

func (an *analysis) display() {
	fmt.Println("Number of repositories :", an.nrepo)
	fmt.Println("Most starred repository :", an.mostStarred, "with", an.highestStars, "stars.")
	fmt.Println("Most forked repository :", an.mostForked, "with", an.highestForks, "forks.")
	fmt.Println("Most watched repository :", an.mostWatched, "with", an.highestWatches, "watchers.")
	fmt.Println()
	fmt.Println("Average stars per repository :", an.avStarsPerRepo)
	fmt.Println("Average forks per repository :", an.avForksPerRepo)
	fmt.Println("Average watchers per repository :", an.avWatchesPerRepo)
	fmt.Println()
	fmt.Println("Total stars :", an.totalStars)
	fmt.Println("Total forks :", an.totalForks)
	fmt.Println("Total watchers :", an.totalWatches)
}

func (an *analysis) analyseRepos(usr string) (err error) {
	var rd allRepos
	err = fetchURL("https://api.github.com/users/"+usr+"/repos", &rd)
	if err != nil {
		return
	}
	an.nrepo = len(rd)
	for _, r := range rd {
		if r.Fork {
			an.forkedRepos++
		} else {
			an.totalStars += r.StargazersCount
			if r.StargazersCount > an.highestStars {
				an.highestStars = r.StargazersCount
				an.mostStarred = r.Name
			}
			an.totalWatches += r.WatchersCount
			if r.WatchersCount > an.highestWatches {
				an.highestWatches = r.WatchersCount
				an.mostWatched = r.Name
			}
			an.totalForks += r.ForksCount
			if r.ForksCount > an.highestForks {
				an.highestForks = r.ForksCount
				an.mostForked = r.Name
			}
		}
	}
	an.avStarsPerRepo = float64(an.totalStars) / float64(an.nrepo)
	an.avForksPerRepo = float64(an.totalForks) / float64(an.nrepo)
	an.avWatchesPerRepo = float64(an.totalWatches) / float64(an.nrepo)
	return
}
