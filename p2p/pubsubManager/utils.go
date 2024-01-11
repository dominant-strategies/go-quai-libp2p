package pubsubManager

import "strings"

// gets the data type from the topic name
func getTopicType(topic string) string {
	return topic[strings.LastIndex(topic, "/")+1:]
}
