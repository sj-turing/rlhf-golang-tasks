	func enrichData(items []Item) []EnrichedItem {
		enrichedItems := make([]EnrichedItem, len(items))
		itemCh := make(chan Item)
		enrichedCh := make(chan EnrichedItem)

		go func() {
			for _, item := range items {
				itemCh <- item
			}
			close(itemCh)
		}()

		var wg sync.WaitGroup
		wg.Add(numWorkers)

		for i := 0; i < numWorkers; i++ {
			go func() {
				for item := range itemCh {
					enrichedItem := enrichItem(item)
					enrichedCh <- enrichedItem
				}
				wg.Done()
			}()
		}

		go func() {
			for enrichedItem := range enrichedCh {
				enrichedItems[enrichedItem.Index] = enrichedItem
			}
		}()

		wg.Wait()
		close(enrichedCh)
		return enrichedItems
	}
