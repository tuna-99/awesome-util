package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Profile represents the data we want.
type Profile struct {
	ID   int
	Name string
}

// Result holds either a profile or an error for a specific ID.
type Result struct {
	ID      int
	Profile *Profile
	Err     error
}

// MockFetch simulates a slow and flaky API.
// It takes random time (up to 1s) and errors 20% of the time.
func MockFetch(ctx context.Context, id int) (*Profile, error) {
	// Simulate latency
	select {
	case <-time.After(time.Duration(rand.Intn(1000)) * time.Millisecond):
	case <-ctx.Done():
		return nil, ctx.Err()
	}

	// Simulate random error
	if rand.Float32() < 0.2 {
		return nil, errors.New("simulated 500 error")
	}

	return &Profile{ID: id, Name: fmt.Sprintf("User-%d", id)}, nil
}

// ---------------------------------------------------------
// YOUR TASK: Implement this function.
// ---------------------------------------------------------

func worker(ctx context.Context, jobs <-chan int, results chan<- Result, wg *sync.WaitGroup) {
	// 1. Fetch the profile for the given ID.
	// 2. Send the result to the 'results' channel.
	defer wg.Done()
	for id := range jobs {
		if ctx.Err() != nil {
			return
		}

		profile, err := MockFetch(ctx, id)
		results <- Result{ID: id, Profile: profile, Err: err}
	}
}

func FetchAllProfiles(ids []int, maxConcurrency int, timeout time.Duration) ([]*Profile, []error) {
	// TODO:
	// 1. Process all IDs concurrently, but max 'maxConcurrency' at a time.
	// 2. Enforce the global 'timeout'.
	// 3. Aggregate results.
	profiles := make([]*Profile, 0, len(ids))
	errors := make([]error, 0, len(ids))
	results := make(chan Result, len(ids))
	jobs := make(chan int, len(ids))
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	// 1. Start workers
	var wg sync.WaitGroup
	for range maxConcurrency {
		wg.Add(1)
		go worker(ctx, jobs, results, &wg)
	}

	// 2. Send jobs
	for _, id := range ids {
		jobs <- id
	}
	close(jobs)

	// 3. Monitor for completion in a separate goroutine
	// This allows us to close the 'results' channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	for {
		select {
		case <-ctx.Done():
			return profiles, errors
		case result, ok := <-results:
			if !ok {
				return profiles, errors
			}

			if result.Err != nil {
				errors = append(errors, result.Err)
			} else {
				profiles = append(profiles, result.Profile)
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	ids := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println("Starting fetch...")
	start := time.Now()

	profiles, errs := FetchAllProfiles(ids, 3, 2*time.Second)

	duration := time.Since(start)

	fmt.Printf("\n--- Report ---\n")
	fmt.Printf("Time Taken: %v\n", duration)
	fmt.Printf("Successful Profiles: %d\n", len(profiles))
	fmt.Printf("Failed IDs: %d\n", len(errs))

	for _, p := range profiles {
		fmt.Printf("[OK] %s\n", p.Name)
	}
	for _, err := range errs {
		fmt.Printf("[ERR] %v\n", err)
	}
}
