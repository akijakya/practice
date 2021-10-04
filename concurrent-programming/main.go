package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	// we query the cache first, if there are no results,
	// we query the database and put the result in the cache

	wg := &sync.WaitGroup{} // init waitgroup
	m := &sync.RWMutex{} // init mutex
	cacheCh := make(chan Book) // create channel for cache
	dbCh := make(chan Book) // create channel for database

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2) // how many tasks are needed to be done
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryCache(id, m); ok {
				ch <- b
			}
			wg.Done() // task 1 is done
		}(id, wg, m, cacheCh)
		go func(id int, wg *sync.WaitGroup, m *sync.RWMutex, ch chan<- Book) {
			if b, ok := queryDatabase(id); ok {
				m.Lock()
				cache[id] = b
				m.Unlock()
				ch <- b
			}
			wg.Done() // task 2 is done
		}(id, wg, m, dbCh)

		// create one goroutine per query to handle response
		go func(cacheCh, dbCh <-chan Book) {
			select {
			case b := <-cacheCh:
				fmt.Println("from cache")
				fmt.Println(b)
				<-dbCh // if we get a hit on the cache channel, there will also be a message in the database channel as well because that always generates message, the cache only occasionally generates message. So if the cache gets hit, we need to "drain" one message from the database channel as well.
			case b := <-dbCh:
				fmt.Println("from database")
				fmt.Println(b)
			}
		}(cacheCh, dbCh)
		time.Sleep(150 * time.Millisecond)
	}

	wg.Wait() // wait for all the tasks to be done

	// // channels
	// wg := &sync.WaitGroup{}
	// ch := make (chan int, 1)

	// wg.Add(2)
	// go func(ch <-chan int, wg *sync.WaitGroup) {
	// 	for msg := range ch {
	// 		fmt.Println(msg)
	// 	}
	// 	wg.Done()
	// }(ch, wg)
	// go func(ch chan<- int, wg *sync.WaitGroup) {
	// 	for i := 0; i < 10; i++ {
	// 		ch <- i
	// 	}
	// 	close(ch)
	// 	wg.Done()
	// }(ch, wg)

	// wg.Wait()
}

// simulate querying the cache
func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	m.RLock()
	b, ok := cache[id]
	m.RUnlock()
	return b, ok
}

// simulate querying a database and put result in cache
func queryDatabase(id int) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			cache[id] = b
			return b, true
		}
	}

	return Book{}, false
}

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title:\t\t%q\n"+
			"Author:\t\t%q\n"+
			"Published:\t%v\n", b.Title, b.Author, b.YearPublished)
}

var books = []Book{
	{
		ID:            1,
		Title:         "Cím 1",
		Author:        "Szerző 1",
		YearPublished: 1999,
	},
	{
		ID:            2,
		Title:         "Cím 2",
		Author:        "Szerző 2",
		YearPublished: 2000,
	},
	{
		ID:            3,
		Title:         "Cím 3",
		Author:        "Szerző 3",
		YearPublished: 2001,
	},
	{
		ID:            4,
		Title:         "Cím 4",
		Author:        "Szerző 4",
		YearPublished: 2002,
	},
	{
		ID:            5,
		Title:         "Cím 5",
		Author:        "Szerző 5",
		YearPublished: 2003,
	},
	{
		ID:            6,
		Title:         "Cím 6",
		Author:        "Szerző 6",
		YearPublished: 2004,
	},
	{
		ID:            7,
		Title:         "Cím 7",
		Author:        "Szerző 7",
		YearPublished: 2005,
	},
	{
		ID:            8,
		Title:         "Cím 8",
		Author:        "Szerző 8",
		YearPublished: 2006,
	},
	{
		ID:            9,
		Title:         "Cím 9",
		Author:        "Szerző 9",
		YearPublished: 2007,
	},
	{
		ID:            10,
		Title:         "Cím 10",
		Author:        "Szerző 10",
		YearPublished: 2008,
	},
}
