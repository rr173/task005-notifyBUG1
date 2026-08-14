package notify

import (
	"testing"
	"time"
)

func TestCloneSentAtPointerIndependence(t *testing.T) {
	s := New()
	now := time.Date(2024, 6, 1, 12, 0, 0, 0, time.UTC)
	s.Create(CreateInput{ID: "T1", Recipient: "u", Content: "c"}, now)

	sentTime := now.Add(time.Hour)
	returned, err := s.MarkSent("T1", sentTime)
	if err != nil {
		t.Fatal(err)
	}

	// 修改返回值的 SentAt 指针指向的值
	*returned.SentAt = time.Date(1999, 1, 1, 0, 0, 0, 0, time.UTC)

	// store 内部不应被影响
	got, _ := s.Get("T1")
	if !got.SentAt.Equal(sentTime) {
		t.Errorf("clone shares SentAt pointer with store: got %v, want %v", got.SentAt, sentTime)
	}
}
