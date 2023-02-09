// Package services
/*
@Coding : utf-8
@Time : 2023/2/9 22:25
@Author : yizhigopher
@Software : GoLand
*/
package services

import (
	"example/models"
	"log"
)

type h interface {
	DoA()
}
type Handler struct {
}

func (h *Handler) DoA() {
	log.Println("do A")
}

type Service struct {
	m MailService
	u UserRepo

	h h
}

func NewService(m MailService, u UserRepo) *Service {
	return &Service{
		m: m,
		u: u,
		h: &Handler{},
	}
}

func (s *Service) UserSignUp() {
	user := models.UserModel{
		Name: "shy hao",
		Age:  18,
	}
	s.u.AddUser(user)
	s.m.Send()

	s.h.DoA()
}
