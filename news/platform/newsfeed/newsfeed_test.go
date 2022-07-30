package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	// feed.Add(Item{"an item","with body"})

	if len(feed.Items) != 1 {
		t.Errorf("Item was not added 1")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{})
	result := feed.GetAll()

	if len(result) != 1 {
		t.Errorf("Item was not added 2")
	}
}
