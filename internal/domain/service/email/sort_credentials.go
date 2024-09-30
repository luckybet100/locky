package email

import "github.com/Locky-Inc/locky/internal/domain"

func (s *Service) SortCredentials() []*domain.EmailProviderCredentials {
	/*
	пока просто возвращаем как есть, в будущем можно будет сюда будет передать несколько провайдеров
	и напилить логику по тому чтобы давать на них равномерное число трафика / детектить отказы и отключать
	*/
	return s.creds 
}
