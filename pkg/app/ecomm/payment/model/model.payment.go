package model

import "time"

//* Si tu sistema requiere flexibilidad, lo más común y escalable es una relación 1:N (uno a muchos), incluso
// * sinormalmente hay solo un pago por pedido.

type PaymentStatus string

const (
	PaymentPending  PaymentStatus = "pending"
	PaymentPaid     PaymentStatus = "paid"
	PaymentFailed   PaymentStatus = "failed"
	PaymentRefunded PaymentStatus = "refunded"
)

type PaymentMethod string

const (
	MethodStripe PaymentMethod = "stripe"
	MethodPaypal PaymentMethod = "paypal"
	MethodApple  PaymentMethod = "apple"
)


// ! Definir los estados del payment pending 
// * para actualizar el pago a estado pagado en el metadata envio el payment id
// * y tambien el order id cual enviar o los dos

// * ver que campos se mostraran en el json o solo el Oder{}
type Payment struct {
	// * oder con payment es 1 a 1  a nose ser que se pueda pagar una pare con stripe
	// * y la otra a credito ahi si seria muchos a muchos y guardar en la tabla
	// * intermedia el monto que se pago por cada método de pago y el estado como seria el modelo
	// * de momento simple
	// * Si es uno a uno no seria mejor llmarolo OrderID en ves de ID ver -
	// * se puede pero el nombre seria _id del bson para que no cree otro
	// * ver que se va a retornar de esta entidad cuando se crea el order al responder en  la api
	// status (pending, paid, failed, refunded, etc.)
	// payment_method (enum: 'stripe', 'paypal', 'apple', etc.)
	// ver el video de siseño de base de datos
	// provider_payment_id (id que devuelve Stripe/PayPal) webhook  me parece
	ID      interface{} `json:"id" bson:"_id"`
	OrderID interface{} `json:"order_id" bson:"order_id"`
	// puede calcular mal mos amount de order y payment ver
	//Amount        *float64      `json:"amount" bson:"amount"` esta en la order collection bueno puede
	Currency      *string       `json:"currency" bson:"currency"` // verificar que concida con el webhook?
	Status        PaymentStatus `json:"status" bson:"status"`     // desde el weehook
	PaymentMethod PaymentMethod `json:"payment_method" bson:"payment_method"`
	// no seria meojr una nueva tabla payment_provider para separa este campo  por que me parece
	// que se va a repetir de momento chill
	ProviderPaymentID *string    `json:"provider_payment_id"` // desde el webhook
	CreatedAt         *time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt         *time.Time `json:"updated_at" bson:"updated_at"`
}

/*
import 'package:flutter/material.dart';
import 'package:flutter_svg/flutter_svg.dart';

class SvgPictureWithFallback extends StatefulWidget {
  final String imageUrl;
  final String fallbackImageUrl;

  const SvgPictureWithFallback({Key? key, required this.imageUrl, required this.fallbackImageUrl}) : super(key: key);

  @override
  State<SvgPictureWithFallback> createState() => _SvgPictureWithFallbackState();
}

class _SvgPictureWithFallbackState extends State<SvgPictureWithFallback> {
  bool _isLoading = true;
  String? _error;

  @override
  Widget build(BuildContext context) {
    return FutureBuilder(
      future: Future.value(true), // Reemplaza con el futuro que retorna la carga de la imagen SVG
      builder: (BuildContext context, AsyncSnapshot<dynamic> snapshot) {
        if (snapshot.hasError) {
          _error = "Error al cargar la imagen";
          _isLoading = false;
          return Center(
            child: SvgPicture.asset(
              widget.fallbackImageUrl, // Imagen de respaldo en caso de error
            ),
          );
        }
        if (snapshot.connectionState == ConnectionState.done && !_isLoading && _error != null) {
          return Center(
            child: Text(_error!),
          );
        }
        if (_isLoading) {
          return Center(
            child: const CircularProgressIndicator(), // Puedes mostrar un indicador de carga
          );
        }

        return Center(
          child: SvgPicture.network(
            widget.imageUrl,
            // height: 100, // Ajusta el tamaño si es necesario
          ),
        );
      },
    );
  }
}

*/
