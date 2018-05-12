package meld

import (
	"time"

	"github.com/go-kit/kit/log"
)

type loggingService struct {
	logger log.Logger
	Service
}

func NewLoggingService(logger log.Logger, s Service) Service {
	return &loggingService{logger, s}
}

func (s *loggingService) CreateAWSEC2() (id string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "create_aws_ec2",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.CreateAWSEC2()
}

func (s *loggingService) ReadAWSEC2() (id string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "read_aws_ec2",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.ReadAWSEC2()
}

func (s *loggingService) UpdateAWSEC2() (id string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "update_aws_ec2",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.UpdateAWSEC2()
}

func (s *loggingService) DeleteAWSEC2() (id string, err error) {
	defer func(begin time.Time) {
		s.logger.Log(
			"method", "delete_aws_ec2",
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.DeleteAWSEC2()
}
