package core

import (
	"context"
)

type Service interface {
	Register(eventName string, callback TriggerCallback)
	Trigger(eventName string, ctx context.Context) []error
}

type TriggerCallback func(ctx context.Context) error

type ServiceImpl struct {
	Events map[string][]TriggerCallback
}

func NewService() Service {
	return &ServiceImpl{}
}

func (s *ServiceImpl) Register(eventName string, callback TriggerCallback) {
	if _, ok := s.Events[eventName]; !ok {
		s.Events[eventName] = []TriggerCallback{}
	}
	s.Events[eventName] = append(s.Events[eventName], callback)
}

func (s *ServiceImpl) Trigger(eventName string, ctx context.Context) []error {
	errs := []error{}
	if triggers, ok := s.Events[eventName]; ok {
		for i := range triggers {
			err := triggers[i](ctx)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	return errs
}
