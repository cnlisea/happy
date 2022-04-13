package vote

import (
	"github.com/cnlisea/happy/delay"
	"testing"
	"time"
)

func TestVote_FullAgree(t *testing.T) {
	v := New(2, 2)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})

	v.Add(1, true)
	t.Log("vote 1")
	v.Add(2, true)
	t.Log("vote 2")
}

func TestVote_OneFail(t *testing.T) {
	v := New(2, 2)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})

	v.Add(1, false)
	t.Log("vote 1")
}

func TestVote_ManyAgree(t *testing.T) {
	v := New(2, 3)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})

	v.Add(1, true)
	t.Log("vote 1")
	v.Add(2, true)
	t.Log("vote 2")
	v.Add(3, true)
	t.Log("vote 3")
}

func TestVote_FullEnd(t *testing.T) {
	v := New(2, 3)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})
	v.FullEnd()

	v.Add(1, true)
	t.Log("vote 1")
	v.Add(2, true)
	t.Log("vote 2")
	v.Add(3, false)
	t.Log("vote 3")
}

func TestVote_Deadline(t *testing.T) {
	v := New(2, 3)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})
	delayInstance := delay.New()
	v.Deadline(delayInstance, 3*time.Second, false, false)
	time.Sleep(3 * time.Second)
	select {
	case <-delayInstance.Done():
		delayInstance.Handler()
	default:
	}

	v.Add(1, true)
	t.Log("vote 1")
	v.Add(2, true)
	t.Log("vote 2")
	v.Add(3, false)
	t.Log("vote 3")
}

func TestVote_DeadlinePass(t *testing.T) {
	v := New(2, 3)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})
	delayInstance := delay.New()
	v.Deadline(delayInstance, time.Second, true, false)
	time.Sleep(time.Second)
	select {
	case <-delayInstance.Done():
		delayInstance.Handler()
	default:
	}

	v.Add(1, true)
	t.Log("vote 1")
	v.Add(2, true)
	t.Log("vote 2")
	v.Add(3, false)
	t.Log("vote 3")
}

func TestVote_DeadlineFirst(t *testing.T) {
	v := New(2, 3)
	v.CallbackPass(func() {
		t.Log("vote pass")
	})
	v.CallbackFail(func() {
		t.Log("vote fail")
	})
	delayInstance := delay.New()
	v.Deadline(delayInstance, time.Second, false, true)

	v.Add(1, true)
	t.Log("vote 1")
	time.Sleep(time.Second)
	select {
	case <-delayInstance.Done():
		delayInstance.Handler()
	default:
	}
	v.Add(2, true)
	t.Log("vote 2")
	v.Add(3, false)
	t.Log("vote 3")
}
