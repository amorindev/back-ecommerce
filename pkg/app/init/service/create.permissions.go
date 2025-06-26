package service

import (
	"context"
	"encoding/json"
	"io"
	"os"

	"com.fernando/pkg/app/permission/errors"
	"com.fernando/pkg/app/permission/model"
)

func (s *Service) CreatePermissions() error {
	permissionNames := []string{"create-product",
		"update-product",
		"delete-product",
		"read-product"}

	var permisions []*model.Permission
	for _, pName := range permissionNames {
		var permision model.Permission
		permision.Name = &pName
		permisions = append(permisions, &permision)
	}

	for _, permision := range permisions {
		p, err := s.PermissionRepo.GetByName(context.Background(), *permision.Name)
		if err != nil && err != errors.ErrPermissionNotFound {
			return err
		}
		if p != nil {
			// si todos existen
			continue
		}
		err = s.PermissionRepo.Insert(context.Background(), permision)
		if err != nil {
			return err
		}
	}

	// * Insertar permissions puede ser creando tambien un slice mas facil
	// * hay una tipo de slice para que cuando se repite no lo duplica
	/* var permissions []*model.Permission
	for _, data := range permissionData.Data {
		for _, permisionName := range data.NamePermissions {
			p, err := s.PermissionRepo.GetByName(context.Background(), permisionName)
			if err != nil {
				if err != errors.ErrPermissionNotFound {
					return err
				}
			}
			if p != nil {
				continue
			}
			var permision model.Permission
			permision.Name = &permisionName
			permissions = append(permissions, &permision)
		}
	}

	// ! Consultar si existe uan de las permission lo mismo para los demas
	// ! DDD veo que dentro del repo lo estan verfiicando si existe
	// ! como manjar si existe per aui noes critico
	t, err := s.PermissionRepo.ExistOne(context.Background(), permissions)
	if err != nil {
		return err
	}
	//fmt.Printf("len: %d\n", len(permissions))
	//fmt.Printf("Exis: %d\n", len(t))

	err = s.PermissionRepo.InsertMany(context.Background(), permissions)
	if err != nil {
		return err
	}
	log.Fatal("test")

	// * Assign Permissions to Role
	//var rolePermissions []*model.RolePermission
	for _, p := range permissionData.Data {
		role, err := s.RoleRepo.GetByName(context.Background(), p.RoleName)
		if err != nil {
			return err
		}
		var perm []*model.Permission // este es para la primera forma
		for _, name := range p.NamePermissions {
			permission, err := s.PermissionRepo.GetByName(context.Background(), name)
			if err != nil {
				return err
			}

			//var rolePermision model.RolePermission // segunda forma
			//rolePermision.RoleID = role.ID
			//rolePermision.PermissionID = permission.ID
			//rolePermissions = append(rolePermissions, &rolePermision) // segunda forma
			perm = append(perm, permission) // segunda forma
		}
		err = s.PermissionRepo.AssignPermissionsToRole(context.Background(), role.ID.(string), perm)
		if err != nil {
			return err
		}
	} */
	// * ver si existe la relacion - de momento solo insertarre

	// * insertarlo todo lo que es assignar como user_role role_permission se debe crear
	// * antes una funcion aparte para verificar que no exista la relacion
	// * de momentto esta todo junto dentro de una sola funcion
	// * Exites dos formas pasando aparete el roleID y la lista de RolePermission{}
	// * (si apotro con roleID podria ser dentro del role anterior)
	// * o enviar RolePermission{} ya pasado ambos voy a probar este ,
	// * ver por que ambos si usas RolePermission debes e crear su repo su DDD
	// * la otra opcion es dentro de los for
	/* err = s.PermissionRepo.AssignPermissionsToRole2(context.Background(),permissions)
	if err != nil {
	  return err
	} */
	return nil
}

/* func (s *Service) CreatePermissions2() error {
	adminP := []string{"create-product", "update-product", "delete-product", "read-product"}
	userP := []string{"read-product"}
	roles := map[string][]string{"ADMIN": adminP, "USER": userP}

	//roles3 := [][]string{"ste":adminP,}

	for key := range roles {
		role, err := s.RoleRepo.GetByName(context.Background(), key)
		if err != nil {
			return err
		}
		// verificar si coincide la permission
		for _, data := range roles {
			for _, value := range data {
				fmt.Printf("Role %+v\n", role)
				fmt.Printf(" Data %v\n", data)
				fmt.Printf(" value %v\n", value)
			}
		}
		//var permission model.Permission
		//permission.
	}

	return nil
}
*/

type InitPermission struct {
	Data []InitPermission2 `json:"data"`
}

type InitPermission2 struct {
	RoleName        string   `json:"role_name"`
	NamePermissions []string `json:"name_permissions"`
}

func (s *Service) CreatePermissionsFromJson() error {
	jsonFile, err := os.Open("pkg/app/init/files/data/insert_permissions.json")
	if err != nil {
		return err
	}

	defer jsonFile.Close()

	byteValue, err := io.ReadAll(jsonFile)
	if err != nil {
		return err
	}

	var permissionData InitPermission

	err = json.Unmarshal(byteValue, &permissionData)
	if err != nil {
		return err
	}

	//fmt.Printf("Permission %v\n", permissionData)
	// * Insertar permissions puede ser creando tambien un slice mas facil
	// * hay una tipo de slice para que cuando se repite no lo duplica
	/* var permissions []*model.Permission
	for _, data := range permissionData.Data {
		for _, permisionName := range data.NamePermissions {
			p, err := s.PermissionRepo.GetByName(context.Background(), permisionName)
			if err != nil {
				if err != errors.ErrPermissionNotFound {
					return err
				}
			}
			if p != nil {
				continue
			}
			var permision model.Permission
			permision.Name = &permisionName
			permissions = append(permissions, &permision)
		}
	}

	// ! Consultar si existe uan de las permission lo mismo para los demas
	// ! DDD veo que dentro del repo lo estan verfiicando si existe
	// ! como manjar si existe per aui noes critico
	t, err := s.PermissionRepo.ExistOne(context.Background(), permissions)
	if err != nil {
		return err
	}
	fmt.Printf("len: %d\n", len(permissions))
	fmt.Printf("Exis: %d\n", len(t))

	err = s.PermissionRepo.InsertMany(context.Background(), permissions)
	if err != nil {
		return err
	}
	log.Fatal("test")

	// * Assign Permissions to Role
	//var rolePermissions []*model.RolePermission
	for _, p := range permissionData.Data {
		role, err := s.RoleRepo.GetByName(context.Background(), p.RoleName)
		if err != nil {
			return err
		}
		var perm []*model.Permission // este es para la primera forma
		for _, name := range p.NamePermissions {
			permission, err := s.PermissionRepo.GetByName(context.Background(), name)
			if err != nil {
				return err
			}

			//var rolePermision model.RolePermission // segunda forma
			//rolePermision.RoleID = role.ID
			//rolePermision.PermissionID = permission.ID
			//rolePermissions = append(rolePermissions, &rolePermision) // segunda forma
			perm = append(perm, permission) // segunda forma
		}
		err = s.PermissionRepo.AssignPermissionsToRole(context.Background(), role.ID.(string), perm)
		if err != nil {
			return err
		}
	} */
	// * ver si existe la relacion - de momento solo insertarre

	// * insertarlo todo lo que es assignar como user_role role_permission se debe crear
	// * antes una funcion aparte para verificar que no exista la relacion
	// * de momentto esta todo junto dentro de una sola funcion
	// * Exites dos formas pasando aparete el roleID y la lista de RolePermission{}
	// * (si apotro con roleID podria ser dentro del role anterior)
	// * o enviar RolePermission{} ya pasado ambos voy a probar este ,
	// * ver por que ambos si usas RolePermission debes e crear su repo su DDD
	// * la otra opcion es dentro de los for
	/* err = s.PermissionRepo.AssignPermissionsToRole2(context.Background(),permissions)
	if err != nil {
	  return err
	} */
	return nil
}
