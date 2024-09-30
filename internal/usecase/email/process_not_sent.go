package email

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/google/uuid"

	"github.com/Locky-Inc/locky/internal/domain"
	emailAdapter "github.com/Locky-Inc/locky/internal/infrastructure/adapters/email"
	"github.com/Locky-Inc/locky/internal/infrastructure/storage"
	"github.com/Locky-Inc/locky/internal/pkg/pointer"
)

func (uc *Usecase) ProcessNotSent(ctx context.Context) error {

	var (
		wg      sync.WaitGroup
		nextKey *uuid.UUID
		tasks   = make(chan *domain.Email, uc.BatchSize)
	)

	// запускаем воркеров, которые будут посылать email-ы
	wg.Add(uc.WorkersCount)
	for i := 0; i < uc.WorkersCount; i++ {
		go func() {
			defer wg.Done()
			for email := range tasks {
				// TODO: добавить логирование ошибок
				_ = uc.ProcessEmail(ctx, email)
			}
		}()
	}

	for {
		// получаем батч из неотправленных email-ов
		emails, err := uc.st.ListEmails(ctx, storage.ListEmailsFilter{
			Status:      pointer.Ref(domain.EmailSentStatusPending),
			Limit:       uc.BatchSize,
			LastPageKey: nextKey,
		})
		if err != nil {
			return fmt.Errorf("failed to get list of emails: %w", err)
		}

		// если email-ов больше нет -- выходим из цикла
		if len(emails) == 0 {
			break
		}

		// проходимся по всем и отправляем
		nextKey = pointer.Ref(emails[len(emails)-1].ID)
		for _, email := range emails {
			tasks <- email
		}
	}

	// закрываем очередь и ожидаем пока не завершится отправка всех email-ов
	close(tasks)
	wg.Wait()

	return nil
}

func (uc *Usecase) ProcessEmail(ctx context.Context, email *domain.Email) error {

	// получаем креды в порядке приоритета
	creds := uc.emailSvc.SortCredentials()

	// перебираем креды
	for _, cred := range creds {

		if err := uc.adapter.SendEmail(ctx, cred, email); err != nil {
			switch {
			case errors.Is(err, emailAdapter.ErrInvalidCredentials) || errors.Is(err, emailAdapter.ErrUnsupportedCredentials):
				continue
			default:
				// TODO добавить проверку на временные ошибки и ретраить их
				// TODO добавить логирование ошибок

				// помечаем собщение как отправленное
				if updateErr := uc.st.UpdateEmailStatus(ctx, email.ID, domain.EmailSentStatusFailedToSend); updateErr != nil {
					return fmt.Errorf("failed to mark email as undeliverable: %w", updateErr)
				}
				return nil
			}
		}

		// помечаем собщение как отправленное
		if updateErr := uc.st.UpdateEmailStatus(ctx, email.ID, domain.EmailSentStatusSent); updateErr != nil {
			return fmt.Errorf("failed to mark email as sent: %w", updateErr)
		}

		return nil
	}

	// TODO добавить публикацию метрик
	return ErrNoWorkingProvider
}
