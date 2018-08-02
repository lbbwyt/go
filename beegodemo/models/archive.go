// archive
package models

import (
	"container/list"
)

type EventType int

const (
	EVENT_JOIN = iota
	EVENT_LEAVE
	EVENT_MESSAGE
)

type Event struct {
	Type      EventType //加入，离开，发送消息
	User      string
	Timestamp int // Unix timestamp (secs)
	Content   string
}

//用来保存服务器上能够保存的消息记录，保存最新的20条
const archiveSize = 20

// 事件归档保存
var archive = list.New()

//保存新的事件到容器中，，若事件的个数已经大于等于20则删除第一个，只保留最新的20个
func NewArchive(event Event) {
	if archive.Len() >= archiveSize {
		archive.Remove(archive.Front())
	}

	archive.PushBack(event)
}

//根据传过来的时间戳返回该时间戳之后的所有事件消息
func GetEvents(lastReceived int) []Event {
	events := make([]Event, 0, archive.Len())
	for event := archive.Front(); event != nil; event = event.Next() {
		e := event.Value.(Event)
		if e.Timestamp > int(lastReceived) {
			events = append(events, e)
		}
	}
	return events
}
