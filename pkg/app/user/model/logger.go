package model

import "github.com/sirupsen/logrus"

func (u User) Logger(layerName string, action string) {
	// capa name - que se esta realizaondo, si es de retorno
	// imprimir el tipo de datos y el valor
	// u.ID = "1" para probar interface{}
	//numSpacios := 6
	//espacios := strings.Repeat(" ", 2)

	greenColor := "\033[32m"
	resetColor := "\033[0m"
	logrus.Infof(
		"%s Layer: %s Action: %s %s",
		greenColor, layerName, action, resetColor,
	)
	logrus.Infof(
		"%s ID: %s %T %s",
		"\033[33m", u.ID, u.ID, resetColor,
	)
	logrus.Infof(
		"%s Email: %s %T %s",
		"\033[33m", u.Email, u.Email, resetColor,
	)
	logrus.Infof(
		"%s Email verified: %v %T %s",
		"\033[33m", u.EmailVerified, u.EmailVerified, resetColor,
	)
	// que pasa en los punteros si viene "" al desreferenciar
	if u.Name == nil {
		// *Con espacios
		/* logrus.Infof(
			"%s Name: %s %v %T %s",
			"\033[33m", espacios, u.Name, u.Name, resetColor,
		) */
		logrus.Infof(
			"%s Name: %v %T %s",
			"\033[33m", u.Name, u.Name, resetColor,
		)

	} else {
		logrus.Infof(
			"%s Name: %v %T %s",
			"\033[33m", *u.Name, u.Name, resetColor,
		)
	}
	if u.CreatedAt == nil {
		logrus.Infof(
			"%s CreatedAt: %v %T %s",
			"\033[33m", u.CreatedAt, u.CreatedAt, resetColor,
		)
	} else {
		logrus.Infof(
			"%s CreatedAt: %v %T %s",
			"\033[33m", *u.CreatedAt, u.CreatedAt, resetColor,
		)
	}

	if u.UpdatedAt == nil {
		logrus.Infof(
			"%s UpdatedAt: %v %T %s ",
			"\033[33m", u.UpdatedAt, u.UpdatedAt, resetColor,
		)
	} else {
		logrus.Infof(
			"%s UpdatedAt: %v %T %s",
			"\033[33m", *u.UpdatedAt, u.UpdatedAt, resetColor,
		)
	}

	// crear el logguer dentro de auth para tambien imprimirlo
	// u.AuthProviderCreate.
}
