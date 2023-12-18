package usecase

import (
	"fiber-upload-file/internal/helper"
	"sync"
)

func (u *Usecase) HitAPIsConcurrently() bool {
	var (
		wg sync.WaitGroup

		apiPaths   = []string{"200", "400", "500"}
		resultBool = true
	)

	for _, path := range apiPaths {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res := helper.SendGetRequest("https://http.cat/" + path)
			resultBool = resultBool && res
		}()
	}

	wg.Wait()

	return resultBool
}
