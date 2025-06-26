 AssignRoleToUser debería estar en el repositorio que maneja la relación entre usuarios y roles

 Si tienes un repositorio específico para roles o relaciones usuario-rol (RoleRepo o UserRoleRepo), entonces la función debería ir allí.
Si no tienes un repositorio separado y los roles se manejan dentro de UserRepo, entonces puede quedarse allí.

Mejor práctica
Si user_roles es una tabla de relación (many-to-many entre users y roles), es recomendable tener un UserRoleRepo o RoleRepo, porque:

Mantienes la separación de responsabilidades.
Evitas que UserRepo maneje demasiadas cosas.
Si RoleRepo ya existe y maneja asignaciones de roles, colócala allí. Si solo UserRepo gestiona estos datos, entonces está bien que permanezca ahí.

Ver si en verda se va usar la tabla intermedia lo mismo para order product bueno eso si tiene datos adicionales