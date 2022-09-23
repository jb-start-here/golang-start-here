package main

import "sync"

func executeAllFuncs(funcs []func(), workerCount int) chan bool {
	done := make(chan bool)

	go func() {
		taskPipeline := make(chan func(), len(funcs))
		for _, task := range funcs {
			taskPipeline <- task
		}
		close(taskPipeline)

		wg := sync.WaitGroup{}

		for i := 0; i < workerCount; i++ {
			wg.Add(1)

			go func() {
				for task := range taskPipeline {
					task()
				}
				wg.Done()
			}()
		}

		wg.Wait()
		done <- true
		close(done)
	}()

	return done
}
