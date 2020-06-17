package libbench

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/dvasilas/proteus/pkg/perf"
)

type workload struct {
	config       *benchmarkConfig
	ops          *operations
	measurements *perf.Perf
}

func newWorkload(conf *benchmarkConfig) (*workload, error) {
	ops, err := newOperations(conf)
	if err != nil {
		return nil, err
	}

	return &workload{
		config:       conf,
		ops:          ops,
		measurements: perf.New(),
	}, nil
}

func (w workload) run(measurementBufferSize int64) (map[string][]time.Duration, map[string]int64, time.Time, time.Time) {
	durations := make(map[string][]time.Duration, measurementBufferSize)
	durations["getHomepage"] = make([]time.Duration, measurementBufferSize)
	durations["vote"] = make([]time.Duration, measurementBufferSize)

	perOpCnt := make(map[string]int64)
	perOpCnt["getHomepage"] = 0
	perOpCnt["vote"] = 0
	var opCnt int64
	var st time.Time
	var respTime time.Duration
	var err error
	timerStarted := false
	warmingUp, warmupTimeout := w.config.Benchmark.DoWarmup, time.After(time.Duration(w.config.Benchmark.Warmup)*time.Second)
	for timeIsUp, timeout := true, time.After(time.Duration(w.config.Benchmark.Runtime)*time.Second); timeIsUp; {

		select {
		case <-timeout:
			timeIsUp = false
		case <-warmupTimeout:
			warmingUp = false
		default:
		}

		if !timerStarted && !warmingUp {
			timerStarted = true
			st = time.Now()
		}
		if opCnt == w.config.Benchmark.OpCount {
			break
		}

		r := rand.Float64()
		if r < w.config.Operations.WriteRatio {
			vote := rand.Float64()
			if vote < w.config.Operations.DownVoteRatio {
				respTime, err = w.downVoteStory(0)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				respTime, err = w.upVoteStory(0)
				if err != nil {
					log.Fatal(err)
				}
			}
			if !warmingUp {
				durations["vote"][perOpCnt["vote"]] = respTime
				perOpCnt["vote"]++
			}
		} else {
			respTime, err = w.getHomepage()
			if err != nil {
				log.Fatal(err)
			}
			if !warmingUp {
				durations["getHomepage"][perOpCnt["getHomepage"]] = respTime
				perOpCnt["getHomepage"]++
			}
		}
		opCnt++
	}
	return durations, perOpCnt, st, time.Now()
}

// Preload ...
func (w workload) preload() error {
	// start from 1 because when MySQL automaticall assigns ids
	// it starts from 1
	// ¯\_(ツ)_/¯
	for i := 1; i <= w.config.Preload.RecordCount.Users; i++ {
		if err := w.addUser(); err != nil {
			return err
		}
	}

	for i := 1; i <= w.config.Preload.RecordCount.Stories; i++ {
		if err := w.addStory(); err != nil {
			return err
		}
		if _, err := w.upVoteStory(i); err != nil {
			return err
		}
	}

	for i := 1; i <= w.config.Preload.RecordCount.Comments; i++ {
		if err := w.addComment(); err != nil {
			return err
		}
		if _, err := w.upVoteComment(i); err != nil {
			return err
		}
	}

	for i := 1; i <= w.config.Preload.RecordCount.StoryVotes; i++ {
		vote := rand.Float64()
		if vote < w.config.Operations.DownVoteRatio {
			if _, err := w.downVoteStory(0); err != nil {
				return err
			}
		} else {
			if _, err := w.upVoteStory(0); err != nil {
				return err
			}
		}
	}

	return nil
}

func (w *workload) getHomepage() (time.Duration, error) {
	st := time.Now()

	_, err := w.ops.getHomepage()

	return time.Since(st), err
}

func (w *workload) addUser() error {
	err := w.ops.addUser()

	w.ops.state.addUser()

	return err
}

func (w *workload) addStory() error {
	userID := w.ops.selectUser()

	err := w.ops.addStory(userID)

	w.ops.state.addStory()

	return err
}

func (w *workload) addComment() error {
	userID := w.ops.selectUser()
	storyID := w.ops.selectStory()

	err := w.ops.addComment(userID, storyID)

	w.ops.state.addComment()

	return err
}

func (w *workload) upVoteStory(storyID int) (time.Duration, error) {
	st := time.Now()

	userID := w.ops.selectUser()
	if storyID == 0 {
		storyID = w.ops.selectStory()
	}

	err := w.ops.upVoteStory(userID, storyID)

	return time.Since(st), err
}

func (w *workload) downVoteStory(storyID int) (time.Duration, error) {
	st := time.Now()

	userID := w.ops.selectUser()
	if storyID == 0 {
		storyID = w.ops.selectStory()
	}

	err := w.ops.downVoteStory(userID, storyID)

	return time.Since(st), err
}

func (w *workload) upVoteComment(commentID int) (time.Duration, error) {
	st := time.Now()

	userID := w.ops.selectUser()
	if commentID == 0 {
		commentID = w.ops.selectStory()
	}

	err := w.ops.upVoteComment(userID, commentID)

	return time.Since(st), err
}

func (w *workload) downVoteComment() (time.Duration, error) {
	st := time.Now()

	userID := w.ops.selectUser()
	commentID := w.ops.selectStory()

	err := w.ops.downVoteComment(userID, commentID)

	return time.Since(st), err
}

func printHomepage(hp homepage) error {
	fmt.Println("----------------------")
	for _, story := range hp.stories {
		fmt.Printf("%s | (%s) \n %s \n %d \n", story.title, story.shortID, story.description, story.voteCount)
	}
	fmt.Println("----------------------")

	return nil
}