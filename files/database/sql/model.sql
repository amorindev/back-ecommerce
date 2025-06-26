-- diseño de base de datos
-- cuando hay un campo id provider, el campo de  password quedara nulo uasriamos un pointer
-- o modelar tabla auth tabla email password y otra tabla auth providers - herencia tiene 
-- su propio nombre en base de datos




-- ? en que tabla pongo created y updated auth o user?
-- eliminas un usario eliminas con cascade o transaccion
-- !cambie a user entonces cambarán las consultas
DROP TABLE IF EXISTS tb_user CASCADE;
CREATE TABLE tb_user(
    id UUID PRIMARY KEY,
    email VARCHAR(150) NOT NULL UNIQUE, -- unique desde el backend?
    --email VARCHAR(150) NOT NULL UNIQUE, OAuth garantiza email verificado google apple
    password VARCHAR(256) NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    --email_verified BOOLEAN NOT NULL DEFAULT false,
    -- ! user table
    name VARCHAR(120) NOT NULL,
    username VARCHAR(120) NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

-- cambiar a tb_auth
-- password hash
-- add provider password, google, apple
-- add provider id - el identificador único proporcionado por el proveedor
-- si es email password es nulo y como manejarlo?
-- auth table debe tener jsonData campo - como por ejemplo al enviar email el codigo alphanumeric  abf1321 
-- debería llevar un metadato AFHD123 para el send email verification
--- ia para documentar y revisa el doc que genera 

-- provider_id: El identificador único del usuario proporcionado por el proveedor (e.g., el sub en el payload de JWT de Google o Apple).
-- derificar si se va cambiar por el email el sub
-- ? UNIQUE y DEFAULT desde el backend?
DROP TABLE IF EXISTS tb_auth CASCADE;
CREATE TABLE tb_auth (
    -- provider provider
    -- ! -- Un usuario puede tener un solo método por proveedor
    id UUID PRIMARY KEY, -- on delete cascate vs transacciones
    --id UUID PRIMARY KEY REFERENCES tb_user(id), -- on delete cascate vs transacciones
    -- token del proveedor , se debe emviar al frontend como supabase?
    provider VARCHAR(90) NOT NULL,  --"local", "google", "apple"
    --provider_id VARCHAR(255) NOT NULL, -- ID único del usuario en el proveedor externo
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
    --password VARCHAR(256), -- Puede ser NULL si usa OAuth
    -- user_id UUID NOT NULL REFERENCES tb_user(id) ON DELETE CASCADE,
    -- * como manejar los delete especialmente los cascade
);

-- hash expires at
-- refresh_token hash
-- filtrar por los token que hansido revocados
-- change to session  
-- agregar el flujo completo con el device
-- refresh token cifrado

DROP TABLE IF EXISTS tb_session CASCADE;
CREATE TABLE tb_session (
    id UUID PRIMARY KEY NOT NULL,
    refresh_token TEXT NOT NULL,
    --device VARCHAR(90) NOT NULL, 
     refresh_token_id TEXT NOT NULL, -- vericar el tipo correcto
    expires_at TIMESTAMP NOT NULL, -- para eliminar por crom job
    created_at TIMESTAMP NOT NULL,
    revoked BOOLEAN NOT NULL,
    user_id UUID REFERENCES tb_auth(id)
);


-- Mejor tenerlo desde una variable una señal u otro tipo
-- al crear una tabla es el historial basicamente
-- tenog que guardar el estado para que el usario puede mantenerse logueado?
-- registrar el estado y enviarlo al frotnen
--MEJOR USAR TABA SEPARADA?
DROP TYPE IF EXISTS auth_type CASCADE;
CREATE TYPE auth_type AS ENUM ('token_refresh','logout', 'provider_switch','role_update', 'send_emailver');
-- deberia eliminarse registros ono? para no acumuar basura 
-- o seria bueno guardarlo 
DROP TABLE IF EXISTS tb_auth_changes CASCADE;
CREATE TABLE tb_auth_changes (
    -- si es integer serial o se asigna desde el backend
    -- es creado desde el backend por que es retornado como por ejemplo el usuario
    -- o quiere dar la responsabilidad al backend de insertar el objeto completo
    id SERIAL PRIMARY KEY,
    type auth_type NOT NULL,
    -- que va aqui?, son opcionales para registrar valores previos y nuevos
    -- ver si sirve esto y la tabla para enviar los server send events
    old_value TEXT NULL,
    new_value TEXT NULL,
    created_at TIMESTAMP NOT NULL,
    user_id UUID REFERENCES tb_auth(id)
);

/* DROP TABLE IF EXISTS tb_task CASCADE;
CREATE TABLE tb_task (
    id INTEGER PRIMARY KEY,
    title VARCHAR(150) NOT NULL,
    description TEXT NOT NULL,
    is_completed BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID REFERENCES tb_auth(id)
);
 */

-- ! Definir que se va hacer desde le backend como asignar valores por defecto o unique
-- ! otras operacciones documentar para que sean aplicados en las dos bases de datos
DROP TABLE IF EXISTS tb_role CASCADE;
CREATE TABLE tb_role (
    id UUID PRIMARY KEY,
    name VARCHAR(150) NOT NULL -- unique desde el back?
);

DROP TABLE IF EXISTS tb_user_role CASCADE;
CREATE TABLE tb_user_role (
    auth_id UUID REFERENCES tb_auth(id), -- on delete cascade desde el backend
    role_id UUID REFERENCES tb_role(id), -- cascade
    PRIMARY KEY (auth_id, role_id)
);

-- tabla de muchos a muchos es primary key
-- que tablas debe usar id  o uuid, CREATED AT, STATE, jsonmetadaso lo de edteam
-- CREAR BASE DE DATOS
-- CREAR USUARIO Y PERMISOS
-- CREAR VISTAS
-- ON DELETE CASCADE

-- cuando crear indices
-- CREATE INDEX idx_user_id ON refresh_tokens(user_id);


/* CREATE TABLE tb_country (
    id SERIAL NOT NULL,
    prefix -- 51 peru
    validate -- el numero 
    number --
    is_verified
); */

--! sera verify-email por que tambien se usara para signin si el suario auno no confirmo su email
/* CREATE TABLE tb_otp(
    id SERIAL PRIMARY KEY, -- serial por que no se va a exponer al usario
    otp_code VARCHAR(10) NOT NULL, -- dependiendo 
    purpose VARCHAR(20) NOT NULL, -- e.g. 'verify_email', 'reset_password', 'sign-up', 'update-email', 'delete-account'
    expires_at TIMESTAMP NOT NULL,
    used BOOLEAN NOT NULL, --se puede dafaut false dejarlo al backend
    created_at TIMESTAMP NOT NULL,
    auth_id INTEGER REFERENCES tb_auth(id) --! verificar si todos los otp codes son enviados cuando este el auth
);
 */

/* CREATE TABLE tb_user(
    id UUID PRIMARY KEY,
    name VARCHAR(120) NOT NULL,
    username VARCHAR(120) NOT NULL
);

CREATE TABLE tb_auth (
    id UUID PRIMARY KEY, 
    email VARCHAR(150) NOT NULL UNIQUE,
    password VARCHAR(256) NOT NULL,
    email_verified BOOLEAN NOT NULL DEFAULT false,
    provider VARCHAR(90) NOT NULL, 
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
); */