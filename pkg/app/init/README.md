Este es un servicio que se va ejecutar cuando arranque laaplicación
Puede crear sus propios puertos y models y repository si es necesario o usar de los
ya creados en otros servicios

Se enfoca en mejoarar la experiencia de desarrollo, no se aún si se puede usar para produccion
o se tenddría que mojorar algo me parece que se puede para el usario admin

Aqui deberia ir migrate funcion de la base de datos
para no afectar en produccion Cuando Corrers Run de init validar if entorno es env - ver

para insertar se usará manualmente o el mismo run u otro?

<!-- ! verificar si los roles ususario cuando se hace un rebuild no se vuelvan a crear en la base de datos>
<!-- los lista de roles el usuario admin>


todavía no se donde poner los files lo dejaré aqui archivos png

<!-- * al guardar el file se debe devolver el url para agregarlo a la base de datos>

<!-- * init cambiar a data laoader>
<!-- * no te parece mejor que dentro de cada uno tenga su interface ProductDataloader CategoryDataLoader>
<!-- * y no uno muy general vertical slicing DDD>
