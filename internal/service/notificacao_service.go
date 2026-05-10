package service

import(
	"fmt"
	"time"
)

type NotificacaoService struct{}

func NewNotificacaoService() *NotificacaoService {
	return &NotificacaoService{}
}

// EnviarBoasVindas simula um envio de notificação demorado
func (s *NotificacaoService) EnviarBoasVindas(nomeCliente string) {
	fmt.Printf("\n[Background] 📧 Preparando e-mail de boas-vindas para: %s...\n", nomeCliente)
	
	// Simula um delay de 5 segundos
	time.Sleep(5 * time.Second)
	
	fmt.Printf("[Background] ✅ E-mail enviado com sucesso para %s!\n", nomeCliente)
}