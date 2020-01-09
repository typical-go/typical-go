package walker

import (
	log "github.com/sirupsen/logrus"
)

// AnnotationDeclListener listen to declaration event for annotation
type AnnotationDeclListener struct {
	name      string
	eventType EventType
	listener  AnnotationListener
}

// NewAnnotationDeclListener return new instance of AnnotationDeclListener
func NewAnnotationDeclListener(name string, eventType EventType, listener AnnotationListener) *AnnotationDeclListener {
	return &AnnotationDeclListener{
		name:      name,
		eventType: eventType,
		listener:  listener,
	}
}

// AnnotationListener handle the annotation
type AnnotationListener interface {
	OnAnnotation(*AnnotationEvent) error
}

// AnnotationEvent is event for annotation listener
type AnnotationEvent struct {
	*DeclEvent
	*Annotation
}

// OnDecl to handle declaration event
func (a *AnnotationDeclListener) OnDecl(e *DeclEvent) (err error) {
	annotation := e.Annotations.Get(a.name)
	if annotation != nil {
		if e.EventType == a.eventType {
			if err = a.listener.OnAnnotation(&AnnotationEvent{
				Annotation: annotation,
				DeclEvent:  e,
			}); err != nil {
				return
			}
		} else {
			log.Warnf("[%s] has no effect to %s:%s", a.name, e.EventType, a.name)
		}
	}
	return
}
