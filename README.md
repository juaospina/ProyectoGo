ProyectoGo
Este es el repositorio del proyecto ProyectoGo, una aplicación web desarrollada en Go que permite manejar una lista de productos. Esta aplicación fue construida utilizando el framework Gin para manejar las rutas y peticiones HTTP.

La aplicación cuenta con las siguientes funcionalidades:

Crear productos: permite crear un nuevo producto enviando una petición POST a la ruta /products con un JSON con la información del producto.
Obtener productos: permite obtener todos los productos existentes enviando una petición GET a la ruta /products.
Obtener producto por ID: permite obtener un producto específico enviando una petición GET a la ruta /products/:id.
Actualizar producto: permite actualizar un producto existente enviando una petición PUT a la ruta /products/:id con un JSON con la información actualizada del producto.
Actualizar producto parcialmente: permite actualizar parcialmente un producto existente enviando una petición PATCH a la ruta /products/:id con un JSON con la información a actualizar del producto.
Eliminar producto: permite eliminar un producto existente enviando una petición DELETE a la ruta /products/:id.
Buscar productos: permite buscar productos con un precio mayor a un valor específico enviando una petición GET a la ruta /products/search?priceGt=100
Además, se incluyen validaciones en las funciones para asegurar que los datos enviados cumplen con las condiciones necesarias.

Para usar este proyecto, es necesario tener Go instalado en tu computadora. Luego, clona este repositorio y ejecuta el comando go run main.go en la carpeta del proyecto. La aplicación se ejecutará en el puerto 3000 por defecto.

Si deseas hacer cambios en el proyecto, te invitamos a crear una rama y hacer tus aportes en ella. Una vez hechos, crea una pull request para que podamos revisar tus cambios y considerarlos para incluirlos en la rama principal.

Si tienes alguna duda o sugerencia, no dudes en crear un issue en este repositorio. ¡Esperamos que encuentres este proyecto útil!