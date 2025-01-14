package monitorapi

// GetNodeRoles extract the node roles from the event message.
func GetNodeRoles(event Interval) string {
	return event.StructuredMessage.Annotations[AnnotationRoles]
}
