package cronjob

// Con expires_at, puedes implementar un proceso automático de limpieza (e.g., un cron job o una tarea programada) que elimine los tokens vencidos y mantenga la base de datos ordenada y liviana.

func Test() {

	// * DELETE FROM refresh_tokens WHERE expires_at < NOW();

	// crear tabla yr relacionar o json en la tabla tb_auth
	// crear el código alphanumerico abc123 guardarlo
	// cronjob elimina los codigos si la tabla tb_auth en su campo
	// email_verified es true
	// crear tablas de logs

	// * Otro ccromjobs para el tema de onboarding 
	// por ejemplo el usaurio por cierto motivo cerro el onbording y tenemos una pestña de onboarding
	// e usurio marco skip y solo vio el primer onboarding
	// asi que hasta que hasta qu no lo vea le aparecerá 
	// para no sobrecargar de onboarding un cron job eliminaría onboarding que estan muy antiguos
	// y asi mejorar la experiencia de usuario
}
