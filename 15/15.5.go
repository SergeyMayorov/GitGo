package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type tMeteo struct {
	t string
	m sync.RWMutex
}

func (m *tMeteo) ReadTemp() string {
	defer m.m.RUnlock()
	m.m.RLock()
	return m.t
}

func (m *tMeteo) ChangeTemp(v string) {
	m.m.Lock()
	m.t = v
	m.m.Unlock()
}

func main() {
	wg := sync.WaitGroup{}
	temp := &tMeteo{t: "0"}
	for i := 0; i < 5; i++ {
		go func() {
			wg.Add(1)
			fmt.Println("Текущая температура: ", temp.ReadTemp())
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			wg.Done()
		}()

		for j := 0; j < 5; j++ {
			go func() {
				wg.Add(1)
				tt := rand.Intn(40) - 20
				temp.ChangeTemp(strconv.Itoa(tt))
				time.Sleep(time.Duration(rand.Intn(100)+50) * time.Millisecond)
				wg.Done()
			}()
		}
	}

	wg.Wait()
}
