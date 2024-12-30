	var cache = map[string]EnrichedItem{}

	func enrichData(items []Item) []EnrichedItem {
		enrichedItems := make([]EnrichedItem, len(items))
		for i, item := range items {
			key := item.Key
			if enrichedItem, ok := cache[key]; ok {
				enrichedItems[i] = enrichedItem
			} else {
				enrichedItem = enrichItem(item)
				cache[key] = enrichedItem
				enrichedItems[i] = enrichedItem
			}
		}
		return enrichedItems
	}
