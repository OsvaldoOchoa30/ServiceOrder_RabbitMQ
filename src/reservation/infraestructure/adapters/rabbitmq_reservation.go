package adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"
    "service_order/src/reservation/domain/entities"

	amqp "github.com/rabbitmq/amqp091-go"
)

type MyExchangeLogs struct {
    ch *amqp.Channel
}

func NewRabbitRepository(ch *amqp.Channel) *MyExchangeLogs {
    if err := ch.ExchangeDeclare(
        "logs",   // Nombre del exchange
        "fanout", // Tipo del exchange
        true,     // Durable
        false,    // Auto-deleted
        false,    // Internal
        false,    // No-wait
        nil,      // Argumentos
    ); err != nil {
        log.Fatalf("Error al declarar el exchange: %v", err)
    }

    return &MyExchangeLogs{ch: ch}
}

// Método Save actualizado para coincidir con la interfaz
func (ch *MyExchangeLogs) Save(reservation *entities.Reservation) error {
    // Crear un objeto con los datos de la reserva


    // Convertir a JSON
    body, err := json.Marshal(reservation)
    if err != nil {
        return fmt.Errorf("error al serializar la reserva: %v", err)
    }

    log.Printf("Enviando mensaje: %s", body)

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    // Publicar el mensaje en RabbitMQ
    if err := ch.ch.PublishWithContext(ctx,
        "logs",  // Exchange
        "",      // Routing key
        false,   // Mandatory
        false,   // Immediate
        amqp.Publishing{
            ContentType: "application/json", // Tipo de contenido
            Body:        body,               // Cuerpo del mensaje
        }); err != nil {
        return fmt.Errorf("error al enviar el mensaje a RabbitMQ: %v", err)
    }

    log.Printf(" [x] Enviado: %s", body)
    return nil
}