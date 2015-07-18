package main

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

	totalCommits     int
	avCommitsPerRepo int
	forkedRepos      int
}

func analyseRepos(usr string) (an analysis, err error) {
	var rd allRepos
	rd, err = fetchReposData(usr)
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
